package internal

import (
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"encoding/binary"
	"fmt"
	"net"
	"sync"
	"time"
)

// SModbusServer Modbus TCP 服务器
type SModbusServer struct {
	config      *SModbusConfig
	handler     *SModbusDeviceHandler
	listener    net.Listener
	ctx         context.Context
	cancel      context.CancelFunc
	isRunning   bool
	mu          sync.RWMutex
	connections map[net.Conn]bool
	connMu      sync.RWMutex
}

// NewModbusServer 创建Modbus服务器
func NewModbusServer(config *SModbusConfig, handler *SModbusDeviceHandler) *SModbusServer {
	return &SModbusServer{
		config:      config,
		handler:     handler,
		connections: make(map[net.Conn]bool),
	}
}

// Start 启动Modbus服务器
func (s *SModbusServer) Start(ctx context.Context) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isRunning {
		c_log.Warning(ctx, "Modbus服务器已经在运行中")
		return nil
	}

	// 创建可取消的上下文，并设置远程协议类型
	s.ctx, s.cancel = context.WithCancel(ctx)
	s.ctx = context.WithValue(s.ctx, c_enum.ELogTypeRemote, c_base.ConstRemoteModbus)

	// 启动TCP监听器
	address := fmt.Sprintf(":%d", s.config.ListenPort)
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return fmt.Errorf("启动Modbus TCP服务器失败: %v", err)
	}

	s.listener = listener
	s.isRunning = true

	// 启动连接处理协程
	go s.acceptConnections()

	c_log.Infof(ctx, "Modbus TCP服务器启动成功，监听端口: %d", s.config.ListenPort)
	return nil
}

// Stop 停止Modbus服务器
func (s *SModbusServer) Stop() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isRunning {
		c_log.Warning(context.Background(), "Modbus服务器未运行")
		return nil
	}

	// 取消上下文
	if s.cancel != nil {
		s.cancel()
	}

	// 关闭监听器
	if s.listener != nil {
		s.listener.Close()
	}

	// 关闭所有连接
	s.connMu.Lock()
	for conn := range s.connections {
		conn.Close()
		delete(s.connections, conn)
	}
	s.connMu.Unlock()

	s.isRunning = false
	c_log.Info(context.Background(), "Modbus TCP服务器已停止")
	return nil
}

// acceptConnections 接受连接
func (s *SModbusServer) acceptConnections() {
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			conn, err := s.listener.Accept()
			if err != nil {
				select {
				case <-s.ctx.Done():
					return
				default:
					c_log.Errorf(s.ctx, "接受连接失败: %v", err)
					continue
				}
			}

			// 记录连接
			s.connMu.Lock()
			s.connections[conn] = true
			s.connMu.Unlock()

			// 处理连接
			go s.handleConnection(conn)
		}
	}
}

// handleConnection 处理连接
func (s *SModbusServer) handleConnection(conn net.Conn) {
	defer func() {
		conn.Close()
		s.connMu.Lock()
		delete(s.connections, conn)
		s.connMu.Unlock()
	}()

	// 设置连接超时
	conn.SetDeadline(time.Now().Add(30 * time.Second))

	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			// 读取Modbus请求
			request, err := s.readModbusRequest(conn)
			if err != nil {
				if err.Error() != "EOF" {
					c_log.Debugf(s.ctx, "读取Modbus请求失败: %v", err)
				}
				return
			}

			// 处理请求
			response, err := s.processModbusRequest(request)
			if err != nil {
				c_log.Debugf(s.ctx, "处理Modbus请求失败: %v", err)
				// 发送异常响应
				s.sendExceptionResponse(conn, request, 0x04) // Slave Device Failure
				continue
			}

			// 发送响应
			if err := s.sendModbusResponse(conn, response); err != nil {
				c_log.Debugf(s.ctx, "发送Modbus响应失败: %v", err)
				return
			}
		}
	}
}

// SModbusRequest Modbus请求结构体
type SModbusRequest struct {
	TransactionID uint16 // 事务ID
	ProtocolID    uint16 // 协议ID
	Length        uint16 // 长度
	UnitID        uint8  // 单元ID
	FunctionCode  uint8  // 功能码
	Data          []byte // 数据
}

// SModbusResponse Modbus响应结构体
type SModbusResponse struct {
	TransactionID uint16 // 事务ID
	ProtocolID    uint16 // 协议ID
	Length        uint16 // 长度
	UnitID        uint8  // 单元ID
	FunctionCode  uint8  // 功能码
	Data          []byte // 数据
}

// readModbusRequest 读取Modbus请求
func (s *SModbusServer) readModbusRequest(conn net.Conn) (*SModbusRequest, error) {
	// 读取MBAP头部（6字节）
	header := make([]byte, 6)
	if _, err := conn.Read(header); err != nil {
		return nil, err
	}

	// 解析MBAP头部
	transactionID := binary.BigEndian.Uint16(header[0:2])
	protocolID := binary.BigEndian.Uint16(header[2:4])
	length := binary.BigEndian.Uint16(header[4:6])
	unitID := header[5]

	// 读取剩余数据
	dataLength := int(length) - 1 // 减去UnitID字段
	if dataLength < 0 {
		return nil, fmt.Errorf("无效的数据长度: %d", length)
	}

	data := make([]byte, dataLength)
	if _, err := conn.Read(data); err != nil {
		return nil, err
	}

	functionCode := data[0]
	requestData := data[1:]

	return &SModbusRequest{
		TransactionID: transactionID,
		ProtocolID:    protocolID,
		Length:        length,
		UnitID:        unitID,
		FunctionCode:  functionCode,
		Data:          requestData,
	}, nil
}

// processModbusRequest 处理Modbus请求
func (s *SModbusServer) processModbusRequest(req *SModbusRequest) (*SModbusResponse, error) {
	// 使用设备处理器处理请求
	responseData, err := s.handler.HandleRequest(req.UnitID, req.FunctionCode, req.Data)
	if err != nil {
		return nil, err
	}

	// 构建响应
	response := &SModbusResponse{
		TransactionID: req.TransactionID,
		ProtocolID:    req.ProtocolID,
		Length:        uint16(2 + len(responseData)), // UnitID + FunctionCode + Data
		UnitID:        req.UnitID,
		FunctionCode:  req.FunctionCode,
		Data:          responseData,
	}

	return response, nil
}

// sendModbusResponse 发送Modbus响应
func (s *SModbusServer) sendModbusResponse(conn net.Conn, resp *SModbusResponse) error {
	// 构建完整的响应数据
	responseData := make([]byte, 6+2+len(resp.Data)) // MBAP + UnitID + FunctionCode + Data

	// MBAP头部
	binary.BigEndian.PutUint16(responseData[0:2], resp.TransactionID)
	binary.BigEndian.PutUint16(responseData[2:4], resp.ProtocolID)
	binary.BigEndian.PutUint16(responseData[4:6], resp.Length)

	// PDU部分
	responseData[6] = resp.UnitID
	responseData[7] = resp.FunctionCode
	copy(responseData[8:], resp.Data)

	// 发送响应
	_, err := conn.Write(responseData)
	return err
}

// sendExceptionResponse 发送异常响应
func (s *SModbusServer) sendExceptionResponse(conn net.Conn, req *SModbusRequest, exceptionCode uint8) {
	response := &SModbusResponse{
		TransactionID: req.TransactionID,
		ProtocolID:    req.ProtocolID,
		Length:        3, // UnitID + FunctionCode + ExceptionCode
		UnitID:        req.UnitID,
		FunctionCode:  req.FunctionCode | 0x80, // 设置异常位
		Data:          []byte{exceptionCode},
	}

	s.sendModbusResponse(conn, response)
}

// IsRunning 检查服务器是否正在运行
func (s *SModbusServer) IsRunning() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.isRunning
}

// GetConnectionCount 获取连接数量
func (s *SModbusServer) GetConnectionCount() int {
	s.connMu.RLock()
	defer s.connMu.RUnlock()
	return len(s.connections)
}

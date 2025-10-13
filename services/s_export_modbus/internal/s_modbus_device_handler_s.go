package internal

import (
	"common"
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_log"
	"context"
	"encoding/binary"
	"fmt"
	"sync"
	"time"
)

// SModbusDeviceHandler Modbus设备处理器
type SModbusDeviceHandler struct {
	deviceMaps map[string]*SDeviceRegisterMap // 设备ID -> 设备映射
	mu         sync.RWMutex                   // 读写锁
}

// NewModbusDeviceHandler 创建设备处理器
func NewModbusDeviceHandler() *SModbusDeviceHandler {
	return &SModbusDeviceHandler{
		deviceMaps: make(map[string]*SDeviceRegisterMap),
	}
}

// UpdateDeviceMaps 更新设备映射
func (h *SModbusDeviceHandler) UpdateDeviceMaps(deviceMaps map[string]*SDeviceRegisterMap) {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.deviceMaps = make(map[string]*SDeviceRegisterMap)
	for deviceId, deviceMap := range deviceMaps {
		h.deviceMaps[deviceId] = deviceMap
	}
}

// HandleRequest 处理Modbus请求
func (h *SModbusDeviceHandler) HandleRequest(unitID uint8, functionCode uint8, data []byte) ([]byte, error) {
	switch functionCode {
	case 0x03: // Read Holding Registers
		return h.handleReadHoldingRegisters(unitID, data)
	default:
		return nil, fmt.Errorf("不支持的功能码: 0x%02X", functionCode)
	}
}

// handleReadHoldingRegisters 处理读保持寄存器请求
func (h *SModbusDeviceHandler) handleReadHoldingRegisters(unitID uint8, data []byte) ([]byte, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	// 查找对应的设备
	deviceMap := h.findDeviceByModbusId(unitID)
	if deviceMap == nil {
		return nil, fmt.Errorf("设备不存在，ModbusId: %d", unitID)
	}

	// 解析请求参数
	if len(data) < 4 {
		return nil, fmt.Errorf("请求数据长度不足")
	}
	startAddr := binary.BigEndian.Uint16(data[0:2])
	quantity := binary.BigEndian.Uint16(data[2:4])

	// 检查地址范围
	if startAddr < deviceMap.StartAddr || startAddr+quantity > deviceMap.StartAddr+deviceMap.TotalRegisters {
		return nil, fmt.Errorf("地址越界: 请求地址 %d-%d, 设备地址范围 %d-%d",
			startAddr, startAddr+quantity-1, deviceMap.StartAddr, deviceMap.StartAddr+deviceMap.TotalRegisters-1)
	}

	// 读取寄存器数据
	registers, err := h.readRegisters(deviceMap, startAddr, quantity)
	if err != nil {
		c_log.Errorf(context.Background(), "读取寄存器失败: %v", err)
		return nil, fmt.Errorf("读取寄存器失败: %v", err)
	}

	// 构建响应数据
	responseData := make([]byte, 1+len(registers)*2)
	responseData[0] = byte(len(registers) * 2) // 字节数

	for i, register := range registers {
		binary.BigEndian.PutUint16(responseData[1+i*2:1+i*2+2], register)
	}

	return responseData, nil
}

// findDeviceByModbusId 根据ModbusId查找设备
func (h *SModbusDeviceHandler) findDeviceByModbusId(modbusId uint8) *SDeviceRegisterMap {
	for _, deviceMap := range h.deviceMaps {
		if deviceMap.ModbusId == modbusId {
			return deviceMap
		}
	}
	return nil
}

// readRegisters 读取寄存器数据
func (h *SModbusDeviceHandler) readRegisters(deviceMap *SDeviceRegisterMap, startAddr, quantity uint16) ([]uint16, error) {
	// 获取设备实例
	device := common.GetDeviceManager().GetDeviceById(deviceMap.DeviceId)
	if device == nil {
		return nil, fmt.Errorf("设备 %s 不存在", deviceMap.DeviceId)
	}

	// 计算相对地址
	relativeStartAddr := startAddr - deviceMap.StartAddr
	relativeEndAddr := relativeStartAddr + quantity

	// 构建寄存器数据
	registers := make([]uint16, quantity)

	// 遍历映射查找对应的点位
	for _, mapping := range deviceMap.Mappings {
		mappingStart := mapping.StartOffset
		mappingEnd := mappingStart + mapping.RegisterCount

		// 检查是否有重叠
		if relativeStartAddr < mappingEnd && relativeEndAddr > mappingStart {
			// 计算重叠范围
			overlapStart := max(relativeStartAddr, mappingStart)
			overlapEnd := min(relativeEndAddr, mappingEnd)

			// 读取重叠部分的数据
			overlapRegisters, err := h.readMappingRegisters(device, mapping, overlapStart-mappingStart, overlapEnd-overlapStart)
			if err != nil {
				return nil, fmt.Errorf("读取点位 %s 失败: %v", mapping.Point.GetKey(), err)
			}

			// 复制到结果中
			for i, register := range overlapRegisters {
				resultIndex := overlapStart - relativeStartAddr + uint16(i)
				if resultIndex < quantity {
					registers[resultIndex] = register
				}
			}
		}
	}

	return registers, nil
}

// readMappingRegisters 读取映射的寄存器数据
func (h *SModbusDeviceHandler) readMappingRegisters(device c_base.IDevice, mapping *SRegisterMapping, offset, count uint16) ([]uint16, error) {
	// 处理固定点位
	if mapping.Point == nil {
		return h.readFixedRegisters(device, offset, count)
	}

	// 处理设备点位
	return h.readDevicePointRegisters(device, mapping, offset, count)
}

// readFixedRegisters 读取固定寄存器数据
func (h *SModbusDeviceHandler) readFixedRegisters(device c_base.IDevice, offset, count uint16) ([]uint16, error) {
	registers := make([]uint16, count)

	// 第一个固定点位：设备在线状态
	if offset == 0 && count >= 1 {
		status := device.GetProtocolStatus()
		var value interface{}
		if status == c_enum.EProtocolConnected {
			value = true
		} else {
			value = false
		}

		// 使用 default 点位进行编码
		encodedRegisters, err := EncodeValueToRegisters(value, c_default.VPointSystemOnlineStatus.ValueType)
		if err != nil {
			return nil, fmt.Errorf("编码在线状态失败: %v", err)
		}

		if len(encodedRegisters) > 0 {
			registers[0] = encodedRegisters[0]
		}
	}

	// 第二个固定点位：通讯时间戳
	if offset <= 1 && offset+count > 1 {
		timestamp := uint32(time.Now().Unix())

		// 使用 default 点位进行编码
		encodedRegisters, err := EncodeValueToRegisters(timestamp, c_default.VPointSystemTimestamp.ValueType)
		if err != nil {
			return nil, fmt.Errorf("编码时间戳失败: %v", err)
		}

		// 复制编码后的寄存器数据
		for i, register := range encodedRegisters {
			if int(offset)+i < int(count) {
				registers[i] = register
			}
		}
	}

	return registers, nil
}

// readDevicePointRegisters 读取设备点位寄存器数据
func (h *SModbusDeviceHandler) readDevicePointRegisters(device c_base.IDevice, mapping *SRegisterMapping, offset, count uint16) ([]uint16, error) {
	// 获取点位值
	pointValue := device.GetProtocolPointValue(&c_base.SProtocolPoint{
		SPoint: &c_base.SPoint{
			Key: mapping.Point.GetKey(),
		},
	})

	if pointValue == nil {
		// 点位值为空，返回默认值
		return make([]uint16, count), nil
	}

	// 将点位值转换为寄存器
	registers, err := EncodeValueToRegisters(pointValue.GetValue(), mapping.Point.GetValueType())
	if err != nil {
		return nil, fmt.Errorf("编码点位值失败: %v", err)
	}

	// 返回请求的部分
	startIndex := int(offset)
	endIndex := startIndex + int(count)
	if startIndex >= len(registers) {
		return make([]uint16, count), nil
	}
	if endIndex > len(registers) {
		endIndex = len(registers)
	}

	return registers[startIndex:endIndex], nil
}

// GetDeviceStatus 获取设备状态
func (h *SModbusDeviceHandler) GetDeviceStatus() []*SModbusDeviceStatus {
	h.mu.RLock()
	defer h.mu.RUnlock()

	var statusList []*SModbusDeviceStatus
	for deviceId, deviceMap := range h.deviceMaps {
		status := &SModbusDeviceStatus{
			DeviceId:       deviceId,
			ModbusId:       deviceMap.ModbusId,
			StartAddr:      deviceMap.StartAddr,
			IsOnline:       deviceMap.IsOnline,
			LastUpdateTime: deviceMap.LastUpdateTime,
		}
		statusList = append(statusList, status)
	}

	return statusList
}

// min 返回两个数中的较小值
func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

// max 返回两个数中的较大值
func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

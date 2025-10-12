package internal

import (
	"common"
	"common/c_log"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"s_db"
	"s_db/s_db_basic"
)

// SModbusManager Modbus管理器
type SModbusManager struct {
	server       *SModbusServer
	handler      *SModbusDeviceHandler
	ctx          context.Context
	cancel       context.CancelFunc
	isRunning    bool
	mu           sync.RWMutex
	deviceMaps   map[string]*SDeviceRegisterMap  // 设备ID -> 设备映射
	deviceStatus map[string]*SModbusDeviceStatus // 设备状态
}

var (
	modbusManagerInstance *SModbusManager
	modbusManagerOnce     sync.Once
)

// GetModbusManager 获取Modbus管理器单例
func GetModbusManager() *SModbusManager {
	modbusManagerOnce.Do(func() {
		modbusManagerInstance = &SModbusManager{
			deviceMaps:   make(map[string]*SDeviceRegisterMap),
			deviceStatus: make(map[string]*SModbusDeviceStatus),
		}
	})
	return modbusManagerInstance
}

// Start 启动Modbus管理器
func (m *SModbusManager) Start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.isRunning {
		c_log.Warning(ctx, "Modbus管理器已经在运行中")
		return nil
	}

	// 创建可取消的上下文
	m.ctx, m.cancel = context.WithCancel(ctx)

	// 加载配置并构建设备映射
	err := m.loadConfigs(m.ctx)
	if err != nil {
		c_log.Errorf(m.ctx, "加载Modbus配置失败: %+v", err)
		return err
	}

	// 创建设备处理器
	m.handler = NewModbusDeviceHandler()
	m.handler.UpdateDeviceMaps(m.deviceMaps)

	// 创建并启动服务器
	m.server = NewModbusServer(m.getConfig(), m.handler)
	err = m.server.Start(m.ctx)
	if err != nil {
		c_log.Errorf(m.ctx, "启动Modbus服务器失败: %+v", err)
		return err
	}

	m.isRunning = true
	c_log.Infof(m.ctx, "Modbus管理器启动成功，共 %d 个设备", len(m.deviceMaps))
	return nil
}

// Stop 停止Modbus管理器
func (m *SModbusManager) Stop(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isRunning {
		c_log.Warning(ctx, "Modbus管理器未运行")
		return nil
	}

	// 停止服务器
	if m.server != nil {
		err := m.server.Stop()
		if err != nil {
			c_log.Errorf(ctx, "停止Modbus服务器失败: %v", err)
		}
	}

	// 取消上下文
	if m.cancel != nil {
		m.cancel()
	}

	// 清空设备映射
	m.deviceMaps = make(map[string]*SDeviceRegisterMap)
	m.deviceStatus = make(map[string]*SModbusDeviceStatus)

	m.isRunning = false
	c_log.Infof(ctx, "Modbus管理器已停止")
	return nil
}

// Reload 重新加载配置
func (m *SModbusManager) Reload(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isRunning {
		c_log.Warning(ctx, "Modbus管理器未运行，无法重载")
		return nil
	}

	// 重新加载配置
	err := m.loadConfigs(ctx)
	if err != nil {
		c_log.Errorf(ctx, "重新加载Modbus配置失败: %+v", err)
		return err
	}

	// 更新设备处理器
	if m.handler != nil {
		m.handler.UpdateDeviceMaps(m.deviceMaps)
	}

	c_log.Infof(ctx, "Modbus配置重载成功，共 %d 个设备", len(m.deviceMaps))
	return nil
}

// loadConfigs 从数据库加载配置
func (m *SModbusManager) loadConfigs(ctx context.Context) error {
	// 获取Modbus配置
	configJson := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingModbusConfig)
	if configJson == "" {
		c_log.Warning(ctx, "Modbus配置为空")
		return nil
	}

	// 解析JSON配置
	var config SModbusConfig
	err := json.Unmarshal([]byte(configJson), &config)
	if err != nil {
		return fmt.Errorf("解析Modbus配置失败: %v", err)
	}

	// 清空现有设备映射
	m.deviceMaps = make(map[string]*SDeviceRegisterMap)
	m.deviceStatus = make(map[string]*SModbusDeviceStatus)

	// 构建设备映射
	err = m.buildDeviceMaps(ctx, &config)
	if err != nil {
		return fmt.Errorf("构建设备映射失败: %v", err)
	}

	return nil
}

// buildDeviceMaps 构建设备映射
func (m *SModbusManager) buildDeviceMaps(ctx context.Context, config *SModbusConfig) error {
	// 遍历配置的设备ID列表
	for _, deviceId := range config.DeviceIds {
		// 获取设备实例
		device := common.GetDeviceManager().GetDeviceById(deviceId)
		if device == nil {
			c_log.Warningf(ctx, "设备 %s 不存在，跳过", deviceId)
			m.deviceStatus[deviceId] = &SModbusDeviceStatus{
				DeviceId: deviceId,
				Error:    "设备不存在",
			}
			continue
		}

		// 获取设备配置
		deviceConfig := common.GetDeviceManager().GetDeviceConfigById(deviceId)
		if deviceConfig == nil {
			c_log.Warningf(ctx, "设备 %s 配置不存在，跳过", deviceId)
			m.deviceStatus[deviceId] = &SModbusDeviceStatus{
				DeviceId: deviceId,
				Error:    "设备配置不存在",
			}
			continue
		}

		// 获取外部参数
		externalParam := deviceConfig.ExternalParam
		if externalParam == nil {
			c_log.Warningf(ctx, "设备 %s 外部参数为空，跳过", deviceId)
			m.deviceStatus[deviceId] = &SModbusDeviceStatus{
				DeviceId: deviceId,
				Error:    "外部参数为空",
			}
			continue
		}

		// 构建设备寄存器映射
		deviceMap, err := BuildDeviceRegisterMap(deviceId, device, externalParam)
		if err != nil {
			c_log.Errorf(ctx, "构建设备 %s 寄存器映射失败: %v", deviceId, err)
			m.deviceStatus[deviceId] = &SModbusDeviceStatus{
				DeviceId: deviceId,
				Error:    fmt.Sprintf("构建设备寄存器映射失败: %v", err),
			}
			continue
		}

		// 保存设备映射
		m.deviceMaps[deviceId] = deviceMap
		m.deviceStatus[deviceId] = &SModbusDeviceStatus{
			DeviceId:       deviceId,
			ModbusId:       deviceMap.ModbusId,
			StartAddr:      deviceMap.StartAddr,
			IsOnline:       deviceMap.IsOnline,
			LastUpdateTime: deviceMap.LastUpdateTime,
		}

		c_log.Infof(ctx, "设备 %s 寄存器映射构建成功: ModbusId=%d, 起始地址=%d, 寄存器数量=%d",
			deviceId, deviceMap.ModbusId, deviceMap.StartAddr, deviceMap.TotalRegisters)
	}

	// 检查地址重叠
	conflictDevices := CheckAddressOverlap(m.deviceMaps)
	if len(conflictDevices) > 0 {
		c_log.Errorf(ctx, "检测到地址重叠的设备: %v", conflictDevices)
		// 标记冲突设备为失败状态
		for _, deviceId := range conflictDevices {
			if status, exists := m.deviceStatus[deviceId]; exists {
				status.Error = "地址重叠冲突"
			}
		}
	}

	return nil
}

// getConfig 获取配置
func (m *SModbusManager) getConfig() *SModbusConfig {
	// 从数据库获取配置
	configJson := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(m.ctx, s_db_basic.SystemSettingModbusConfig)
	if configJson == "" {
		return &SModbusConfig{
			ListenPort: 502,
			DeviceIds:  []string{},
		}
	}

	var config SModbusConfig
	json.Unmarshal([]byte(configJson), &config)
	return &config
}

// GetDeviceCount 获取设备数量
func (m *SModbusManager) GetDeviceCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.deviceMaps)
}

// IsRunning 检查是否正在运行
func (m *SModbusManager) IsRunning() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.isRunning
}

// GetAllDeviceStatus 获取所有设备状态
func (m *SModbusManager) GetAllDeviceStatus() []*SModbusDeviceStatus {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var statusList []*SModbusDeviceStatus
	for _, status := range m.deviceStatus {
		statusList = append(statusList, status)
	}

	return statusList
}

// GetServerStatus 获取服务器状态
func (m *SModbusManager) GetServerStatus() (bool, int, int) {
	m.mu.RLock()
	defer m.mu.RUnlock()

	if m.server == nil {
		return false, 0, 0
	}

	return m.server.IsRunning(), m.server.GetConnectionCount(), len(m.deviceMaps)
}

// GetAllDeviceMaps 获取所有设备映射信息
func (m *SModbusManager) GetAllDeviceMaps() map[string]*SDeviceRegisterMap {
	m.mu.RLock()
	defer m.mu.RUnlock()

	// 返回设备映射的副本，避免外部修改
	deviceMapsCopy := make(map[string]*SDeviceRegisterMap)
	for deviceId, deviceMap := range m.deviceMaps {
		deviceMapsCopy[deviceId] = deviceMap
	}

	return deviceMapsCopy
}

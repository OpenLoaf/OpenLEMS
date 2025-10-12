package internal

import (
	"context"
	"encoding/json"
	"sync"

	"common/c_log"
	"s_db"
	"s_db/s_db_basic"
)

// SMqttExportManager MQTT导出管理器
type SMqttExportManager struct {
	clients   map[int]*SMqttClient // 配置索引 -> 客户端
	mu        sync.RWMutex         // 读写锁
	ctx       context.Context      // 上下文
	cancel    context.CancelFunc   // 取消函数
	isRunning bool                 // 运行状态
}

var (
	mqttExportManagerInstance *SMqttExportManager
	mqttExportManagerOnce     sync.Once
)

// GetMqttExportManager 获取MQTT导出管理器单例
func GetMqttExportManager() *SMqttExportManager {
	mqttExportManagerOnce.Do(func() {
		mqttExportManagerInstance = &SMqttExportManager{
			clients: make(map[int]*SMqttClient),
		}
	})
	return mqttExportManagerInstance
}

// Start 启动MQTT导出管理器
func (m *SMqttExportManager) Start(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.isRunning {
		c_log.Warning(ctx, "MQTT导出管理器已经在运行中")
		return nil
	}

	// 创建可取消的上下文
	m.ctx, m.cancel = context.WithCancel(ctx)

	// 加载配置并启动客户端
	err := m.loadConfigs(m.ctx)
	if err != nil {
		c_log.Errorf(m.ctx, "加载MQTT配置失败: %+v", err)
		return err
	}

	m.isRunning = true
	c_log.Infof(m.ctx, "MQTT导出管理器启动成功，共 %d 个客户端", len(m.clients))
	return nil
}

// Stop 停止MQTT导出管理器
func (m *SMqttExportManager) Stop(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isRunning {
		c_log.Warning(ctx, "MQTT导出管理器未运行")
		return nil
	}

	// 停止所有客户端
	for index, client := range m.clients {
		if err := client.Stop(); err != nil {
			c_log.Errorf(ctx, "停止MQTT客户端失败: 索引=%d, 错误=%v", index, err)
		}
	}

	// 清空客户端映射
	m.clients = make(map[int]*SMqttClient)

	// 取消上下文
	if m.cancel != nil {
		m.cancel()
	}

	m.isRunning = false
	c_log.Infof(ctx, "MQTT导出管理器已停止")
	return nil
}

// Reload 重新加载配置
func (m *SMqttExportManager) Reload(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isRunning {
		c_log.Warning(ctx, "MQTT导出管理器未运行，无法重载")
		return nil
	}

	// 停止所有现有客户端
	for index, client := range m.clients {
		if err := client.Stop(); err != nil {
			c_log.Errorf(ctx, "停止MQTT客户端失败: 索引=%d, 错误=%v", index, err)
		}
	}

	// 清空客户端映射
	m.clients = make(map[int]*SMqttClient)

	// 重新加载配置
	err := m.loadConfigs(ctx)
	if err != nil {
		c_log.Errorf(ctx, "重新加载MQTT配置失败: %+v", err)
		return err
	}

	c_log.Infof(ctx, "MQTT配置重载成功，共 %d 个客户端", len(m.clients))
	return nil
}

// loadConfigs 从数据库加载配置
func (m *SMqttExportManager) loadConfigs(ctx context.Context) error {
	// 获取MQTT配置列表
	configJson := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingMqttConfigList)
	if configJson == "" {
		c_log.Warning(ctx, "MQTT配置为空")
		return nil
	}

	// 解析JSON配置
	var configs []SMqttExportConfig
	err := json.Unmarshal([]byte(configJson), &configs)
	if err != nil {
		return err
	}

	// 获取系统序列号
	systemNumber := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingSystemNumber)

	// 为每个启用的配置创建客户端
	for index, config := range configs {
		if !config.Enabled {
			c_log.Debugf(ctx, "跳过禁用的MQTT配置: 索引=%d, 服务器=%s:%d", index, config.ServerAddress, config.ServerPort)
			continue
		}

		// 创建格式化器
		formatter := m.createFormatter(config.ServiceStandard)
		if formatter == nil {
			c_log.Errorf(ctx, "创建格式化器失败: 标准=%s", config.ServiceStandard)
			continue
		}

		// 创建MQTT客户端
		client := NewMqttClient(&config, formatter, systemNumber)

		// 启动客户端
		err := client.Start(m.ctx)
		if err != nil {
			c_log.Errorf(ctx, "启动MQTT客户端失败: 索引=%d, 服务器=%s:%d, 错误=%v", index, config.ServerAddress, config.ServerPort, err)
			continue
		}

		// 保存客户端
		m.clients[index] = client
		c_log.Infof(ctx, "MQTT客户端启动成功: 索引=%d, 服务器=%s:%d, 设备数量=%d", index, config.ServerAddress, config.ServerPort, len(config.DeviceIds))
	}

	return nil
}

// createFormatter 根据服务标准创建格式化器
func (m *SMqttExportManager) createFormatter(standard string) IDataFormatter {
	switch standard {
	case "standard":
		return &SStandardFormatter{}
	default:
		// 默认使用standard格式化器
		return &SStandardFormatter{}
	}
}

// GetClientCount 获取客户端数量
func (m *SMqttExportManager) GetClientCount() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return len(m.clients)
}

// IsRunning 检查是否正在运行
func (m *SMqttExportManager) IsRunning() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.isRunning
}

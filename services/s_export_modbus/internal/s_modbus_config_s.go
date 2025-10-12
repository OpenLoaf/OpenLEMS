package internal

// SModbusConfig Modbus配置结构体
type SModbusConfig struct {
	Enabled    bool     `json:"enabled"`    // 是否启用Modbus服务
	ListenPort int      `json:"listenPort"` // Modbus TCP服务器监听端口
	DeviceIds  []string `json:"deviceIds"`  // 设备ID列表
}

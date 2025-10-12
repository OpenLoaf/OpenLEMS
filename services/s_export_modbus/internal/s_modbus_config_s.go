package internal

// SModbusConfig Modbus配置结构体
type SModbusConfig struct {
	ListenPort int      `json:"listenPort"` // Modbus TCP服务器监听端口
	DeviceIds  []string `json:"deviceIds"`  // 设备ID列表
}

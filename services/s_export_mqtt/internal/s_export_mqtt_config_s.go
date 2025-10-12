package internal

// SMqttExportConfig MQTT导出配置结构体
type SMqttExportConfig struct {
	ServerAddress    string   `json:"serverAddress"`    // MQTT服务器地址
	ServerPort       int      `json:"serverPort"`       // MQTT服务器端口
	ServiceStandard  string   `json:"serviceStandard"`  // 服务标准（如"standard"）
	AllowControl     bool     `json:"allowControl"`     // 是否允许控制（本期不实现）
	Enabled          bool     `json:"enabled"`          // 是否启用
	DeviceIds        []string `json:"deviceIds"`        // 设备ID列表
	RewriteChannel   bool     `json:"rewriteChannel"`   // 是否重写通道
	PushChannel      string   `json:"pushChannel"`      // 推送通道（topic）
	SubscribeChannel string   `json:"subscribeChannel"` // 订阅通道（本期不实现）
	UploadPeriod     int      `json:"uploadPeriod"`     // 上传周期（秒）
}

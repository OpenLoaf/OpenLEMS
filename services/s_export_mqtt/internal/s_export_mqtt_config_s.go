package internal

// SMqttExportConfig MQTT配置结构体
type SMqttExportConfig struct {
	ServerAddress      string   `json:"serverAddress"`      // MQTT服务器地址
	ServerPort         int      `json:"serverPort"`         // MQTT服务器端口
	Username           string   `json:"username"`           // MQTT用户名（可选）
	Password           string   `json:"password"`           // MQTT密码（可选）
	UseSSL             bool     `json:"useSSL"`             // 是否使用SSL连接
	InsecureSkipVerify bool     `json:"insecureSkipVerify"` // 是否跳过SSL证书验证
	ConnectTimeout     int      `json:"connectTimeout"`     // 连接超时时间（秒）
	ReconnectInterval  int      `json:"reconnectInterval"`  // 重连间隔时间（秒）
	KeepAliveTimeout   int      `json:"keepAliveTimeout"`   // 保活超时时间（秒）
	ServiceStandard    string   `json:"serviceStandard"`    // 服务标准（如"standard"）
	AllowControl       bool     `json:"allowControl"`       // 是否允许控制（本期不实现）
	Enabled            bool     `json:"enabled"`            // 是否启用
	DeviceIds          []string `json:"deviceIds"`          // 设备ID列表
	RewriteChannel     bool     `json:"rewriteChannel"`     // 是否重写通道
	PushChannel        string   `json:"pushChannel"`        // 推送通道（topic）
	SubscribeChannel   string   `json:"subscribeChannel"`   // 订阅通道（本期不实现）
	UploadPeriod       int      `json:"uploadPeriod"`       // 上传周期（秒）
}

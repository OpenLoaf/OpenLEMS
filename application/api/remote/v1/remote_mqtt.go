package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetMqttStatusReq struct {
	g.Meta `path:"/remote/mqtt/status" method:"get" tags:"远程管理" summary:"获取MQTT服务状态列表" role:"user"`
}

type ReloadMqttReq struct {
	g.Meta `path:"/remote/mqtt/reload" method:"post" tags:"远程管理" summary:"重新加载MQTT服务配置" role:"admin"`
}

type ReloadMqttRes struct {
}

type GetMqttStatusRes struct {
	IsRunning    bool                `json:"isRunning" dc:"MQTT服务是否正在运行"`
	ClientCount  int                 `json:"clientCount" dc:"MQTT客户端数量"`
	ClientStatus []*MqttClientStatus `json:"clientStatus" dc:"MQTT客户端状态列表"`
}

// MqttClientStatus MQTT客户端状态结构体
type MqttClientStatus struct {
	Config        *MqttConfig `json:"config" dc:"MQTT配置信息"`
	IsConnected   bool        `json:"isConnected" dc:"是否已连接"`
	IsRunning     bool        `json:"isRunning" dc:"是否正在运行"`
	ClientID      string      `json:"clientId" dc:"客户端ID"`
	SystemNumber  string      `json:"systemNumber" dc:"系统序列号"`
	Topic         string      `json:"topic" dc:"当前使用的Topic"`
	DeviceCount   int         `json:"deviceCount" dc:"设备数量"`
	UploadPeriod  int         `json:"uploadPeriod" dc:"上传周期（秒）"`
	LastPublishAt *time.Time  `json:"lastPublishAt" dc:"最后发布时间"`
	PublishCount  int64       `json:"publishCount" dc:"发布消息总数"`
	ErrorCount    int64       `json:"errorCount" dc:"错误次数"`
	StartTime     *time.Time  `json:"startTime" dc:"启动时间"`
}

// MqttConfig MQTT配置结构体
type MqttConfig struct {
	ServerAddress      string   `json:"serverAddress" dc:"MQTT服务器地址"`
	ServerPort         int      `json:"serverPort" dc:"MQTT服务器端口"`
	Username           string   `json:"username" dc:"MQTT用户名"`
	Password           string   `json:"password" dc:"MQTT密码"`
	UseSSL             bool     `json:"useSSL" dc:"是否使用SSL连接"`
	InsecureSkipVerify bool     `json:"insecureSkipVerify" dc:"是否跳过SSL证书验证"`
	ConnectTimeout     int      `json:"connectTimeout" dc:"连接超时时间（秒）"`
	ReconnectInterval  int      `json:"reconnectInterval" dc:"重连间隔时间（秒）"`
	KeepAliveTimeout   int      `json:"keepAliveTimeout" dc:"保活超时时间（秒）"`
	ServiceStandard    string   `json:"serviceStandard" dc:"服务标准"`
	AllowControl       bool     `json:"allowControl" dc:"是否允许控制"`
	Enabled            bool     `json:"enabled" dc:"是否启用"`
	DeviceIds          []string `json:"deviceIds" dc:"设备ID列表"`
	RewriteChannel     bool     `json:"rewriteChannel" dc:"是否重写通道"`
	PushChannel        string   `json:"pushChannel" dc:"推送通道（topic）"`
	SubscribeChannel   string   `json:"subscribeChannel" dc:"订阅通道"`
	UploadPeriod       int      `json:"uploadPeriod" dc:"上传周期（秒）"`
}

package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type MqttSwitchReq struct {
	g.Meta  `path:"/remote/mqtt/enabled/{enabled}" method:"put" tags:"远程管理" summary:"MQTT服务开关" role:"admin"`
	Enabled bool `json:"enabled" dc:"是否启用MQTT服务" v:"required#启用状态不能为空"`
}

type MqttSwitchRes struct {
}

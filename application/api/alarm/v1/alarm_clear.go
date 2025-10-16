package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 清除告警 ====================

type ClearAlarmReq struct {
	g.Meta   `path:"/alarm/clear" method:"post" tags:"告警" summary:"清除告警" role:"admin"`
	Level    string `json:"level" v:"in:None,Warn,Alert,Error" dc:"告警类型，必填"`
	DeviceId string `json:"deviceId" dc:"设备ID，可选；为空则清除所有设备的指定类型告警"`
}

type ClearAlarmRes struct{}

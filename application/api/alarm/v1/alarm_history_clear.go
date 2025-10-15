package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 清除告警历史 ====================

type ClearAlarmHistoryReq struct {
	g.Meta   `path:"/alarm/history/clear" method:"post" tags:"告警" summary:"清除告警历史"`
	DeviceId string `json:"deviceId" dc:"设备ID，可选；为空则清除所有设备的告警历史"`
	Level    string `json:"level" dc:"告警类型，可选；为空表示全部级别"`
}

type ClearAlarmHistoryRes struct{}

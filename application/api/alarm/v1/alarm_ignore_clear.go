package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ==================== 清除忽略告警 ====================

type ClearAlarmIgnoreReq struct {
	g.Meta   `path:"/alarm/ignore/clear" method:"post" tags:"告警" summary:"清除忽略告警"`
	DeviceId string `json:"deviceId" dc:"设备ID，可选；为空则清除所有设备的忽略"`
	Point    string `json:"point" dc:"点位，可选；为空表示全部点位"`
}

type ClearAlarmIgnoreRes struct{}

package v1

import (
	"application/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetRealDeviceCacheReq struct {
	g.Meta   `path:"/device/real/cache" method:"post" tags:"设备相关" summary:"获取设备缓存值" role:"user"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备Key"`

	ShowTelemetryOnly bool `json:"showTelemetryOnly" dc:"是否只显示遥测点位，默认false显示所有点位"`
}

type GetRealDeviceCacheRes struct {
	AlarmLevel        string                       `json:"alarmLevel" dc:"告警级别"`
	DeviceServerState string                       `json:"deviceServerState,omitempty" dc:"设备服务状态"`
	LastUpdateTime    *time.Time                   `json:"lastUpdateTime" dc:"数据最后更新时间"`
	Values            []*entity.SSingleDeviceValue `json:"values" dc:"点位数值列表"`
}

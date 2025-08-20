package v1

import (
	"application/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type GetRealDeviceCacheReq struct {
	g.Meta           `path:"/device/real/cache" method:"post" tags:"设备相关" summary:"获取设备缓存值"`
	DeviceId         string   `json:"deviceId" v:"required" dc:"设备Key"`
	TelemetryKeyList []string `json:"telemetryKeyList"  dc:"遥测列表（空查询所有）"`
}

type GetRealDeviceCacheRes struct {
	AlarmLevel        string                       `json:"alarmLevel" dc:"告警级别"`
	DeviceServerState string                       `json:"deviceServerState,omitempty" dc:"设备服务状态"`
	LastUpdateTime    *time.Time                   `json:"lastUpdateTime" dc:"数据最后更新时间"`
	Values            []*entity.SSingleDeviceGroup `json:"values" dc:"数值"`
}

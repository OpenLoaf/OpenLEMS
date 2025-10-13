package v1

import (
	"application/internal/model/entity"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type GetVirtualDeviceCacheReq struct {
	g.Meta           `path:"/device/virtual/cache" method:"post" tags:"设备相关" summary:"获取虚拟设备缓存值"`
	DeviceId         string   `json:"deviceId" v:"required" dc:"设备Key"`
	TelemetryKeyList []string `json:"telemetryKeyList"  dc:"遥测列表（空查询所有）"`
}

type GetVirtualDeviceCacheRes struct {
	AlarmLevel        string                       `json:"alarmLevel" dc:"告警级别"`
	DeviceServerState string                       `json:"deviceServerState,omitempty" dc:"设备服务状态"`
	LastUpdateTime    *time.Time                   `json:"lastUpdateTime" dc:"数据最后更新时间"`
	Values            []*entity.SSingleDeviceValue `json:"values" dc:"点位数值列表"`
}

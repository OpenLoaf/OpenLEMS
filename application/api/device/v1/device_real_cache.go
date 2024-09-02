package v1

import (
	"application/internal/model/entity"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
)

type GetRealDeviceCacheReq struct {
	g.Meta           `path:"/device/real/cache" method:"post" tags:"设备相关" summary:"获取设备缓存值"`
	DeviceId         string   `json:"deviceId" v:"required" dc:"设备Key"`
	TelemetryKeyList []string `json:"telemetryKeyList"  dc:"遥测列表"`
}

type GetRealDeviceCacheRes struct {
	DeviceId       string                       `json:"deviceId" dc:"设备Key"`
	DeviceType     c_base.EDeviceType           `json:"deviceType" dc:"设备类型"`
	DeviceName     string                       `json:"deviceName" dc:"设备名称"`
	LastUpdateTime string                       `json:"lastUpdateTime" dc:"数据最后更新时间"`
	Values         []*entity.SSingleDeviceValue `json:"values" dc:"数值"`
}

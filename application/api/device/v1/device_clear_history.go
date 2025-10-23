package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ClearDeviceHistoryReq struct {
	g.Meta   `path:"/device/clear-history" method:"delete" tags:"设备相关" summary:"清除设备历史数据" role:"admin"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type ClearDeviceHistoryRes struct {
}

package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteDeviceReq struct {
	g.Meta   `path:"/device/delete" method:"delete" tags:"设备相关" summary:"删除设备"`
	DeviceId string `json:"deviceId" dc:"设备Id"`
}

type DeleteDeviceRes struct {
	DeviceId string `json:"deviceId" dc:"设备Id"`
}

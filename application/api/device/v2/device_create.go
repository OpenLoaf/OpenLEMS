package v2

import (
	"github.com/gogf/gf/v2/frame/g"
)

type CreateDeviceReq struct {
	g.Meta `path:"/device/create" method:"post" tags:"设备相关" summary:"创建设备"`
	DeviceName   string `json:"deviceName" dc:"设备名称"`
	DevicePId    string `json:"devicePId" dc:"父设备Id"`
}

type CreateDeviceRes struct {
	DeviceId string `json:"deviceId" dc:"设备Id"`
}

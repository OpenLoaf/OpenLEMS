package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteDeviceReq struct {
	g.Meta   `path:"/device/delete" method:"delete" tags:"设备相关" summary:"删除设备" role:"admin"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type DeleteDeviceRes struct {
}

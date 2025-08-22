package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceNameListReq struct {
	g.Meta `path:"/device/name-list" method:"get" tags:"设备相关" summary:"获取所有设备的名称列表"`
}

type GetDeviceNameListRes struct {
	DeviceNames map[string]string `json:"deviceNames" dc:"设备名称列表，格式为{设备ID:设备名称}"`
}

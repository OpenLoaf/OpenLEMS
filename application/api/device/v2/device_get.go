package v2

import (
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceTreeReq struct {
	g.Meta         `path:"/device/tree" method:"get" tags:"设备相关" summary:"获取设备树"`
	ActiveRootOnly bool `json:"activeRootOnly" default:"false" dc:"是否只显示激活的设备下的设备"`
	RunningOnly    bool `json:"runningOnly" default:"false" dc:"是否只显示运行中的设备"`
}

type GetDeviceTreeRes struct {
	DeviceTree []*entity.SDeviceTree `json:"deviceTree" dc:"设备树"`
}

package v1

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
	//DeviceTree []*c_base.SDeviceConfig `json:"deviceTree"`
}

type DisableDeviceReq struct {
	g.Meta   `path:"/device/disable" method:"post" tags:"设备相关" summary:"停用设备"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type DisableDeviceRes struct {
}

type EnableDeviceReq struct {
	g.Meta   `path:"/device/enable" method:"post" tags:"设备相关" summary:"启用设备"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type EnableDeviceRes struct {
}

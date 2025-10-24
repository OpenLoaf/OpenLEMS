package v1

import (
	"application/internal/model/entity"
	"common/c_enum"

	"github.com/gogf/gf/v2/frame/g"
)

type GetDeviceTreeReq struct {
	g.Meta         `path:"/device/tree" method:"get" tags:"设备相关" summary:"获取设备树" role:"user"`
	ActiveRootOnly bool                         `json:"activeRootOnly" default:"false" dc:"是否只显示激活的设备下的设备"`
	RunningStatus  *c_enum.EDeviceRunningStatus `json:"runningStatus" dc:"设备运行状态过滤，nil表示查询所有，running表示只显示运行中的设备，stopped表示只显示已停止的设备"`
	Enabled        *bool                        `json:"enabled" dc:"设备启用状态过滤，nil表示查询所有，true表示只显示启用的设备，false表示只显示未启用的设备"`
}

type GetDeviceTreeRes struct {
	DeviceTree []*entity.SDeviceTree `json:"deviceTree" dc:"设备树"`
	//DeviceTree []*c_base.SDeviceConfig `json:"deviceTree"`
}

type DisableDeviceReq struct {
	g.Meta   `path:"/device/disable" method:"post" tags:"设备相关" summary:"停用设备" role:"admin"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type DisableDeviceRes struct {
}

type EnableDeviceReq struct {
	g.Meta   `path:"/device/enable" method:"post" tags:"设备相关" summary:"启用设备" role:"admin"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type EnableDeviceRes struct {
}

type GetDeviceTreeByIdReq struct {
	g.Meta   `path:"/device/tree/{deviceId}" method:"get" tags:"设备相关" summary:"根据设备ID获取设备树" role:"user"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type GetDeviceTreeByIdRes struct {
	DeviceTree *entity.SDeviceTree `json:"deviceTree" dc:"设备树"`
}

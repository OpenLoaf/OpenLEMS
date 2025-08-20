package v1

import (
	"common/c_base"

	"github.com/gogf/gf/v2/frame/g"
)

type ControlDeviceReq struct {
	g.Meta      `path:"/control/device" method:"post" tags:"设备控制" summary:"控制设备"`
	DeviceId    string                 `json:"deviceId" v:"required" dc:"设备ID"`
	CommandName string                 `json:"commandName" v:"required" dc:"指令名称"`
	Parameters  map[string]interface{} `json:"parameters" dc:"参数值"`
}

type ControlDeviceRes struct {
}

type GetCustomServicesReq struct {
	g.Meta   `path:"/control/custom-services" method:"get" tags:"设备控制" summary:"获取设备自定义服务列表"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

type GetCustomServicesRes struct {
	Services []*c_base.SCustomService `json:"services" dc:"自定义服务列表"`
}

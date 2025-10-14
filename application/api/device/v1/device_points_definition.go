package v1

import (
	"common/c_base"

	"github.com/gogf/gf/v2/frame/g"
)

// GetDevicePointsDefinitionReq 获取设备全部点位定义请求
type GetDevicePointsDefinitionReq struct {
	g.Meta   `path:"/device/{deviceId}/points/definitions" method:"get" tags:"设备相关" summary:"获取设备全部点位定义"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

// GetDevicePointsDefinitionRes 获取设备全部点位定义响应
type GetDevicePointsDefinitionRes struct {
	IsVirtualDevice bool                       `json:"isVirtualDevice" dc:"是否是虚拟设备"`
	Fields          []*c_base.SFieldDefinition `json:"fields" dc:"点位定义列表"`
}

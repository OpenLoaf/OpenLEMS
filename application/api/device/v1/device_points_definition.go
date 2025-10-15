package v1

import (
	"common/c_base"
	"common/c_enum"

	"github.com/gogf/gf/v2/frame/g"
)

// GetDevicePointsDefinitionReq 获取设备全部点位定义请求
type GetDevicePointsDefinitionReq struct {
	g.Meta   `path:"/device/{deviceId}/points/definitions" method:"get" tags:"设备相关" summary:"获取设备全部点位定义"`
	DeviceId string `json:"deviceId" v:"required" dc:"设备ID"`
}

// GetDevicePointsDefinitionRes 获取设备全部点位定义响应
type GetDevicePointsDefinitionRes struct {
	IsVirtualDevice bool                  `json:"isVirtualDevice" dc:"是否是虚拟设备"`
	Groups          []*c_base.SPointGroup `json:"groups" dc:"分组列表，按GroupSort排序"`
	Fields          []*SDevicePointField  `json:"fields" dc:"点位定义列表（使用groupKey）"`
}

// SDevicePointField 设备点位字段（响应专用），从 SFieldDefinition 拷贝，使用 groupKey
type SDevicePointField struct {
	Key                string                            `json:"key"`
	Name               string                            `json:"name,omitempty"`
	GroupKey           string                            `json:"groupKey,omitempty"`
	ValueType          c_enum.EConfigFieldsValueType     `json:"valueType"`
	ComponentType      c_enum.EConfigFieldsComponentType `json:"componentType"`
	Step               *float32                          `json:"step,omitempty"`
	Required           bool                              `json:"required"`
	Unit               *string                           `json:"unit,omitempty"`
	Min                *int64                            `json:"min,omitempty"`
	Max                *int64                            `json:"max,omitempty"`
	Default            *string                           `json:"default,omitempty"`
	ValueExplain       []*c_base.SFieldExplain           `json:"valueExplain,omitempty"`
	ParamExplain       []*c_base.SFieldExplain           `json:"paramExplain,omitempty"`
	Regex              *string                           `json:"regex,omitempty"`
	RegexFailedMessage *string                           `json:"regexFailedMessage,omitempty"`
	Description        string                            `json:"description,omitempty"`
}

package v1

import (
	"common/c_base"

	"github.com/gogf/gf/v2/frame/g"
)

type UpdateDeviceReq struct {
	g.Meta        `path:"/device/update" method:"put" tags:"设备相关" summary:"更新设备"`
	DeviceId      string                 `json:"deviceId" v:"required" dc:"设备ID"`
	Name          string                 `json:"name" dc:"设备名称"`
	ProtocolId    string                 `json:"protocolId" dc:"协议ID"`
	Driver        string                 `json:"driver" dc:"驱动名称"`
	EnableDebug   bool                   `json:"enableDebug" dc:"启用调试模式"`
	ManualMode    *bool                  `json:"manualMode" dc:"是否手动模式"`
	Enabled       *bool                  `json:"enabled" dc:"是否启用"`
	Sort          *int                   `json:"sort" dc:"排序"`
	Params        string                 `json:"params" dc:"参数"`
	ExternalParam *c_base.SExternalParam `json:"externalParam" dc:"对外参数，存储JSON格式数据"`
}

type UpdateDeviceRes struct {
	DeviceId string `json:"deviceId" dc:"设备ID"`
}

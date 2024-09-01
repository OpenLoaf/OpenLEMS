package v1

import (
	"application/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetTelemetryDescriptionReq struct {
	g.Meta `path:"/telemetry/description" method:"get" tags:"遥测相关" summary:"获取的描述信息"`
}

type GetTelemetryDescriptionRes struct {
	Entrance  *TelemetryDescriptionObj `json:"entrance,omitempty" dc:"电网"` // 允许为空
	Ess       *TelemetryDescriptionObj `json:"ess,omitempty" dc:"储能"`
	Load      *TelemetryDescriptionObj `json:"load,omitempty" dc:"负荷"`
	Pv        *TelemetryDescriptionObj `json:"pv,omitempty" dc:"光伏"`
	Charge    *TelemetryDescriptionObj `json:"charge,omitempty" dc:"充电桩"`
	Generator *TelemetryDescriptionObj `json:"generator,omitempty" dc:"发电机"`
}

type TelemetryDescriptionObj struct {
	Name     string                    `json:"name" dc:"名称"`
	Children []*entity.DeviceTelemetry `json:"children,omitempty" dc:"子节点"`
}

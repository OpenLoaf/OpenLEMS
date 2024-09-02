package v1

import (
	"application/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetStationEssChildrenTelemetryReq struct {
	g.Meta `path:"/station/ess/children/telemetry" method:"get" tags:"储能相关" summary:"获取储能组中的每个设备的状态"`
}

type GetStationEssChildrenTelemetryRes struct {
	EssStatusList []*entity.EnergyStoreStatus `json:"essStatusList,omitempty" dc:"储能组中的每个设备的状态"`
}

package v1

import (
	"application/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

type GetStationEssTelemetryReq struct {
	g.Meta `path:"/station/ess/telemetry" method:"get" tags:"储能相关" summary:"获取总的储能数据"`
}

type GetStationEssTelemetryRes struct {
	*entity.EnergyStoreStatus
}

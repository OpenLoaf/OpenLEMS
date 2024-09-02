package v1

import (
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
)

type PostStationEssSetStatusReq struct {
	g.Meta `path:"/station/ess/set/status" method:"post" tags:"储能相关" summary:"设置储能组状态"`
	Status c_base.EEnergyStoreStatus `json:"status" dc:"储能组的状态"`
}

type PostStationEssSetStatusRes struct {
}

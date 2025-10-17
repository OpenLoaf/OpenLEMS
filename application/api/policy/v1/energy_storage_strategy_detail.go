package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetEnergyStorageStrategyDetailReq struct {
	g.Meta `path:"/strategy/energy-storage/{id}" method:"get" tags:"策略相关" summary:"获取储能策略详情"`
	Id     int `json:"id" in:"path"`
}

type GetEnergyStorageStrategyDetailRes = EnergyStorageStrategy

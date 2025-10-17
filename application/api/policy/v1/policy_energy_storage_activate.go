package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ActivateEnergyStorageStrategyReq struct {
	g.Meta `path:"/strategy/energy-storage/activate" method:"post" tags:"策略相关" summary:"激活或停用储能策略" role:"admin"`
	Id     int  `json:"id"`
	Active bool `json:"active"`
}

type ActivateEnergyStorageStrategyRes struct{}

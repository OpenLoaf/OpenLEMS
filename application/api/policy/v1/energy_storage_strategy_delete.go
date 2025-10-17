package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DeleteEnergyStorageStrategyReq struct {
	g.Meta `path:"/strategy/energy-storage/{id}" method:"delete" tags:"策略相关" summary:"删除储能策略" role:"admin"`
	Id     int `json:"id" in:"path"`
}

type DeleteEnergyStorageStrategyRes struct{}

package v1

import (
	"common/c_enum"

	"github.com/gogf/gf/v2/frame/g"
)

type ActivateEnergyStorageStrategyReq struct {
	g.Meta `path:"/strategy/energy-storage/activate" method:"post" tags:"策略相关" summary:"激活或停用储能策略" role:"admin"`
	Id     int            `json:"id" v:"required|min:1" dc:"策略ID"`
	Status c_enum.EStatus `json:"status" v:"required|in:Enable,Disable" dc:"状态：Enable-启用，Disable-停用"`
}

type ActivateEnergyStorageStrategyRes struct{}

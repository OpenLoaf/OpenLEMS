package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type DetectEnergyStorageStrategyConflictsReq struct {
	g.Meta      `path:"/strategy/energy-storage/conflicts" method:"post" tags:"策略相关" summary:"储能策略冲突检测" role:"admin"`
	StrategyIds []int                    `json:"strategyIds"`
	Candidates  []*EnergyStorageStrategy `json:"candidates"`
}

type DetectEnergyStorageStrategyConflictsRes struct {
	Conflicts []struct {
		StrategyId    int      `json:"strategyId"`
		ConflictWith  []int    `json:"conflictWith"`
		ConflictDates []string `json:"conflictDates"`
	} `json:"conflicts"`
}

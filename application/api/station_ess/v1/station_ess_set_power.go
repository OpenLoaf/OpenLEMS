package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type PostStationEssSetPowerReq struct {
	g.Meta `path:"/station/ess/set/power" method:"post" tags:"储能相关" summary:"设置储能组的功率"`
	Power  int32 `json:"power" dc:"储能组的功率"`
}

type PostStationEssSetPowerRes struct {
}

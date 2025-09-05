package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetPolicyConfigDescReq struct {
	g.Meta `path:"/policy/config/desc" method:"get" tags:"策略相关" summary:"获取策略配置描述"`
}

type GetPolicyConfigDescRes struct {
}

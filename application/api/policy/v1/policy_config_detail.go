package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetPolicyConfigReq struct {
	g.Meta   `path:"/policy/config" method:"get" tags:"策略相关" summary:"获取策略配置"`
	PolicyId string `json:"policyId" dc:"空代表查询默认的"`
}

type GetPolicyConfigRes struct {
	SettingId string `json:"settingId" dc:"设置ID"`
	Config    any    `json:"config" dc:"策略配置JSON对象"`
}

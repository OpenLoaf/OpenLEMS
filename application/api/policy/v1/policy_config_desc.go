package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetPolicyConfigDescReq struct {
	g.Meta   `path:"/policy/config/desc" method:"get" tags:"策略相关" summary:"获取策略配置描述"`
	PolicyId string `json:"policyId" dc:"空代表查询默认的"`
}

type GetPolicyConfigDescRes struct {
	SettingId string `json:"settingId" dc:"设置ID"`
	Config    any    `json:"config" dc:"策略配置JSON对象"`
}

type UpdatePolicyConfigReq struct {
	g.Meta       `path:"/policy/config/update" method:"post" tags:"策略相关" summary:"更新策略配置（双合一：存在则更新，不存在则新建）"`
	PolicyId     string `json:"policyId" v:"required|regex:^policy_.*#策略ID不能为空|策略ID必须以policy_开头" dc:"策略ID，必须以policy_开头"`
	PolicyConfig string `json:"policyConfig" v:"required#策略配置不能为空" dc:"策略配置内容"`
}

type UpdatePolicyConfigRes struct {
}

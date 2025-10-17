package p_policy_mircogrid

import (
	"common/c_base"
	"p_policy_mircogrid/internal"
)

// NewPolicyMircogrid 创建微电网策略实例
func NewPolicyMircogrid() c_base.IPolicy {
	return internal.NewPolicyMircogrid()
}

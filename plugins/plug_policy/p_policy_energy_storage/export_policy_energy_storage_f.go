package p_policy_energy_storage

import (
	"common/c_base"
	"p_policy_energy_storage/internal"
)

// NewPolicyEnergyStorage 创建储能站策略实例
func NewPolicyEnergyStorage() c_base.IPolicy {
	return internal.NewPolicyEnergyStorage()
}

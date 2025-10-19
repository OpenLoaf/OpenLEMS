package common

import (
	"common/c_base"
	"common/c_enum"
	"context"
)

// IPolicyManager 策略管理器接口
type IPolicyManager interface {
	// Start 启动策略管理器
	Start(ctx context.Context) error

	// Shutdown 关闭策略管理器
	Shutdown()

	// Status 获取运行状态
	Status() c_enum.EServerState

	// GetActivePolicyId 获取当前激活的策略ID
	GetActivePolicyId() string

	// GetActivePolicy 获取当前激活的策略实例
	GetActivePolicy() c_base.IPolicy

	// SwitchPolicy 切换策略
	SwitchPolicy(ctx context.Context, policyId string) error

	// RegisterPolicy 注册策略插件
	RegisterPolicy(policyId string, policy c_base.IPolicy) error

	// GetAllRegisteredPolicies 获取所有已注册的策略信息
	GetAllRegisteredPolicies() []*c_base.SPolicyInfo
}

var policyManager IPolicyManager

// RegisterPolicyManager 注册策略管理器
func RegisterPolicyManager(pm IPolicyManager) {
	policyManager = pm
}

// GetPolicyManager 获取策略管理器
func GetPolicyManager() IPolicyManager {
	if policyManager == nil {
		panic("policy manager is nil !")
	}
	return policyManager
}


package c_base

import "context"

// IPolicy 策略接口，所有策略插件必须实现此接口
type IPolicy interface {
	// Init 初始化策略
	Init(ctx context.Context) error

	// Shutdown 关闭策略，释放资源
	Shutdown()

	// Run 执行策略（每分钟调用一次）
	Run(ctx context.Context) error

	// GetPolicyId 获取策略ID
	GetPolicyId() string

	// GetPolicyName 获取策略名称
	GetPolicyName() string

	// GetConfig 获取策略配置
	GetConfig() interface{}

	// GetPolicyInfo 获取策略详细信息
	GetPolicyInfo() *SPolicyInfo
}

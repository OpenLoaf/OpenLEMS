package internal

import (
	"common/c_base"
	"common/c_log"
	"context"
	"encoding/json"
	"s_db"
	"s_policy"

	"github.com/pkg/errors"
)

// sPolicyEnergyStorage 储能站策略实现
type sPolicyEnergyStorage struct {
	policyId   string
	policyName string
	config     *SPolicyEnergyStorageConfig
	ctx        context.Context
}

// NewPolicyEnergyStorage 创建储能站策略实例
func NewPolicyEnergyStorage() c_base.IPolicy {
	return &sPolicyEnergyStorage{
		policyId:   "policy_ess",
		policyName: "储能站策略",
		config:     &SPolicyEnergyStorageConfig{},
	}
}

// Init 初始化策略
func (s *sPolicyEnergyStorage) Init(ctx context.Context) error {
	s.ctx = ctx
	c_log.Infof(ctx, "正在初始化储能站策略...")

	// 从setting表加载配置
	configStrPtr := s_db.GetSettingService().GetSettingValueById(ctx, "policy_ess")
	if configStrPtr == nil {
		c_log.Warning(ctx, "储能站策略配置为空，策略将不执行")
		s.config = nil
		return nil
	}

	// 解析JSON配置到 s.config
	err := json.Unmarshal([]byte(*configStrPtr), s.config)
	if err != nil {
		return errors.Wrap(err, "解析储能站策略配置失败")
	}

	c_log.Infof(ctx, "储能站策略初始化成功")
	return nil
}

// Run 执行策略（每分钟调用一次）
func (s *sPolicyEnergyStorage) Run(ctx context.Context) error {
	// 配置为空时不执行
	if s.config == nil {
		c_log.Debug(ctx, "储能站策略配置为空，跳过执行")
		return nil
	}

	c_log.Debug(ctx, "储能站策略开始执行...")

	// 获取当前激活的储能定时任务
	task, err := s_policy.GetActiveEnergyStorageTask(ctx)
	if err != nil {
		c_log.Warningf(ctx, "获取储能定时任务失败: %v", err)
		return nil // 没有任务时不报错，只记录日志
	}

	c_log.Infof(ctx, "当前激活的储能定时任务: ID=%d, Name=%s", task.Id, task.Name)

	// TODO: 实现具体业务逻辑
	// 1. 解析任务配置
	// 2. 获取储能设备数据
	// 3. 执行储能任务（充电/放电）
	// 4. 下发控制指令

	c_log.Debug(ctx, "储能站策略执行完成")
	return nil
}

// Shutdown 关闭策略
func (s *sPolicyEnergyStorage) Shutdown() {
	if s.ctx != nil {
		c_log.Info(s.ctx, "储能站策略关闭")
	}
}

// GetPolicyId 获取策略ID
func (s *sPolicyEnergyStorage) GetPolicyId() string {
	return s.policyId
}

// GetPolicyName 获取策略名称
func (s *sPolicyEnergyStorage) GetPolicyName() string {
	return s.policyName
}

// GetConfig 获取策略配置
func (s *sPolicyEnergyStorage) GetConfig() interface{} {
	return s.config
}

// GetPolicyInfo 获取策略详细信息
func (s *sPolicyEnergyStorage) GetPolicyInfo() *c_base.SPolicyInfo {
	return &c_base.SPolicyInfo{
		PolicyId:               s.policyId,
		PolicyName:             s.policyName,
		Description:            "储能站能量管理策略，根据定时任务执行充放电操作",
		ConfigFieldDefinitions: nil, // TODO: 通过反射从SPolicyEnergyStorageConfig生成
	}
}

package internal

import (
	"common/c_base"
	"common/c_log"
	"context"
	"encoding/json"
	"s_db"

	"github.com/pkg/errors"
)

// sPolicyMircogrid 微电网策略实现
type sPolicyMircogrid struct {
	policyId   string
	policyName string
	config     *SPolicyMircogridConfig
	ctx        context.Context
}

// NewPolicyMircogrid 创建微电网策略实例
func NewPolicyMircogrid() c_base.IPolicy {
	return &sPolicyMircogrid{
		policyId:   "policy_microgrid",
		policyName: "微电网策略",
		config:     &SPolicyMircogridConfig{},
	}
}

// Init 初始化策略
func (s *sPolicyMircogrid) Init(ctx context.Context) error {
	s.ctx = ctx
	c_log.Infof(ctx, "正在初始化微电网策略...")

	// 从setting表加载配置
	configStrPtr := s_db.GetSettingService().GetSettingValueById(ctx, "policy_microgrid")
	if configStrPtr == nil {
		c_log.Warning(ctx, "微电网策略配置为空，策略将不执行")
		s.config = nil
		return nil
	}

	// 解析JSON配置到 s.config
	err := json.Unmarshal([]byte(*configStrPtr), s.config)
	if err != nil {
		return errors.Wrap(err, "解析微电网策略配置失败")
	}

	c_log.Infof(ctx, "微电网策略初始化成功")
	return nil
}

// Run 执行策略（每分钟调用一次）
func (s *sPolicyMircogrid) Run(ctx context.Context) error {
	// 配置为空时不执行
	if s.config == nil {
		c_log.Debug(ctx, "微电网策略配置为空，跳过执行")
		return nil
	}

	c_log.Debug(ctx, "微电网策略开始执行...")

	// TODO: 实现具体业务逻辑
	// 1. 获取当前激活的储能定时任务
	// 2. 获取设备数据（电网、光伏、负荷、储能等）
	// 3. 执行优化算法
	// 4. 下发控制指令

	c_log.Debug(ctx, "微电网策略执行完成")
	return nil
}

// Shutdown 关闭策略
func (s *sPolicyMircogrid) Shutdown() {
	if s.ctx != nil {
		c_log.Info(s.ctx, "微电网策略关闭")
	}
}

// GetPolicyId 获取策略ID
func (s *sPolicyMircogrid) GetPolicyId() string {
	return s.policyId
}

// GetPolicyName 获取策略名称
func (s *sPolicyMircogrid) GetPolicyName() string {
	return s.policyName
}

// GetConfig 获取策略配置
func (s *sPolicyMircogrid) GetConfig() interface{} {
	return s.config
}

// GetPolicyInfo 获取策略详细信息
func (s *sPolicyMircogrid) GetPolicyInfo() *c_base.SPolicyInfo {
	fields, err := c_base.BuildConfigStructFields(s.config)
	if err != nil {
		c_log.BizErrorf(s.ctx, "加载[%s]策略失败", s.policyName)
		return nil
	}
	return &c_base.SPolicyInfo{
		PolicyId:               s.policyId,
		PolicyName:             s.policyName,
		Description:            "微电网能量管理策略，支持需量控制、削峰填谷、动态扩容等功能",
		ConfigFieldDefinitions: fields,
	}
}

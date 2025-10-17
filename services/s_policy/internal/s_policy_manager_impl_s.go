package internal

import (
	"common"
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"s_db"
	"s_db/s_db_basic"
	"sync"
	"time"

	"github.com/pkg/errors"
)

// sPolicyManagerImpl 策略管理器实现
type sPolicyManagerImpl struct {
	policies       map[string]c_base.IPolicy // 策略插件注册表
	activePolicy   c_base.IPolicy            // 当前激活的策略
	activePolicyId string                    // 当前激活的策略ID
	ticker         *time.Ticker              // 定时器（每分钟触发）
	status         c_enum.EServerState       // 运行状态
	mu             sync.RWMutex              // 读写锁
	ctx            context.Context           // 上下文
	cancelFunc     context.CancelFunc        // 取消函数
}

// NewPolicyManager 创建策略管理器实例
func NewPolicyManager(ctx context.Context) common.IPolicyManager {
	managerCtx, cancelFunc := context.WithCancel(ctx)
	return &sPolicyManagerImpl{
		policies:   make(map[string]c_base.IPolicy),
		status:     c_enum.EStateStopped,
		ctx:        managerCtx,
		cancelFunc: cancelFunc,
	}
}

// Start 启动策略管理器
func (pm *sPolicyManagerImpl) Start(ctx context.Context) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pm.status == c_enum.EStateRunning {
		return errors.New("策略管理器已经在运行中")
	}

	c_log.Info(ctx, "正在启动策略管理器...")

	// 从setting表获取激活的策略ID
	activePolicyIdPtr := s_db.GetSettingService().GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingActivePolicyId)
	var activePolicyId string
	if activePolicyIdPtr != nil {
		activePolicyId = *activePolicyIdPtr
	}
	c_log.Infof(ctx, "从配置中读取到激活的策略ID: %s", activePolicyId)

	// 如果策略ID不为空，加载并初始化对应策略
	if activePolicyId != "" {
		policy, exists := pm.policies[activePolicyId]
		if !exists {
			c_log.Warningf(ctx, "策略 %s 未注册，跳过初始化", activePolicyId)
		} else {
			err := policy.Init(ctx)
			if err != nil {
				c_log.Errorf(ctx, "初始化策略 %s 失败: %+v", activePolicyId, err)
				return errors.Wrapf(err, "初始化策略 %s 失败", activePolicyId)
			}
			pm.activePolicy = policy
			pm.activePolicyId = activePolicyId
			c_log.Infof(ctx, "策略 %s 初始化成功", activePolicyId)
		}
	} else {
		c_log.Info(ctx, "未配置激活的策略，策略管理器将以空闲模式运行")
	}

	// 启动定时器，每分钟执行一次策略的Run方法
	pm.ticker = time.NewTicker(1 * time.Minute)
	pm.status = c_enum.EStateRunning

	// 在goroutine中监听定时器和上下文取消信号
	go pm.runLoop()

	c_log.Info(ctx, "策略管理器启动成功")
	return nil
}

// runLoop 定时执行循环
func (pm *sPolicyManagerImpl) runLoop() {
	defer func() {
		if r := recover(); r != nil {
			c_log.Errorf(pm.ctx, "策略管理器运行循环发生panic: %v", r)
		}
	}()

	for {
		select {
		case <-pm.ticker.C:
			pm.mu.RLock()
			activePolicy := pm.activePolicy
			activePolicyId := pm.activePolicyId
			pm.mu.RUnlock()

			if activePolicy != nil {
				c_log.Debugf(pm.ctx, "开始执行策略: %s", activePolicyId)
				err := activePolicy.Run(pm.ctx)
				if err != nil {
					c_log.Errorf(pm.ctx, "策略 %s 执行失败: %+v", activePolicyId, err)
				} else {
					c_log.Debugf(pm.ctx, "策略 %s 执行成功", activePolicyId)
				}
			}

		case <-pm.ctx.Done():
			c_log.Info(pm.ctx, "策略管理器接收到停止信号")
			return
		}
	}
}

// Shutdown 关闭策略管理器
func (pm *sPolicyManagerImpl) Shutdown() {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if pm.status != c_enum.EStateRunning {
		return
	}

	c_log.Info(pm.ctx, "正在关闭策略管理器...")

	// 停止定时器
	if pm.ticker != nil {
		pm.ticker.Stop()
	}

	// 调用当前激活策略的Shutdown方法
	if pm.activePolicy != nil {
		c_log.Infof(pm.ctx, "正在关闭策略: %s", pm.activePolicyId)
		pm.activePolicy.Shutdown()
		pm.activePolicy = nil
		pm.activePolicyId = ""
	}

	// 取消上下文
	if pm.cancelFunc != nil {
		pm.cancelFunc()
	}

	pm.status = c_enum.EStateStopped
	c_log.Info(pm.ctx, "策略管理器已关闭")
}

// Status 获取运行状态
func (pm *sPolicyManagerImpl) Status() c_enum.EServerState {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.status
}

// GetActivePolicyId 获取当前激活的策略ID
func (pm *sPolicyManagerImpl) GetActivePolicyId() string {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.activePolicyId
}

// GetActivePolicy 获取当前激活的策略实例
func (pm *sPolicyManagerImpl) GetActivePolicy() c_base.IPolicy {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.activePolicy
}

// SwitchPolicy 切换策略
func (pm *sPolicyManagerImpl) SwitchPolicy(ctx context.Context, policyId string) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	c_log.Infof(ctx, "正在切换策略: %s -> %s", pm.activePolicyId, policyId)

	// 停止当前策略
	if pm.activePolicy != nil {
		c_log.Infof(ctx, "正在停止当前策略: %s", pm.activePolicyId)
		pm.activePolicy.Shutdown()
		pm.activePolicy = nil
		pm.activePolicyId = ""
	}

	// 从注册表中查找新策略
	newPolicy, exists := pm.policies[policyId]
	if !exists {
		return errors.Errorf("策略 %s 未注册", policyId)
	}

	// 初始化新策略
	err := newPolicy.Init(ctx)
	if err != nil {
		c_log.Errorf(ctx, "初始化策略 %s 失败: %+v", policyId, err)
		return errors.Wrapf(err, "初始化策略 %s 失败", policyId)
	}

	// 更新激活策略
	pm.activePolicy = newPolicy
	pm.activePolicyId = policyId

	// 更新setting表中的active_policy_id
	err = s_db.GetSettingService().SetSettingValueById(ctx, s_db_basic.SystemSettingActivePolicyId.Id, policyId)
	if err != nil {
		c_log.Errorf(ctx, "更新配置表中的激活策略ID失败: %+v", err)
		return errors.Wrap(err, "更新配置表中的激活策略ID失败")
	}

	c_log.Infof(ctx, "策略切换成功: %s", policyId)
	return nil
}

// RegisterPolicy 注册策略插件
func (pm *sPolicyManagerImpl) RegisterPolicy(policyId string, policy c_base.IPolicy) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if _, exists := pm.policies[policyId]; exists {
		return errors.Errorf("策略 %s 已经注册", policyId)
	}

	pm.policies[policyId] = policy
	c_log.Infof(pm.ctx, "策略 %s 注册成功", policyId)
	return nil
}

// GetAllRegisteredPolicies 获取所有已注册的策略信息
func (pm *sPolicyManagerImpl) GetAllRegisteredPolicies() []*c_base.SPolicyInfo {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	policyInfos := make([]*c_base.SPolicyInfo, 0, len(pm.policies))
	for _, policy := range pm.policies {
		policyInfo := policy.GetPolicyInfo()
		if policyInfo == nil {
			continue
		}
		policyInfos = append(policyInfos, policyInfo)
	}

	return policyInfos
}

package internal

import (
	"context"
	"sort"
	"sync"
	"time"

	"common/c_enum"
	"common/c_log"
	"s_db"
	"s_db/s_db_model"

	"github.com/pkg/errors"
)

// IEnergyManageManager 储能策略管理器接口
type IEnergyManageManager interface {
	Start(ctx context.Context, interval time.Duration) error
	Stop(ctx context.Context) error
	Restart(ctx context.Context, interval time.Duration) error
	ReloadStrategies(ctx context.Context) error
}

// SEnergyManageManager 储能策略管理器实现
type SEnergyManageManager struct {
	strategies map[string]*SEnergyManageStrategy // key: strategyId
	mu         sync.RWMutex
	ctx        context.Context
	cancel     context.CancelFunc
	ticker     *time.Ticker
	isRunning  bool
	interval   time.Duration
}

var (
	energyManageManagerInstance IEnergyManageManager
	energyManageManagerOnce     sync.Once
)

// GetEnergyManageManager 获取储能策略管理器单例
func GetEnergyManageManager() IEnergyManageManager {
	energyManageManagerOnce.Do(func() {
		energyManageManagerInstance = &SEnergyManageManager{
			strategies: make(map[string]*SEnergyManageStrategy),
		}
	})
	return energyManageManagerInstance
}

// Start 启动储能策略管理器
func (m *SEnergyManageManager) Start(ctx context.Context, interval time.Duration) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.isRunning {
		c_log.Warning(ctx, "储能策略管理器已经在运行中")
		return nil
	}

	// 创建可取消的上下文
	m.ctx, m.cancel = context.WithCancel(ctx)
	m.interval = interval

	// 加载所有策略
	err := m.loadAllStrategies(m.ctx)
	if err != nil {
		c_log.Errorf(m.ctx, "加载储能策略失败: %+v", err)
		return err
	}

	// 启动定时器
	m.ticker = time.NewTicker(interval)
	m.isRunning = true

	// 启动执行协程
	go m.executionLoop()

	c_log.Infof(m.ctx, "储能策略管理器启动成功，检查间隔: %v", interval)
	return nil
}

// Stop 停止储能策略管理器
func (m *SEnergyManageManager) Stop(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if !m.isRunning {
		c_log.Warning(ctx, "储能策略管理器未在运行")
		return nil
	}

	// 停止定时器
	if m.ticker != nil {
		m.ticker.Stop()
	}

	// 取消上下文
	if m.cancel != nil {
		m.cancel()
	}

	m.isRunning = false
	c_log.Info(ctx, "储能策略管理器已停止")
	return nil
}

// Restart 重启储能策略管理器
func (m *SEnergyManageManager) Restart(ctx context.Context, interval time.Duration) error {
	c_log.Info(ctx, "开始重启储能策略管理器")

	// 先停止服务
	err := m.Stop(ctx)
	if err != nil {
		c_log.Errorf(ctx, "停止储能策略管理器失败: %+v", err)
		return err
	}

	// 等待一小段时间确保资源释放
	time.Sleep(500 * time.Millisecond)

	// 重新启动服务
	err = m.Start(ctx, interval)
	if err != nil {
		c_log.Errorf(ctx, "启动储能策略管理器失败: %+v", err)
		return err
	}

	c_log.Info(ctx, "储能策略管理器重启成功")
	return nil
}

// ReloadStrategies 重新加载所有策略
func (m *SEnergyManageManager) ReloadStrategies(ctx context.Context) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	return m.loadAllStrategies(ctx)
}

// executionLoop 执行循环
func (m *SEnergyManageManager) executionLoop() {
	for {
		select {
		case <-m.ctx.Done():
			c_log.Info(m.ctx, "储能策略执行循环已停止")
			return
		case <-m.ticker.C:
			m.executeStrategies()
		}
	}
}

// executeStrategies 执行所有启用的储能策略
func (m *SEnergyManageManager) executeStrategies() {
	m.mu.RLock()
	strategies := make([]*SEnergyManageStrategy, 0)
	for _, strategy := range m.strategies {
		if strategy.Status == "active" {
			strategies = append(strategies, strategy)
		}
	}
	m.mu.RUnlock()

	// 过滤当前时间生效的策略
	now := time.Now()
	activeStrategies := make([]*SEnergyManageStrategy, 0)
	for _, strategy := range strategies {
		// 只处理启用状态的策略
		if c_enum.ParseEnergyStorageStrategyStatus(strategy.Status) == c_enum.EStatusEnable && strategy.IsActive(now) {
			activeStrategies = append(activeStrategies, strategy)
		}
	}

	if len(activeStrategies) == 0 {
		return
	}

	// 按设备分组，每个设备取优先级最高（priority最小）的策略
	deviceStrategyMap := make(map[string][]*SEnergyManageStrategy)
	for _, strategy := range activeStrategies {
		for _, deviceId := range strategy.EssDeviceIdsList {
			deviceStrategyMap[deviceId] = append(deviceStrategyMap[deviceId], strategy)
		}
	}

	// 每个设备执行优先级最高的策略
	for deviceId, deviceStrategies := range deviceStrategyMap {
		// 按 priority 排序，取第一个
		sort.Slice(deviceStrategies, func(i, j int) bool {
			return deviceStrategies[i].Priority < deviceStrategies[j].Priority
		})

		bestStrategy := deviceStrategies[0]
		c_log.Debugf(m.ctx, "设备 %s 执行策略: %s (优先级: %d)", deviceId, bestStrategy.Name, bestStrategy.Priority)

		// 执行该策略
		err := executeStrategy(m.ctx, bestStrategy)
		if err != nil {
			c_log.Errorf(m.ctx, "执行储能策略失败: %+v", err)
		}
	}
}

// loadAllStrategies 从数据库加载所有储能策略
func (m *SEnergyManageManager) loadAllStrategies(ctx context.Context) error {
	// 获取所有储能策略
	strategies, _, err := s_db.GetEnergyStorageStrategyService().GetEnergyStorageStrategyPage(ctx, 1, 1000, map[string]interface{}{})
	if err != nil {
		return errors.Wrap(err, "获取储能策略列表失败")
	}

	// 清空现有缓存
	m.strategies = make(map[string]*SEnergyManageStrategy)

	// 加载到内存缓存
	for _, strategy := range strategies {
		parsedStrategy, err := m.parseStrategy(strategy)
		if err != nil {
			c_log.Errorf(ctx, "解析储能策略失败，ID: %s, 错误: %+v", strategy.Id, err)
			continue
		}

		// 验证策略配置
		if err := ValidateStrategy(parsedStrategy.DateRangeParsed, parsedStrategy.TimeRangeParsed, parsedStrategy.ConfigParsed, parsedStrategy.EssDeviceIdsList); err != nil {
			c_log.Errorf(ctx, "策略配置验证失败，策略名: %s, ID: %s, 错误: %+v", strategy.Name, strategy.Id, err)
			continue // 跳过无效策略
		}

		m.strategies[strategy.Id] = parsedStrategy
	}

	c_log.Infof(ctx, "成功加载 %d 个储能策略", len(strategies))
	return nil
}

// parseStrategy 解析策略模型为储能策略
func (m *SEnergyManageManager) parseStrategy(model *s_db_model.SEnergyStorageStrategyModel) (*SEnergyManageStrategy, error) {
	return ParseStrategy(model)
}

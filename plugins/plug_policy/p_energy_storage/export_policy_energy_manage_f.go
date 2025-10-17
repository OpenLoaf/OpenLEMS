package p_energy_storage

import (
	"context"
	"p_energy_storage/internal"
)

// ExecuteStrategy 执行储能策略
func ExecuteStrategy(ctx context.Context, targetPower int32, essDeviceIds []string) error {
	return internal.ExecuteStrategy(ctx, targetPower, essDeviceIds)
}

// ValidateStrategy 验证储能策略配置
func ValidateStrategy(dateRange *SDateRange, timeRange *STimeRange, config *SStrategyConfig) error {
	return internal.ValidateStrategy(dateRange, timeRange, config)
}

package p_energy_manage

import (
	"context"
	"p_energy_manage/internal"
	"time"
)

// Init 初始化储能策略服务
func Init() {
	// 初始化储能策略管理器
	manager := GetEnergyManageManager()
	// 这里可以根据需要添加其他初始化逻辑
	_ = manager
}

// GetEnergyManageManager 获取储能策略管理器
func GetEnergyManageManager() internal.IEnergyManageManager {
	return internal.GetEnergyManageManager()
}

// Start 启动储能策略服务
func Start(ctx context.Context, interval time.Duration) error {
	manager := GetEnergyManageManager()
	return manager.Start(ctx, interval)
}

// Stop 停止储能策略服务
func Stop(ctx context.Context) error {
	manager := GetEnergyManageManager()
	return manager.Stop(ctx)
}

// ValidateStrategy 验证储能策略配置
func ValidateStrategy(dateRange *SDateRange, timeRange *STimeRange, config *SStrategyConfig, essDeviceIds []string) error {
	// 转换为内部类型进行验证
	internalDateRange := &internal.SDateRange{
		StartDate:  dateRange.StartDate,
		EndDate:    dateRange.EndDate,
		IsLongTerm: dateRange.IsLongTerm,
	}
	internalTimeRange := &internal.STimeRange{
		Type:         timeRange.Type,
		WeekdayType:  timeRange.WeekdayType,
		CustomDays:   timeRange.CustomDays,
		CustomMonths: timeRange.CustomMonths,
	}
	internalConfig := &internal.SStrategyConfig{
		SocMinRatio:              config.SocMinRatio,
		SocMaxRatio:              config.SocMaxRatio,
		EnableHealthOptimization: config.EnableHealthOptimization,
		MonthlyChargeDay:         config.MonthlyChargeDay,
		Points:                   config.Points,
	}
	return internal.ValidateStrategy(internalDateRange, internalTimeRange, internalConfig, essDeviceIds)
}

// Restart 重启储能策略服务
func Restart(ctx context.Context, interval time.Duration) error {
	manager := GetEnergyManageManager()
	return manager.Restart(ctx, interval)
}

package s_price

import (
	"common"
	"context"
	"s_price/internal"
)

// 导出内部类型
type SPriceSegment = internal.SPriceSegment
type SDateRange = internal.SDateRange
type STimeRange = internal.STimeRange
type SPriceInfo = internal.SPriceInfo
type SCurrentPrice = internal.SCurrentPrice

// NewPriceManager 创建电价管理器实例
func NewPriceManager(ctx context.Context) common.IPriceManager {
	return internal.NewPriceManager(ctx)
}

// GetCurrentActivePrice 获取当前激活的电价
func GetCurrentActivePrice(ctx context.Context) (*SPriceInfo, error) {
	return internal.GetCurrentActivePrice(ctx)
}

// GetCurrentPriceSegment 获取当前时间对应的电价时段
func GetCurrentPriceSegment(ctx context.Context) (*SPriceSegment, error) {
	return internal.GetCurrentPriceSegment(ctx)
}

// RefreshPriceCache 刷新电价缓存
func RefreshPriceCache(ctx context.Context) error {
	cache := internal.GetPriceCache()
	return cache.RefreshCache(ctx)
}

// SaveCurrentPriceToStorage 保存当前电价到Storage
func SaveCurrentPriceToStorage(ctx context.Context) error {
	return internal.SaveCurrentPriceToStorage(ctx)
}

// StartHourlyPriceSaveTimer 启动每小时电价保存定时器
func StartHourlyPriceSaveTimer(ctx context.Context) error {
	return internal.StartHourlyPriceSaveTimer(ctx)
}

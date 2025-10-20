package common

import (
	"common/c_base"
	"context"
)

var (
	priceManager IPriceManager
)

// IPriceManager 电价管理器接口
type IPriceManager interface {
	// GetCurrentPrice 获取当前激活的电价信息
	GetCurrentPrice(ctx context.Context) (*c_base.SPriceInfo, error)

	// RefreshCache 刷新电价缓存
	RefreshCache(ctx context.Context) error

	// Start 启动电价管理器
	Start(ctx context.Context) error

	// Shutdown 关闭电价管理器
	Shutdown()
}

// RegisterPriceManager 注册电价管理器
func RegisterPriceManager(pm IPriceManager) {
	priceManager = pm
}

// GetPriceManager 获取电价管理器
func GetPriceManager() IPriceManager {
	if priceManager == nil {
		panic("price manager is nil!")
	}
	return priceManager
}

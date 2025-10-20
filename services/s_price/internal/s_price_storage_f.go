package internal

import (
	"common"
	"common/c_log"
	"context"
	"time"

	"github.com/pkg/errors"
)

// SaveCurrentPriceToStorage 保存当前电价到Storage
func SaveCurrentPriceToStorage(ctx context.Context) error {
	// 获取当前激活的电价
	activePrice, err := GetCurrentActivePrice(ctx)
	if err != nil {
		return errors.Wrap(err, "获取当前激活电价失败")
	}

	if activePrice == nil {
		c_log.Debug(ctx, "没有激活的电价，跳过Storage保存")
		return nil
	}

	// 获取当前时间对应的电价时段
	currentSegment, err := GetCurrentPriceSegment(ctx)
	if err != nil {
		return errors.Wrap(err, "获取当前电价时段失败")
	}

	if currentSegment == nil {
		c_log.Debug(ctx, "当前时间没有匹配的电价时段，跳过Storage保存")
		return nil
	}

	// 构建保存数据
	currentPrice := &SCurrentPrice{
		PriceId:   activePrice.Id,
		Price:     currentSegment.Price,
		PriceType: string(currentSegment.PriceType),
		Timestamp: time.Now().Unix(),
	}

	// 使用新的电价保存方法
	err = common.GetStorageInstance().SavePriceData(
		currentPrice.PriceId,
		currentPrice.Price,
		currentPrice.PriceType,
		currentPrice.Timestamp,
	)
	if err != nil {
		return errors.Wrap(err, "保存电价数据到Storage失败")
	}

	c_log.Debugf(ctx, "电价数据已保存到Storage - 电价ID: %d, 价格: %.4f, 类型: %s",
		currentPrice.PriceId, currentPrice.Price, currentPrice.PriceType)

	return nil
}

// StartHourlyPriceSaveTimer 启动每小时电价保存定时器
func StartHourlyPriceSaveTimer(ctx context.Context) error {
	c_log.Info(ctx, "启动每小时电价保存定时器...")

	// 立即执行一次
	if err := SaveCurrentPriceToStorage(ctx); err != nil {
		c_log.Errorf(ctx, "首次保存电价到Storage失败: %+v", err)
	}

	// 启动定时器，每小时执行一次
	go func() {
		ticker := time.NewTicker(1 * time.Hour)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				c_log.Info(ctx, "电价保存定时器已停止")
				return
			case <-ticker.C:
				if err := SaveCurrentPriceToStorage(ctx); err != nil {
					c_log.Errorf(ctx, "定时保存电价到Storage失败: %+v", err)
				}
			}
		}
	}()

	c_log.Info(ctx, "每小时电价保存定时器已启动")
	return nil
}

// GetCurrentPriceFromStorage 从Storage获取当前电价
func GetCurrentPriceFromStorage(ctx context.Context) (*SCurrentPrice, error) {
	// 这里可以实现从Storage读取电价数据的逻辑
	// 由于Storage接口的限制，这里暂时返回nil
	// 实际实现需要根据具体的Storage接口来调整
	c_log.Debug(ctx, "从Storage获取当前电价（功能待实现）")
	return nil, nil
}

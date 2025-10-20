package internal

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"s_db"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/pkg/errors"
)

// 电价缓存管理器
type sPriceCache struct {
	prices   []*SPriceInfo // 所有电价信息
	mutex    sync.RWMutex  // 读写锁
	lastLoad time.Time     // 最后加载时间
}

var (
	priceCache *sPriceCache
	cacheOnce  sync.Once
)

// GetPriceCache 获取电价缓存实例
func GetPriceCache() *sPriceCache {
	cacheOnce.Do(func() {
		priceCache = &sPriceCache{
			prices: make([]*SPriceInfo, 0),
		}
	})
	return priceCache
}

// LoadAllPrices 加载所有电价到缓存
func (c *sPriceCache) LoadAllPrices(ctx context.Context) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c_log.Info(ctx, "开始加载电价数据到缓存...")

	// 从数据库获取所有电价
	priceModels, err := s_db.GetPriceService().GetAllPrices(ctx)
	if err != nil {
		return errors.Wrap(err, "从数据库加载电价数据失败")
	}

	// 转换为缓存格式
	prices := make([]*SPriceInfo, 0, len(priceModels))
	for _, model := range priceModels {
		priceInfo, err := c.convertModelToInfo(ctx, model)
		if err != nil {
			c_log.Errorf(ctx, "转换电价模型失败 - ID: %d, 错误: %+v", model.Id, err)
			continue
		}
		prices = append(prices, priceInfo)
	}

	c.prices = prices
	c.lastLoad = time.Now()

	c_log.Infof(ctx, "电价缓存加载完成，共加载 %d 条记录", len(prices))
	return nil
}

// RefreshCache 刷新缓存
func (c *sPriceCache) RefreshCache(ctx context.Context) error {
	c_log.Info(ctx, "刷新电价缓存...")
	return c.LoadAllPrices(ctx)
}

// GetAllPrices 获取所有电价
func (c *sPriceCache) GetAllPrices() []*SPriceInfo {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	// 返回副本，避免外部修改
	result := make([]*SPriceInfo, len(c.prices))
	copy(result, c.prices)
	return result
}

// GetEnabledPrices 获取所有启用的电价
func (c *sPriceCache) GetEnabledPrices() []*SPriceInfo {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	var enabled []*SPriceInfo
	for _, price := range c.prices {
		if price.Status == c_enum.EStatusEnable {
			enabled = append(enabled, price)
		}
	}
	return enabled
}

// GetPriceById 根据ID获取电价
func (c *sPriceCache) GetPriceById(id int) *SPriceInfo {
	c.mutex.RLock()
	defer c.mutex.RUnlock()

	for _, price := range c.prices {
		if price.Id == id {
			return price
		}
	}
	return nil
}

// convertModelToInfo 将数据库模型转换为缓存信息
func (c *sPriceCache) convertModelToInfo(ctx context.Context, model *s_db_model.SPriceModel) (*SPriceInfo, error) {
	info := &SPriceInfo{
		Id:          model.Id,
		Description: model.Description,
		Priority:    model.Priority,
		Status:      c_enum.ParseStatus(model.Status),
		CreatedBy:   model.CreatedBy,
	}

	// 时间字段转换
	if model.CreatedAt != nil {
		info.CreatedAt = &model.CreatedAt.Time
	}
	if model.UpdatedAt != nil {
		info.UpdatedAt = &model.UpdatedAt.Time
	}

	// 远程ID处理
	if model.RemoteId != "" {
		info.RemoteId = &model.RemoteId
	}

	// JSON字段反序列化
	if model.DateRange != "" {
		var dateRange SDateRange
		if err := gjson.DecodeTo(model.DateRange, &dateRange); err == nil {
			info.DateRange = &dateRange
		}
	}

	if model.TimeRange != "" {
		var timeRange STimeRange
		if err := gjson.DecodeTo(model.TimeRange, &timeRange); err == nil {
			info.TimeRange = &timeRange
		}
	}

	if model.PriceSegments != "" {
		var segments []*SPriceSegment
		if err := gjson.DecodeTo(model.PriceSegments, &segments); err == nil {
			info.PriceSegments = segments
		}
	}

	// 计算当前是否激活（基于日期/时间范围，不包含启用状态）
	info.IsActive = c.isActiveAtTime(time.Now(), info.DateRange, info.TimeRange)

	return info, nil
}

// isActiveAtTime 判断在给定时间点是否命中日期/时间范围
func (c *sPriceCache) isActiveAtTime(now time.Time, dr *SDateRange, tr *STimeRange) bool {
	// 简化的时间范围判断逻辑
	if dr == nil && tr == nil {
		return true
	}

	// 日期范围判断
	if dr != nil {
		if !dr.IsLongTerm && dr.StartDate != "" {
			if startDate, err := time.Parse("2006-01-02", dr.StartDate); err == nil && now.Before(startDate) {
				return false
			}
		}
		if !dr.IsLongTerm && dr.EndDate != "" {
			if endDate, err := time.Parse("2006-01-02", dr.EndDate); err == nil && now.After(endDate) {
				return false
			}
		}
	}

	// 时间范围判断（简化版）
	if tr != nil {
		switch tr.Type {
		case "weekday":
			// 工作日判断
			if tr.WeekdayType == "workday" && (now.Weekday() == time.Saturday || now.Weekday() == time.Sunday) {
				return false
			}
			if tr.WeekdayType == "weekend" && now.Weekday() != time.Saturday && now.Weekday() != time.Sunday {
				return false
			}
		}
	}

	return true
}

// GetLastLoadTime 获取最后加载时间
func (c *sPriceCache) GetLastLoadTime() time.Time {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	return c.lastLoad
}

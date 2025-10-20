package internal

import (
	"common"
	"common/c_base"
	"common/c_log"
	"context"

	"github.com/pkg/errors"
)

// sPriceManagerImpl 电价管理器实现
type sPriceManagerImpl struct {
	ctx    context.Context
	cancel context.CancelFunc
}

// NewPriceManager 创建电价管理器实例
func NewPriceManager(ctx context.Context) common.IPriceManager {
	managerCtx, cancel := context.WithCancel(ctx)
	return &sPriceManagerImpl{
		ctx:    managerCtx,
		cancel: cancel,
	}
}

// GetCurrentPrice 获取当前激活的电价信息
func (m *sPriceManagerImpl) GetCurrentPrice(ctx context.Context) (*c_base.SPriceInfo, error) {
	// 获取当前激活的电价
	activePrice, err := GetCurrentActivePrice(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "获取当前激活电价失败")
	}

	if activePrice == nil {
		return nil, nil
	}

	// 转换为通用格式
	priceInfo := &c_base.SPriceInfo{
		Id:            activePrice.Id,
		Description:   activePrice.Description,
		Priority:      activePrice.Priority,
		Status:        activePrice.Status,
		IsActive:      activePrice.IsActive,
		DateRange:     convertDateRange(activePrice.DateRange),
		TimeRange:     convertTimeRange(activePrice.TimeRange),
		PriceSegments: convertPriceSegments(activePrice.PriceSegments),
		RemoteId:      activePrice.RemoteId,
		CreatedAt:     activePrice.CreatedAt,
		UpdatedAt:     activePrice.UpdatedAt,
		CreatedBy:     activePrice.CreatedBy,
	}

	return priceInfo, nil
}

// RefreshCache 刷新电价缓存
func (m *sPriceManagerImpl) RefreshCache(ctx context.Context) error {
	cache := GetPriceCache()
	return cache.RefreshCache(ctx)
}

// Start 启动电价管理器
func (m *sPriceManagerImpl) Start(ctx context.Context) error {
	c_log.Info(ctx, "启动电价管理器...")

	// 加载电价数据到缓存
	if err := m.RefreshCache(ctx); err != nil {
		return errors.Wrap(err, "初始化电价缓存失败")
	}

	// 启动每小时保存定时器
	if err := StartHourlyPriceSaveTimer(ctx); err != nil {
		return errors.Wrap(err, "启动电价保存定时器失败")
	}

	c_log.Info(ctx, "电价管理器启动成功")
	return nil
}

// Shutdown 关闭电价管理器
func (m *sPriceManagerImpl) Shutdown() {
	if m.cancel != nil {
		m.cancel()
	}
	c_log.Info(m.ctx, "电价管理器已关闭")
}

// convertDateRange 转换日期范围
func convertDateRange(dr *SDateRange) *c_base.SDateRange {
	if dr == nil {
		return nil
	}
	return &c_base.SDateRange{
		StartDate:  dr.StartDate,
		EndDate:    dr.EndDate,
		IsLongTerm: dr.IsLongTerm,
	}
}

// convertTimeRange 转换时间范围
func convertTimeRange(tr *STimeRange) *c_base.STimeRange {
	if tr == nil {
		return nil
	}
	return &c_base.STimeRange{
		Type:         tr.Type,
		WeekdayType:  tr.WeekdayType,
		CustomDays:   tr.CustomDays,
		CustomMonths: tr.CustomMonths,
	}
}

// convertPriceSegments 转换电价时段
func convertPriceSegments(segments []*SPriceSegment) []*c_base.SPriceSegment {
	if segments == nil {
		return nil
	}

	result := make([]*c_base.SPriceSegment, len(segments))
	for i, segment := range segments {
		result[i] = &c_base.SPriceSegment{
			StartTime: segment.StartTime,
			EndTime:   segment.EndTime,
			PriceType: segment.PriceType,
			Price:     segment.Price,
		}
	}
	return result
}

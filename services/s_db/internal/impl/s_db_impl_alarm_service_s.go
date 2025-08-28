package impl

import (
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/gogf/gf/v2/frame/g"
)

type sAlarmServiceImpl struct {
	// 缓存相关字段
	ignoreCache     map[string]bool // 缓存忽略状态，key: deviceId:sourceDeviceId:point
	cacheMutex      sync.RWMutex
	cacheExpireTime time.Duration
	lastCleanTime   time.Time
}

var (
	alarmServiceInstance s_db_basic.IAlarmService
	alarmServiceOnce     sync.Once
)

func GetAlarmService() s_db_basic.IAlarmService {
	alarmServiceOnce.Do(func() {
		alarmServiceInstance = &sAlarmServiceImpl{
			ignoreCache:     make(map[string]bool),
			cacheExpireTime: 5 * time.Minute, // 缓存5分钟
			lastCleanTime:   time.Now(),
		}
	})
	return alarmServiceInstance
}

// 生成缓存key
func (s *sAlarmServiceImpl) getCacheKey(deviceId, sourceDeviceId, point string) string {
	return deviceId + ":" + sourceDeviceId + ":" + point
}

// 清理过期缓存
func (s *sAlarmServiceImpl) cleanExpiredCache() {
	now := time.Now()
	if now.Sub(s.lastCleanTime) < s.cacheExpireTime {
		return
	}

	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	// 清理所有缓存（简单策略：定期清理所有缓存）
	s.ignoreCache = make(map[string]bool)
	s.lastCleanTime = now
	g.Log().Debugf(context.Background(), "告警忽略缓存已清理")
}

// 从缓存获取忽略状态
func (s *sAlarmServiceImpl) getFromCache(deviceId, sourceDeviceId, point string) (bool, bool) {
	s.cacheMutex.RLock()
	defer s.cacheMutex.RUnlock()

	cacheKey := s.getCacheKey(deviceId, sourceDeviceId, point)
	if isIgnored, exists := s.ignoreCache[cacheKey]; exists {
		return isIgnored, true
	}
	return false, false
}

// 设置缓存
func (s *sAlarmServiceImpl) setCache(deviceId, sourceDeviceId, point string, isIgnored bool) {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	cacheKey := s.getCacheKey(deviceId, sourceDeviceId, point)
	s.ignoreCache[cacheKey] = isIgnored
}

// 清除指定设备的缓存（仅按目标设备清除其相关缓存前缀）
func (s *sAlarmServiceImpl) clearDeviceCache(deviceId string) {
	s.cacheMutex.Lock()
	defer s.cacheMutex.Unlock()

	prefix := deviceId + ":"
	for key := range s.ignoreCache {
		if len(key) > len(prefix) && key[:len(prefix)] == prefix {
			delete(s.ignoreCache, key)
		}
	}
}

// ==================== 告警历史相关方法 ====================

// CreateAlarmHistory 创建告警历史记录
func (s *sAlarmServiceImpl) CreateAlarmHistory(ctx context.Context, deviceId, sourceDeviceId, point, level, title, detail string, triggerAt *time.Time) error {
	if triggerAt == nil {
		return errors.Errorf("deviceId:%s point:%s level:%s title:%s detail:%s 创建告警记录失败，createAt为空", deviceId, point, level, title, detail)
	}
	now := time.Now()
	alarmHistory := &s_db_model.SAlarmHistoryModel{
		DeviceId:       deviceId,
		SourceDeviceId: sourceDeviceId,
		Point:          point,
		Level:          level,
		Title:          title,
		Detail:         detail,
		TriggerAt:      triggerAt,
		ClearAt:        &now,
	}

	err := alarmHistory.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建告警历史记录失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, point, err)
		return err
	}

	return nil
}

// GetAlarmHistoryByDeviceId 根据设备ID获取告警历史记录
func (s *sAlarmServiceImpl) GetAlarmHistoryByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SAlarmHistoryModel, error) {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	records, err := alarmHistory.GetByDeviceId(ctx, deviceId)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警历史记录失败 - 设备ID: %s, 错误: %+v", deviceId, err)
		return nil, err
	}

	return records, nil
}

// GetAlarmHistoryByDeviceIdAndPoint 根据设备ID和点位获取告警历史记录
func (s *sAlarmServiceImpl) GetAlarmHistoryByDeviceIdAndPoint(ctx context.Context, deviceId, point string) ([]*s_db_model.SAlarmHistoryModel, error) {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	records, err := alarmHistory.GetByDeviceIdAndPoint(ctx, deviceId, point)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警历史记录失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, point, err)
		return nil, err
	}

	return records, nil
}

// DeleteAlarmHistoryByDeviceId 根据设备ID删除告警历史记录
func (s *sAlarmServiceImpl) DeleteAlarmHistoryByDeviceId(ctx context.Context, deviceId string) error {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	err := alarmHistory.DeleteByDeviceId(ctx, deviceId)
	if err != nil {
		g.Log().Errorf(ctx, "删除告警历史记录失败 - 设备ID: %s, 错误: %+v", deviceId, err)
		return err
	}

	g.Log().Infof(ctx, "成功删除告警历史记录 - 设备ID: %s", deviceId)
	return nil
}

// GetAllAlarmHistory 获取所有告警历史记录
func (s *sAlarmServiceImpl) GetAllAlarmHistory(ctx context.Context) ([]*s_db_model.SAlarmHistoryModel, error) {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	records, err := alarmHistory.GetAll(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取所有告警历史记录失败 - 错误: %+v", err)
		return nil, err
	}

	return records, nil
}

// GetAlarmHistoryPage 分页获取告警历史记录
func (s *sAlarmServiceImpl) GetAlarmHistoryPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SAlarmHistoryModel, int, error) {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	records, total, err := alarmHistory.GetPage(ctx, page, pageSize, filters)
	if err != nil {
		g.Log().Errorf(ctx, "分页获取告警历史记录失败 - 页码: %d, 页大小: %d, 过滤条件: %+v, 错误: %+v", page, pageSize, filters, err)
		return nil, 0, err
	}

	return records, total, nil
}

// ClearAllAlarmHistory 清除所有告警历史记录
func (s *sAlarmServiceImpl) ClearAllAlarmHistory(ctx context.Context) error {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	err := alarmHistory.ClearAll(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "清除所有告警历史记录失败 - 错误: %+v", err)
		return err
	}

	g.Log().Infof(ctx, "成功清除所有告警历史记录并执行VACUUM")
	return nil
}

// GetAlarmHistoryCount 获取告警历史表记录总数
func (s *sAlarmServiceImpl) GetAlarmHistoryCount(ctx context.Context) int {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	count, err := alarmHistory.GetCount(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警历史表记录总数失败 - 错误: %+v", err)
		return 0
	}
	return count
}

// ==================== 告警忽略相关方法 ====================

// CreateAlarmIgnore 创建告警忽略记录
func (s *sAlarmServiceImpl) CreateAlarmIgnore(ctx context.Context, deviceId, sourceDeviceId, point string) error {
	now := time.Now()
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{
		DeviceId:       deviceId,
		SourceDeviceId: sourceDeviceId,
		Point:          point,
		CreatedAt:      &now,
	}

	err := alarmIgnore.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建告警忽略记录失败 - 设备ID: %s, 源设备ID: %s, 点位: %s, 错误: %+v", deviceId, sourceDeviceId, point, err)
		return err
	}

	// 清除相关缓存，确保下次查询能获取到最新状态
	s.clearDeviceCache(deviceId)

	return nil
}

// GetAlarmIgnoreByDeviceId 根据设备ID获取告警忽略记录
func (s *sAlarmServiceImpl) GetAlarmIgnoreByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SAlarmIgnoreModel, error) {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	records, err := alarmIgnore.GetByDeviceId(ctx, deviceId)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警忽略记录失败 - 设备ID: %s, 错误: %+v", deviceId, err)
		return nil, err
	}

	return records, nil
}

// IsAlarmIgnored 检查告警是否被忽略（设备+源设备+点位）
func (s *sAlarmServiceImpl) IsAlarmIgnored(ctx context.Context, deviceId, sourceDeviceId, point string) (bool, error) {
	// 定期清理过期缓存
	s.cleanExpiredCache()

	// 优先从缓存获取
	if isIgnored, found := s.getFromCache(deviceId, sourceDeviceId, point); found {
		g.Log().Debugf(ctx, "从缓存获取告警忽略状态 - 设备ID: %s, 源设备ID: %s, 点位: %s, 状态: %t", deviceId, sourceDeviceId, point, isIgnored)
		return isIgnored, nil
	}

	// 缓存未命中，从数据库获取
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	isIgnored, err := alarmIgnore.IsIgnored(ctx, deviceId, sourceDeviceId, point)
	if err != nil {
		g.Log().Errorf(ctx, "检查告警忽略状态失败 - 设备ID: %s, 源设备ID: %s, 点位: %s, 错误: %+v", deviceId, sourceDeviceId, point, err)
		return false, err
	}

	// 将结果设置到缓存
	s.setCache(deviceId, sourceDeviceId, point, isIgnored)

	if isIgnored {
		g.Log().Infof(ctx, "告警已被忽略 - 设备ID: %s, 源设备ID: %s, 点位: %s", deviceId, sourceDeviceId, point)
	} else {
		g.Log().Debugf(ctx, "告警未被忽略 - 设备ID: %s, 源设备ID: %s, 点位: %s", deviceId, sourceDeviceId, point)
	}

	return isIgnored, nil
}

// DeleteAlarmIgnoreByDeviceId 根据设备ID删除告警忽略记录
func (s *sAlarmServiceImpl) DeleteAlarmIgnoreByDeviceId(ctx context.Context, deviceId string) error {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	err := alarmIgnore.DeleteByDeviceId(ctx, deviceId)
	if err != nil {
		g.Log().Errorf(ctx, "删除告警忽略记录失败 - 设备ID: %s, 错误: %+v", deviceId, err)
		return err
	}

	g.Log().Infof(ctx, "成功删除告警忽略记录 - 设备ID: %s", deviceId)
	s.clearDeviceCache(deviceId) // 清除指定设备的缓存
	return nil
}

// DeleteAlarmIgnoreByDeviceIdAndPoint 根据设备ID和点位删除告警忽略记录
func (s *sAlarmServiceImpl) DeleteAlarmIgnoreByDeviceIdAndPoint(ctx context.Context, deviceId, point string) error {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	err := alarmIgnore.DeleteByDeviceIdAndPoint(ctx, deviceId, point)
	if err != nil {
		g.Log().Errorf(ctx, "删除告警忽略记录失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, point, err)
		return err
	}

	g.Log().Infof(ctx, "成功删除告警忽略记录 - 设备ID: %s, 点位: %s", deviceId, point)
	s.clearDeviceCache(deviceId) // 清除指定设备的缓存
	return nil
}

// GetAllAlarmIgnore 获取所有告警忽略记录
func (s *sAlarmServiceImpl) GetAllAlarmIgnore(ctx context.Context) ([]*s_db_model.SAlarmIgnoreModel, error) {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	records, err := alarmIgnore.GetAll(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取所有告警忽略记录失败 - 错误: %+v", err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取所有告警忽略记录，共 %d 条记录", len(records))
	return records, nil
}

// GetAlarmIgnorePage 分页获取告警忽略记录
func (s *sAlarmServiceImpl) GetAlarmIgnorePage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SAlarmIgnoreModel, int, error) {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	records, total, err := alarmIgnore.GetPage(ctx, page, pageSize, filters)
	if err != nil {
		g.Log().Errorf(ctx, "分页获取告警忽略记录失败 - 页码: %d, 页大小: %d, 过滤条件: %+v, 错误: %+v", page, pageSize, filters, err)
		return nil, 0, err
	}

	return records, total, nil
}

// GetAlarmIgnoreCount 获取告警忽略表记录总数
func (s *sAlarmServiceImpl) GetAlarmIgnoreCount(ctx context.Context) int {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	count, err := alarmIgnore.GetCount(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警忽略表记录总数失败 - 错误: %+v", err)
		return 0
	}
	return count
}

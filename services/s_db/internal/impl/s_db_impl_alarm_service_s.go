package impl

import (
	"common/c_base"
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/pkg/errors"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

type sAlarmServiceImpl struct {
	// 缓存相关字段
	ignoreCache   *gcache.Cache // 缓存忽略状态，key: deviceId:sourceDeviceId:point
	countCache    *gcache.Cache // 缓存计数信息，包括历史总数和忽略总数
	cacheDuration time.Duration // 缓存过期时间
}

var (
	alarmServiceInstance s_db_basic.IAlarmService
	alarmServiceOnce     sync.Once
)

func GetAlarmService() s_db_basic.IAlarmService {
	alarmServiceOnce.Do(func() {
		alarmServiceInstance = &sAlarmServiceImpl{
			ignoreCache:   gcache.New(),    // 创建忽略状态缓存
			countCache:    gcache.New(),    // 创建计数缓存
			cacheDuration: 5 * time.Minute, // 缓存5分钟
		}
	})
	return alarmServiceInstance
}

// 生成缓存key
func (s *sAlarmServiceImpl) getCacheKey(deviceId, sourceDeviceId, point string) string {
	return deviceId + ":" + sourceDeviceId + ":" + point
}

// 清除指定设备的缓存
func (s *sAlarmServiceImpl) clearDeviceCache(ctx context.Context, deviceId string) {
	prefix := deviceId + ":"
	// 使用gcache的Keys方法找到所有匹配的键并删除
	keys, err := s.ignoreCache.Keys(ctx)
	if err != nil {
		g.Log().Warningf(ctx, "获取缓存键失败 - 错误: %+v", err)
		return
	}
	for _, key := range keys {
		if keyStr, ok := key.(string); ok {
			if len(keyStr) > len(prefix) && keyStr[:len(prefix)] == prefix {
				s.ignoreCache.Remove(ctx, key)
			}
		}
	}
	g.Log().Debugf(ctx, "已清除设备[%s]的告警忽略缓存", deviceId)
}

// 清除计数缓存
func (s *sAlarmServiceImpl) clearCountCache(ctx context.Context) {
	s.countCache.Clear(ctx)
	g.Log().Debugf(ctx, "已清除告警计数缓存")
}

// ==================== 告警历史相关方法 ====================

// CreateAlarmHistory 创建告警历史记录
func (s *sAlarmServiceImpl) CreateAlarmHistory(ctx context.Context, deviceId, sourceDeviceId string, meta *c_base.Meta, detail string, triggerAt *time.Time) error {
	if triggerAt == nil {
		return errors.Errorf("deviceId:%s point:%s level:%s title:%s detail:%s 创建告警记录失败，createAt为空", deviceId, meta.Name, meta.Level.String(), meta.Cn, detail)
	}
	now := time.Now()
	alarmHistory := &s_db_model.SAlarmHistoryModel{
		DeviceId:       deviceId,
		SourceDeviceId: sourceDeviceId,
		Point:          meta.Name,
		Level:          meta.Level.String(),
		PointName:      meta.Cn,
		Detail:         detail,
		TriggerAt:      triggerAt,
		ClearAt:        &now,
	}

	err := alarmHistory.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建告警历史记录失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, meta.Name, err)
		return err
	}

	// 清除计数缓存，确保下次查询能获取到最新状态
	s.clearCountCache(ctx)

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

	// 清除计数缓存，确保下次查询能获取到最新状态
	s.clearCountCache(ctx)

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

	// 清除计数缓存，确保下次查询能获取到最新状态
	s.clearCountCache(ctx)

	g.Log().Infof(ctx, "成功清除所有告警历史记录并执行VACUUM")
	return nil
}

// GetAlarmHistoryCount 获取告警历史表记录总数
func (s *sAlarmServiceImpl) GetAlarmHistoryCount(ctx context.Context) int {
	cacheKey := "history_count"

	// 优先从缓存获取
	cachedCount, err := s.countCache.Get(ctx, cacheKey)
	if err == nil && cachedCount != nil {
		if count := cachedCount.Int(); count != 0 || cachedCount.String() != "" {
			g.Log().Debugf(ctx, "从缓存获取告警历史总数 - 总数: %d", count)
			return count
		}
	}

	// 缓存未命中，从数据库获取
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	count, err := alarmHistory.GetCount(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警历史表记录总数失败 - 错误: %+v", err)
		return 0
	}

	// 将结果设置到缓存
	err = s.countCache.Set(ctx, cacheKey, count, s.cacheDuration)
	if err != nil {
		g.Log().Warningf(ctx, "设置告警历史总数缓存失败 - 错误: %+v", err)
	}

	g.Log().Debugf(ctx, "从数据库获取告警历史总数并缓存 - 总数: %d", count)
	return count
}

// ==================== 告警忽略相关方法 ====================

// CreateAlarmIgnore 创建告警忽略记录
func (s *sAlarmServiceImpl) CreateAlarmIgnore(ctx context.Context, deviceId, sourceDeviceId, point, pointName string) error {
	now := time.Now()
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{
		DeviceId:       deviceId,
		SourceDeviceId: sourceDeviceId,
		Point:          point,
		PointName:      pointName,
		CreatedAt:      &now,
	}

	err := alarmIgnore.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建告警忽略记录失败 - 设备ID: %s, 源设备ID: %s, 点位: %s, 错误: %+v", deviceId, sourceDeviceId, point, err)
		return err
	}

	// 清除相关缓存，确保下次查询能获取到最新状态
	s.clearDeviceCache(ctx, deviceId)
	s.clearCountCache(ctx)

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
	cacheKey := s.getCacheKey(deviceId, sourceDeviceId, point)

	// 优先从缓存获取
	cachedResult, err := s.ignoreCache.Get(ctx, cacheKey)
	if err == nil && cachedResult != nil {
		isIgnored := cachedResult.Bool()
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
	err = s.ignoreCache.Set(ctx, cacheKey, isIgnored, s.cacheDuration)
	if err != nil {
		g.Log().Warningf(ctx, "设置告警忽略缓存失败 - 设备ID: %s, 源设备ID: %s, 点位: %s, 错误: %+v", deviceId, sourceDeviceId, point, err)
	}

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
	s.clearDeviceCache(ctx, deviceId) // 清除指定设备的缓存
	s.clearCountCache(ctx)            // 清除计数缓存
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
	s.clearDeviceCache(ctx, deviceId) // 清除指定设备的缓存
	s.clearCountCache(ctx)            // 清除计数缓存
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
	cacheKey := "ignore_count"

	// 优先从缓存获取
	cachedCount, err := s.countCache.Get(ctx, cacheKey)
	if err == nil && cachedCount != nil {
		if count := cachedCount.Int(); count != 0 || cachedCount.String() != "" {
			g.Log().Debugf(ctx, "从缓存获取告警忽略总数 - 总数: %d", count)
			return count
		}
	}

	// 缓存未命中，从数据库获取
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	count, err := alarmIgnore.GetCount(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警忽略表记录总数失败 - 错误: %+v", err)
		return 0
	}

	// 将结果设置到缓存
	err = s.countCache.Set(ctx, cacheKey, count, s.cacheDuration)
	if err != nil {
		g.Log().Warningf(ctx, "设置告警忽略总数缓存失败 - 错误: %+v", err)
	}

	g.Log().Debugf(ctx, "从数据库获取告警忽略总数并缓存 - 总数: %d", count)
	return count
}

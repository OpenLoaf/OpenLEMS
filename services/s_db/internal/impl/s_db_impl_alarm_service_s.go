package impl

import (
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

type sAlarmServiceImpl struct {
}

var (
	alarmServiceInstance s_db_basic.IAlarmService
	alarmServiceOnce     sync.Once
)

func GetAlarmService() s_db_basic.IAlarmService {
	alarmServiceOnce.Do(func() {
		alarmServiceInstance = &sAlarmServiceImpl{}
	})
	return alarmServiceInstance
}

// ==================== 告警历史相关方法 ====================

// CreateAlarmHistory 创建告警历史记录
func (s *sAlarmServiceImpl) CreateAlarmHistory(ctx context.Context, deviceId, point, level, title, detail string) error {
	alarmHistory := &s_db_model.SAlarmHistoryModel{
		DeviceId: deviceId,
		Point:    point,
		Level:    level,
		Title:    title,
		Detail:   detail,
	}

	err := alarmHistory.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建告警历史记录失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, point, err)
		return err
	}

	g.Log().Infof(ctx, "成功创建告警历史记录 - 设备ID: %s, 点位: %s, 等级: %s", deviceId, point, level)
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

	g.Log().Infof(ctx, "成功获取告警历史记录 - 设备ID: %s, 共 %d 条记录", deviceId, len(records))
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

	g.Log().Infof(ctx, "成功获取告警历史记录 - 设备ID: %s, 点位: %s, 共 %d 条记录", deviceId, point, len(records))
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

	g.Log().Infof(ctx, "成功获取所有告警历史记录，共 %d 条记录", len(records))
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

	g.Log().Infof(ctx, "成功分页获取告警历史记录 - 页码: %d, 页大小: %d, 总数: %d, 当前页记录数: %d, 过滤条件: %+v", page, pageSize, total, len(records), filters)
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
func (s *sAlarmServiceImpl) GetAlarmHistoryCount(ctx context.Context) (int, error) {
	alarmHistory := &s_db_model.SAlarmHistoryModel{}
	count, err := alarmHistory.GetCount(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取告警历史表记录总数失败 - 错误: %+v", err)
		return 0, err
	}

	g.Log().Infof(ctx, "成功获取告警历史表记录总数: %d", count)
	return count, nil
}

// ==================== 告警忽略相关方法 ====================

// CreateAlarmIgnore 创建告警忽略记录
func (s *sAlarmServiceImpl) CreateAlarmIgnore(ctx context.Context, deviceId, point string) error {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{
		DeviceId: deviceId,
		Point:    point,
	}

	err := alarmIgnore.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建告警忽略记录失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, point, err)
		return err
	}

	g.Log().Infof(ctx, "成功创建告警忽略记录 - 设备ID: %s, 点位: %s", deviceId, point)
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

	g.Log().Infof(ctx, "成功获取告警忽略记录 - 设备ID: %s, 共 %d 条记录", deviceId, len(records))
	return records, nil
}

// IsAlarmIgnored 检查告警是否被忽略
func (s *sAlarmServiceImpl) IsAlarmIgnored(ctx context.Context, deviceId, point string) (bool, error) {
	alarmIgnore := &s_db_model.SAlarmIgnoreModel{}
	isIgnored, err := alarmIgnore.IsIgnored(ctx, deviceId, point)
	if err != nil {
		g.Log().Errorf(ctx, "检查告警忽略状态失败 - 设备ID: %s, 点位: %s, 错误: %+v", deviceId, point, err)
		return false, err
	}

	if isIgnored {
		g.Log().Infof(ctx, "告警已被忽略 - 设备ID: %s, 点位: %s", deviceId, point)
	} else {
		g.Log().Infof(ctx, "告警未被忽略 - 设备ID: %s, 点位: %s", deviceId, point)
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

	g.Log().Infof(ctx, "成功分页获取告警忽略记录 - 页码: %d, 页大小: %d, 总数: %d, 当前页记录数: %d, 过滤条件: %+v", page, pageSize, total, len(records), filters)
	return records, total, nil
}

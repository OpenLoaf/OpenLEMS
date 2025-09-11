package impl

import (
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type sLogServiceImpl struct {
}

var (
	logServiceInstance s_db_basic.ILogService
	logServiceOnce     sync.Once
)

func GetLogService() s_db_basic.ILogService {
	logServiceOnce.Do(func() {
		logServiceInstance = &sLogServiceImpl{}
	})
	return logServiceInstance
}

// CreateLog 创建日志记录
func (s *sLogServiceImpl) CreateLog(ctx context.Context, logType, deviceId, level, content string) error {
	now := time.Now()
	log := &s_db_model.SLogModel{
		Type:      logType,
		DeviceId:  deviceId,
		Level:     level,
		Content:   content,
		CreatedAt: &now,
	}

	err := log.Create(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "创建日志记录失败 - 类型: %s, 设备ID: %s, 等级: %s, 错误: %+v", logType, deviceId, level, err)
		return err
	}

	g.Log().Debugf(ctx, "成功创建日志记录 - 类型: %s, 设备ID: %s, 等级: %s", logType, deviceId, level)
	return nil
}

// GetLogByDeviceId 根据设备ID获取日志记录
func (s *sLogServiceImpl) GetLogByDeviceId(ctx context.Context, deviceId string) ([]*s_db_model.SLogModel, error) {
	log := &s_db_model.SLogModel{}
	records, err := log.GetByDeviceId(ctx, deviceId)
	if err != nil {
		g.Log().Errorf(ctx, "获取日志记录失败 - 设备ID: %s, 错误: %+v", deviceId, err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取日志记录 - 设备ID: %s, 共 %d 条记录", deviceId, len(records))
	return records, nil
}

// GetLogByType 根据日志类型获取日志记录
func (s *sLogServiceImpl) GetLogByType(ctx context.Context, logType string) ([]*s_db_model.SLogModel, error) {
	log := &s_db_model.SLogModel{}
	records, err := log.GetByType(ctx, logType)
	if err != nil {
		g.Log().Errorf(ctx, "获取日志记录失败 - 类型: %s, 错误: %+v", logType, err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取日志记录 - 类型: %s, 共 %d 条记录", logType, len(records))
	return records, nil
}

// GetLogByLevel 根据日志等级获取日志记录
func (s *sLogServiceImpl) GetLogByLevel(ctx context.Context, level string) ([]*s_db_model.SLogModel, error) {
	log := &s_db_model.SLogModel{}
	records, err := log.GetByLevel(ctx, level)
	if err != nil {
		g.Log().Errorf(ctx, "获取日志记录失败 - 等级: %s, 错误: %+v", level, err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取日志记录 - 等级: %s, 共 %d 条记录", level, len(records))
	return records, nil
}

// GetLogByDeviceIdAndType 根据设备ID和日志类型获取日志记录
func (s *sLogServiceImpl) GetLogByDeviceIdAndType(ctx context.Context, deviceId, logType string) ([]*s_db_model.SLogModel, error) {
	log := &s_db_model.SLogModel{}
	records, err := log.GetByDeviceIdAndType(ctx, deviceId, logType)
	if err != nil {
		g.Log().Errorf(ctx, "获取日志记录失败 - 设备ID: %s, 类型: %s, 错误: %+v", deviceId, logType, err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取日志记录 - 设备ID: %s, 类型: %s, 共 %d 条记录", deviceId, logType, len(records))
	return records, nil
}

// DeleteLogByDeviceId 根据设备ID删除日志记录
func (s *sLogServiceImpl) DeleteLogByDeviceId(ctx context.Context, deviceId string) error {
	log := &s_db_model.SLogModel{}
	err := log.DeleteByDeviceId(ctx, deviceId)
	if err != nil {
		g.Log().Errorf(ctx, "删除日志记录失败 - 设备ID: %s, 错误: %+v", deviceId, err)
		return err
	}

	g.Log().Infof(ctx, "成功删除日志记录 - 设备ID: %s", deviceId)
	return nil
}

// DeleteLogByType 根据日志类型删除日志记录
func (s *sLogServiceImpl) DeleteLogByType(ctx context.Context, logType string) error {
	log := &s_db_model.SLogModel{}
	err := log.DeleteByType(ctx, logType)
	if err != nil {
		g.Log().Errorf(ctx, "删除日志记录失败 - 类型: %s, 错误: %+v", logType, err)
		return err
	}

	g.Log().Infof(ctx, "成功删除日志记录 - 类型: %s", logType)
	return nil
}

// DeleteLogByFilters 根据过滤条件删除日志记录
func (s *sLogServiceImpl) DeleteLogByFilters(ctx context.Context, filters map[string]interface{}) (int, error) {
	log := &s_db_model.SLogModel{}
	deletedCount, err := log.DeleteByFilters(ctx, filters)
	if err != nil {
		g.Log().Errorf(ctx, "根据条件删除日志记录失败 - 过滤条件: %+v, 错误: %+v", filters, err)
		return 0, err
	}

	g.Log().Infof(ctx, "成功根据条件删除日志记录 - 过滤条件: %+v, 删除数量: %d", filters, deletedCount)
	return deletedCount, nil
}

// GetAllLog 获取所有日志记录
func (s *sLogServiceImpl) GetAllLog(ctx context.Context) ([]*s_db_model.SLogModel, error) {
	log := &s_db_model.SLogModel{}
	records, err := log.GetAll(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取所有日志记录失败 - 错误: %+v", err)
		return nil, err
	}

	g.Log().Infof(ctx, "成功获取所有日志记录，共 %d 条记录", len(records))
	return records, nil
}

// GetLogPage 分页获取日志记录
func (s *sLogServiceImpl) GetLogPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SLogModel, int, error) {
	log := &s_db_model.SLogModel{}
	records, total, err := log.GetPage(ctx, page, pageSize, filters)
	if err != nil {
		g.Log().Errorf(ctx, "分页获取日志记录失败 - 页码: %d, 页大小: %d, 过滤条件: %+v, 错误: %+v", page, pageSize, filters, err)
		return nil, 0, err
	}

	g.Log().Infof(ctx, "成功分页获取日志记录 - 页码: %d, 页大小: %d, 总数: %d, 当前页记录数: %d, 过滤条件: %+v", page, pageSize, total, len(records), filters)
	return records, total, nil
}

// ClearAllLog 清除所有日志记录
func (s *sLogServiceImpl) ClearAllLog(ctx context.Context) error {
	log := &s_db_model.SLogModel{}
	err := log.ClearAll(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "清除所有日志记录失败 - 错误: %+v", err)
		return err
	}

	g.Log().Infof(ctx, "成功清除所有日志记录并执行VACUUM")
	return nil
}

// GetLogCount 获取日志表记录总数
func (s *sLogServiceImpl) GetLogCount(ctx context.Context) (int, error) {
	log := &s_db_model.SLogModel{}
	count, err := log.GetCount(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取日志表记录总数失败 - 错误: %+v", err)
		return 0, err
	}

	g.Log().Infof(ctx, "成功获取日志表记录总数: %d", count)
	return count, nil
}

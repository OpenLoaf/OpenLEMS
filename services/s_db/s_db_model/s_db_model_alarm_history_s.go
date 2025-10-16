package s_db_model

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableAlarmHistory = "alarm_history"

	// 告警历史表特有字段
	FieldAlarmHistoryDeviceId       = "device_id"
	FieldAlarmHistorySourceDeviceId = "source_device_id"
	FieldAlarmHistoryPoint          = "point"
	FieldAlarmHistoryPointName      = "point_name"
	FieldAlarmHistoryLevel          = "level"
	FieldAlarmHistoryTitle          = "title"
	FieldAlarmHistoryDetail         = "detail"
	FieldAlarmHistoryTriggerAt      = "trigger_at"
	FieldAlarmHistoryClearAt        = "clear_at"
)

// 告警历史表结构
type SAlarmHistoryModel struct {
	g.Meta         `orm:"table:alarm_history"`
	Id             int        `json:"id" orm:"id,primary,auto_increment"`
	DeviceId       string     `json:"deviceId" orm:"device_id"`
	SourceDeviceId string     `json:"sourceDeviceId" orm:"source_device_id"` // 原设备ID
	Point          string     `json:"point" orm:"point"`
	Level          string     `json:"level" orm:"level"`
	PointName      string     `json:"pointName" orm:"point_name"`
	Detail         string     `json:"detail" orm:"detail"`
	TriggerAt      *time.Time `json:"triggerAt" orm:"trigger_at"`
	ClearAt        *time.Time `json:"clearAt" orm:"clear_at,auto_now_add"`
}

// Create 创建告警历史记录
func (a *SAlarmHistoryModel) Create(ctx context.Context) error {
	// 排除ID字段，让数据库自动生成
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).FieldsEx("id").Insert(a)
	return err
}

// GetById 根据ID获取告警历史记录
func (a *SAlarmHistoryModel) GetById(ctx context.Context, id int) error {
	return g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldId, id).Scan(a)
}

// GetByDeviceId 根据设备ID获取告警历史记录
func (a *SAlarmHistoryModel) GetByDeviceId(ctx context.Context, deviceId string) ([]*SAlarmHistoryModel, error) {
	var records []*SAlarmHistoryModel
	err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldAlarmHistoryDeviceId, deviceId).Scan(&records)
	return records, err
}

// GetByDeviceIdAndPoint 根据设备ID和点位获取告警历史记录
func (a *SAlarmHistoryModel) GetByDeviceIdAndPoint(ctx context.Context, deviceId, point string) ([]*SAlarmHistoryModel, error) {
	var records []*SAlarmHistoryModel
	err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldAlarmHistoryDeviceId, deviceId).Where(FieldAlarmHistoryPoint, point).Scan(&records)
	return records, err
}

// Update 更新告警历史记录
func (a *SAlarmHistoryModel) Update(ctx context.Context) error {
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldId, a.Id).Update(a)
	return err
}

// Delete 删除告警历史记录
func (a *SAlarmHistoryModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldId, a.Id).Delete()
	return err
}

// DeleteByDeviceId 根据设备ID删除告警历史记录
func (a *SAlarmHistoryModel) DeleteByDeviceId(ctx context.Context, deviceId string) error {
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldAlarmHistoryDeviceId, deviceId).Delete()
	return err
}

// DeleteByLevel 根据级别删除告警历史记录
func (a *SAlarmHistoryModel) DeleteByLevel(ctx context.Context, level string) error {
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldAlarmHistoryLevel, level).Delete()
	return err
}

// DeleteByDeviceIdAndLevel 根据设备ID与级别删除告警历史记录
func (a *SAlarmHistoryModel) DeleteByDeviceIdAndLevel(ctx context.Context, deviceId, level string) error {
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).
		Where(FieldAlarmHistoryDeviceId, deviceId).
		Where(FieldAlarmHistoryLevel, level).
		Delete()
	return err
}

// DeleteByFilters 根据过滤条件删除告警历史记录
// deviceId 为空表示所有设备，level 为空表示所有级别
func (a *SAlarmHistoryModel) DeleteByFilters(ctx context.Context, deviceId, level string) error {
	// 当两个过滤条件都为空时，明确执行全表清空（遵循 ClearAll 的逻辑）
	if deviceId == "" && level == "" {
		return a.ClearAll(ctx)
	}

	model := g.Model(TableAlarmHistory).Ctx(ctx)

	if deviceId != "" {
		model = model.Where(FieldAlarmHistoryDeviceId, deviceId)
	}
	if level != "" {
		model = model.Where(FieldAlarmHistoryLevel, level)
	}

	_, err := model.Delete()
	return err
}

// GetAll 获取所有告警历史记录
func (a *SAlarmHistoryModel) GetAll(ctx context.Context) ([]*SAlarmHistoryModel, error) {
	var records []*SAlarmHistoryModel
	err := g.Model(TableAlarmHistory).Ctx(ctx).Scan(&records)
	return records, err
}

// GetPage 分页获取告警历史记录
func (a *SAlarmHistoryModel) GetPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*SAlarmHistoryModel, int, error) {
	model := g.Model(TableAlarmHistory).Ctx(ctx)

	// 应用过滤条件
	if filters != nil {
		// 日期过滤 (yyyy-MM-dd)
		if date, ok := filters["date"].(string); ok && date != "" {
			model = model.Where("DATE(trigger_at) = ?", date)
		}

		// 设备ID过滤
		if deviceId, ok := filters[FieldAlarmHistoryDeviceId].(string); ok && deviceId != "" {
			model = model.Where(FieldAlarmHistoryDeviceId, deviceId)
		}

		// 级别过滤
		if level, ok := filters[FieldAlarmHistoryLevel].(string); ok && level != "" {
			model = model.Where(FieldAlarmHistoryLevel, level)
		}

		// 点位过滤
		if point, ok := filters[FieldAlarmHistoryPoint].(string); ok && point != "" {
			model = model.Where(FieldAlarmHistoryPoint, point)
		}

		// 标题模糊搜索
		if title, ok := filters[FieldAlarmHistoryTitle].(string); ok && title != "" {
			model = model.Where(FieldAlarmHistoryTitle+" LIKE ?", "%"+title+"%")
		}
	}

	// 计算总数
	total, err := model.Count()
	if err != nil {
		return nil, 0, err
	}

	// 计算偏移量
	offset := (page - 1) * pageSize

	// 分页查询
	var records []*SAlarmHistoryModel
	err = model.Order("trigger_at DESC").Limit(offset, pageSize).Scan(&records)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// ClearAll 清除所有告警历史记录并执行VACUUM
func (a *SAlarmHistoryModel) ClearAll(ctx context.Context) error {
	// 删除所有记录 - 使用 WHERE 1=1 来满足安全要求
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Where("1=1").Delete()
	if err != nil {
		return err
	}

	// 执行VACUUM操作回收空间
	_, err = g.DB().Exec(ctx, "VACUUM")
	return err
}

// GetCount 获取告警历史表记录总数
func (a *SAlarmHistoryModel) GetCount(ctx context.Context) (int, error) {
	count, err := g.Model(TableAlarmHistory).Ctx(ctx).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetCountByDeviceId 根据设备ID获取告警历史表记录总数
func (a *SAlarmHistoryModel) GetCountByDeviceId(ctx context.Context, deviceId string) (int, error) {
	count, err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldAlarmHistoryDeviceId, deviceId).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

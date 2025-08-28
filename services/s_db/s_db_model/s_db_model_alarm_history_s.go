package s_db_model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableAlarmHistory = "alarm_history"

	// 字段名
	FieldDeviceId = "device_id"
	FieldPoint    = "point"
	FieldLevel    = "level"
	FieldTitle    = "title"
	FieldDetail   = "detail"
)

// 告警历史表结构
type SAlarmHistoryModel struct {
	g.Meta    `orm:"table:alarm_history"`
	Id        int    `json:"id" orm:"id"`
	DeviceId  string `json:"device_id" orm:"device_id"`
	Point     string `json:"point" orm:"point"`
	Level     string `json:"level" orm:"level"`
	Title     string `json:"title" orm:"title"`
	Detail    string `json:"detail" orm:"detail"`
	CreatedAt string `json:"created_at" orm:"created_at"`
}

// Create 创建告警历史记录
func (a *SAlarmHistoryModel) Create(ctx context.Context) error {
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Insert(a)
	return err
}

// GetById 根据ID获取告警历史记录
func (a *SAlarmHistoryModel) GetById(ctx context.Context, id int) error {
	return g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldId, id).Scan(a)
}

// GetByDeviceId 根据设备ID获取告警历史记录
func (a *SAlarmHistoryModel) GetByDeviceId(ctx context.Context, deviceId string) ([]*SAlarmHistoryModel, error) {
	var records []*SAlarmHistoryModel
	err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldDeviceId, deviceId).Scan(&records)
	return records, err
}

// GetByDeviceIdAndPoint 根据设备ID和点位获取告警历史记录
func (a *SAlarmHistoryModel) GetByDeviceIdAndPoint(ctx context.Context, deviceId, point string) ([]*SAlarmHistoryModel, error) {
	var records []*SAlarmHistoryModel
	err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldDeviceId, deviceId).Where(FieldPoint, point).Scan(&records)
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
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Where(FieldDeviceId, deviceId).Delete()
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
			model = model.Where("DATE(created_at) = ?", date)
		}

		// 设备ID过滤
		if deviceId, ok := filters["device_id"].(string); ok && deviceId != "" {
			model = model.Where(FieldDeviceId, deviceId)
		}

		// 级别过滤
		if level, ok := filters["level"].(string); ok && level != "" {
			model = model.Where(FieldLevel, level)
		}

		// 点位过滤
		if point, ok := filters["point"].(string); ok && point != "" {
			model = model.Where(FieldPoint, point)
		}

		// 标题模糊搜索
		if title, ok := filters["title"].(string); ok && title != "" {
			model = model.Where(FieldTitle+" LIKE ?", "%"+title+"%")
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
	err = model.Order("created_at DESC").Limit(offset, pageSize).Scan(&records)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// ClearAll 清除所有告警历史记录并执行VACUUM
func (a *SAlarmHistoryModel) ClearAll(ctx context.Context) error {
	// 删除所有记录
	_, err := g.Model(TableAlarmHistory).Ctx(ctx).Delete()
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

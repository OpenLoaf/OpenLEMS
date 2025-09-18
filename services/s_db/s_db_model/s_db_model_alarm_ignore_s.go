package s_db_model

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableAlarmIgnore = "alarm_ignore"

	// 告警忽略表特有字段
	FieldAlarmIgnoreDeviceId       = "device_id"
	FieldAlarmIgnoreSourceDeviceId = "source_device_id"
	FieldAlarmIgnorePoint          = "point"
	FieldAlarmIgnorePointName      = "point_name"
)

// 告警忽略表结构
type SAlarmIgnoreModel struct {
	g.Meta         `orm:"table:alarm_ignore"`
	Id             int        `json:"id" orm:"id,primary,auto_increment"`
	DeviceId       string     `json:"deviceId" orm:"device_id"`
	SourceDeviceId string     `json:"sourceDeviceId" orm:"source_device_id"` // 原设备ID
	Point          string     `json:"point" orm:"point"`
	PointName      string     `json:"pointName" orm:"point_name"`
	CreatedAt      *time.Time `json:"createdAt" orm:"created_at,auto_now_add"`
}

// Create 创建告警忽略记录
func (a *SAlarmIgnoreModel) Create(ctx context.Context) error {
	// 排除ID字段，让数据库自动生成
	_, err := g.Model(TableAlarmIgnore).Ctx(ctx).FieldsEx("id").Insert(a)
	return err
}

// GetById 根据ID获取告警忽略记录
func (a *SAlarmIgnoreModel) GetById(ctx context.Context, id int) error {
	return g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldId, id).Scan(a)
}

// GetByDeviceId 根据设备ID获取告警忽略记录
func (a *SAlarmIgnoreModel) GetByDeviceId(ctx context.Context, deviceId string) ([]*SAlarmIgnoreModel, error) {
	var records []*SAlarmIgnoreModel
	err := g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldAlarmIgnoreDeviceId, deviceId).Scan(&records)
	return records, err
}

// GetByDeviceIdAndPoint 根据设备ID和点位获取告警忽略记录
func (a *SAlarmIgnoreModel) GetByDeviceIdAndPoint(ctx context.Context, deviceId, point string) error {
	return g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldAlarmIgnoreDeviceId, deviceId).Where(FieldAlarmIgnorePoint, point).Scan(a)
}

// IsIgnored 检查是否被忽略（设备+源设备+点位）
func (a *SAlarmIgnoreModel) IsIgnored(ctx context.Context, deviceId, sourceDeviceId, point string) (bool, error) {
	count, err := g.Model(TableAlarmIgnore).Ctx(ctx).
		Where(FieldAlarmIgnoreDeviceId, deviceId).
		Where(FieldAlarmIgnoreSourceDeviceId, sourceDeviceId).
		Where(FieldAlarmIgnorePoint, point).Count()
	return count > 0, err
}

// Update 更新告警忽略记录
func (a *SAlarmIgnoreModel) Update(ctx context.Context) error {
	_, err := g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldId, a.Id).Update(a)
	return err
}

// Delete 删除告警忽略记录
func (a *SAlarmIgnoreModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldId, a.Id).Delete()
	return err
}

// DeleteByDeviceId 根据设备ID删除告警忽略记录
func (a *SAlarmIgnoreModel) DeleteByDeviceId(ctx context.Context, deviceId string) error {
	_, err := g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldAlarmIgnoreDeviceId, deviceId).Delete()
	return err
}

// DeleteByDeviceIdAndPoint 根据设备ID和点位删除告警忽略记录
func (a *SAlarmIgnoreModel) DeleteByDeviceIdAndPoint(ctx context.Context, deviceId, point string) error {
	_, err := g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldAlarmIgnoreDeviceId, deviceId).Where(FieldAlarmIgnorePoint, point).Delete()
	return err
}

// GetAll 获取所有告警忽略记录
func (a *SAlarmIgnoreModel) GetAll(ctx context.Context) ([]*SAlarmIgnoreModel, error) {
	var records []*SAlarmIgnoreModel
	err := g.Model(TableAlarmIgnore).Ctx(ctx).Scan(&records)
	return records, err
}

// GetPage 分页获取告警忽略记录
func (a *SAlarmIgnoreModel) GetPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*SAlarmIgnoreModel, int, error) {
	model := g.Model(TableAlarmIgnore).Ctx(ctx)

	// 应用过滤条件
	if filters != nil {
		// 日期过滤 (yyyy-MM-dd)
		if date, ok := filters["date"].(string); ok && date != "" {
			model = model.Where("DATE(created_at) = ?", date)
		}

		// 设备ID过滤
		if deviceId, ok := filters[FieldAlarmIgnoreDeviceId].(string); ok && deviceId != "" {
			model = model.Where(FieldAlarmIgnoreDeviceId, deviceId)
		}

		// 点位过滤
		if point, ok := filters[FieldAlarmIgnorePoint].(string); ok && point != "" {
			model = model.Where(FieldAlarmIgnorePoint, point)
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
	var records []*SAlarmIgnoreModel
	err = model.Order("created_at DESC").Limit(offset, pageSize).Scan(&records)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// GetCount 获取告警忽略表记录总数
func (a *SAlarmIgnoreModel) GetCount(ctx context.Context) (int, error) {
	count, err := g.Model(TableAlarmIgnore).Ctx(ctx).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetCountByDeviceId 根据设备ID获取告警忽略表记录总数
func (a *SAlarmIgnoreModel) GetCountByDeviceId(ctx context.Context, deviceId string) (int, error) {
	count, err := g.Model(TableAlarmIgnore).Ctx(ctx).Where(FieldAlarmIgnoreDeviceId, deviceId).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

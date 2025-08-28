package s_db_model

import (
	"context"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableLog = "log"

	// 字段名
	FieldContent = "content"
)

// 日志表结构
type SLogModel struct {
	g.Meta    `orm:"table:log"`
	Id        int        `json:"id" orm:"id,primary,auto_increment"`
	Type      string     `json:"type" orm:"type"`
	DeviceId  string     `json:"device_id" orm:"device_id"`
	Level     string     `json:"level" orm:"level"`
	Content   string     `json:"content" orm:"content"`
	CreatedAt *time.Time `json:"created_at" orm:"created_at,auto_now_add"`
}

// Create 创建日志记录
func (l *SLogModel) Create(ctx context.Context) error {
	// 排除ID字段，让数据库自动生成
	_, err := g.Model(TableLog).Ctx(ctx).FieldsEx("id").Insert(l)
	return err
}

// GetById 根据ID获取日志记录
func (l *SLogModel) GetById(ctx context.Context, id int) error {
	return g.Model(TableLog).Ctx(ctx).Where(FieldId, id).Scan(l)
}

// GetByDeviceId 根据设备ID获取日志记录
func (l *SLogModel) GetByDeviceId(ctx context.Context, deviceId string) ([]*SLogModel, error) {
	var records []*SLogModel
	err := g.Model(TableLog).Ctx(ctx).Where(FieldDeviceId, deviceId).Scan(&records)
	return records, err
}

// GetByType 根据日志类型获取日志记录
func (l *SLogModel) GetByType(ctx context.Context, logType string) ([]*SLogModel, error) {
	var records []*SLogModel
	err := g.Model(TableLog).Ctx(ctx).Where(FieldType, logType).Scan(&records)
	return records, err
}

// GetByLevel 根据日志等级获取日志记录
func (l *SLogModel) GetByLevel(ctx context.Context, level string) ([]*SLogModel, error) {
	var records []*SLogModel
	err := g.Model(TableLog).Ctx(ctx).Where(FieldLevel, level).Scan(&records)
	return records, err
}

// GetByDeviceIdAndType 根据设备ID和日志类型获取日志记录
func (l *SLogModel) GetByDeviceIdAndType(ctx context.Context, deviceId, logType string) ([]*SLogModel, error) {
	var records []*SLogModel
	err := g.Model(TableLog).Ctx(ctx).Where(FieldDeviceId, deviceId).Where(FieldType, logType).Scan(&records)
	return records, err
}

// Update 更新日志记录
func (l *SLogModel) Update(ctx context.Context) error {
	_, err := g.Model(TableLog).Ctx(ctx).Where(FieldId, l.Id).Update(l)
	return err
}

// Delete 删除日志记录
func (l *SLogModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableLog).Ctx(ctx).Where(FieldId, l.Id).Delete()
	return err
}

// DeleteByDeviceId 根据设备ID删除日志记录
func (l *SLogModel) DeleteByDeviceId(ctx context.Context, deviceId string) error {
	_, err := g.Model(TableLog).Ctx(ctx).Where(FieldDeviceId, deviceId).Delete()
	return err
}

// DeleteByType 根据日志类型删除日志记录
func (l *SLogModel) DeleteByType(ctx context.Context, logType string) error {
	_, err := g.Model(TableLog).Ctx(ctx).Where(FieldType, logType).Delete()
	return err
}

// GetAll 获取所有日志记录
func (l *SLogModel) GetAll(ctx context.Context) ([]*SLogModel, error) {
	var records []*SLogModel
	err := g.Model(TableLog).Ctx(ctx).Scan(&records)
	return records, err
}

// GetPage 分页获取日志记录
func (l *SLogModel) GetPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*SLogModel, int, error) {
	model := g.Model(TableLog).Ctx(ctx)

	// 应用过滤条件
	if filters != nil {
		// 日期过滤 (yyyy-MM-dd)
		if date, ok := filters["date"].(string); ok && date != "" {
			model = model.Where("DATE(created_at) = ?", date)
		}

		// 类型过滤
		if logType, ok := filters["type"].(string); ok && logType != "" {
			model = model.Where(FieldType, logType)
		}

		// 级别过滤
		if level, ok := filters["level"].(string); ok && level != "" {
			model = model.Where(FieldLevel, level)
		}

		// 设备ID过滤
		if deviceId, ok := filters["device_id"].(string); ok && deviceId != "" {
			model = model.Where(FieldDeviceId, deviceId)
		}

		// 内容模糊搜索
		if content, ok := filters["content"].(string); ok && content != "" {
			model = model.Where(FieldContent+" LIKE ?", "%"+content+"%")
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
	var records []*SLogModel
	err = model.Order("created_at DESC").Limit(offset, pageSize).Scan(&records)
	if err != nil {
		return nil, 0, err
	}

	return records, total, nil
}

// ClearAll 清除所有日志记录并执行VACUUM
func (l *SLogModel) ClearAll(ctx context.Context) error {
	// 删除所有记录
	_, err := g.Model(TableLog).Ctx(ctx).Delete()
	if err != nil {
		return err
	}

	// 执行VACUUM操作回收空间
	_, err = g.DB().Exec(ctx, "VACUUM")
	return err
}

// GetCount 获取日志表记录总数
func (l *SLogModel) GetCount(ctx context.Context) (int, error) {
	count, err := g.Model(TableLog).Ctx(ctx).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

package s_db_model

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableSetting = "setting"
)

// 设置表结构
type SSettingModel struct {
	g.Meta `orm:"table:setting"`
	SDatabaseBasic
	Value    string `json:"value" orm:"value"`
	IsPublic bool   `json:"isPublic" orm:"is_public"`
	Enabled  bool   `json:"enable" orm:"enabled"`
	Remark   string `json:"remark" orm:"remark"`
	Sort     int    `json:"sort" orm:"sort"`
}

// GetValue 获取设置值
func (s *SSettingModel) GetValue() string {
	return s.Value
}

// SetValue 设置值
func (s *SSettingModel) SetValue(value string) {
	s.Value = value
}

// Create 创建设置记录
func (s *SSettingModel) Create(ctx context.Context) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Insert(s)
	return err
}

// GetById 根据ID获取设置记录
func (s *SSettingModel) GetById(ctx context.Context, id string) error {
	return g.Model(TableSetting).Ctx(ctx).Where(FieldId, id).Scan(s)
}

// Update 更新设置记录
func (s *SSettingModel) Update(ctx context.Context) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(FieldId, s.Id).Update(s)
	return err
}

// UpdateFields 更新指定字段
func (s *SSettingModel) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(FieldId, s.Id).Update(data)
	return err
}

// Delete 删除设置记录
func (s *SSettingModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(FieldId, s.Id).Delete()
	return err
}

// DeleteById 根据ID删除设置记录
func DeleteSettingById(ctx context.Context, id uint) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(FieldId, id).Delete()
	return err
}

// DeleteByName 根据名称删除设置记录
func DeleteSettingByName(ctx context.Context, name string) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(FieldName, name).Delete()
	return err
}

// GetAll 获取所有设置记录
func GetAllSettings(ctx context.Context) ([]*SSettingModel, error) {
	var settings []*SSettingModel
	err := g.Model(TableSetting).Ctx(ctx).Scan(&settings)
	return settings, err
}

// GetByCondition 根据条件获取设置记录
func GetSettingsByCondition(ctx context.Context, condition g.Map) ([]*SSettingModel, error) {
	var settings []*SSettingModel
	err := g.Model(TableSetting).Ctx(ctx).Where(condition).Scan(&settings)
	return settings, err
}

// GetEnabledSettings 获取所有启用的设置
func GetEnabledSettings(ctx context.Context) ([]*SSettingModel, error) {
	var settings []*SSettingModel
	err := g.Model(TableSetting).Ctx(ctx).Where(FieldEnable, true).Scan(&settings)
	return settings, err
}

// Count 获取设置总数
func CountSettings(ctx context.Context) (int, error) {
	count, err := g.Model(TableSetting).Ctx(ctx).Count()
	return count, err
}

// CountByCondition 根据条件获取设置数量
func CountSettingsByCondition(ctx context.Context, condition g.Map) (int, error) {
	count, err := g.Model(TableSetting).Ctx(ctx).Where(condition).Count()
	return count, err
}

// Paginate 分页获取设置列表
func PaginateSettings(ctx context.Context, page, pageSize int) ([]*SSettingModel, error) {
	var settings []*SSettingModel
	err := g.Model(TableSetting).Ctx(ctx).Page(page, pageSize).Scan(&settings)
	return settings, err
}

// IsEnabled 检查设置是否启用
func (s *SSettingModel) IsEnabled() bool {
	return s.Enabled
}

// SetEnabled 设置启用状态
func (s *SSettingModel) SetEnabled(ctx context.Context, enabled bool) error {
	return s.UpdateFields(ctx, g.Map{FieldEnable: enabled})
}

// UpdateValue 更新设置值
func (s *SSettingModel) UpdateValue(ctx context.Context, value string) error {
	s.Value = value
	return s.UpdateFields(ctx, g.Map{FieldValue: value})
}

// UpdateRemark 更新备注
func (s *SSettingModel) UpdateRemark(ctx context.Context, remark string) error {
	s.Remark = remark
	return s.UpdateFields(ctx, g.Map{FieldRemark: remark})
}

// GetAllSettingsOrderBySort 获取所有设置记录，按sort字段排序
func GetAllSettingsOrderBySort(ctx context.Context) ([]*SSettingModel, error) {
	var settings []*SSettingModel
	err := g.Model(TableSetting).Ctx(ctx).Order(FieldSort).Scan(&settings)
	return settings, err
}

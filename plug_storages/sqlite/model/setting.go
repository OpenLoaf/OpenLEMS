package model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableSetting = "setting"

	// 字段名
	SettingFieldId     = "id"
	SettingFieldName   = "name"
	SettingFieldValue  = "value"
	SettingFieldEnable = "enable"
	SettingFieldRemark = "remark"
)

// 设置表结构
type Setting struct {
	g.Meta `orm:"table:setting"`
	Id     uint   `json:"id" orm:"id,primary"`
	Name   string `json:"name" orm:"name"`
	Value  string `json:"value" orm:"value"`
	Enable bool   `json:"enable" orm:"enable"`
	Remark string `json:"remark" orm:"remark"`
}

// GetValue 获取设置值
func (s *Setting) GetValue() string {
	return s.Value
}

// SetValue 设置值
func (s *Setting) SetValue(value string) {
	s.Value = value
}

// Create 创建设置记录
func (s *Setting) Create(ctx context.Context) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Insert(s)
	return err
}

// GetById 根据ID获取设置记录
func (s *Setting) GetById(ctx context.Context, id uint) error {
	return g.Model(TableSetting).Ctx(ctx).Where(SettingFieldId, id).Scan(s)
}

// GetByName 根据名称获取设置记录
func (s *Setting) GetByName(ctx context.Context, name string) error {
	return g.Model(TableSetting).Ctx(ctx).Where(SettingFieldName, name).Scan(s)
}

// Update 更新设置记录
func (s *Setting) Update(ctx context.Context) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(SettingFieldId, s.Id).Update(s)
	return err
}

// UpdateFields 更新指定字段
func (s *Setting) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(SettingFieldId, s.Id).Update(data)
	return err
}

// Delete 删除设置记录
func (s *Setting) Delete(ctx context.Context) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(SettingFieldId, s.Id).Delete()
	return err
}

// DeleteById 根据ID删除设置记录
func DeleteSettingById(ctx context.Context, id uint) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(SettingFieldId, id).Delete()
	return err
}

// DeleteByName 根据名称删除设置记录
func DeleteSettingByName(ctx context.Context, name string) error {
	_, err := g.Model(TableSetting).Ctx(ctx).Where(SettingFieldName, name).Delete()
	return err
}

// GetAll 获取所有设置记录
func GetAllSettings(ctx context.Context) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model(TableSetting).Ctx(ctx).Scan(&settings)
	return settings, err
}

// GetByCondition 根据条件获取设置记录
func GetSettingsByCondition(ctx context.Context, condition g.Map) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model(TableSetting).Ctx(ctx).Where(condition).Scan(&settings)
	return settings, err
}

// GetEnabledSettings 获取所有启用的设置
func GetEnabledSettings(ctx context.Context) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model(TableSetting).Ctx(ctx).Where(SettingFieldEnable, true).Scan(&settings)
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
func PaginateSettings(ctx context.Context, page, pageSize int) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model(TableSetting).Ctx(ctx).Page(page, pageSize).Scan(&settings)
	return settings, err
}

// IsEnabled 检查设置是否启用
func (s *Setting) IsEnabled() bool {
	return s.Enable
}

// SetEnabled 设置启用状态
func (s *Setting) SetEnabled(ctx context.Context, enabled bool) error {
	return s.UpdateFields(ctx, g.Map{SettingFieldEnable: enabled})
}

// UpdateValue 更新设置值
func (s *Setting) UpdateValue(ctx context.Context, value string) error {
	s.Value = value
	return s.UpdateFields(ctx, g.Map{SettingFieldValue: value})
}

// UpdateRemark 更新备注
func (s *Setting) UpdateRemark(ctx context.Context, remark string) error {
	s.Remark = remark
	return s.UpdateFields(ctx, g.Map{SettingFieldRemark: remark})
}

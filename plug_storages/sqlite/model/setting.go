package model

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
)

// 设置表结构
type Setting struct {
	g.Meta `orm:"table:setting"`
	Id     uint   `json:"id" orm:"id,primary"`
	Name   string `json:"name" orm:"name"`
	Params string `json:"params" orm:"params"`
	Enable bool   `json:"enable" orm:"enable"`
}

// GetParamsMap 获取参数的map格式
func (s *Setting) GetParamsMap() (map[string]string, error) {
	if s.Params == "" || s.Params == "null" {
		return map[string]string{}, nil
	}

	// 先反序列化为 map[string]interface{} 来处理混合类型
	var paramsMapInterface map[string]interface{}
	err := gjson.DecodeTo(s.Params, &paramsMapInterface)
	if err != nil {
		return nil, err
	}

	// 转换为 map[string]string
	paramsMap := make(map[string]string)
	for key, value := range paramsMapInterface {
		paramsMap[key] = fmt.Sprintf("%v", value)
	}

	return paramsMap, nil
}

// SetParamsFromMap 从map设置参数
func (s *Setting) SetParamsFromMap(paramsMap g.Map) error {
	if paramsMap == nil {
		s.Params = ""
		return nil
	}

	paramsJSON, err := gjson.Encode(paramsMap)
	if err != nil {
		return err
	}

	s.Params = string(paramsJSON)
	return nil
}

// Create 创建设置记录
func (s *Setting) Create(ctx context.Context) error {
	_, err := g.Model("setting").Ctx(ctx).Insert(s)
	return err
}

// GetById 根据ID获取设置记录
func (s *Setting) GetById(ctx context.Context, id uint) error {
	return g.Model("setting").Ctx(ctx).Where("id", id).Scan(s)
}

// GetByName 根据名称获取设置记录
func (s *Setting) GetByName(ctx context.Context, name string) error {
	return g.Model("setting").Ctx(ctx).Where("name", name).Scan(s)
}

// Update 更新设置记录
func (s *Setting) Update(ctx context.Context) error {
	_, err := g.Model("setting").Ctx(ctx).Where("id", s.Id).Update(s)
	return err
}

// UpdateFields 更新指定字段
func (s *Setting) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model("setting").Ctx(ctx).Where("id", s.Id).Update(data)
	return err
}

// Delete 删除设置记录
func (s *Setting) Delete(ctx context.Context) error {
	_, err := g.Model("setting").Ctx(ctx).Where("id", s.Id).Delete()
	return err
}

// DeleteById 根据ID删除设置记录
func DeleteSettingById(ctx context.Context, id uint) error {
	_, err := g.Model("setting").Ctx(ctx).Where("id", id).Delete()
	return err
}

// GetAll 获取所有设置记录
func GetAllSettings(ctx context.Context) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model("setting").Ctx(ctx).Scan(&settings)
	return settings, err
}

// GetByCondition 根据条件获取设置记录
func GetSettingsByCondition(ctx context.Context, condition g.Map) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model("setting").Ctx(ctx).Where(condition).Scan(&settings)
	return settings, err
}

// GetEnabledSettings 获取所有启用的设置
func GetEnabledSettings(ctx context.Context) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model("setting").Ctx(ctx).Where("enable", true).Scan(&settings)
	return settings, err
}

// GetDisabledSettings 获取所有禁用的设置
func GetDisabledSettings(ctx context.Context) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model("setting").Ctx(ctx).Where("enable", false).Scan(&settings)
	return settings, err
}

// Count 获取设置总数
func CountSettings(ctx context.Context) (int, error) {
	count, err := g.Model("setting").Ctx(ctx).Count()
	return count, err
}

// CountByCondition 根据条件获取设置数量
func CountSettingsByCondition(ctx context.Context, condition g.Map) (int, error) {
	count, err := g.Model("setting").Ctx(ctx).Where(condition).Count()
	return count, err
}

// Paginate 分页获取设置列表
func PaginateSettings(ctx context.Context, page, pageSize int) ([]*Setting, error) {
	var settings []*Setting
	err := g.Model("setting").Ctx(ctx).Page(page, pageSize).Scan(&settings)
	return settings, err
}

// Exists 检查设置是否存在
func (s *Setting) Exists(ctx context.Context) (bool, error) {
	count, err := g.Model("setting").Ctx(ctx).Where("id", s.Id).Count()
	return count > 0, err
}

// ExistsByName 根据名称检查设置是否存在
func ExistsSettingByName(ctx context.Context, name string) (bool, error) {
	count, err := g.Model("setting").Ctx(ctx).Where("name", name).Count()
	return count > 0, err
}

// ToggleEnable 切换启用状态
func (s *Setting) ToggleEnable(ctx context.Context) error {
	s.Enable = !s.Enable
	return s.Update(ctx)
}

// SetEnable 设置启用状态
func (s *Setting) SetEnable(ctx context.Context, enable bool) error {
	return s.UpdateFields(ctx, g.Map{"enable": enable})
}

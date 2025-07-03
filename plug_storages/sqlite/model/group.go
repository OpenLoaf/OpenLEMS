package model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
)

// 数据库相关常量
const (
	// 表名
	TableGroup = "group"
	
	// 字段名
	GroupFieldId      = "id"
	GroupFieldName    = "name"
	GroupFieldVersion = "version"
)

// 分组表结构
type Group struct {
	g.Meta  `orm:"table:group"`
	Id      uint   `json:"id" orm:"id,primary"`
	Name    string `json:"name" orm:"name"`
	Version string `json:"version" orm:"version"`
}

// Create 创建分组记录
func (group *Group) Create(ctx context.Context) error {
	_, err := g.Model(TableGroup).Ctx(ctx).Insert(group)
	return err
}

// GetById 根据ID获取分组记录
func (group *Group) GetById(ctx context.Context, id uint) error {
	return g.Model(TableGroup).Ctx(ctx).Where(GroupFieldId, id).Scan(group)
}

// GetByName 根据名称获取分组记录
func (group *Group) GetByName(ctx context.Context, name string) error {
	return g.Model(TableGroup).Ctx(ctx).Where(GroupFieldName, name).Scan(group)
}

// Update 更新分组记录
func (group *Group) Update(ctx context.Context) error {
	_, err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldId, group.Id).Update(group)
	return err
}

// UpdateFields 更新指定字段
func (group *Group) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldId, group.Id).Update(data)
	return err
}

// Delete 删除分组记录
func (group *Group) Delete(ctx context.Context) error {
	_, err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldId, group.Id).Delete()
	return err
}

// DeleteById 根据ID删除分组记录
func DeleteGroupById(ctx context.Context, id uint) error {
	_, err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldId, id).Delete()
	return err
}

// GetAll 获取所有分组记录
func GetAllGroups(ctx context.Context) ([]*Group, error) {
	var groups []*Group
	err := g.Model(TableGroup).Ctx(ctx).Scan(&groups)
	return groups, err
}

// GetByCondition 根据条件获取分组记录
func GetGroupsByCondition(ctx context.Context, condition g.Map) ([]*Group, error) {
	var groups []*Group
	err := g.Model(TableGroup).Ctx(ctx).Where(condition).Scan(&groups)
	return groups, err
}

// GetByVersion 根据版本获取分组列表
func GetGroupsByVersion(ctx context.Context, version string) ([]*Group, error) {
	var groups []*Group
	err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldVersion, version).Scan(&groups)
	return groups, err
}

// Count 获取分组总数
func CountGroups(ctx context.Context) (int, error) {
	count, err := g.Model(TableGroup).Ctx(ctx).Count()
	return count, err
}

// CountByCondition 根据条件获取分组数量
func CountGroupsByCondition(ctx context.Context, condition g.Map) (int, error) {
	count, err := g.Model(TableGroup).Ctx(ctx).Where(condition).Count()
	return count, err
}

// Paginate 分页获取分组列表
func PaginateGroups(ctx context.Context, page, pageSize int) ([]*Group, error) {
	var groups []*Group
	err := g.Model(TableGroup).Ctx(ctx).Page(page, pageSize).Scan(&groups)
	return groups, err
}

// Exists 检查分组是否存在
func (group *Group) Exists(ctx context.Context) (bool, error) {
	count, err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldId, group.Id).Count()
	return count > 0, err
}

// ExistsByName 根据名称检查分组是否存在
func ExistsGroupByName(ctx context.Context, name string) (bool, error) {
	count, err := g.Model(TableGroup).Ctx(ctx).Where(GroupFieldName, name).Count()
	return count > 0, err
}

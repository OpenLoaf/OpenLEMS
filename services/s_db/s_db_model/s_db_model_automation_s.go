package s_db_model

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// 数据库相关常量
const (
	// 表名
	TableAutomation = "automation"
)

// 自动化表结构
type SAutomationModel struct {
	g.Meta         `orm:"table:automation"`
	Id             int         `json:"id" orm:"id"`
	StartTime      *gtime.Time `json:"startTime" orm:"start_time"`
	EndTime        *gtime.Time `json:"endTime" orm:"end_time"`
	TimeRangeType  string      `json:"timeRangeType" orm:"time_range_type"`
	TimeRangeValue string      `json:"timeRangeValue" orm:"time_range_value"`
	TriggerRule    string      `json:"triggerRule" orm:"trigger_rule"`
	ExecuteRule    string      `json:"executeRule" orm:"execute_rule"`
	Enable         bool        `json:"enable" orm:"enable"`
	CreatedAt      *gtime.Time `json:"createdAt" orm:"created_at"`
	UpdatedAt      *gtime.Time `json:"updatedAt" orm:"updated_at"`
}

// Getter/Setter 方法
func (s *SAutomationModel) GetId() int {
	return s.Id
}

func (s *SAutomationModel) SetId(id int) {
	s.Id = id
}

func (s *SAutomationModel) GetStartTime() *gtime.Time {
	return s.StartTime
}

func (s *SAutomationModel) SetStartTime(startTime *gtime.Time) {
	s.StartTime = startTime
}

func (s *SAutomationModel) GetEndTime() *gtime.Time {
	return s.EndTime
}

func (s *SAutomationModel) SetEndTime(endTime *gtime.Time) {
	s.EndTime = endTime
}

func (s *SAutomationModel) GetTimeRangeType() string {
	return s.TimeRangeType
}

func (s *SAutomationModel) SetTimeRangeType(timeRangeType string) {
	s.TimeRangeType = timeRangeType
}

func (s *SAutomationModel) GetTimeRangeValue() string {
	return s.TimeRangeValue
}

func (s *SAutomationModel) SetTimeRangeValue(timeRangeValue string) {
	s.TimeRangeValue = timeRangeValue
}

func (s *SAutomationModel) GetTriggerRule() string {
	return s.TriggerRule
}

func (s *SAutomationModel) SetTriggerRule(triggerRule string) {
	s.TriggerRule = triggerRule
}

func (s *SAutomationModel) GetExecuteRule() string {
	return s.ExecuteRule
}

func (s *SAutomationModel) SetExecuteRule(executeRule string) {
	s.ExecuteRule = executeRule
}

func (s *SAutomationModel) GetEnable() bool {
	return s.Enable
}

func (s *SAutomationModel) SetEnable(enable bool) {
	s.Enable = enable
}

func (s *SAutomationModel) GetCreatedAt() *gtime.Time {
	return s.CreatedAt
}

func (s *SAutomationModel) SetCreatedAt(createdAt *gtime.Time) {
	s.CreatedAt = createdAt
}

func (s *SAutomationModel) GetUpdatedAt() *gtime.Time {
	return s.UpdatedAt
}

func (s *SAutomationModel) SetUpdatedAt(updatedAt *gtime.Time) {
	s.UpdatedAt = updatedAt
}

// CRUD 方法
func (s *SAutomationModel) Create(ctx context.Context) error {
	_, err := g.Model(TableAutomation).Ctx(ctx).Insert(s)
	return err
}

func (s *SAutomationModel) GetById(ctx context.Context, id int) error {
	return g.Model(TableAutomation).Ctx(ctx).Where(FieldId, id).Scan(s)
}

func (s *SAutomationModel) Update(ctx context.Context) error {
	_, err := g.Model(TableAutomation).Ctx(ctx).Where(FieldId, s.Id).Update(s)
	return err
}

func (s *SAutomationModel) UpdateFields(ctx context.Context, data g.Map) error {
	_, err := g.Model(TableAutomation).Ctx(ctx).Where(FieldId, s.Id).Update(data)
	return err
}

func (s *SAutomationModel) Delete(ctx context.Context) error {
	_, err := g.Model(TableAutomation).Ctx(ctx).Where(FieldId, s.Id).Delete()
	return err
}

// 静态方法
func DeleteAutomationById(ctx context.Context, id int) error {
	_, err := g.Model(TableAutomation).Ctx(ctx).Where(FieldId, id).Delete()
	return err
}

func GetAllAutomations(ctx context.Context) ([]*SAutomationModel, error) {
	var automations []*SAutomationModel
	err := g.Model(TableAutomation).Ctx(ctx).Scan(&automations)
	return automations, err
}

func GetAutomationsByCondition(ctx context.Context, condition g.Map) ([]*SAutomationModel, error) {
	var automations []*SAutomationModel
	err := g.Model(TableAutomation).Ctx(ctx).Where(condition).Scan(&automations)
	return automations, err
}

func GetAutomationsByTimeRangeType(ctx context.Context, timeRangeType string) ([]*SAutomationModel, error) {
	var automations []*SAutomationModel
	err := g.Model(TableAutomation).Ctx(ctx).Where(FieldTimeRangeType, timeRangeType).Scan(&automations)
	return automations, err
}

func GetEnabledAutomations(ctx context.Context) ([]*SAutomationModel, error) {
	var automations []*SAutomationModel
	err := g.Model(TableAutomation).Ctx(ctx).Where(FieldEnable, true).Scan(&automations)
	return automations, err
}

func CountAutomations(ctx context.Context) (int, error) {
	count, err := g.Model(TableAutomation).Ctx(ctx).Count()
	return count, err
}

func PaginateAutomations(ctx context.Context, page, pageSize int) ([]*SAutomationModel, error) {
	var automations []*SAutomationModel
	err := g.Model(TableAutomation).Ctx(ctx).Page(page, pageSize).Scan(&automations)
	return automations, err
}

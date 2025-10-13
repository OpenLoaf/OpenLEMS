package impl

import (
	"context"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"s_db/s_db_basic"
	"s_db/s_db_model"
)

type sEnergyStorageStrategyServiceImpl struct{}

var (
	energyStorageStrategyServiceInstance s_db_basic.IEnergyStorageStrategyService
	energyStorageStrategyServiceOnce     sync.Once
)

func GetEnergyStorageStrategyService() s_db_basic.IEnergyStorageStrategyService {
	energyStorageStrategyServiceOnce.Do(func() {
		energyStorageStrategyServiceInstance = &sEnergyStorageStrategyServiceImpl{}
	})
	return energyStorageStrategyServiceInstance
}

// 校验输入
func (s *sEnergyStorageStrategyServiceImpl) validateAndNormalizeInput(data map[string]interface{}) error {
	name, _ := data["name"].(string)
	if l := len(strings.TrimSpace(name)); l < 2 || l > 50 {
		return errors.Errorf("name 长度需在2-50之间")
	}
	// priority 1..5
	if v, ok := data["priority"].(int); !ok || v < 1 || v > 5 {
		return errors.Errorf("priority 必须在1-5之间")
	}
	// essDeviceIds 至少一个
	var ids []string
	switch raw := data["essDeviceIds"].(type) {
	case []string:
		ids = raw
	case []interface{}:
		for _, it := range raw {
			if s, ok := it.(string); ok && s != "" {
				ids = append(ids, s)
			}
		}
	}
	if len(ids) == 0 {
		return errors.Errorf("essDeviceIds 至少包含1个设备ID")
	}

	// config 校验（socMinRatio<=socMaxRatio, monthlyChargeDay 1..28）
	cfgStr, _ := data["config"].(string)
	if cfgStr == "" {
		// 也支持 map -> string
		if m, ok := data["config"].(map[string]interface{}); ok {
			b, _ := gjson.Encode(m)
			cfgStr = string(b)
			data["config"] = cfgStr
		}
	}
	if cfgStr != "" {
		var cfg map[string]interface{}
		if err := gjson.DecodeTo(cfgStr, &cfg); err == nil {
			minv, _ := cfg["socMinRatio"].(float64)
			maxv, _ := cfg["socMaxRatio"].(float64)
			if maxv < minv {
				return errors.Errorf("config.socMinRatio 不能大于 socMaxRatio")
			}
			if mv, ok := cfg["monthlyChargeDay"].(float64); ok {
				if mv < 1 || mv > 28 {
					return errors.Errorf("config.monthlyChargeDay 必须在1-28之间")
				}
			}
		}
	}

	// dateRange/timeRange 必须存在（字符串或对象）
	for _, k := range []string{"dateRange", "timeRange"} {
		if _, ok := data[k]; !ok {
			return errors.Errorf("%s 不能为空", k)
		}
		if _, ok := data[k].(string); !ok {
			if m, ok2 := data[k].(map[string]interface{}); ok2 {
				b, _ := gjson.Encode(m)
				data[k] = string(b)
			}
		}
	}

	// 统一 essDeviceIds 为 JSON 字符串
	if _, ok := data["essDeviceIds"].(string); !ok {
		b, _ := gjson.Encode(ids)
		data["essDeviceIds"] = string(b)
	}

	return nil
}

func (s *sEnergyStorageStrategyServiceImpl) CreateEnergyStorageStrategy(ctx context.Context, data map[string]interface{}) (string, error) {
	if err := s.validateAndNormalizeInput(data); err != nil {
		return "", err
	}
	id := uuid.NewString()
	m := &s_db_model.SEnergyStorageStrategyModel{
		Id:           id,
		Name:         strings.TrimSpace(data["name"].(string)),
		Description:  strings.TrimSpace(gconvString(data["description"])),
		Priority:     data["priority"].(int),
		Status:       gconvString(data["status"]),
		IsDefault:    gconvBool(data["isDefault"]),
		DateRange:    gconvString(data["dateRange"]),
		TimeRange:    gconvString(data["timeRange"]),
		Config:       gconvString(data["config"]),
		EssDeviceIds: gconvString(data["essDeviceIds"]),
		CreatedBy:    gconvString(data["createdBy"]),
	}
	if err := m.Create(ctx); err != nil {
		return "", errors.Wrap(err, "创建储能策略失败")
	}
	return id, nil
}

func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageStrategyById(ctx context.Context, id string) (*s_db_model.SEnergyStorageStrategyModel, error) {
	m := &s_db_model.SEnergyStorageStrategyModel{}
	if err := m.GetById(ctx, id); err != nil {
		return nil, err
	}
	if m.Id == "" {
		return nil, errors.Errorf("记录不存在: %s", id)
	}
	return m, nil
}

func (s *sEnergyStorageStrategyServiceImpl) UpdateEnergyStorageStrategy(ctx context.Context, id string, data map[string]interface{}) error {
	if err := s.validateAndNormalizeInput(data); err != nil {
		return err
	}
	m := &s_db_model.SEnergyStorageStrategyModel{Id: id}
	return m.UpdateFields(ctx, g.Map{
		s_db_model.FieldEssName:        strings.TrimSpace(gconvString(data["name"])),
		s_db_model.FieldEssDescription: strings.TrimSpace(gconvString(data["description"])),
		s_db_model.FieldEssPriority:    gconvInt(data["priority"]),
		s_db_model.FieldEssStatus:      gconvString(data["status"]),
		s_db_model.FieldEssIsDefault:   gconvBool(data["isDefault"]),
		s_db_model.FieldEssDateRange:   gconvString(data["dateRange"]),
		s_db_model.FieldEssTimeRange:   gconvString(data["timeRange"]),
		s_db_model.FieldEssConfig:      gconvString(data["config"]),
		s_db_model.FieldEssDeviceIds:   gconvString(data["essDeviceIds"]),
		s_db_model.FieldEssCreatedBy:   gconvString(data["createdBy"]),
	})
}

func (s *sEnergyStorageStrategyServiceImpl) DeleteEnergyStorageStrategy(ctx context.Context, id string) error {
	m := &s_db_model.SEnergyStorageStrategyModel{Id: id}
	return m.Delete(ctx)
}

func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageStrategyPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SEnergyStorageStrategyModel, int, error) {
	model := g.Model(s_db_model.TableEnergyStorageStrategy).Ctx(ctx)
	if v, ok := filters["status"].(string); ok && v != "" && v != "all" {
		model = model.Where(s_db_model.FieldEssStatus, v)
	}
	if v, ok := filters["priority"].(int); ok && v >= 1 && v <= 5 {
		model = model.Where(s_db_model.FieldEssPriority, v)
	}
	if v, ok := filters["keyword"].(string); ok && strings.TrimSpace(v) != "" {
		like := "%" + strings.TrimSpace(v) + "%"
		model = model.WhereOr(model.Builder().WhereLike(s_db_model.FieldEssName, like).WhereLike(s_db_model.FieldEssDescription, like))
	}
	total, err := model.Clone().Count()
	if err != nil {
		return nil, 0, err
	}
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	var list []*s_db_model.SEnergyStorageStrategyModel
	err = model.Page(page, pageSize).OrderDesc(s_db_model.FieldCreatedAt).Scan(&list)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageStrategiesByIds(ctx context.Context, ids []string) ([]*s_db_model.SEnergyStorageStrategyModel, error) {
	if len(ids) == 0 {
		return []*s_db_model.SEnergyStorageStrategyModel{}, nil
	}
	var list []*s_db_model.SEnergyStorageStrategyModel
	if err := g.Model(s_db_model.TableEnergyStorageStrategy).Ctx(ctx).WhereIn(s_db_model.FieldId, ids).Scan(&list); err != nil {
		return nil, err
	}
	return list, nil
}

func (s *sEnergyStorageStrategyServiceImpl) SetEnergyStorageStrategyActive(ctx context.Context, id string, active bool) error {
	status := "inactive"
	if active {
		status = "active"
	}
	_, err := g.Model(s_db_model.TableEnergyStorageStrategy).Ctx(ctx).Where(s_db_model.FieldId, id).Update(g.Map{s_db_model.FieldEssStatus: status})
	return err
}

func (s *sEnergyStorageStrategyServiceImpl) SetEnergyStorageStrategyDefault(ctx context.Context, id string, isDefault bool) error {
	if isDefault {
		// 先清除其他默认
		if _, err := g.Model(s_db_model.TableEnergyStorageStrategy).Ctx(ctx).Where(s_db_model.FieldEssIsDefault, true).Update(g.Map{s_db_model.FieldEssIsDefault: false}); err != nil {
			return err
		}
	}
	_, err := g.Model(s_db_model.TableEnergyStorageStrategy).Ctx(ctx).Where(s_db_model.FieldId, id).Update(g.Map{s_db_model.FieldEssIsDefault: isDefault})
	return err
}

// DetectConflictsByIds 按ID列表检测冲突
func (s *sEnergyStorageStrategyServiceImpl) DetectConflictsByIds(ctx context.Context, ids []string) ([]map[string]interface{}, error) {
	list, err := s.GetEnergyStorageStrategiesByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	return detectConflicts(list)
}

// DetectConflictsForCandidates 候选策略冲突检测
func (s *sEnergyStorageStrategyServiceImpl) DetectConflictsForCandidates(ctx context.Context, candidates []map[string]interface{}) ([]map[string]interface{}, error) {
	// 将候选转为模型列表（仅用于检测，不入库）
	var list []*s_db_model.SEnergyStorageStrategyModel
	for _, c := range candidates {
		// 简化：只转换与时间/设备相关字段
		dr, _ := toJSONString(c["dateRange"])
		tr, _ := toJSONString(c["timeRange"])
		idsJSON, _ := toJSONString(c["essDeviceIds"])
		list = append(list, &s_db_model.SEnergyStorageStrategyModel{
			Id:           gconvString(c["id"]),
			DateRange:    dr,
			TimeRange:    tr,
			EssDeviceIds: idsJSON,
		})
	}
	return detectConflicts(list)
}

// 冲突检测实现：日期交集 + 时间交集 + 设备交集
func detectConflicts(list []*s_db_model.SEnergyStorageStrategyModel) ([]map[string]interface{}, error) {
	// 解析并逐对比较（此处提供占位实现，后续可细化到具体日期列表）
	var result []map[string]interface{}
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if overlapStrategy(list[i], list[j]) {
				result = append(result, map[string]interface{}{
					"strategyId":    list[i].Id,
					"conflictWith":  []string{list[j].Id},
					"conflictDates": []string{},
				})
			}
		}
	}
	return result, nil
}

func overlapStrategy(a, b *s_db_model.SEnergyStorageStrategyModel) bool {
	// 设备交集
	var aIds, bIds []string
	_ = gjson.DecodeTo(a.EssDeviceIds, &aIds)
	_ = gjson.DecodeTo(b.EssDeviceIds, &bIds)
	if !hasIntersection(aIds, bIds) {
		return false
	}
	// 时间范围与日期范围简单占位判断：字符串非空即认为可能重叠
	if strings.TrimSpace(a.DateRange) == "" || strings.TrimSpace(b.DateRange) == "" {
		return false
	}
	if strings.TrimSpace(a.TimeRange) == "" || strings.TrimSpace(b.TimeRange) == "" {
		return false
	}
	return true
}

func hasIntersection(a, b []string) bool {
	m := map[string]struct{}{}
	for _, x := range a {
		m[x] = struct{}{}
	}
	for _, y := range b {
		if _, ok := m[y]; ok {
			return true
		}
	}
	return false
}

// 工具转换
func gconvString(v interface{}) string {
	s, _ := gjson.EncodeString(v)
	if s == "" {
		if vv, ok := v.(string); ok {
			return vv
		}
	}
	if vv, ok := v.(string); ok {
		return vv
	}
	return s
}
func gconvInt(v interface{}) int {
	if i, ok := v.(int); ok {
		return i
	}
	if f, ok := v.(float64); ok {
		return int(f)
	}
	return 0
}
func gconvBool(v interface{}) bool {
	if b, ok := v.(bool); ok {
		return b
	}
	return false
}
func toJSONString(v interface{}) (string, error) {
	if s, ok := v.(string); ok {
		return s, nil
	}
	b, err := gjson.Encode(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

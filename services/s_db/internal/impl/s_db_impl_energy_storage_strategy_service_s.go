package impl

import (
	"context"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"

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

// 验证模型
func (s *sEnergyStorageStrategyServiceImpl) validateModel(ctx context.Context, model *s_db_model.SEnergyStorageModel) error {
	// 基础字段验证（通过标签自动验证）
	if err := g.Validator().Data(model).Run(ctx); err != nil {
		return errors.Wrap(err.FirstError(), "字段验证失败")
	}

	// 复杂验证逻辑（手动验证）
	if model.Config != "" {
		var cfg struct {
			SocMinRatio      float64 `json:"socMinRatio"`
			SocMaxRatio      float64 `json:"socMaxRatio"`
			MonthlyChargeDay int     `json:"monthlyChargeDay"`
		}
		if err := gjson.DecodeTo(model.Config, &cfg); err == nil {
			// 验证 SOC 范围
			if cfg.SocMaxRatio < cfg.SocMinRatio {
				return errors.New("config.socMinRatio 不能大于 socMaxRatio")
			}
			// 验证月度充电日
			if cfg.MonthlyChargeDay < 1 || cfg.MonthlyChargeDay > 28 {
				return errors.New("config.monthlyChargeDay 必须在1-28之间")
			}
		}
	}

	return nil
}

// CreateEnergyStorageStrategy 创建储能策略
func (s *sEnergyStorageStrategyServiceImpl) CreateEnergyStorageStrategy(ctx context.Context, model *s_db_model.SEnergyStorageModel) (int, error) {
	// 验证模型
	if err := s.validateModel(ctx, model); err != nil {
		return 0, err
	}

	// 自动设置 createdBy 为 admin
	model.CreatedBy = "admin"

	// 创建记录（时间戳和ID由数据库自动处理）
	if err := model.Create(ctx); err != nil {
		return 0, errors.Wrap(err, "创建储能策略失败")
	}

	// 获取刚插入的记录ID
	lastIdValue, err := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx).
		Fields(s_db_model.FieldId).
		Order("id DESC").
		Limit(1).
		Value()
	if err != nil {
		return 0, errors.Wrap(err, "获取新创建的策略ID失败")
	}

	return cvt.IntE(lastIdValue)
}

// GetEnergyStorageStrategyById 根据ID获取储能策略
func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageStrategyById(ctx context.Context, id int) (*s_db_model.SEnergyStorageModel, error) {
	model := &s_db_model.SEnergyStorageModel{}
	if err := model.GetById(ctx, id); err != nil {
		return nil, errors.Wrap(err, "获取储能策略失败")
	}
	if model.Id == 0 {
		return nil, errors.Errorf("记录不存在: %d", id)
	}
	return model, nil
}

// UpdateEnergyStorageStrategy 更新储能策略
func (s *sEnergyStorageStrategyServiceImpl) UpdateEnergyStorageStrategy(ctx context.Context, model *s_db_model.SEnergyStorageModel) error {
	// 验证模型
	if err := s.validateModel(ctx, model); err != nil {
		return err
	}

	// 确保 createdBy 不为空
	if model.CreatedBy == "" {
		model.CreatedBy = "admin"
	}

	// 更新记录
	return model.Update(ctx)
}

// DeleteEnergyStorageStrategy 删除储能策略
func (s *sEnergyStorageStrategyServiceImpl) DeleteEnergyStorageStrategy(ctx context.Context, id int) error {
	model := &s_db_model.SEnergyStorageModel{Id: id}
	return model.Delete(ctx)
}

// GetEnergyStorageStrategyPage 分页查询储能策略
func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageStrategyPage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SEnergyStorageModel, int, error) {
	model := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx)

	// 处理过滤条件
	if status, ok := filters[s_db_model.FieldEssStatus].(string); ok && status != "" && status != "all" {
		model = model.Where(s_db_model.FieldEssStatus, status)
	}

	if priority, ok := filters[s_db_model.FieldEssPriority]; ok {
		p := cvt.Int(priority)
		if p >= 1 && p <= 5 {
			model = model.Where(s_db_model.FieldEssPriority, p)
		}
	}

	if keyword, ok := filters["keyword"].(string); ok && strings.TrimSpace(keyword) != "" {
		like := "%" + strings.TrimSpace(keyword) + "%"
		model = model.WhereOr(
			model.Builder().
				WhereLike(s_db_model.FieldName, like).
				WhereOrLike(s_db_model.FieldEssDescription, like),
		)
	}

	// 获取总数
	total, err := model.Clone().Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, "统计总数失败")
	}

	// 分页参数校验
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	// 查询列表
	var list []*s_db_model.SEnergyStorageModel
	err = model.Page(page, pageSize).OrderDesc(s_db_model.FieldCreatedAt).Scan(&list)
	if err != nil {
		return nil, 0, errors.Wrap(err, "查询列表失败")
	}

	return list, total, nil
}

// GetEnergyStorageStrategiesByIds 根据ID列表获取储能策略
func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageStrategiesByIds(ctx context.Context, ids []int) ([]*s_db_model.SEnergyStorageModel, error) {
	if len(ids) == 0 {
		return []*s_db_model.SEnergyStorageModel{}, nil
	}

	var list []*s_db_model.SEnergyStorageModel
	err := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx).
		WhereIn(s_db_model.FieldId, ids).
		Scan(&list)
	if err != nil {
		return nil, errors.Wrap(err, "批量查询策略失败")
	}

	return list, nil
}

// SetEnergyStorageStrategyActive 设置策略激活状态
func (s *sEnergyStorageStrategyServiceImpl) SetEnergyStorageStrategyActive(ctx context.Context, id int, active bool) error {
	status := "inactive"
	if active {
		status = "active"
	}

	_, err := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx).
		Where(s_db_model.FieldId, id).
		Update(g.Map{s_db_model.FieldEssStatus: status})

	return errors.Wrap(err, "设置策略状态失败")
}

// SetEnergyStorageStrategyDefault 设置默认策略
func (s *sEnergyStorageStrategyServiceImpl) SetEnergyStorageStrategyDefault(ctx context.Context, id int, isDefault bool) error {
	if isDefault {
		// 先清除其他默认策略
		if _, err := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx).
			Where(s_db_model.FieldEssIsDefault, true).
			Update(g.Map{s_db_model.FieldEssIsDefault: false}); err != nil {
			return errors.Wrap(err, "清除其他默认策略失败")
		}
	}

	_, err := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx).
		Where(s_db_model.FieldId, id).
		Update(g.Map{s_db_model.FieldEssIsDefault: isDefault})

	return errors.Wrap(err, "设置默认策略失败")
}

// DetectConflictsByIds 按ID列表检测冲突
func (s *sEnergyStorageStrategyServiceImpl) DetectConflictsByIds(ctx context.Context, ids []int) ([]map[string]interface{}, error) {
	list, err := s.GetEnergyStorageStrategiesByIds(ctx, ids)
	if err != nil {
		return nil, err
	}
	return detectConflicts(list)
}

// DetectConflictsForCandidates 候选策略冲突检测
func (s *sEnergyStorageStrategyServiceImpl) DetectConflictsForCandidates(ctx context.Context, candidates []map[string]interface{}) ([]map[string]interface{}, error) {
	// 将候选转为模型列表（仅用于检测，不入库）
	var list []*s_db_model.SEnergyStorageModel
	for _, c := range candidates {
		id := cvt.Int(c["id"])
		dateRange := cvt.String(c["dateRange"])
		timeRange := cvt.String(c["timeRange"])

		// 如果是对象，需要序列化为JSON字符串
		if dr, ok := c["dateRange"].(map[string]interface{}); ok {
			if b, err := gjson.Encode(dr); err == nil {
				dateRange = string(b)
			}
		}
		if tr, ok := c["timeRange"].(map[string]interface{}); ok {
			if b, err := gjson.Encode(tr); err == nil {
				timeRange = string(b)
			}
		}

		list = append(list, &s_db_model.SEnergyStorageModel{
			Id:        id,
			DateRange: dateRange,
			TimeRange: timeRange,
		})
	}
	return detectConflicts(list)
}

// 冲突检测实现：日期交集 + 时间交集
func detectConflicts(list []*s_db_model.SEnergyStorageModel) ([]map[string]interface{}, error) {
	var result []map[string]interface{}

	// 逐对比较
	for i := 0; i < len(list); i++ {
		for j := i + 1; j < len(list); j++ {
			if overlapStrategy(list[i], list[j]) {
				result = append(result, map[string]interface{}{
					"strategyId":    list[i].Id,
					"conflictWith":  []int{list[j].Id},
					"conflictDates": []string{},
				})
			}
		}
	}

	return result, nil
}

// 判断两个策略是否重叠
func overlapStrategy(a, b *s_db_model.SEnergyStorageModel) bool {
	// 时间范围与日期范围简单占位判断：字符串非空即认为可能重叠
	if strings.TrimSpace(a.DateRange) == "" || strings.TrimSpace(b.DateRange) == "" {
		return false
	}
	if strings.TrimSpace(a.TimeRange) == "" || strings.TrimSpace(b.TimeRange) == "" {
		return false
	}
	return true
}

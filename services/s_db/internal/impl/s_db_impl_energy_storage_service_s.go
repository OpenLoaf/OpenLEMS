package impl

import (
	"context"
	"strings"
	"sync"

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
	return nil
}

// CreateEnergyStorageStrategy 创建储能策略
func (s *sEnergyStorageStrategyServiceImpl) CreateEnergyStorage(ctx context.Context, model *s_db_model.SEnergyStorageModel) (int, error) {
	// 验证模型
	if err := s.validateModel(ctx, model); err != nil {
		return 0, errors.Wrap(err, "验证失败")
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

	return lastIdValue.Int(), nil
}

// GetEnergyStorageStrategyById 根据ID获取储能策略
func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageById(ctx context.Context, id int) (*s_db_model.SEnergyStorageModel, error) {
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
func (s *sEnergyStorageStrategyServiceImpl) UpdateEnergyStorage(ctx context.Context, model *s_db_model.SEnergyStorageModel) error {
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
func (s *sEnergyStorageStrategyServiceImpl) DeleteEnergyStorage(ctx context.Context, id int) error {
	model := &s_db_model.SEnergyStorageModel{Id: id}
	return model.Delete(ctx)
}

// GetEnergyStorageStrategyPage 分页查询储能策略
func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStoragePage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SEnergyStorageModel, int, error) {
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
func (s *sEnergyStorageStrategyServiceImpl) GetEnergyStorageByIds(ctx context.Context, ids []int) ([]*s_db_model.SEnergyStorageModel, error) {
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

// GetEnabledEnergyStorages 获取所有启用的储能定时任务
func (s *sEnergyStorageStrategyServiceImpl) GetEnabledEnergyStorages(ctx context.Context) ([]*s_db_model.SEnergyStorageModel, error) {
	var list []*s_db_model.SEnergyStorageModel
	err := g.Model(s_db_model.TableEnergyStorage).Ctx(ctx).
		WhereIn(s_db_model.FieldEssStatus, []string{"Enable", "Enabled"}).
		OrderDesc(s_db_model.FieldEssPriority).
		OrderDesc(s_db_model.FieldCreatedAt).
		Scan(&list)
	if err != nil {
		return nil, errors.Wrap(err, "查询启用的储能定时任务失败")
	}

	return list, nil
}

// SetEnergyStorageStrategyDefault 设置默认策略

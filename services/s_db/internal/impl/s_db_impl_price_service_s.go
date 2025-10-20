package impl

import (
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
)

type sPriceServiceImpl struct{}

var (
	priceServiceInstance s_db_basic.IPriceService
	priceServiceOnce     sync.Once
)

// GetPriceService 获取电价服务实例
func GetPriceService() s_db_basic.IPriceService {
	priceServiceOnce.Do(func() {
		priceServiceInstance = &sPriceServiceImpl{}
	})
	return priceServiceInstance
}

// CreatePrice 创建电价
func (s *sPriceServiceImpl) CreatePrice(ctx context.Context, model *s_db_model.SPriceModel) (int, error) {
	// 验证模型
	if err := s.validateModel(ctx, model); err != nil {
		return 0, err
	}

	// 自动设置 createdBy 为 admin
	model.CreatedBy = "admin"

	// 创建记录（时间戳和ID由数据库自动处理）
	if err := model.Create(ctx); err != nil {
		return 0, errors.Wrap(err, "创建电价记录失败")
	}

	// 获取刚插入的记录ID
	lastIdValue, err := g.Model(s_db_model.TablePrice).Ctx(ctx).
		Fields(s_db_model.FieldId).
		Order("id DESC").
		Limit(1).
		Value()
	if err != nil {
		return 0, errors.Wrap(err, "获取新创建的电价记录ID失败")
	}

	return gconv.Int(lastIdValue), nil
}

// GetPriceById 根据ID获取电价
func (s *sPriceServiceImpl) GetPriceById(ctx context.Context, id int) (*s_db_model.SPriceModel, error) {
	model := &s_db_model.SPriceModel{}
	if err := model.GetById(ctx, id); err != nil {
		return nil, errors.Wrap(err, "获取电价记录失败")
	}
	return model, nil
}

// UpdatePrice 更新电价
func (s *sPriceServiceImpl) UpdatePrice(ctx context.Context, model *s_db_model.SPriceModel) error {
	// 验证模型
	if err := s.validateModel(ctx, model); err != nil {
		return err
	}

	// 更新记录（排除ID和CreatedAt）
	if err := model.Update(ctx); err != nil {
		return errors.Wrap(err, "更新电价记录失败")
	}

	return nil
}

// DeletePrice 删除电价
func (s *sPriceServiceImpl) DeletePrice(ctx context.Context, id int) error {
	model := &s_db_model.SPriceModel{Id: id}
	if err := model.Delete(ctx); err != nil {
		return errors.Wrap(err, "删除电价记录失败")
	}
	return nil
}

// GetPricePage 分页获取电价列表
func (s *sPriceServiceImpl) GetPricePage(ctx context.Context, page, pageSize int, filters map[string]interface{}) ([]*s_db_model.SPriceModel, int, error) {
	model := g.Model(s_db_model.TablePrice).Ctx(ctx)

	// 应用过滤条件
	if filters != nil {
		if description, ok := filters[s_db_model.FieldPriceDescription].(string); ok && description != "" {
			model = model.WhereLike(s_db_model.FieldPriceDescription, "%"+description+"%")
		}
		if priority, ok := filters[s_db_model.FieldPricePriority]; ok {
			p := gconv.Int(priority)
			if p >= 1 && p <= 5 {
				model = model.Where(s_db_model.FieldPricePriority, p)
			}
		}
		if status, ok := filters[s_db_model.FieldPriceStatus].(string); ok && status != "" {
			model = model.Where(s_db_model.FieldPriceStatus, status)
		}
		if remoteId, ok := filters[s_db_model.FieldPriceRemoteId].(string); ok && remoteId != "" {
			model = model.Where(s_db_model.FieldPriceRemoteId, remoteId)
		}
	}

	// 获取总数
	total, err := model.Count()
	if err != nil {
		return nil, 0, errors.Wrap(err, "获取电价记录总数失败")
	}

	// 分页查询
	var records []*s_db_model.SPriceModel
	err = model.Order(s_db_model.FieldPricePriority+" ASC, "+s_db_model.FieldCreatedAt+" DESC").
		Page(page, pageSize).
		Scan(&records)
	if err != nil {
		return nil, 0, errors.Wrap(err, "分页查询电价记录失败")
	}

	return records, total, nil
}

// GetAllPrices 获取所有电价
func (s *sPriceServiceImpl) GetAllPrices(ctx context.Context) ([]*s_db_model.SPriceModel, error) {
	var records []*s_db_model.SPriceModel
	err := g.Model(s_db_model.TablePrice).Ctx(ctx).
		Order(s_db_model.FieldPricePriority + " ASC, " + s_db_model.FieldCreatedAt + " DESC").
		Scan(&records)
	if err != nil {
		return nil, errors.Wrap(err, "获取所有电价记录失败")
	}
	return records, nil
}

// GetEnabledPrices 获取所有启用的电价
func (s *sPriceServiceImpl) GetEnabledPrices(ctx context.Context) ([]*s_db_model.SPriceModel, error) {
	var records []*s_db_model.SPriceModel
	err := g.Model(s_db_model.TablePrice).Ctx(ctx).
		Where(s_db_model.FieldPriceStatus, "Enable").
		Order(s_db_model.FieldPricePriority + " ASC, " + s_db_model.FieldCreatedAt + " DESC").
		Scan(&records)
	if err != nil {
		return nil, errors.Wrap(err, "获取启用的电价记录失败")
	}
	return records, nil
}

// GetPricesByRemoteId 根据远程ID获取电价
func (s *sPriceServiceImpl) GetPricesByRemoteId(ctx context.Context, remoteId string) ([]*s_db_model.SPriceModel, error) {
	var records []*s_db_model.SPriceModel
	err := g.Model(s_db_model.TablePrice).Ctx(ctx).
		Where(s_db_model.FieldPriceRemoteId, remoteId).
		Order(s_db_model.FieldPricePriority + " ASC, " + s_db_model.FieldCreatedAt + " DESC").
		Scan(&records)
	if err != nil {
		return nil, errors.Wrap(err, "根据远程ID获取电价记录失败")
	}
	return records, nil
}

// GetPriceCount 获取电价记录总数
func (s *sPriceServiceImpl) GetPriceCount(ctx context.Context) (int, error) {
	count, err := g.Model(s_db_model.TablePrice).Ctx(ctx).Count()
	if err != nil {
		return 0, errors.Wrap(err, "获取电价记录总数失败")
	}
	return count, nil
}

// validateModel 验证电价模型
func (s *sPriceServiceImpl) validateModel(ctx context.Context, model *s_db_model.SPriceModel) error {
	// 基础字段验证（通过标签自动验证）
	if err := g.Validator().Data(model).Run(ctx); err != nil {
		return errors.Wrap(err.FirstError(), "字段验证失败")
	}

	// 复杂验证逻辑（手动验证）
	if model.PriceSegments != "" {
		// 验证JSON格式
		if !gjson.Valid(model.PriceSegments) {
			return errors.New("price_segments 必须是有效的JSON格式")
		}
	}

	return nil
}

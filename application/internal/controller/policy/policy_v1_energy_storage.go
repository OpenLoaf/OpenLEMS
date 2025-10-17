package policy

import (
	v1 "application/api/policy/v1"
	"context"
	"s_db"
	"s_db/s_db_model"

	"p_energy_storage"

	"github.com/gogf/gf/v2/encoding/gjson"
)

// marshalJSON 将结构体序列化为JSON字符串
func marshalJSON(v interface{}) string {
	if v == nil {
		return ""
	}
	b, _ := gjson.Encode(v)
	return string(b)
}

// CreateEnergyStorageStrategy 创建储能策略
func (c *ControllerV1) CreateEnergyStorageStrategy(ctx context.Context, req *v1.CreateEnergyStorageStrategyReq) (res *v1.CreateEnergyStorageStrategyRes, err error) {
	// 使用插件层验证
	if err := p_energy_storage.ValidateStrategy(req.DateRange, req.TimeRange, req.Config); err != nil {
		return nil, err
	}

	// 构造 Model
	model := &s_db_model.SEnergyStorageModel{
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status.String(),
		DateRange:   marshalJSON(req.DateRange),
		TimeRange:   marshalJSON(req.TimeRange),
		Config:      marshalJSON(req.Config),
	}

	// 调用服务层
	_, err = s_db.GetEnergyStorageStrategyService().CreateEnergyStorage(ctx, model)
	if err != nil {
		return nil, err
	}

	return &v1.CreateEnergyStorageStrategyRes{}, nil
}

// UpdateEnergyStorageStrategy 更新储能策略
func (c *ControllerV1) UpdateEnergyStorageStrategy(ctx context.Context, req *v1.UpdateEnergyStorageStrategyReq) (res *v1.UpdateEnergyStorageStrategyRes, err error) {
	// 使用插件层验证
	if err := p_energy_storage.ValidateStrategy(req.DateRange, req.TimeRange, req.Config); err != nil {
		return nil, err
	}

	// 构造 Model
	model := &s_db_model.SEnergyStorageModel{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status.String(),
		DateRange:   marshalJSON(req.DateRange),
		TimeRange:   marshalJSON(req.TimeRange),
		Config:      marshalJSON(req.Config),
	}

	// 调用服务层
	err = s_db.GetEnergyStorageStrategyService().UpdateEnergyStorage(ctx, model)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateEnergyStorageStrategyRes{}, nil
}

// DeleteEnergyStorageStrategy 删除储能策略
func (c *ControllerV1) DeleteEnergyStorageStrategy(ctx context.Context, req *v1.DeleteEnergyStorageStrategyReq) (res *v1.DeleteEnergyStorageStrategyRes, err error) {
	err = s_db.GetEnergyStorageStrategyService().DeleteEnergyStorage(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteEnergyStorageStrategyRes{}, nil
}

// GetEnergyStorageStrategyList 查询储能策略列表
func (c *ControllerV1) GetEnergyStorageStrategyList(ctx context.Context, req *v1.GetEnergyStorageStrategyListReq) (res *v1.GetEnergyStorageStrategyListRes, err error) {
	// 构造过滤条件
	filters := make(map[string]interface{})

	if req.Status != nil {
		filters[s_db_model.FieldEssStatus] = *req.Status
	}
	if req.Priority != nil {
		filters[s_db_model.FieldEssPriority] = *req.Priority
	}
	if req.Keyword != nil {
		filters["keyword"] = *req.Keyword
	}

	// 调用服务层
	list, total, err := s_db.GetEnergyStorageStrategyService().GetEnergyStoragePage(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		return nil, err
	}

	// 转换为 DTO
	dtoList := make([]*v1.EnergyStorage, 0, len(list))
	for _, m := range list {
		dto := &v1.EnergyStorage{}
		if err := dto.UnmarshalValue(m); err != nil {
			return nil, err
		}
		dtoList = append(dtoList, dto)
	}

	return &v1.GetEnergyStorageStrategyListRes{List: dtoList, Total: total}, nil
}

// GetEnergyStorageStrategyDetail 获取储能策略详情
func (c *ControllerV1) GetEnergyStorageStrategyDetail(ctx context.Context, req *v1.GetEnergyStorageStrategyDetailReq) (res *v1.GetEnergyStorageStrategyDetailRes, err error) {
	m, err := s_db.GetEnergyStorageStrategyService().GetEnergyStorageById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dto := &v1.EnergyStorage{}
	if err := dto.UnmarshalValue(m); err != nil {
		return nil, err
	}

	return (*v1.GetEnergyStorageStrategyDetailRes)(dto), nil
}

// ActivateEnergyStorageStrategy 激活/停用储能策略
func (c *ControllerV1) ActivateEnergyStorageStrategy(ctx context.Context, req *v1.ActivateEnergyStorageStrategyReq) (res *v1.ActivateEnergyStorageStrategyRes, err error) {
	err = s_db.GetEnergyStorageStrategyService().SetEnergyStorageActive(ctx, req.Id, req.Active)
	if err != nil {
		return nil, err
	}
	return &v1.ActivateEnergyStorageStrategyRes{}, nil
}

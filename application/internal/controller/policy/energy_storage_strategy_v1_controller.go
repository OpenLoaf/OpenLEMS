package policy

import (
	v1 "application/api/policy/v1"
	"context"
	"s_db"
	"s_db/s_db_model"

	"p_energy_manage"

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
	if err := p_energy_manage.ValidateStrategy(req.DateRange, req.TimeRange, req.Config); err != nil {
		return nil, err
	}

	// 构造 Model
	model := &s_db_model.SEnergyStorageStrategyModel{
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status.String(),
		IsDefault:   req.IsDefault,
		DateRange:   marshalJSON(req.DateRange),
		TimeRange:   marshalJSON(req.TimeRange),
		Config:      marshalJSON(req.Config),
	}

	// 调用服务层
	_, err = s_db.GetEnergyStorageStrategyService().CreateEnergyStorageStrategy(ctx, model)
	if err != nil {
		return nil, err
	}

	return &v1.CreateEnergyStorageStrategyRes{}, nil
}

// UpdateEnergyStorageStrategy 更新储能策略
func (c *ControllerV1) UpdateEnergyStorageStrategy(ctx context.Context, req *v1.UpdateEnergyStorageStrategyReq) (res *v1.UpdateEnergyStorageStrategyRes, err error) {
	// 使用插件层验证
	if err := p_energy_manage.ValidateStrategy(req.DateRange, req.TimeRange, req.Config); err != nil {
		return nil, err
	}

	// 构造 Model
	model := &s_db_model.SEnergyStorageStrategyModel{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		Priority:    req.Priority,
		Status:      req.Status.String(),
		IsDefault:   req.IsDefault,
		DateRange:   marshalJSON(req.DateRange),
		TimeRange:   marshalJSON(req.TimeRange),
		Config:      marshalJSON(req.Config),
	}

	// 调用服务层
	err = s_db.GetEnergyStorageStrategyService().UpdateEnergyStorageStrategy(ctx, model)
	if err != nil {
		return nil, err
	}

	return &v1.UpdateEnergyStorageStrategyRes{}, nil
}

// DeleteEnergyStorageStrategy 删除储能策略
func (c *ControllerV1) DeleteEnergyStorageStrategy(ctx context.Context, req *v1.DeleteEnergyStorageStrategyReq) (res *v1.DeleteEnergyStorageStrategyRes, err error) {
	err = s_db.GetEnergyStorageStrategyService().DeleteEnergyStorageStrategy(ctx, req.Id)
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
	list, total, err := s_db.GetEnergyStorageStrategyService().GetEnergyStorageStrategyPage(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		return nil, err
	}

	// 转换为 DTO
	dtoList := make([]*v1.EnergyStorageStrategy, 0, len(list))
	for _, m := range list {
		dto := &v1.EnergyStorageStrategy{}
		if err := dto.UnmarshalValue(m); err != nil {
			return nil, err
		}
		dtoList = append(dtoList, dto)
	}

	return &v1.GetEnergyStorageStrategyListRes{List: dtoList, Total: total}, nil
}

// GetEnergyStorageStrategyDetail 获取储能策略详情
func (c *ControllerV1) GetEnergyStorageStrategyDetail(ctx context.Context, req *v1.GetEnergyStorageStrategyDetailReq) (res *v1.GetEnergyStorageStrategyDetailRes, err error) {
	m, err := s_db.GetEnergyStorageStrategyService().GetEnergyStorageStrategyById(ctx, req.Id)
	if err != nil {
		return nil, err
	}

	dto := &v1.EnergyStorageStrategy{}
	if err := dto.UnmarshalValue(m); err != nil {
		return nil, err
	}

	return (*v1.GetEnergyStorageStrategyDetailRes)(dto), nil
}

// DetectEnergyStorageStrategyConflicts 储能策略冲突检测
func (c *ControllerV1) DetectEnergyStorageStrategyConflicts(ctx context.Context, req *v1.DetectEnergyStorageStrategyConflictsReq) (res *v1.DetectEnergyStorageStrategyConflictsRes, err error) {
	var out []map[string]interface{}

	if len(req.StrategyIds) > 0 {
		// 根据ID列表检测冲突
		out, err = s_db.GetEnergyStorageStrategyService().DetectConflictsByIds(ctx, req.StrategyIds)
	} else {
		// 根据候选策略检测冲突
		var candidates []map[string]interface{}
		for _, c := range req.Candidates {
			b, _ := gjson.Encode(c)
			var m map[string]interface{}
			_ = gjson.DecodeTo(b, &m)
			candidates = append(candidates, m)
		}
		out, err = s_db.GetEnergyStorageStrategyService().DetectConflictsForCandidates(ctx, candidates)
	}

	if err != nil {
		return nil, err
	}

	// 转换响应格式
	res = &v1.DetectEnergyStorageStrategyConflictsRes{}
	for _, item := range out {
		sId, ok1 := item["strategyId"].(int)
		if !ok1 {
			// 尝试转换其他类型
			sId = int(item["strategyId"].(float64))
		}

		var cw []int
		if cwRaw, ok := item["conflictWith"]; ok {
			if cwSlice, ok2 := cwRaw.([]int); ok2 {
				cw = cwSlice
			} else if cwSlice2, ok3 := cwRaw.([]interface{}); ok3 {
				for _, v := range cwSlice2 {
					cw = append(cw, int(v.(float64)))
				}
			}
		}

		var cd []string
		if cdRaw, ok := item["conflictDates"]; ok {
			if cdSlice, ok2 := cdRaw.([]string); ok2 {
				cd = cdSlice
			}
		}

		res.Conflicts = append(res.Conflicts, struct {
			StrategyId    int      `json:"strategyId"`
			ConflictWith  []int    `json:"conflictWith"`
			ConflictDates []string `json:"conflictDates"`
		}{
			StrategyId:    sId,
			ConflictWith:  cw,
			ConflictDates: cd,
		})
	}

	return res, nil
}

// ActivateEnergyStorageStrategy 激活/停用储能策略
func (c *ControllerV1) ActivateEnergyStorageStrategy(ctx context.Context, req *v1.ActivateEnergyStorageStrategyReq) (res *v1.ActivateEnergyStorageStrategyRes, err error) {
	err = s_db.GetEnergyStorageStrategyService().SetEnergyStorageStrategyActive(ctx, req.Id, req.Active)
	if err != nil {
		return nil, err
	}
	return &v1.ActivateEnergyStorageStrategyRes{}, nil
}

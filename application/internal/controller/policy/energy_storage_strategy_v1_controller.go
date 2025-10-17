package policy

import (
	v1 "application/api/policy/v1"
	"context"
	"s_db"

	"p_energy_manage"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/util/gconv"
)

type EnergyStorageStrategyControllerV1 struct{}

// CreateEnergyStorageStrategy 创建策略
func (c *EnergyStorageStrategyControllerV1) CreateEnergyStorageStrategy(ctx context.Context, req *v1.CreateEnergyStorageStrategyReq) (res *v1.CreateEnergyStorageStrategyRes, err error) {
	// 使用插件层验证
	if err := p_energy_manage.ValidateStrategy(req.DateRange, req.TimeRange, req.Config, req.EssDeviceIds); err != nil {
		return nil, err
	}

	// 使用 gconv.Map 自动转换结构体为 map
	data := gconv.Map(req)
	// 特殊处理：将枚举转为字符串
	data["status"] = req.Status.String()

	_, e := s_db.GetEnergyStorageStrategyService().CreateEnergyStorageStrategy(ctx, data)
	if e != nil {
		return nil, e
	}
	return &v1.CreateEnergyStorageStrategyRes{}, nil
}

// UpdateEnergyStorageStrategy 更新策略
func (c *EnergyStorageStrategyControllerV1) UpdateEnergyStorageStrategy(ctx context.Context, req *v1.UpdateEnergyStorageStrategyReq) (res *v1.UpdateEnergyStorageStrategyRes, err error) {
	// 使用插件层验证
	if err := p_energy_manage.ValidateStrategy(req.DateRange, req.TimeRange, req.Config, req.EssDeviceIds); err != nil {
		return nil, err
	}

	// 使用 gconv.Map 自动转换结构体为 map
	data := gconv.Map(req)
	// 特殊处理：将枚举转为字符串
	data["status"] = req.Status.String()
	// 移除 id 字段（id 通过单独参数传递）
	delete(data, "id")

	if e := s_db.GetEnergyStorageStrategyService().UpdateEnergyStorageStrategy(ctx, req.Id, data); e != nil {
		return nil, e
	}
	return &v1.UpdateEnergyStorageStrategyRes{}, nil
}

// DeleteEnergyStorageStrategy 删除策略
func (c *EnergyStorageStrategyControllerV1) DeleteEnergyStorageStrategy(ctx context.Context, req *v1.DeleteEnergyStorageStrategyReq) (res *v1.DeleteEnergyStorageStrategyRes, err error) {
	if e := s_db.GetEnergyStorageStrategyService().DeleteEnergyStorageStrategy(ctx, req.Id); e != nil {
		return nil, e
	}
	return &v1.DeleteEnergyStorageStrategyRes{}, nil
}

// GetEnergyStorageStrategyList 查询列表
func (c *EnergyStorageStrategyControllerV1) GetEnergyStorageStrategyList(ctx context.Context, req *v1.GetEnergyStorageStrategyListReq) (res *v1.GetEnergyStorageStrategyListRes, err error) {
	// 使用 gconv.Map 自动转换结构体为 map，然后移除分页参数
	filters := gconv.Map(req)
	delete(filters, "page")
	delete(filters, "pageSize")

	// 移除 nil 值
	for key, value := range filters {
		if value == nil {
			delete(filters, key)
		}
	}

	list, total, e := s_db.GetEnergyStorageStrategyService().GetEnergyStorageStrategyPage(ctx, req.Page, req.PageSize, filters)
	if e != nil {
		return nil, e
	}
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

// GetEnergyStorageStrategyDetail 获取详情
func (c *EnergyStorageStrategyControllerV1) GetEnergyStorageStrategyDetail(ctx context.Context, req *v1.GetEnergyStorageStrategyDetailReq) (res *v1.GetEnergyStorageStrategyDetailRes, err error) {
	m, e := s_db.GetEnergyStorageStrategyService().GetEnergyStorageStrategyById(ctx, req.Id)
	if e != nil {
		return nil, e
	}
	dto := &v1.EnergyStorageStrategy{}
	if err := dto.UnmarshalValue(m); err != nil {
		return nil, err
	}
	return (*v1.GetEnergyStorageStrategyDetailRes)(dto), nil
}

// DetectEnergyStorageStrategyConflicts 冲突检测
func (c *EnergyStorageStrategyControllerV1) DetectEnergyStorageStrategyConflicts(ctx context.Context, req *v1.DetectEnergyStorageStrategyConflictsReq) (res *v1.DetectEnergyStorageStrategyConflictsRes, err error) {
	var out []map[string]interface{}
	var e error
	if len(req.StrategyIds) > 0 {
		out, e = s_db.GetEnergyStorageStrategyService().DetectConflictsByIds(ctx, req.StrategyIds)
	} else {
		// candidates 转 map 传入
		var cs []map[string]interface{}
		for _, c := range req.Candidates {
			b, _ := gjson.Encode(c)
			var m map[string]interface{}
			_ = gjson.DecodeTo(b, &m)
			cs = append(cs, m)
		}
		out, e = s_db.GetEnergyStorageStrategyService().DetectConflictsForCandidates(ctx, cs)
	}
	if e != nil {
		return nil, e
	}
	// 直接透传
	res = &v1.DetectEnergyStorageStrategyConflictsRes{}
	// 简单映射
	for _, item := range out {
		sId, _ := item["strategyId"].(string)
		cw, _ := item["conflictWith"].([]string)
		cd, _ := item["conflictDates"].([]string)
		res.Conflicts = append(res.Conflicts, struct {
			StrategyId    string   `json:"strategyId"`
			ConflictWith  []string `json:"conflictWith"`
			ConflictDates []string `json:"conflictDates"`
		}{StrategyId: sId, ConflictWith: cw, ConflictDates: cd})
	}
	return res, nil
}

// ActivateEnergyStorageStrategy 激活/停用
func (c *EnergyStorageStrategyControllerV1) ActivateEnergyStorageStrategy(ctx context.Context, req *v1.ActivateEnergyStorageStrategyReq) (res *v1.ActivateEnergyStorageStrategyRes, err error) {
	if e := s_db.GetEnergyStorageStrategyService().SetEnergyStorageStrategyActive(ctx, req.Id, req.Active); e != nil {
		return nil, e
	}
	return &v1.ActivateEnergyStorageStrategyRes{}, nil
}

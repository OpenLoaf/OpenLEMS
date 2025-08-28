package alarm

import (
	v1 "application/api/alarm/v1"
	"context"
)

// GetCurrentAlarms 获取当前告警分页列表（占位实现，无业务逻辑）
func (c *ControllerV1) GetCurrentAlarms(ctx context.Context, req *v1.GetCurrentAlarmsReq) (res *v1.GetCurrentAlarmsRes, err error) {
	// 规范化分页参数
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 返回空列表与总数0（先不上业务逻辑）
	return &v1.GetCurrentAlarmsRes{
		Total: 0,
		Items: []v1.CurrentAlarmItem{},
	}, nil
}

package alarm

import (
	v1 "application/api/alarm/v1"
	"context"
	"s_db"
	"strings"
)

// GetHistoryAlarms 历史告警分页查询
func (c *ControllerV1) GetHistoryAlarms(ctx context.Context, req *v1.GetHistoryAlarmsReq) (res *v1.GetHistoryAlarmsRes, err error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 20
	}

	filters := make(map[string]interface{})
	if req.DeviceId != "" {
		filters["device_id"] = req.DeviceId
	}
	if req.Level != "" && req.Level != "ALL" {
		filters["level"] = req.Level
	}
	if req.Point != "" {
		filters["point"] = req.Point
	}
	if req.Title != "" {
		filters["title"] = req.Title
	}
	if req.Date != "" {
		filters["date"] = req.Date
	}

	records, total, err := s_db.GetAlarmService().GetAlarmHistoryPage(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]v1.HistoryAlarmItem, 0, len(records))
	for _, r := range records {
		items = append(items, v1.HistoryAlarmItem{
			Id:        r.Id,
			DeviceId:  r.DeviceId,
			Point:     r.Point,
			Level:     r.Level,
			Title:     r.Title,
			Detail:    r.Detail,
			CreatedAt: r.CreatedAt,
		})
	}
	return &v1.GetHistoryAlarmsRes{Total: total, Items: items}, nil
}

// CreateAlarmIgnore 创建忽略告警
func (c *ControllerV1) CreateAlarmIgnore(ctx context.Context, req *v1.CreateAlarmIgnoreReq) (res *v1.CreateAlarmIgnoreRes, err error) {
	if strings.TrimSpace(req.DeviceId) == "" {
		return &v1.CreateAlarmIgnoreRes{Success: false, Message: "设备ID不能为空"}, nil
	}
	if strings.TrimSpace(req.Point) == "" {
		return &v1.CreateAlarmIgnoreRes{Success: false, Message: "告警点位名称不能为空"}, nil
	}

	svc := s_db.GetAlarmService()
	ignored, err := svc.IsAlarmIgnored(ctx, req.DeviceId, req.Point)
	if err != nil {
		return &v1.CreateAlarmIgnoreRes{Success: false, Message: "检查告警忽略状态失败"}, nil
	}
	if ignored {
		return &v1.CreateAlarmIgnoreRes{Success: false, Message: "该告警点位已被忽略"}, nil
	}
	if err := svc.CreateAlarmIgnore(ctx, req.DeviceId, req.Point); err != nil {
		return &v1.CreateAlarmIgnoreRes{Success: false, Message: "创建忽略告警失败"}, nil
	}
	return &v1.CreateAlarmIgnoreRes{Success: true, Message: "成功创建忽略告警"}, nil
}

// DeleteAlarmIgnore 删除忽略告警
func (c *ControllerV1) DeleteAlarmIgnore(ctx context.Context, req *v1.DeleteAlarmIgnoreReq) (res *v1.DeleteAlarmIgnoreRes, err error) {
	if strings.TrimSpace(req.DeviceId) == "" {
		return &v1.DeleteAlarmIgnoreRes{Success: false, Message: "设备ID不能为空"}, nil
	}
	if strings.TrimSpace(req.Point) == "" {
		return &v1.DeleteAlarmIgnoreRes{Success: false, Message: "告警点位名称不能为空"}, nil
	}

	svc := s_db.GetAlarmService()
	ignored, err := svc.IsAlarmIgnored(ctx, req.DeviceId, req.Point)
	if err != nil {
		return &v1.DeleteAlarmIgnoreRes{Success: false, Message: "检查告警忽略状态失败"}, nil
	}
	if !ignored {
		return &v1.DeleteAlarmIgnoreRes{Success: false, Message: "该告警点位未被忽略"}, nil
	}
	if err := svc.DeleteAlarmIgnoreByDeviceIdAndPoint(ctx, req.DeviceId, req.Point); err != nil {
		return &v1.DeleteAlarmIgnoreRes{Success: false, Message: "删除忽略告警失败"}, nil
	}
	return &v1.DeleteAlarmIgnoreRes{Success: true, Message: "成功删除忽略告警"}, nil
}

// GetAlarmIgnore 忽略告警分页查询
func (c *ControllerV1) GetAlarmIgnore(ctx context.Context, req *v1.GetAlarmIgnoreReq) (res *v1.GetAlarmIgnoreRes, err error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 || req.PageSize > 100 {
		req.PageSize = 20
	}

	filters := make(map[string]interface{})
	if req.DeviceId != "" {
		filters["device_id"] = req.DeviceId
	}
	if req.Point != "" {
		filters["point"] = req.Point
	}
	if req.Date != "" {
		filters["date"] = req.Date
	}

	records, total, err := s_db.GetAlarmService().GetAlarmIgnorePage(ctx, req.Page, req.PageSize, filters)
	if err != nil {
		return nil, err
	}

	items := make([]v1.AlarmIgnoreItem, 0, len(records))
	for _, r := range records {
		items = append(items, v1.AlarmIgnoreItem{Id: r.Id, DeviceId: r.DeviceId, Point: r.Point, CreatedAt: r.CreatedAt})
	}
	return &v1.GetAlarmIgnoreRes{Total: total, Items: items}, nil
}

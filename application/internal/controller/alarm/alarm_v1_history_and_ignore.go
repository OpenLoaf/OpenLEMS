package alarm

import (
	v1 "application/api/alarm/v1"
	"common"
	"common/c_base"
	"context"
	"s_db"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
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
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	items := make([]v1.HistoryAlarmItem, 0, len(records))
	for _, r := range records {
		items = append(items, v1.HistoryAlarmItem{
			Id:               r.Id,
			DeviceId:         r.DeviceId,
			DeviceName:       common.GetDeviceManager().GetDeviceNameById(r.DeviceId),
			SourceDeviceId:   r.SourceDeviceId,
			SourceDeviceName: common.GetDeviceManager().GetDeviceNameById(r.SourceDeviceId),
			Point:            r.Point,
			Level:            r.Level,
			Title:            r.PointName,
			Detail:           r.Detail,
			TriggerAt:        r.TriggerAt,
			ClearAt:          r.ClearAt,
		})
	}
	return &v1.GetHistoryAlarmsRes{Total: total, Items: items}, nil
}

// CreateAlarmIgnore 创建忽略告警
func (c *ControllerV1) CreateAlarmIgnore(ctx context.Context, req *v1.CreateAlarmIgnoreReq) (res *v1.CreateAlarmIgnoreRes, err error) {
	if strings.TrimSpace(req.DeviceId) == "" || strings.TrimSpace(req.SourceDeviceId) == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	if strings.TrimSpace(req.Point) == "" || strings.TrimSpace(req.PointName) == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}

	svc := s_db.GetAlarmService()
	ignored, err := svc.IsAlarmIgnored(ctx, req.DeviceId, req.SourceDeviceId, req.Point)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	if ignored {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed)
	}
	if err := svc.CreateAlarmIgnore(ctx, req.DeviceId, req.SourceDeviceId, req.Point, req.PointName); err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	common.GetDeviceManager().IteratorParentDevicesById(req.DeviceId, func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if device == nil {
			return true
		}
		device.IgnoreClearAlarm(req.SourceDeviceId, req.Point)
		return true
	})

	return &v1.CreateAlarmIgnoreRes{}, nil
}

// DeleteAlarmIgnore 删除忽略告警
func (c *ControllerV1) DeleteAlarmIgnore(ctx context.Context, req *v1.DeleteAlarmIgnoreReq) (res *v1.DeleteAlarmIgnoreRes, err error) {
	if strings.TrimSpace(req.DeviceId) == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}
	if strings.TrimSpace(req.Point) == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter)
	}

	svc := s_db.GetAlarmService()
	if err := svc.DeleteAlarmIgnoreByDeviceIdAndPoint(ctx, req.DeviceId, req.Point); err != nil {
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}
	return &v1.DeleteAlarmIgnoreRes{}, nil
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
		return nil, gerror.NewCode(gcode.CodeInternalError)
	}

	items := make([]v1.AlarmIgnoreItem, 0, len(records))
	for _, r := range records {
		items = append(items, v1.AlarmIgnoreItem{
			Id:               r.Id,
			DeviceId:         r.DeviceId,
			DeviceName:       common.GetDeviceManager().GetDeviceNameById(r.DeviceId),
			SourceDeviceId:   r.SourceDeviceId,
			SourceDeviceName: common.GetDeviceManager().GetDeviceNameById(r.SourceDeviceId),
			Point:            r.Point,
			PointName:        r.PointName,
			CreatedAt:        r.CreatedAt,
		})
	}
	return &v1.GetAlarmIgnoreRes{
		Total: total,
		Items: items,
	}, nil
}

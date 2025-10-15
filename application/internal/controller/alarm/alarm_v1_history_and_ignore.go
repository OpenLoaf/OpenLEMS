package alarm

import (
	v1 "application/api/alarm/v1"
	"common"
	"common/c_base"
	"common/c_log"
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
	if req.Level != "" && req.Level != "All" {
		filters["level"] = req.Level
	}
	if req.Point != "" {
		filters["point"] = req.Point
	}
	if req.PointName != "" {
		filters["title"] = req.PointName
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
			PointName:        r.PointName,
			Detail:           r.Detail,
			TriggerAt:        r.TriggerAt,
			ClearAt:          r.ClearAt,
		})
	}
	return &v1.GetHistoryAlarmsRes{Total: total, Items: items}, nil
}

// ClearAlarmHistory 清除告警历史
func (c *ControllerV1) ClearAlarmHistory(ctx context.Context, req *v1.ClearAlarmHistoryReq) (res *v1.ClearAlarmHistoryRes, err error) {
	// 若 deviceId 为空，则清除全部历史；否则仅清除该设备的历史
	// 可选 level：若提供则按级别过滤清除（当前 s_db 暂无按级别清除接口，先实现设备/全量清除）

	// 统计清除数量用于业务日志（通过计数接口获取前置总数）
	beforeCount := s_db.GetAlarmService().GetAlarmHistoryCount(ctx, req.DeviceId)

	if strings.TrimSpace(req.DeviceId) == "" {
		if err := s_db.GetAlarmService().ClearAllAlarmHistory(ctx); err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError)
		}
		c_log.BizInfof(ctx, "清除所有设备的告警历史完成，受影响记录数(约): %d", beforeCount)
	} else {
		if err := s_db.GetAlarmService().DeleteAlarmHistoryByDeviceId(ctx, req.DeviceId); err != nil {
			return nil, gerror.NewCode(gcode.CodeInternalError)
		}
		c_log.BizInfof(ctx, "清除设备[%s]的告警历史完成，受影响记录数(约): %d", req.DeviceId, beforeCount)
	}

	return &v1.ClearAlarmHistoryRes{}, nil
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

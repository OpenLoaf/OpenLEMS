package alarm

import (
	v1 "application/api/alarm/v1"
	"common"
	"common/c_base"
	"context"
	"fmt"
	"s_db"
	"strings"
)

// GetCurrentAlarms 获取当前告警分页列表（手动分页，迭代时过滤与分页）
func (c *ControllerV1) GetCurrentAlarms(ctx context.Context, req *v1.GetCurrentAlarmsReq) (res *v1.GetCurrentAlarmsRes, err error) {
	// 1) 分页参数规范化
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}
	start := (page - 1) * pageSize
	end := start + pageSize

	// 2) 迭代设备-告警，边过滤边做分页窗口
	items := make([]*v1.CurrentAlarmItem, 0, pageSize)
	matchIndex := 0 // 记录过滤后命中的全局索引，用于 Total 与分页

	common.GetDeviceManager().IteratorAllDevices(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if device == nil {
			return true
		}
		alarmList := device.GetAlarmList()
		if len(alarmList) == 0 {
			return true
		}
		for _, alarm := range alarmList {
			// 过滤条件
			if req.DeviceId != "" && alarm.DeviceId != req.DeviceId {
				continue
			}
			if req.Level != "" && req.Level != "ALL" && !strings.EqualFold(alarm.Level.String(), req.Level) {
				continue
			}
			if req.Point != "" && (alarm.Meta == nil || alarm.Meta.Name != req.Point) {
				continue
			}

			// 命中计数（用于 Total 与分页窗口计算）
			matchIndex++
			if matchIndex <= start {
				continue
			}
			if matchIndex > end {
				continue
			}

			// 收集当前页数据
			items = append(items, &v1.CurrentAlarmItem{
				DeviceId:         config.Id,
				DeviceName:       config.Name,
				SourceDeviceId:   alarm.DeviceId,
				SourceDeviceName: config.Name,
				Point:            alarm.Meta.Name,
				PointName:        alarm.Meta.Cn,
				Level:            alarm.Level.String(),
				Title:            alarm.Meta.Cn,
				Detail:           fmt.Sprintf("[%s]触发！值为: %v", alarm.Meta.Cn, alarm.Value),
				CreatedAt:        alarm.HappenTime,
			})
		}
		return true
	})

	// 3) 返回分页数据（Total 为命中条目总数）
	return &v1.GetCurrentAlarmsRes{
		Total:        matchIndex,
		HistoryTotal: s_db.GetAlarmService().GetAlarmHistoryCount(ctx),
		IgnoreTotal:  s_db.GetAlarmService().GetAlarmIgnoreCount(ctx),
		Items:        items,
	}, nil
}

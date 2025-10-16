package alarm

import (
	v1 "application/api/alarm/v1"
	"common"
	"common/c_base"
	"context"
	"fmt"
	"s_db"
	"sort"
	"strings"
)

// GetCurrentAlarms 获取当前告警分页列表（先收集所有告警，排序后分页）
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

	// 2) 收集所有符合条件的告警
	allItems := make([]*v1.CurrentAlarmItem, 0)

	common.GetDeviceManager().IteratorChildDevicesById(req.DeviceId, func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if device == nil {
			return true
		}
		alarmList := device.GetAlarmList()
		if len(alarmList) == 0 {
			return true
		}
		for _, alarm := range alarmList {
			if req.Level != "" && req.Level != "All" && !strings.EqualFold(alarm.GetLevel().String(), req.Level) {
				continue
			}
			if req.Point != "" && (alarm.IPoint == nil || alarm.IPoint.GetKey() != req.Point) {
				continue
			}

			var detail string
			pointName := alarm.IPoint.GetName()
			pointValue := alarm.GetValue()

			// 尝试获取值的解释
			if explain, err := c_base.ExplainPointValue(alarm.IPoint, pointValue); err == nil && explain != "" {
				detail = fmt.Sprintf("[%s]触发！值为: %s", pointName, explain)
			} else {
				detail = fmt.Sprintf("[%s]触发！值为: %v", pointName, pointValue)
			}

			createAt := alarm.GetHappenTime()
			// 收集所有符合条件的告警
			allItems = append(allItems, &v1.CurrentAlarmItem{
				DeviceId:         config.Id,
				DeviceName:       config.Name,
				SourceDeviceId:   alarm.GetDeviceId(),
				SourceDeviceName: config.Name,
				Point:            alarm.IPoint.GetKey(),
				PointName:        alarm.IPoint.GetName(),
				Level:            alarm.GetLevel().String(),
				Detail:           detail,
				CreatedAt:        &createAt,
			})
		}
		return true
	})

	// 3) 对告警列表进行排序
	// 排序规则：先按时间倒序（最新的在前），时间相同的按点位名称正序（字母顺序）
	sort.Slice(allItems, func(i, j int) bool {
		// 先比较时间（倒序）
		if allItems[i].CreatedAt != nil && allItems[j].CreatedAt != nil {
			if !allItems[i].CreatedAt.Equal(*allItems[j].CreatedAt) {
				return allItems[i].CreatedAt.After(*allItems[j].CreatedAt)
			}
		} else if allItems[i].CreatedAt != nil {
			return true // i有时间，j没有时间，i排在前面
		} else if allItems[j].CreatedAt != nil {
			return false // j有时间，i没有时间，j排在前面
		}

		// 时间相同或都为空时，按点位名称正序排序
		return allItems[i].PointName < allItems[j].PointName
	})

	// 4) 计算分页
	total := len(allItems)
	start := (page - 1) * pageSize
	end := start + pageSize

	// 确保不超出范围
	if start >= total {
		start = total
	}
	if end > total {
		end = total
	}

	// 5) 提取当前页的数据
	var items []*v1.CurrentAlarmItem
	if start < end {
		items = allItems[start:end]
	} else {
		items = make([]*v1.CurrentAlarmItem, 0)
	}

	// 6) 返回分页数据
	return &v1.GetCurrentAlarmsRes{
		Total:        total,
		HistoryTotal: s_db.GetAlarmService().GetAlarmHistoryCount(ctx, req.DeviceId),
		IgnoreTotal:  s_db.GetAlarmService().GetAlarmIgnoreCount(ctx, req.DeviceId),
		Items:        items,
	}, nil
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package alarm

import (
	"context"

	v1 "application/api/alarm/v1"
)

type IAlarmV1 interface {
	// 当前告警分页查询
	GetCurrentAlarms(ctx context.Context, req *v1.GetCurrentAlarmsReq) (res *v1.GetCurrentAlarmsRes, err error)

	// 历史告警分页查询
	GetHistoryAlarms(ctx context.Context, req *v1.GetHistoryAlarmsReq) (res *v1.GetHistoryAlarmsRes, err error)

	// 创建忽略告警
	CreateAlarmIgnore(ctx context.Context, req *v1.CreateAlarmIgnoreReq) (res *v1.CreateAlarmIgnoreRes, err error)

	// 删除忽略告警
	DeleteAlarmIgnore(ctx context.Context, req *v1.DeleteAlarmIgnoreReq) (res *v1.DeleteAlarmIgnoreRes, err error)

	// 忽略告警分页查询
	GetAlarmIgnore(ctx context.Context, req *v1.GetAlarmIgnoreReq) (res *v1.GetAlarmIgnoreRes, err error)
}

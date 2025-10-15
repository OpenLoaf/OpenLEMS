// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package alarm

import (
	"context"

	v1 "application/api/alarm/v1"
)

type IAlarmV1 interface {
	GetCurrentAlarms(ctx context.Context, req *v1.GetCurrentAlarmsReq) (res *v1.GetCurrentAlarmsRes, err error)
	GetHistoryAlarms(ctx context.Context, req *v1.GetHistoryAlarmsReq) (res *v1.GetHistoryAlarmsRes, err error)
	CreateAlarmIgnore(ctx context.Context, req *v1.CreateAlarmIgnoreReq) (res *v1.CreateAlarmIgnoreRes, err error)
	DeleteAlarmIgnore(ctx context.Context, req *v1.DeleteAlarmIgnoreReq) (res *v1.DeleteAlarmIgnoreRes, err error)
	GetAlarmIgnore(ctx context.Context, req *v1.GetAlarmIgnoreReq) (res *v1.GetAlarmIgnoreRes, err error)
	ClearAlarm(ctx context.Context, req *v1.ClearAlarmReq) (res *v1.ClearAlarmRes, err error)
}

// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package system

import (
	"context"

	v1 "application/api/system/v1"
)

type ISystemV1 interface {
	GetSummary(ctx context.Context, req *v1.GetSummaryReq) (res *v1.GetSummaryRes, err error)
	GetNetworkTraffic(ctx context.Context, req *v1.GetNetworkTrafficReq) (res *v1.GetNetworkTrafficRes, err error)
	GetTimeInfo(ctx context.Context, req *v1.GetTimeInfoReq) (res *v1.GetTimeInfoRes, err error)
	GetSystemInfo(ctx context.Context, req *v1.GetSystemInfoReq) (res *v1.GetSystemInfoRes, err error)
	GetNow(ctx context.Context, req *v1.GetNowReq) (res *v1.GetNowRes, err error)
	UpdateHostname(ctx context.Context, req *v1.UpdateHostnameReq) (res *v1.UpdateHostnameRes, err error)
	UpdateSystemTime(ctx context.Context, req *v1.UpdateSystemTimeReq) (res *v1.UpdateSystemTimeRes, err error)
	RebootApply(ctx context.Context, req *v1.RebootApplyReq) (res *v1.RebootApplyRes, err error)
	RebootExecute(ctx context.Context, req *v1.RebootExecuteReq) (res *v1.RebootExecuteRes, err error)
	GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *v1.GetSettingRes, err error)
	UpdateStorageTime(ctx context.Context, req *v1.UpdateStorageTimeReq) (res *v1.UpdateStorageTimeRes, err error)
}

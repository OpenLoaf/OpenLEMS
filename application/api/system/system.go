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
}

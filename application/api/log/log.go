// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package log

import (
	"context"

	"application/api/log/v1"
)

type ILogV1 interface {
	DeleteBizLog(ctx context.Context, req *v1.DeleteBizLogReq) (res *v1.DeleteBizLogRes, err error)
	GetBizLog(ctx context.Context, req *v1.GetBizLogReq) (res *v1.GetBizLogRes, err error)
	QueryBizLogInfo(ctx context.Context, req *v1.QueryBizLogInfoReq) (res *v1.QueryBizLogInfoRes, err error)
	QueryBizLogStatistics(ctx context.Context, req *v1.QueryBizLogStatisticsReq) (res *v1.QueryBizLogStatisticsRes, err error)
}

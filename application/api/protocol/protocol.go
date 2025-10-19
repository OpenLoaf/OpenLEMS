// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package protocol

import (
	"context"

	v1 "application/api/protocol/v1"
)

type IProtocolV1 interface {
	CreateProtocol(ctx context.Context, req *v1.CreateProtocolReq) (res *v1.CreateProtocolRes, err error)
	DeleteProtocol(ctx context.Context, req *v1.DeleteProtocolReq) (res *v1.DeleteProtocolRes, err error)
	GetProtocolList(ctx context.Context, req *v1.GetProtocolListReq) (res *v1.GetProtocolListRes, err error)
	UpdateProtocol(ctx context.Context, req *v1.UpdateProtocolReq) (res *v1.UpdateProtocolRes, err error)
	PostProtocolMonitorOverview(ctx context.Context, req *v1.PostProtocolMonitorOverviewReq) (res *v1.PostProtocolMonitorOverviewRes, err error)
	PostProtocolMonitorTrend(ctx context.Context, req *v1.PostProtocolMonitorTrendReq) (res *v1.PostProtocolMonitorTrendRes, err error)
	PostProtocolMonitorDistribution(ctx context.Context, req *v1.PostProtocolMonitorDistributionReq) (res *v1.PostProtocolMonitorDistributionRes, err error)
	PostProtocolMonitorMetrics(ctx context.Context, req *v1.PostProtocolMonitorMetricsReq) (res *v1.PostProtocolMonitorMetricsRes, err error)
}

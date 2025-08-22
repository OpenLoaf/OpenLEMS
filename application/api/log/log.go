// =================================================================================
// Code generated and maintained similar to GoFrame CLI style. DO NOT EDIT casually.
// =================================================================================

package log

import (
	"context"

	v1 "application/api/log/v1"
)

type ILogV1 interface {
	GetBizLog(ctx context.Context, req *v1.GetBizLogReq) (res *v1.GetBizLogRes, err error)
}

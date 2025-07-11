// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package driver

import (
	"context"

	"application/api/driver/v1"
)

type IDriverV1 interface {
	GetDriverList(ctx context.Context, req *v1.GetDriverListReq) (res *v1.GetDriverListRes, err error)
}

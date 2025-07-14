// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package protocol

import (
	"context"

	"application/api/protocol/v1"
)

type IProtocolV1 interface {
	GetProtocolList(ctx context.Context, req *v1.GetProtocolListReq) (res *v1.GetProtocolListRes, err error)
}

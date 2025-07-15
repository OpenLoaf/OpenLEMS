// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package protocol

import (
	"context"

	"application/api/protocol/v1"
)

type IProtocolV1 interface {
	CreateProtocol(ctx context.Context, req *v1.CreateProtocolReq) (res *v1.CreateProtocolRes, err error)
	DeleteProtocol(ctx context.Context, req *v1.DeleteProtocolReq) (res *v1.DeleteProtocolRes, err error)
	GetProtocolList(ctx context.Context, req *v1.GetProtocolListReq) (res *v1.GetProtocolListRes, err error)
	UpdateProtocol(ctx context.Context, req *v1.UpdateProtocolReq) (res *v1.UpdateProtocolRes, err error)
}

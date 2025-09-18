// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package network

import (
	"context"

	"application/api/network/v1"
)

type INetworkV1 interface {
	GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error)
	UpdateNetworkInterface(ctx context.Context, req *v1.UpdateNetworkInterfaceReq) (res *v1.UpdateNetworkInterfaceRes, err error)
	SetInterfaceState(ctx context.Context, req *v1.SetInterfaceStateReq) (res *v1.SetInterfaceStateRes, err error)
	Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error)
}

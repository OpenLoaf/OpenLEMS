// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device

import (
	"context"

	"application/api/device/v1"
	"application/api/device/v2"
)

type IDeviceV1 interface {
	GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error)
	GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error)
}

type IDeviceV2 interface {
	CreateDevice(ctx context.Context, req *v2.CreateDeviceReq) (res *v2.CreateDeviceRes, err error)
	DeleteDevice(ctx context.Context, req *v2.DeleteDeviceReq) (res *v2.DeleteDeviceRes, err error)
	GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error)
}

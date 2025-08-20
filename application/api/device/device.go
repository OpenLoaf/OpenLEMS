// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device

import (
	"context"

	v1 "application/api/device/v1"
	v2 "application/api/device/v2"
)

type IDeviceV1 interface {
	GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error)
	GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error)
}

type IDeviceV2 interface {
	GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error)
	DisableDevice(ctx context.Context, req *v2.DisableDeviceReq) (res *v2.DisableDeviceRes, err error)
	EnableDevice(ctx context.Context, req *v2.EnableDeviceReq) (res *v2.EnableDeviceRes, err error)
}

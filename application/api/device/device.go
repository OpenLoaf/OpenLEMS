// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device

import (
	"context"

	v1 "application/api/device/v1"
)

type IDeviceV1 interface {
	GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error)
	GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error)
	GetDeviceHistory(ctx context.Context, req *v1.PostDeviceHistoryReq) (res *v1.PostDeviceHistoryRes, err error)
	GetDeviceNameList(ctx context.Context, req *v1.GetDeviceNameListReq) (res *v1.GetDeviceNameListRes, err error)
}

type IDeviceV2 interface {
	GetDeviceTree(ctx context.Context, req *v1.GetDeviceTreeReq) (res *v1.GetDeviceTreeRes, err error)
	GetDeviceTreeById(ctx context.Context, req *v1.GetDeviceTreeByIdReq) (res *v1.GetDeviceTreeByIdRes, err error)
	DisableDevice(ctx context.Context, req *v1.DisableDeviceReq) (res *v1.DisableDeviceRes, err error)
	EnableDevice(ctx context.Context, req *v1.EnableDeviceReq) (res *v1.EnableDeviceRes, err error)
	UpdateDevice(ctx context.Context, req *v1.UpdateDeviceReq) (res *v1.UpdateDeviceRes, err error)
}

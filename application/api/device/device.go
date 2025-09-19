// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device

import (
	"context"

	v1 "application/api/device/v1"
)

type IDeviceV1 interface {
	GetDeviceDetailedList(ctx context.Context, req *v1.GetDeviceDetailedListReq) (res *v1.GetDeviceDetailedListRes, err error)
	GetDeviceTree(ctx context.Context, req *v1.GetDeviceTreeReq) (res *v1.GetDeviceTreeRes, err error)
	DisableDevice(ctx context.Context, req *v1.DisableDeviceReq) (res *v1.DisableDeviceRes, err error)
	EnableDevice(ctx context.Context, req *v1.EnableDeviceReq) (res *v1.EnableDeviceRes, err error)
	GetDeviceTreeById(ctx context.Context, req *v1.GetDeviceTreeByIdReq) (res *v1.GetDeviceTreeByIdRes, err error)
	PostDeviceHistory(ctx context.Context, req *v1.PostDeviceHistoryReq) (res *v1.PostDeviceHistoryRes, err error)
	GetDeviceNameList(ctx context.Context, req *v1.GetDeviceNameListReq) (res *v1.GetDeviceNameListRes, err error)
	GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error)
	GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error)
	GetDeviceTelemetry(ctx context.Context, req *v1.GetDeviceTelemetryReq) (res *v1.GetDeviceTelemetryRes, err error)
	GetDeviceStatus(ctx context.Context, req *v1.GetDeviceStatusReq) (res *v1.GetDeviceStatusRes, err error)
	GetDeviceTelemetryService(ctx context.Context, req *v1.GetDeviceTelemetryServiceReq) (res *v1.GetDeviceTelemetryServiceRes, err error)
	UpdateDevice(ctx context.Context, req *v1.UpdateDeviceReq) (res *v1.UpdateDeviceRes, err error)
	GetVirtualDeviceCache(ctx context.Context, req *v1.GetVirtualDeviceCacheReq) (res *v1.GetVirtualDeviceCacheRes, err error)
}

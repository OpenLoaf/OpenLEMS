// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package device

import (
	"context"

	"application/api/device/v1"
)

type IDeviceV1 interface {
	GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error)
	GetRealDeviceList(ctx context.Context, req *v1.GetRealDeviceListReq) (res *v1.GetRealDeviceListRes, err error)
}

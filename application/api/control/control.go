// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package control

import (
	"context"

	"application/api/control/v1"
)

type IControlV1 interface {
	ControlDevice(ctx context.Context, req *v1.ControlDeviceReq) (res *v1.ControlDeviceRes, err error)
	GetCustomServices(ctx context.Context, req *v1.GetCustomServicesReq) (res *v1.GetCustomServicesRes, err error)
}

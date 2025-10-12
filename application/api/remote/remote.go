// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package remote

import (
	"context"

	v1 "application/api/remote/v1"
)

type IRemoteV1 interface {
	GetMqttStatus(ctx context.Context, req *v1.GetMqttStatusReq) (res *v1.GetMqttStatusRes, err error)
}

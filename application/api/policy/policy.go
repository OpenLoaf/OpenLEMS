// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package policy

import (
	"context"

	"application/api/policy/v1"
)

type IPolicyV1 interface {
	GetPolicyConfigDesc(ctx context.Context, req *v1.GetPolicyConfigDescReq) (res *v1.GetPolicyConfigDescRes, err error)
	UpdatePolicyConfig(ctx context.Context, req *v1.UpdatePolicyConfigReq) (res *v1.UpdatePolicyConfigRes, err error)
}

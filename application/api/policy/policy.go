// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package policy

import (
	"context"

	"application/api/policy/v1"
)

type IPolicyV1 interface {
	GetPolicyConfig(ctx context.Context, req *v1.GetPolicyConfigReq) (res *v1.GetPolicyConfigRes, err error)
}

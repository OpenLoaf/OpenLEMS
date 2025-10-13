// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package policy

import (
	"context"

	v1 "application/api/policy/v1"
)

type IPolicyV1 interface {
	GetPolicyConfig(ctx context.Context, req *v1.GetPolicyConfigReq) (res *v1.GetPolicyConfigRes, err error)

	CreateEnergyStorageStrategy(ctx context.Context, req *v1.CreateEnergyStorageStrategyReq) (res *v1.CreateEnergyStorageStrategyRes, err error)
	UpdateEnergyStorageStrategy(ctx context.Context, req *v1.UpdateEnergyStorageStrategyReq) (res *v1.UpdateEnergyStorageStrategyRes, err error)
	DeleteEnergyStorageStrategy(ctx context.Context, req *v1.DeleteEnergyStorageStrategyReq) (res *v1.DeleteEnergyStorageStrategyRes, err error)
	GetEnergyStorageStrategyList(ctx context.Context, req *v1.GetEnergyStorageStrategyListReq) (res *v1.GetEnergyStorageStrategyListRes, err error)
	GetEnergyStorageStrategyDetail(ctx context.Context, req *v1.GetEnergyStorageStrategyDetailReq) (res *v1.GetEnergyStorageStrategyDetailRes, err error)
	DetectEnergyStorageStrategyConflicts(ctx context.Context, req *v1.DetectEnergyStorageStrategyConflictsReq) (res *v1.DetectEnergyStorageStrategyConflictsRes, err error)
	ActivateEnergyStorageStrategy(ctx context.Context, req *v1.ActivateEnergyStorageStrategyReq) (res *v1.ActivateEnergyStorageStrategyRes, err error)
}

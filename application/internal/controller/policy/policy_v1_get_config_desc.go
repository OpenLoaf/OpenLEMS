package policy

import (
	v1 "application/api/policy/v1"
	"context"
)

func (c *ControllerV1) GetPolicyConfigDesc(ctx context.Context, req *v1.GetPolicyConfigDescReq) (res *v1.GetPolicyConfigDescRes, err error) {
	// TODO: 实现获取策略配置描述的逻辑
	return &v1.GetPolicyConfigDescRes{}, nil
}

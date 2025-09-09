package policy

import (
	v1 "application/api/policy/v1"
	"context"
	"encoding/json"
	"errors"
	"s_db"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) GetPolicyConfigDesc(ctx context.Context, req *v1.GetPolicyConfigDescReq) (res *v1.GetPolicyConfigDescRes, err error) {
	var policyId string

	// 确定要查询的策略ID
	if req.PolicyId == "" {
		policyId = s_db.GetSettingService().GetRootPolicyId(ctx)
		if policyId == "" {
			return &v1.GetPolicyConfigDescRes{
				Config: nil,
			}, nil
		}
	} else {
		policyId = req.PolicyId
	}

	// 获取策略配置字符串
	policyConfigStr := s_db.GetSettingService().GetSettingValueById(ctx, policyId)
	if policyConfigStr == "" {
		return &v1.GetPolicyConfigDescRes{
			Config: nil,
		}, nil
	}

	// 将策略配置字符串解析为JSON对象
	var policyConfig any
	if err := json.Unmarshal([]byte(policyConfigStr), &policyConfig); err != nil {
		g.Log().Errorf(ctx, "解析策略配置JSON失败 - 策略ID: %s, 配置内容: %s, 错误: %+v", policyId, policyConfigStr, err)
		return nil, errors.New("策略配置格式错误，无法解析为JSON")
	}

	g.Log().Infof(ctx, "成功获取策略配置 - 策略ID: %s", policyId)
	return &v1.GetPolicyConfigDescRes{
		Config: policyConfig,
	}, nil
}

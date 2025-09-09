package policy

import (
	v1 "application/api/policy/v1"
	"common/c_base"
	"context"
	"errors"
	"strings"

	"s_db"
	"s_db/s_db_basic"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) UpdatePolicyConfig(ctx context.Context, req *v1.UpdatePolicyConfigReq) (res *v1.UpdatePolicyConfigRes, err error) {
	// 验证策略ID格式
	if !strings.HasPrefix(req.PolicyId, "policy_") {
		g.Log().Errorf(ctx, "策略ID格式错误 - 策略ID必须以 'policy_' 开头: %s", req.PolicyId)
		return nil, errors.New("策略ID必须以 'policy_' 开头")
	}

	// 验证策略配置不能为空
	if strings.TrimSpace(req.PolicyConfig) == "" {
		g.Log().Errorf(ctx, "策略配置不能为空 - 策略ID: %s", req.PolicyId)
		return nil, errors.New("策略配置不能为空")
	}

	// 获取设置服务
	settingService := s_db.GetSettingService()

	// 创建或更新策略配置（使用GetSettingValueByIdWithDefaultValue来自动创建）
	oldValue := settingService.GetSettingValueByIdWithDefaultValue(ctx, req.PolicyId, c_base.ESettingGroupPolicy, req.PolicyConfig)
	if oldValue != "" {
		// 更新
		err = settingService.SetSettingValueById(ctx, req.PolicyId, req.PolicyConfig)
		if err != nil {
			return nil, err
		}
	}

	// 设置当前策略为激活策略
	settingService.GetSettingValueByIdWithDefaultValue(ctx, s_db_basic.SettingActivePolicyIdKey, c_base.ESettingGroupPolicy, req.PolicyId)

	g.Log().Infof(ctx, "成功更新策略配置并设置为激活策略 - 策略ID: %s", req.PolicyId)
	return &v1.UpdatePolicyConfigRes{}, nil
}

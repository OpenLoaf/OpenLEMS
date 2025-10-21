package policy

import (
	v1 "application/api/policy/v1"
	"common/c_enum"
	"common/c_log"
	"context"
	"encoding/json"
	"errors"
	"s_db"
)

func (c *ControllerV1) GetPolicyConfig(ctx context.Context, req *v1.GetPolicyConfigReq) (res *v1.GetPolicyConfigRes, err error) {
	var policyId string
	var settingId string

	// 确定要查询的策略ID
	if req.PolicyId == "" {
		rootPolicyId := s_db.GetSettingService().GetRootPolicyId(ctx)
		if rootPolicyId != nil {
			policyId = *rootPolicyId
		}
		if policyId == "" {
			// 如果rootPolicyId也为空，查询setting表中group为c_enum.ESettingGroupPolicy的所有配置的第一条
			settings, err := s_db.GetSettingService().GetAllSettingsByGroup(ctx, c_enum.ESettingGroupPolicy)
			if err != nil {
				c_log.Errorf(ctx, "获取策略分组设置失败 - 错误: %+v", err)
				return nil, errors.New("获取策略配置失败")
			}

			if len(settings) == 0 {
				return &v1.GetPolicyConfigRes{
					SettingId: "",
					Config:    nil,
				}, nil
			}

			// 返回第一条配置
			settingId = settings[0].Id
			policyConfigStr := settings[0].GetValue()

			if policyConfigStr == "" {
				return &v1.GetPolicyConfigRes{
					SettingId: settingId,
					Config:    nil,
				}, nil
			}

			// 将策略配置字符串解析为JSON对象
			var policyConfig any
			if err := json.Unmarshal([]byte(policyConfigStr), &policyConfig); err != nil {
				c_log.Errorf(ctx, "解析策略配置JSON失败 - 设置ID: %s, 配置内容: %s, 错误: %+v", settingId, policyConfigStr, err)
				return nil, errors.New("策略配置格式错误，无法解析为JSON")
			}

			c_log.Infof(ctx, "成功获取策略配置 - 设置ID: %s", settingId)
			return &v1.GetPolicyConfigRes{
				SettingId: settingId,
				Config:    policyConfig,
			}, nil
		}
		settingId = policyId
	} else {
		policyId = req.PolicyId
		settingId = policyId
	}

	// 获取策略配置字符串
	policyConfigStrPtr := s_db.GetSettingService().GetSettingValueById(ctx, policyId)
	if policyConfigStrPtr == nil {
		return &v1.GetPolicyConfigRes{
			SettingId: settingId,
			Config:    nil,
		}, nil
	}

	policyConfigStr := *policyConfigStrPtr
	// 将策略配置字符串解析为JSON对象
	var policyConfig any
	if err := json.Unmarshal([]byte(policyConfigStr), &policyConfig); err != nil {
		c_log.Errorf(ctx, "解析策略配置JSON失败 - 策略ID: %s, 配置内容: %s, 错误: %+v", policyId, policyConfigStr, err)
		return nil, errors.New("策略配置格式错误，无法解析为JSON")
	}

	c_log.Infof(ctx, "成功获取策略配置 - 策略ID: %s", policyId)
	return &v1.GetPolicyConfigRes{
		SettingId: settingId,
		Config:    policyConfig,
	}, nil
}

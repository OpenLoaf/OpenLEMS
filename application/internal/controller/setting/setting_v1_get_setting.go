package setting

import (
	"context"
	"s_db"

	v1 "application/api/setting/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetSetting 获取公开设置信息
func (c *ControllerV1) GetSetting(ctx context.Context, req *v1.GetSettingReq) (res *v1.GetSettingRes, err error) {
	// 获取公开且启用的设置信息
	settings, err := s_db.GetSettingService().GetPublicEnabledSettings(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取公开设置信息失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取公开设置信息失败")
	}

	// 构建设置列表
	var settingsList []v1.SettingItem
	for _, setting := range settings {
		settingsList = append(settingsList, v1.SettingItem{
			Id:        setting.Id,
			Value:     setting.Value,
			GroupName: setting.Group,
			Remark:    setting.Remark,
		})
	}

	g.Log().Infof(ctx, "成功获取 %d 条公开设置信息", len(settingsList))
	return &v1.GetSettingRes{
		Settings: settingsList,
	}, nil
}

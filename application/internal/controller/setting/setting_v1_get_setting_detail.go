package setting

import (
	"context"
	"errors"
	"s_db"

	v1 "application/api/setting/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetSettingDetail 获取设置详情
func (c *ControllerV1) GetSettingDetail(ctx context.Context, req *v1.GetSettingDetailReq) (res *v1.GetSettingDetailRes, err error) {
	// 参数验证
	if req.Id == "" {
		g.Log().Error(ctx, "设置ID不能为空")
		return nil, errors.New("设置ID不能为空")
	}

	// 调用服务层获取设置详情
	setting, err := s_db.GetSettingService().GetSettingById(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "获取设置详情失败 - 设置ID: %s, 错误: %+v", req.Id, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取设置详情失败")
	}
	if setting == nil {
		return nil, nil
	}

	// 转换为响应格式
	res = &v1.GetSettingDetailRes{
		Id:        setting.Id,
		Value:     setting.Value,
		IsPublic:  setting.IsPublic,
		Enabled:   setting.Enabled,
		Remark:    setting.Remark,
		Sort:      setting.Sort,
		Group:     setting.Group,
		CreatedAt: &setting.CreatedAt.Time,
		UpdatedAt: &setting.UpdatedAt.Time,
	}

	g.Log().Infof(ctx, "成功获取设置详情 - 设置ID: %s", req.Id)
	return res, nil
}

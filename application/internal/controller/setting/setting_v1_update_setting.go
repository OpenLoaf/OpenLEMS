package setting

import (
	"context"
	"errors"
	"s_db"
	"s_db/s_db_model"

	v1 "application/api/setting/v1"

	"common/c_log"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpdateSetting 更新设置信息
func (c *ControllerV1) UpdateSetting(ctx context.Context, req *v1.UpdateSettingReq) (res *v1.UpdateSettingRes, err error) {
	// 参数验证
	if req.Id == "" {
		return nil, errors.New("设置ID不能为空")
	}
	if req.Group == "" {
		return nil, errors.New("分组不能为空")
	}

	// 设置默认值
	isPublic := false
	if req.IsPublic != nil {
		isPublic = *req.IsPublic
	}

	enabled := true
	if req.Enabled != nil {
		enabled = *req.Enabled
	}

	// 尝试获取现有设置
	existingSetting, err := s_db.GetSettingService().GetSettingById(ctx, req.Id)
	if err != nil || existingSetting == nil {
		// 如果设置不存在，创建新设置
		c_log.Infof(ctx, "设置ID %s 不存在，创建新设置", req.Id)

		// 创建新的设置记录
		newSetting := &s_db_model.SSettingModel{}
		newSetting.Id = req.Id
		newSetting.Value = req.Value
		newSetting.IsPublic = isPublic
		newSetting.Enabled = enabled
		newSetting.Remark = req.Remark
		newSetting.Sort = req.Sort
		newSetting.Group = req.Group
		newSetting.CreatedAt = gtime.Now()
		newSetting.UpdatedAt = gtime.Now()

		// 保存到数据库
		err = newSetting.Create(ctx)
		if err != nil {
			c_log.Errorf(ctx, "创建设置失败 - 设置ID: %s, 错误: %+v", req.Id, err)
			return nil, gerror.WrapCode(gcode.CodeInternalError, err, "创建设置失败")
		}

		c_log.Infof(ctx, "成功创建设置 - 设置ID: %s", req.Id)
		return &v1.UpdateSettingRes{}, nil
	}

	// 设置存在，更新现有设置
	c_log.Infof(ctx, "设置ID %s 已存在，更新设置", req.Id)

	// 更新设置字段
	existingSetting.Value = req.Value
	existingSetting.IsPublic = isPublic
	existingSetting.Enabled = enabled
	existingSetting.Remark = req.Remark
	existingSetting.Sort = req.Sort
	existingSetting.Group = req.Group
	existingSetting.UpdatedAt = gtime.Now()

	// 保存更新到数据库
	err = existingSetting.Update(ctx)
	if err != nil {
		c_log.Errorf(ctx, "更新设置失败 - 设置ID: %s, 错误: %+v", req.Id, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新设置失败")
	}

	c_log.Infof(ctx, "成功更新设置 - 设置ID: %s", req.Id)
	return &v1.UpdateSettingRes{}, nil
}

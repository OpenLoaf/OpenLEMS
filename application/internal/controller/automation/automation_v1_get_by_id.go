package automation

import (
	v1 "application/api/automation/v1"
	"application/internal/model/entity"
	"context"
	"errors"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetAutomationById 根据ID获取自动化任务详情
func (c *Controller) GetAutomationById(ctx context.Context, req *v1.GetAutomationByIdReq) (res *v1.GetAutomationByIdRes, err error) {
	g.Log().Infof(ctx, "获取自动化任务详情 - ID: %d", req.Id)

	// 参数验证
	if req.Id <= 0 {
		return nil, errors.New("自动化任务ID必须大于0")
	}

	// 调用服务层获取数据
	automation, err := s_db.GetAutomationService().GetAutomationById(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "获取自动化任务失败 - ID: %d, 错误: %+v", req.Id, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取自动化任务失败")
	}

	// 检查记录是否存在
	if automation == nil {
		g.Log().Warningf(ctx, "自动化任务不存在 - ID: %d", req.Id)
		return nil, gerror.NewCode(gcode.CodeNotFound, "自动化任务不存在")
	}

	// 转换为实体对象
	var automationEntity entity.SAutomation
	if err := automationEntity.UnmarshalValue(automation); err != nil {
		g.Log().Errorf(ctx, "转换自动化实体失败 - ID: %d, 错误: %+v", req.Id, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "转换自动化实体失败")
	}

	res = &v1.GetAutomationByIdRes{
		Automation: &automationEntity,
	}

	g.Log().Infof(ctx, "成功获取自动化任务详情 - ID: %d", req.Id)
	return res, nil
}

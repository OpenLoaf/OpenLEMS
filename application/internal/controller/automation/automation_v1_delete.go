package automation

import (
	v1 "application/api/automation/v1"
	"application/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// DeleteAutomation 删除自动化任务
func (c *Controller) DeleteAutomation(ctx context.Context, req *v1.DeleteAutomationReq) (res *v1.DeleteAutomationRes, err error) {
	g.Log().Infof(ctx, "删除自动化任务 - ID: %d", req.Id)

	// 参数验证
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "自动化任务ID必须大于0")
	}

	// 调用服务层删除自动化任务
	err = service.Automation().DeleteAutomation(ctx, req.Id)
	if err != nil {
		g.Log().Errorf(ctx, "删除自动化任务失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "删除自动化任务失败")
	}

	res = &v1.DeleteAutomationRes{}

	g.Log().Infof(ctx, "成功删除自动化任务 - ID: %d", req.Id)
	return res, nil
}

package automation

import (
	v1 "application/api/automation/v1"
	"application/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// ToggleAutomation 开启/停用自动化任务
func (c *Controller) ToggleAutomation(ctx context.Context, req *v1.ToggleAutomationReq) (res *v1.ToggleAutomationRes, err error) {
	g.Log().Infof(ctx, "切换自动化任务状态 - ID: %d, 启用: %t", req.Id, req.Enable)

	// 参数验证
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "自动化任务ID必须大于0")
	}

	// 构建更新数据
	updateData := map[string]interface{}{
		"enabled": req.Enable,
	}

	// 调用服务层更新自动化任务状态
	err = service.Automation().UpdateAutomation(ctx, req.Id, updateData)
	if err != nil {
		g.Log().Errorf(ctx, "切换自动化任务状态失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "切换自动化任务状态失败")
	}

	res = &v1.ToggleAutomationRes{
		Enabled: req.Enable,
	}

	g.Log().Infof(ctx, "成功切换自动化任务状态 - ID: %d, 启用: %t", req.Id, req.Enable)
	return res, nil
}

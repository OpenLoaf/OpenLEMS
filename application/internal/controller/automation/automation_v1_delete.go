package automation

import (
	v1 "application/api/automation/v1"
	"common/c_log"
	"context"
	"errors"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// DeleteAutomation 删除自动化任务
func (c *Controller) DeleteAutomation(ctx context.Context, req *v1.DeleteAutomationReq) (res *v1.DeleteAutomationRes, err error) {
	c_log.Infof(ctx, "删除自动化任务 - ID: %d", req.Id)

	// 参数验证
	if req.Id <= 0 {
		return nil, errors.New("自动化任务ID必须大于0")
	}

	// 调用服务层删除自动化任务
	err = s_db.GetAutomationService().DeleteAutomation(ctx, req.Id)
	if err != nil {
		c_log.Errorf(ctx, "删除自动化任务失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "删除自动化任务失败")
	}

	res = &v1.DeleteAutomationRes{}

	c_log.Infof(ctx, "成功删除自动化任务 - ID: %d", req.Id)
	return res, nil
}

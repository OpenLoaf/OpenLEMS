package automation

import (
	v1 "application/api/automation/v1"
	"application/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UpdateAutomation 更新自动化任务
func (c *Controller) UpdateAutomation(ctx context.Context, req *v1.UpdateAutomationReq) (res *v1.UpdateAutomationRes, err error) {
	g.Log().Infof(ctx, "更新自动化任务 - ID: %d", req.Id)

	// 参数验证
	if req.Id <= 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "自动化任务ID必须大于0")
	}

	// 构建更新数据
	updateData := make(map[string]interface{})

	if req.StartTime != nil {
		updateData["startTime"] = gtime.New(req.StartTime)
	}
	if req.EndTime != nil {
		updateData["endTime"] = gtime.New(req.EndTime)
	}
	if req.TimeRangeType != "" {
		updateData["timeRangeType"] = req.TimeRangeType
	}
	if req.TimeRangeValue != "" {
		updateData["timeRangeValue"] = req.TimeRangeValue
	}
	if req.TriggerRule != "" {
		updateData["triggerRule"] = req.TriggerRule
	}
	if req.ExecuteRule != "" {
		updateData["executeRule"] = req.ExecuteRule
	}
	if req.Enabled != nil {
		updateData["enabled"] = *req.Enabled
	}

	// 检查是否有需要更新的字段
	if len(updateData) == 0 {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "没有需要更新的字段")
	}

	// 调用服务层更新自动化任务
	err = service.Automation().UpdateAutomation(ctx, req.Id, updateData)
	if err != nil {
		g.Log().Errorf(ctx, "更新自动化任务失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新自动化任务失败")
	}

	res = &v1.UpdateAutomationRes{}

	g.Log().Infof(ctx, "成功更新自动化任务 - ID: %d", req.Id)
	return res, nil
}

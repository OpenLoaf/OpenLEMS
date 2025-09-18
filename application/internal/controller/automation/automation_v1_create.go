package automation

import (
	v1 "application/api/automation/v1"
	"context"
	"encoding/json"
	"errors"
	"s_db"
	"time"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// CreateAutomation 创建自动化任务
func (c *Controller) CreateAutomation(ctx context.Context, req *v1.CreateAutomationReq) (res *v1.CreateAutomationRes, err error) {
	g.Log().Infof(ctx, "创建自动化任务 - 时间范围类型: %s, 启用状态: %t", req.TimeRangeType, req.Enabled)

	// 参数验证
	if req.TriggerConfig == nil {
		return nil, errors.New("触发配置不能为空")
	}
	if req.ExecuteRule == "" {
		return nil, errors.New("执行规则不能为空")
	}

	// 转换时间字段为 *time.Time 类型
	var startTime, endTime *time.Time
	if req.StartTime != nil {
		startTime = &req.StartTime.Time
	}
	if req.EndTime != nil {
		endTime = &req.EndTime.Time
	}

	// 将触发配置序列化为 JSON 字符串
	triggerRuleJson, err := json.Marshal(req.TriggerConfig)
	if err != nil {
		g.Log().Errorf(ctx, "序列化触发配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "序列化触发配置失败")
	}

	// 调用服务层创建自动化任务
	id, err := s_db.GetAutomationService().CreateAutomation(ctx,
		req.Name,
		startTime,
		endTime,
		req.TimeRangeType,
		req.TimeRangeValue,
		string(triggerRuleJson),
		req.ExecuteRule,
		req.TriggerConfig.ExecutionInterval)
	if err != nil {
		g.Log().Errorf(ctx, "创建自动化任务失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "创建自动化任务失败")
	}

	res = &v1.CreateAutomationRes{
		Id: id,
	}

	g.Log().Infof(ctx, "成功创建自动化任务 - ID: %d", id)
	return res, nil
}

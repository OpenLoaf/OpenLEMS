package automation

import (
	v1 "application/api/automation/v1"
	"application/internal/model/entity"
	"common/c_log"
	"context"
	"s_db"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// GetAutomationPage 获取自动化分页列表
func (c *Controller) GetAutomationPage(ctx context.Context, req *v1.GetAutomationPageReq) (res *v1.GetAutomationPageRes, err error) {
	c_log.Infof(ctx, "获取自动化分页列表 - 页码: %d, 每页数量: %d, 设备ID: %s", req.Page, req.PageSize, req.DeviceId)

	// 构建过滤条件
	filters := make(map[string]interface{})
	if req.DeviceId != "" {
		// 这里需要根据实际的业务逻辑来过滤设备相关的自动化任务
		// 可能需要从触发规则或执行规则中解析设备ID
		// 暂时不进行设备过滤，因为自动化任务表中没有直接的设备ID字段
	}
	if req.Enabled != nil {
		filters[s_db_model.FieldEnabled] = *req.Enabled
	}

	// 调用服务层获取数据
	automations, total, err := s_db.GetAutomationService().GetAutomationPage(ctx, req.Page, req.PageSize, req.DeviceId, filters)
	if err != nil {
		c_log.Errorf(ctx, "获取自动化分页列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取自动化列表失败")
	}

	// 转换为实体对象
	var automationList []*entity.SAutomation
	for _, automation := range automations {
		var automationEntity entity.SAutomation
		if err := automationEntity.UnmarshalValue(automation); err != nil {
			c_log.Errorf(ctx, "转换自动化实体失败: %+v", err)
			continue
		}
		automationList = append(automationList, &automationEntity)
	}

	res = &v1.GetAutomationPageRes{
		List:     automationList,
		Total:    total,
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	c_log.Infof(ctx, "成功获取自动化分页列表 - 总数: %d", total)
	return res, nil
}

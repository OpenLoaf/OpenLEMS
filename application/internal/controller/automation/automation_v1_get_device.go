package automation

import (
	v1 "application/api/automation/v1"
	"application/internal/model/entity"
	"context"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetAutomationsByDevice 获取设备自动化列表
func (c *Controller) GetAutomationsByDevice(ctx context.Context, req *v1.GetAutomationsByDeviceReq) (res *v1.GetAutomationsByDeviceRes, err error) {
	g.Log().Infof(ctx, "获取设备自动化列表 - 设备ID: %s", req.DeviceId)

	// 构建过滤条件
	filters := make(map[string]interface{})
	// 注意：这里需要根据实际的业务逻辑来过滤设备相关的自动化任务
	// 可能需要从触发规则或执行规则中解析设备ID
	// 暂时使用空的过滤条件，获取所有自动化任务
	filters = nil

	// 调用服务层获取数据
	automations, err := s_db.GetAutomationService().GetAutomationsByFilters(ctx, req.DeviceId, filters)
	if err != nil {
		g.Log().Errorf(ctx, "获取设备自动化列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取设备自动化列表失败")
	}

	// 转换为实体对象
	var automationList []*entity.SAutomation
	for _, automation := range automations {
		var automationEntity entity.SAutomation
		if err := automationEntity.UnmarshalValue(automation); err != nil {
			g.Log().Errorf(ctx, "转换自动化实体失败: %+v", err)
			continue
		}
		automationList = append(automationList, &automationEntity)
	}

	res = &v1.GetAutomationsByDeviceRes{
		Automations: automationList,
		Count:       len(automationList),
	}

	g.Log().Infof(ctx, "成功获取设备自动化列表 - 设备ID: %s, 数量: %d", req.DeviceId, len(automationList))
	return res, nil
}

package network

import (
	"context"

	v1 "application/api/network/v1"
	"t_network_manager"
	"t_network_manager/public"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// SetInterfaceState 设置网络接口状态（启用/禁用）
func (c *ControllerV1) SetInterfaceState(ctx context.Context, req *v1.SetInterfaceStateReq) (res *v1.SetInterfaceStateRes, err error) {
	// 1. 参数验证
	if err := c.validateSetStateRequest(req); err != nil {
		return nil, err
	}

	// 2. 获取网络管理器实例
	networkManager := t_network_manager.GetInstance()
	if networkManager == nil {
		g.Log().Errorf(ctx, "获取网络管理器实例失败")
		return nil, gerror.NewCode(gcode.CodeInternalError, "网络管理器初始化失败")
	}

	// 3. 设置接口状态
	interfaceID := public.InterfaceID(req.Name)
	if err := networkManager.SetInterfaceState(ctx, interfaceID, req.Up); err != nil {
		g.Log().Errorf(ctx, "设置网络接口 %s 状态失败: %+v", req.Name, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "设置网络接口状态失败")
	}

	// 4. 记录操作日志
	action := "禁用"
	if req.Up {
		action = "启用"
	}
	g.Log().Infof(ctx, "网络接口 %s %s成功", req.Name, action)

	return &v1.SetInterfaceStateRes{}, nil
}

// validateSetStateRequest 验证设置状态请求参数
func (c *ControllerV1) validateSetStateRequest(req *v1.SetInterfaceStateReq) error {
	if req.Name == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "接口名称不能为空")
	}

	// 可以添加更多验证逻辑，比如检查接口名称格式等
	// 这里保持简单，只验证必填字段

	return nil
}

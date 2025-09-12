package network

import (
	"context"
	"strings"
	"t_network_manager/public"

	v1 "application/api/network/v1"
	"t_network_manager"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetNetworkInterfaceList 获取本机网络接口列表
func (c *ControllerV1) GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error) {
	// 1. 获取网络管理器实例
	networkManager := t_network_manager.GetInstance()
	if networkManager == nil {
		g.Log().Errorf(ctx, "获取网络管理器实例失败")
		return nil, gerror.NewCode(gcode.CodeInternalError, "网络管理器初始化失败")
	}

	// 2. 获取所有网络接口
	interfaces, err := networkManager.GetAllInterfaces(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取网络接口列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取网络接口列表失败")
	}

	// 3. 数据转换和过滤
	var result []*public.InterfaceSummary
	for _, iface := range interfaces {
		// 根据参数决定是否包含回环接口
		if !req.IncludeLoopback && c.isLoopbackInterface(iface.Name) {
			continue
		}
		if iface.MAC == "" {
			continue
		}
		result = append(result, iface)
	}

	// 4. 构建响应
	g.Log().Infof(ctx, "成功获取 %d 个网络接口", len(result))
	return &v1.GetNetworkInterfaceListRes{
		Interfaces: result,
	}, nil
}

// isLoopbackInterface 判断是否为回环接口
func (c *ControllerV1) isLoopbackInterface(name string) bool {
	// 常见的回环接口名称
	loopbackNames := []string{
		"lo", "lo0", "Loopback", "Loopback Pseudo-Interface 1",
	}

	name = strings.ToLower(name)
	for _, loopbackName := range loopbackNames {
		if strings.Contains(name, strings.ToLower(loopbackName)) {
			return true
		}
	}

	return false
}

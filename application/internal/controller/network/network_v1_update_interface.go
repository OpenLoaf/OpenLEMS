package network

import (
	"context"
	"fmt"
	"net"

	v1 "application/api/network/v1"
	"t_network_manager"
	"t_network_manager/public"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateNetworkInterface 更新网络接口配置
func (c *ControllerV1) UpdateNetworkInterface(ctx context.Context, req *v1.UpdateNetworkInterfaceReq) (res *v1.UpdateNetworkInterfaceRes, err error) {
	// 1. 参数验证
	if err := c.validateUpdateRequest(req); err != nil {
		return nil, err
	}

	// 2. 获取网络管理器实例
	networkManager := t_network_manager.GetInstance()
	if networkManager == nil {
		g.Log().Errorf(ctx, "获取网络管理器实例失败")
		return nil, gerror.NewCode(gcode.CodeInternalError, "网络管理器初始化失败")
	}

	// 3. 更新接口配置
	interfaceID := public.InterfaceID(req.Name)
	if err := networkManager.UpdateInterfaceConfig(ctx, interfaceID, *req.Config); err != nil {
		g.Log().Errorf(ctx, "更新网络接口 %s 配置失败: %+v", req.Name, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新网络接口配置失败")
	}

	g.Log().Infof(ctx, "网络接口 %s 配置更新成功", req.Name)

	// 4. 清除网络接口缓存，确保下次获取最新数据
	c.clearNetworkInterfaceCache(ctx)

	return &v1.UpdateNetworkInterfaceRes{}, nil
}

// validateUpdateRequest 验证更新请求参数
func (c *ControllerV1) validateUpdateRequest(req *v1.UpdateNetworkInterfaceReq) error {
	if req.Name == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "接口名称不能为空")
	}

	if req.Config == nil {
		return gerror.NewCode(gcode.CodeInvalidParameter, "配置参数不能为空")
	}

	// 如果使用静态IP，验证必要参数
	if req.Config.DHCP != nil && !*req.Config.DHCP {
		if len(req.Config.IPv4) == 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "静态IP模式下必须提供IP地址")
		}

		// 验证IP地址格式
		for _, ipv4Config := range req.Config.IPv4 {
			if ipv4Config == nil {
				continue
			}
			if net.ParseIP(ipv4Config.IPv4) == nil {
				return gerror.NewCode(gcode.CodeInvalidParameter, fmt.Sprintf("无效的IP地址: %s", ipv4Config.IPv4))
			}

			// 验证子网掩码
			if ipv4Config.SubnetMask != "" {
				if net.ParseIP(ipv4Config.SubnetMask) == nil {
					return gerror.NewCode(gcode.CodeInvalidParameter, "无效的子网掩码格式")
				}
			}
		}
	}

	// 验证网关地址
	for _, gateway := range req.Config.Gateway {
		if net.ParseIP(gateway) == nil {
			return gerror.NewCode(gcode.CodeInvalidParameter, fmt.Sprintf("无效的网关地址: %s", gateway))
		}
	}

	// 验证DNS服务器地址
	for _, dns := range req.Config.DNS {
		if net.ParseIP(dns) == nil {
			return gerror.NewCode(gcode.CodeInvalidParameter, fmt.Sprintf("无效的DNS服务器地址: %s", dns))
		}
	}

	return nil
}

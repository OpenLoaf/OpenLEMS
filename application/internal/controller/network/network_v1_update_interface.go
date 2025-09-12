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

	// 3. 构建配置对象
	config := c.buildInterfaceConfig(req)

	// 4. 更新接口配置
	interfaceID := public.InterfaceID(req.Name)
	if err := networkManager.UpdateInterfaceConfig(ctx, interfaceID, config); err != nil {
		g.Log().Errorf(ctx, "更新网络接口 %s 配置失败: %+v", req.Name, err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新网络接口配置失败")
	}

	g.Log().Infof(ctx, "网络接口 %s 配置更新成功", req.Name)
	return &v1.UpdateNetworkInterfaceRes{}, nil
}

// validateUpdateRequest 验证更新请求参数
func (c *ControllerV1) validateUpdateRequest(req *v1.UpdateNetworkInterfaceReq) error {
	if req.Name == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "接口名称不能为空")
	}

	// 如果使用静态IP，验证必要参数
	if !req.DHCP {
		if len(req.IPAddresses) == 0 {
			return gerror.NewCode(gcode.CodeInvalidParameter, "静态IP模式下必须提供IP地址")
		}

		// 验证IP地址格式
		for _, ip := range req.IPAddresses {
			if net.ParseIP(ip) == nil {
				return gerror.NewCode(gcode.CodeInvalidParameter, fmt.Sprintf("无效的IP地址: %s", ip))
			}
		}

		// 验证子网掩码
		if req.Netmask != "" {
			if net.ParseIP(req.Netmask) == nil {
				return gerror.NewCode(gcode.CodeInvalidParameter, "无效的子网掩码格式")
			}
		}

		// 验证网关
		if req.Gateway != "" {
			if net.ParseIP(req.Gateway) == nil {
				return gerror.NewCode(gcode.CodeInvalidParameter, "无效的网关地址")
			}
		}
	}

	// 验证DNS服务器地址
	for _, dns := range req.DNS {
		if net.ParseIP(dns) == nil {
			return gerror.NewCode(gcode.CodeInvalidParameter, fmt.Sprintf("无效的DNS服务器地址: %s", dns))
		}
	}

	return nil
}

// buildInterfaceConfig 构建接口配置对象
func (c *ControllerV1) buildInterfaceConfig(req *v1.UpdateNetworkInterfaceReq) public.InterfaceConfig {
	config := public.InterfaceConfig{
		DHCP: &req.DHCP,
		DNS:  req.DNS,
	}

	// 如果不是DHCP模式，设置静态IP配置
	if !req.DHCP && len(req.IPAddresses) > 0 {
		var ipv4Configs []*public.Ipv4Config
		for _, ip := range req.IPAddresses {
			ipv4Config := &public.Ipv4Config{
				IPv4:       ip,
				SubnetMask: req.Netmask,
			}
			ipv4Configs = append(ipv4Configs, ipv4Config)
		}
		config.IPv4 = ipv4Configs
	}

	// 设置网关
	if req.Gateway != "" {
		config.Gateway = []string{req.Gateway}
	}

	return config
}

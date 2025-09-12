package network

import (
	"context"
	"fmt"
	"net"
	"time"

	v1 "application/api/network/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateNetworkInterface 更新网络接口配置（仅支持Linux）
func (c *ControllerV1) UpdateNetworkInterface(ctx context.Context, req *v1.UpdateNetworkInterfaceReq) (res *v1.UpdateNetworkInterfaceRes, err error) {
	start := time.Now()
	g.Log().Infof(ctx, "开始更新网络接口配置: %s", req.Name)

	// 创建验证器和网络管理器
	validator := NewNetworkValidator()
	networkManager := NewNetworkManager()

	// 构建更新请求
	updateReq := &UpdateInterfaceRequest{
		Name:        req.Name,
		DHCP:        req.DHCP,
		IPAddresses: req.IPAddresses,
		Netmask:     req.Netmask,
		Gateway:     req.Gateway,
		DNS:         req.DNS,
	}

	// 验证请求参数
	if err := validator.ValidateUpdateRequest(updateReq); err != nil {
		g.Log().Errorf(ctx, "参数验证失败: %+v", err)
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, err.Error())
	}

	configMode := "静态配置"
	if req.DHCP {
		configMode = "DHCP配置"
	}
	g.Log().Infof(ctx, "应用网络配置: name=%s mode=%s ipAddresses=%v mask=%s gw=%s dns=%v", req.Name, configMode, req.IPAddresses, req.Netmask, req.Gateway, req.DNS)

	// 更新网络接口配置
	if err := networkManager.UpdateInterface(ctx, updateReq); err != nil {
		g.Log().Errorf(ctx, "更新网络接口配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "更新网络配置失败")
	}

	// 验证配置结果
	if err := c.verifyInterfaceConfiguration(ctx, updateReq); err != nil {
		g.Log().Warningf(ctx, "配置验证失败: %+v", err)
		// 配置验证失败不阻断返回，只记录警告
	}

	g.Log().Infof(ctx, "网络接口配置更新完成: %s，耗时: %v", req.Name, time.Since(start))
	return &v1.UpdateNetworkInterfaceRes{}, nil
}

// verifyInterfaceConfiguration 验证接口配置结果
func (c *ControllerV1) verifyInterfaceConfiguration(ctx context.Context, req *UpdateInterfaceRequest) error {
	// 等待配置生效
	time.Sleep(2 * time.Second)

	// 读取系统当前状态进行比对
	ifi, err := net.InterfaceByName(req.Name)
	if err != nil {
		return err
	}

	addrs, _ := ifi.Addrs()
	var configuredIPs []string

	// 收集当前配置的IP地址
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && ipNet.IP.To4() != nil {
			configuredIPs = append(configuredIPs, ipNet.IP.String())
		}
	}

	if req.DHCP {
		// DHCP模式：检查接口是否启用
		up := ifi.Flags&net.FlagUp != 0
		if !up {
			return fmt.Errorf("接口未启用")
		}

		if len(configuredIPs) == 0 {
			g.Log().Infof(ctx, "DHCP配置已应用（接口已启用，等待网络连接获取IP地址）")
		} else {
			g.Log().Infof(ctx, "DHCP配置已应用（已获取IP地址: %v）", configuredIPs)
		}
	} else {
		// 静态模式：检查所有请求的IP地址是否都已配置
		for _, requestedIP := range req.IPAddresses {
			found := false
			for _, configuredIP := range configuredIPs {
				if configuredIP == requestedIP {
					found = true
					break
				}
			}
			if !found {
				return fmt.Errorf("IP地址 %s 未生效", requestedIP)
			}
		}
		g.Log().Infof(ctx, "静态配置已应用")
	}

	return nil
}

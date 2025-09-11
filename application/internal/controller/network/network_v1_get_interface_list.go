package network

import (
	"context"
	"net"
	"time"

	v1 "application/api/network/v1"
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetNetworkInterfaceList 获取本机网络接口列表
func (c *ControllerV1) GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error) {
	start := time.Now()
	g.Log().Infof(ctx, "开始获取网络接口列表，参数: includeLoopback=%v", req.IncludeLoopback)

	// 获取所有网络接口
	netInterfaces, err := net.Interfaces()
	if err != nil {
		g.Log().Errorf(ctx, "获取网络接口列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取网络接口列表失败")
	}

	g.Log().Infof(ctx, "系统中共有 %d 个网络接口", len(netInterfaces))

	var interfaces []*entity.SNetworkInterface
	for i, netIface := range netInterfaces {
		g.Log().Debugf(ctx, "处理接口 %d: 名称=%s, MAC=%s, 标志=%v, MTU=%d, 索引=%d",
			i, netIface.Name, netIface.HardwareAddr.String(), netIface.Flags, netIface.MTU, netIface.Index)

		// 只返回有MAC地址的接口
		if len(netIface.HardwareAddr) == 0 {
			g.Log().Debugf(ctx, "跳过接口 %s: 没有MAC地址", netIface.Name)
			continue
		}

		// 过滤条件检查
		if !req.IncludeLoopback && netIface.Name == "lo" {
			g.Log().Debugf(ctx, "跳过接口 %s: 回环接口且未包含回环", netIface.Name)
			continue
		}

		// 构建接口信息
		iface := &entity.SNetworkInterface{
			Name:  netIface.Name,
			Type:  getInterfaceType(netIface),
			MAC:   netIface.HardwareAddr.String(),
			Up:    netIface.Flags&net.FlagUp != 0,
			MTU:   netIface.MTU,
			Index: netIface.Index,
		}

		g.Log().Debugf(ctx, "构建接口信息: 名称=%s, 类型=%s, MAC=%s, 启用=%v",
			iface.Name, iface.Type, iface.MAC, iface.Up)

		// 获取IP地址信息
		addrs, err := netIface.Addrs()
		if err != nil {
			g.Log().Warningf(ctx, "获取接口 %s 的IP地址失败: %+v", netIface.Name, err)
		} else {
			g.Log().Debugf(ctx, "接口 %s 有 %d 个IP地址", netIface.Name, len(addrs))

			var ipv4Addrs, ipv6Addrs []string
			for j, addr := range addrs {
				g.Log().Debugf(ctx, "处理地址 %d: %s (类型: %T)", j, addr.String(), addr)

				ipNet, ok := addr.(*net.IPNet)
				if !ok {
					g.Log().Debugf(ctx, "跳过地址 %s: 不是IPNet类型", addr.String())
					continue
				}

				ipStr := ipNet.IP.String()
				if ipNet.IP.To4() != nil {
					ipv4Addrs = append(ipv4Addrs, ipStr)
					if iface.IPv4 == "" {
						iface.IPv4 = ipStr
						iface.Netmask = getNetmaskFromIPNet(ipNet)
						g.Log().Debugf(ctx, "设置IPv4: %s, 子网掩码: %s", ipStr, iface.Netmask)
					}
				} else {
					ipv6Addrs = append(ipv6Addrs, ipStr)
					if iface.IPv6 == "" {
						iface.IPv6 = ipStr
						g.Log().Debugf(ctx, "设置IPv6: %s", ipStr)
					}
				}
			}
			iface.IPAddresses = append(ipv4Addrs, ipv6Addrs...)
			g.Log().Debugf(ctx, "接口 %s 最终IP地址: IPv4=%s, IPv6=%s, 总数=%d",
				netIface.Name, iface.IPv4, iface.IPv6, len(iface.IPAddresses))
		}

		// 判断连接状态
		iface.Connected = iface.Up && len(iface.IPAddresses) > 0
		g.Log().Debugf(ctx, "接口 %s 连接状态: 启用=%v, 有IP=%v, 连接=%v",
			netIface.Name, iface.Up, len(iface.IPAddresses) > 0, iface.Connected)

		interfaces = append(interfaces, iface)
		g.Log().Infof(ctx, "成功添加接口: %s", netIface.Name)
	}

	g.Log().Infof(ctx, "获取网络接口列表完成，共 %d 个接口，耗时: %v", len(interfaces), time.Since(start))

	if len(interfaces) == 0 {
		g.Log().Warningf(ctx, "没有找到符合条件的网络接口，请检查过滤条件")
	}

	// 打印最终结果摘要
	for i, iface := range interfaces {
		g.Log().Infof(ctx, "结果接口 %d: 名称=%s, 类型=%s, MAC=%s, IPv4=%s, 连接=%v",
			i, iface.Name, iface.Type, iface.MAC, iface.IPv4, iface.Connected)
	}

	return &v1.GetNetworkInterfaceListRes{
		Interfaces: interfaces,
		Total:      len(interfaces),
	}, nil
}

// getInterfaceType 获取接口类型
func getInterfaceType(iface net.Interface) string {
	// 根据接口名称和标志判断类型
	if iface.Flags&net.FlagLoopback != 0 {
		return "loopback"
	}
	if iface.Flags&net.FlagPointToPoint != 0 {
		return "point-to-point"
	}
	if iface.Flags&net.FlagBroadcast != 0 {
		return "ethernet"
	}
	return "unknown"
}

// getNetmaskFromIPNet 从IPNet中提取子网掩码
func getNetmaskFromIPNet(ipNet *net.IPNet) string {
	return ipNet.Mask.String()
}

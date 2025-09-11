package network

import (
	"context"
	"net"
	"os"
	"runtime"
	"strings"
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

			var ipv4Addrs []string
			for j, addr := range addrs {
				g.Log().Debugf(ctx, "处理地址 %d: %s (类型: %T)", j, addr.String(), addr)

				ipNet, ok := addr.(*net.IPNet)
				if !ok {
					g.Log().Debugf(ctx, "跳过地址 %s: 不是IPNet类型", addr.String())
					continue
				}

				ipStr := ipNet.IP.String()
				if ipNet.IP.To4() != nil {
					// 只处理IPv4地址
					ipv4Addrs = append(ipv4Addrs, ipStr)
					if iface.Netmask == "" {
						iface.Netmask = getNetmaskFromIPNet(ipNet)
						g.Log().Debugf(ctx, "设置子网掩码: %s", iface.Netmask)
					}
				} else {
					// 忽略IPv6地址
					g.Log().Debugf(ctx, "忽略IPv6地址: %s", ipStr)
				}
			}
			iface.IPAddresses = ipv4Addrs
			g.Log().Debugf(ctx, "接口 %s 最终IP地址: 总数=%d",
				netIface.Name, len(iface.IPAddresses))
		}

		interfaces = append(interfaces, iface)
		g.Log().Infof(ctx, "成功添加接口: %s", netIface.Name)
	}

	// 获取DNS服务器地址列表
	dnsServers := getDNSServers(ctx)

	g.Log().Infof(ctx, "获取网络接口列表完成，共 %d 个接口，耗时: %v", len(interfaces), time.Since(start))

	if len(interfaces) == 0 {
		g.Log().Warningf(ctx, "没有找到符合条件的网络接口，请检查过滤条件")
	}

	// 打印最终结果摘要
	for i, iface := range interfaces {
		g.Log().Infof(ctx, "结果接口 %d: 名称=%s, 类型=%s, MAC=%s, IP地址数=%d",
			i, iface.Name, iface.Type, iface.MAC, len(iface.IPAddresses))
	}

	g.Log().Infof(ctx, "DNS服务器列表: %v", dnsServers)

	return &v1.GetNetworkInterfaceListRes{
		Interfaces: interfaces,
		Total:      len(interfaces),
		DNS:        dnsServers,
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

// getDNSServers 获取系统DNS服务器地址列表（仅Linux系统）
func getDNSServers(ctx context.Context) []string {
	// 只在Linux系统上获取DNS配置
	if runtime.GOOS != "linux" {
		return []string{}
	}

	var dnsServers []string

	// 尝试从 /etc/resolv.conf 文件读取DNS配置
	if resolvConf, err := os.ReadFile("/etc/resolv.conf"); err == nil {
		lines := strings.Split(string(resolvConf), "\n")
		for _, line := range lines {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "nameserver ") {
				dns := strings.TrimSpace(strings.TrimPrefix(line, "nameserver "))
				if dns != "" && net.ParseIP(dns) != nil {
					dnsServers = append(dnsServers, dns)
					g.Log().Debugf(ctx, "从 /etc/resolv.conf 读取到DNS服务器: %s", dns)
				}
			}
		}
	} else {
		g.Log().Warningf(ctx, "无法读取 /etc/resolv.conf 文件: %+v", err)
	}

	// 如果没有找到DNS服务器，添加一些常用的公共DNS
	if len(dnsServers) == 0 {
		g.Log().Infof(ctx, "未找到系统DNS配置，使用默认公共DNS服务器")
		dnsServers = []string{"8.8.8.8", "8.8.4.4", "1.1.1.1"}
	}

	g.Log().Infof(ctx, "获取到 %d 个DNS服务器", len(dnsServers))
	return dnsServers
}

package network

import (
	"context"
	"net"
	"os"
	"os/exec"
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

	var interfaces []*entity.SNetworkInterface
	for _, netIface := range netInterfaces {
		// 只返回有MAC地址的接口
		if len(netIface.HardwareAddr) == 0 {
			continue
		}

		// 过滤条件检查
		if !req.IncludeLoopback && netIface.Name == "lo" {
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

		// 获取网关地址
		iface.Gateway = getGatewayForInterface(ctx, netIface.Name)

		// 获取IP地址信息
		addrs, err := netIface.Addrs()
		if err != nil {
			g.Log().Warningf(ctx, "获取接口 %s 的IP地址失败: %+v", netIface.Name, err)
		} else {
			var ipv4Addrs []string
			for _, addr := range addrs {
				ipNet, ok := addr.(*net.IPNet)
				if !ok {
					continue
				}

				ipStr := ipNet.IP.String()
				if ipNet.IP.To4() != nil {
					// 只处理IPv4地址
					ipv4Addrs = append(ipv4Addrs, ipStr)
					if iface.Netmask == "" {
						iface.Netmask = getNetmaskFromIPNet(ipNet)
					}
				}
			}
			iface.IPAddresses = ipv4Addrs
		}

		interfaces = append(interfaces, iface)
	}

	// 获取DNS服务器地址列表
	dnsServers := getDNSServers(ctx)

	g.Log().Infof(ctx, "获取网络接口列表完成，共 %d 个接口，耗时: %v", len(interfaces), time.Since(start))

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
				}
			}
		}
	}

	// 检查是否使用了systemd-resolved等本地解析器
	if len(dnsServers) == 1 && dnsServers[0] == "127.0.0.53" {
		realDNSServers := getDNSServersFromResolvectl(ctx)
		if len(realDNSServers) > 0 {
			dnsServers = realDNSServers
		}
	}

	return dnsServers
}

// getGatewayForInterface 获取指定网络接口的网关地址
func getGatewayForInterface(ctx context.Context, interfaceName string) string {
	switch runtime.GOOS {
	case "linux":
		return getGatewayLinux(interfaceName)
	case "darwin": // macOS
		return getGatewayMacOS(interfaceName)
	case "windows":
		return getGatewayWindows(interfaceName)
	default:
		return ""
	}
}

// getGatewayLinux 在Linux系统上获取网关地址
func getGatewayLinux(interfaceName string) string {
	// 执行 ip route show 命令获取路由信息
	cmd := exec.Command("ip", "route", "show", "dev", interfaceName)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 查找默认路由 (0.0.0.0/0 或 default)
		if strings.HasPrefix(line, "default") || strings.HasPrefix(line, "0.0.0.0/0") {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// 查找 "via" 关键字后的网关地址
				for i, field := range fields {
					if field == "via" && i+1 < len(fields) {
						gateway := fields[i+1]
						// 验证是否为有效的IP地址
						if net.ParseIP(gateway) != nil {
							return gateway
						}
					}
				}
			}
		}
	}

	return ""
}

// getGatewayMacOS 在macOS系统上获取网关地址
func getGatewayMacOS(interfaceName string) string {
	// 执行 route -n get default 命令获取默认路由
	cmd := exec.Command("route", "-n", "get", "default")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 查找 "gateway:" 行
		if strings.HasPrefix(line, "gateway:") {
			parts := strings.Split(line, "gateway:")
			if len(parts) > 1 {
				gateway := strings.TrimSpace(parts[1])
				// 验证是否为有效的IP地址
				if net.ParseIP(gateway) != nil {
					return gateway
				}
			}
		}
	}

	return ""
}

// getGatewayWindows 在Windows系统上获取网关地址
func getGatewayWindows(interfaceName string) string {
	// 执行 route print 命令获取路由信息
	cmd := exec.Command("route", "print")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	inActiveRoutes := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// 检查是否进入活动路由部分
		if strings.Contains(line, "Active Routes:") {
			inActiveRoutes = true
			continue
		}

		// 如果不在活动路由部分，跳过
		if !inActiveRoutes {
			continue
		}

		// 查找默认路由 (0.0.0.0 开头的行)
		if strings.HasPrefix(line, "0.0.0.0") {
			fields := strings.Fields(line)
			if len(fields) >= 3 {
				// 第二个字段通常是网关地址
				gateway := fields[2]
				// 验证是否为有效的IP地址
				if net.ParseIP(gateway) != nil {
					return gateway
				}
			}
		}
	}

	return ""
}

// getDNSServersFromResolvectl 使用resolvectl命令获取真实的DNS服务器
func getDNSServersFromResolvectl(ctx context.Context) []string {
	var dnsServers []string

	// 执行 resolvectl status 命令
	cmd := exec.Command("resolvectl", "status")
	output, err := cmd.Output()
	if err != nil {
		return dnsServers
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 查找包含 "DNS Servers:" 的行
		if strings.Contains(line, "DNS Servers:") {
			// 提取DNS服务器地址
			parts := strings.Split(line, "DNS Servers:")
			if len(parts) > 1 {
				dnsPart := strings.TrimSpace(parts[1])
				// 按空格分割多个DNS服务器
				servers := strings.Fields(dnsPart)
				for _, server := range servers {
					// 验证是否为有效的IP地址
					if net.ParseIP(server) != nil {
						dnsServers = append(dnsServers, server)
					}
				}
			}
			break
		}
	}

	return dnsServers
}

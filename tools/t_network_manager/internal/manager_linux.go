//go:build linux

package internal

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"t_network_manager/public"
	"time"

	"github.com/vishvananda/netlink"
)

// SManager 为 Linux 的 NetworkManager 实现
type SManager struct{}

// NewSManager 创建 Linux 网络管理器实例
func NewSManager() public.NetworkManager {
	return &SManager{}
}

// GetAllInterfaces 获取所有网络接口信息
func (m *SManager) GetAllInterfaces(ctx context.Context) ([]*public.InterfaceSummary, error) {
	// 获取所有网络链接
	links, err := netlink.LinkList()
	if err != nil {
		return nil, fmt.Errorf("获取网络接口列表失败: %w", err)
	}

	var interfaces []*public.InterfaceSummary
	for _, link := range links {
		summary, err := m.getInterfaceDetails(ctx, link)
		if err != nil {
			continue // 跳过有问题的接口
		}
		interfaces = append(interfaces, summary)
	}

	return interfaces, nil
}

// getInterfaceDetails 获取单个接口的详细信息
func (m *SManager) getInterfaceDetails(ctx context.Context, link netlink.Link) (*public.InterfaceSummary, error) {
	attrs := link.Attrs()

	summary := &public.InterfaceSummary{
		ID:   public.InterfaceID(attrs.Name),
		Name: attrs.Name,
		MAC:  attrs.HardwareAddr.String(),
		IsUp: attrs.Flags&net.FlagUp != 0,
		MTU:  attrs.MTU,
	}

	// 获取IP地址信息
	if err := m.getIPConfig(ctx, link, summary); err != nil {
		return nil, err
	}

	// 获取路由信息（网关）
	if err := m.getGatewayInfo(ctx, link, summary); err != nil {
		return nil, err
	}

	// 获取DNS配置
	if err := m.getDNSConfig(ctx, summary); err != nil {
		return nil, err
	}

	// 获取DHCP状态
	if err := m.getDHCPStatus(ctx, link, summary); err != nil {
		return nil, err
	}

	return summary, nil
}

// getIPConfig 获取IP配置信息
func (m *SManager) getIPConfig(ctx context.Context, link netlink.Link, summary *public.InterfaceSummary) error {
	// 获取接口的IP地址
	addresses, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return fmt.Errorf("获取IP地址失败: %w", err)
	}

	// 解析IPv4地址信息
	for _, addr := range addresses {
		if addr.IP.To4() != nil { // 确保是IPv4地址
			ipv4Config := &public.Ipv4Config{
				IPv4:       addr.IP.String(),
				SubnetMask: net.IP(addr.Mask).String(),
			}
			summary.IPv4 = append(summary.IPv4, ipv4Config)
		}
	}

	return nil
}

// getGatewayInfo 获取网关信息
func (m *SManager) getGatewayInfo(ctx context.Context, link netlink.Link, summary *public.InterfaceSummary) error {
	// 获取路由表
	routes, err := netlink.RouteList(link, netlink.FAMILY_V4)
	if err != nil {
		return fmt.Errorf("获取路由信息失败: %w", err)
	}

	// 查找默认路由（网关）
	for _, route := range routes {
		if route.Gw != nil {
			summary.Gateway = append(summary.Gateway, route.Gw.String())
		}
	}

	return nil
}

// getDNSConfig 获取DNS配置
func (m *SManager) getDNSConfig(ctx context.Context, summary *public.InterfaceSummary) error {
	// 首先尝试通过nmcli获取DNS配置
	if m.isNetworkManagerAvailable() {
		dnsServers, err := m.getDNSWithNmcli()
		if err == nil && len(dnsServers) > 0 {
			summary.DNS = dnsServers
			return nil
		}
	}

	// 如果nmcli不可用或获取失败，尝试读取 /etc/resolv.conf
	dnsServers, err := m.getDNSFromResolvConf()
	if err == nil {
		summary.DNS = dnsServers
	} else {
		summary.DNS = []string{} // 默认返回空列表
	}

	return nil
}

// getDNSWithNmcli 通过nmcli获取DNS配置
func (m *SManager) getDNSWithNmcli() ([]string, error) {
	// 获取活动连接的DNS配置
	cmd := exec.Command("nmcli", "-t", "-f", "IP4.DNS", "connection", "show", "--active")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("获取DNS配置失败: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	var dnsServers []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || line == "--" {
			continue
		}

		// DNS服务器可能以空格分隔
		servers := strings.Fields(line)
		for _, server := range servers {
			server = strings.TrimSpace(server)
			if server != "" && net.ParseIP(server) != nil {
				dnsServers = append(dnsServers, server)
			}
		}
	}

	return dnsServers, nil
}

// getDNSFromResolvConf 从 /etc/resolv.conf 获取DNS配置
func (m *SManager) getDNSFromResolvConf() ([]string, error) {
	content, err := os.ReadFile("/etc/resolv.conf")
	if err != nil {
		return nil, fmt.Errorf("读取 /etc/resolv.conf 失败: %w", err)
	}

	lines := strings.Split(string(content), "\n")
	var dnsServers []string

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "nameserver ") {
			server := strings.TrimSpace(strings.TrimPrefix(line, "nameserver "))
			if server != "" && net.ParseIP(server) != nil {
				dnsServers = append(dnsServers, server)
			}
		}
	}

	return dnsServers, nil
}

// getDHCPStatus 获取DHCP状态
func (m *SManager) getDHCPStatus(ctx context.Context, link netlink.Link, summary *public.InterfaceSummary) error {
	interfaceName := link.Attrs().Name

	// 尝试多种方法检测DHCP状态
	dhcpStatus := m.detectDHCPStatus(interfaceName)
	summary.DHCP = dhcpStatus

	return nil
}

// detectDHCPStatus 检测DHCP状态
func (m *SManager) detectDHCPStatus(interfaceName string) bool {
	// 方法1: 检查 /etc/network/interfaces
	if m.checkNetworkInterfaces(interfaceName) {
		return true
	}

	// 方法2: 检查 systemd-networkd 配置
	if m.checkSystemdNetworkd(interfaceName) {
		return true
	}

	// 方法3: 检查 dhclient 租约文件
	if m.checkDhclientLease(interfaceName) {
		return true
	}

	// 方法4: 检查 NetworkManager 状态（如果可用）
	if m.checkNetworkManager(interfaceName) {
		return true
	}

	// 默认返回false（静态配置）
	return false
}

// checkNetworkInterfaces 检查 /etc/network/interfaces 文件
func (m *SManager) checkNetworkInterfaces(interfaceName string) bool {
	content, err := os.ReadFile("/etc/network/interfaces")
	if err != nil {
		return false
	}

	lines := strings.Split(string(content), "\n")
	inInterface := false

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// 检查是否进入目标接口配置段
		if strings.HasPrefix(line, "iface "+interfaceName) {
			inInterface = true
			continue
		}

		// 如果遇到新的接口配置段，退出当前接口
		if strings.HasPrefix(line, "iface ") && !strings.HasPrefix(line, "iface "+interfaceName) {
			inInterface = false
			continue
		}

		// 在目标接口配置段中查找DHCP配置
		if inInterface && strings.Contains(line, "dhcp") {
			return true
		}
	}

	return false
}

// checkSystemdNetworkd 检查 systemd-networkd 配置
func (m *SManager) checkSystemdNetworkd(interfaceName string) bool {
	// 检查 /etc/systemd/network/ 目录下的配置文件
	configPath := fmt.Sprintf("/etc/systemd/network/%s.network", interfaceName)
	content, err := os.ReadFile(configPath)
	if err != nil {
		// 如果特定接口配置文件不存在，检查通用配置
		configPath = "/etc/systemd/network/50-dhcp.network"
		content, err = os.ReadFile(configPath)
		if err != nil {
			return false
		}
	}

	// 检查配置文件中是否包含DHCP配置
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "DHCP=yes") || strings.Contains(line, "DHCP=ipv4") {
			return true
		}
	}

	return false
}

// checkDhclientLease 检查 dhclient 租约文件
func (m *SManager) checkDhclientLease(interfaceName string) bool {
	// 检查 /var/lib/dhcp/ 目录下的租约文件
	leasePath := fmt.Sprintf("/var/lib/dhcp/dhclient.%s.leases", interfaceName)
	_, err := os.Stat(leasePath)
	if err == nil {
		return true
	}

	// 检查 /var/lib/dhclient/ 目录（某些发行版使用此路径）
	leasePath = fmt.Sprintf("/var/lib/dhclient/dhclient.%s.leases", interfaceName)
	_, err = os.Stat(leasePath)
	if err == nil {
		return true
	}

	return false
}

// checkNetworkManager 检查 NetworkManager 状态
func (m *SManager) checkNetworkManager(interfaceName string) bool {
	// 这里可以扩展使用 NetworkManager 的 D-Bus 接口
	// 或者检查 NetworkManager 的配置文件
	// 目前返回false，表示未检测到NetworkManager的DHCP配置
	return false
}

// UpdateInterfaceConfig 更新接口配置
func (m *SManager) UpdateInterfaceConfig(ctx context.Context, id public.InterfaceID, cfg public.InterfaceConfig) error {
	interfaceName := string(id)

	// 获取网络接口
	link, err := netlink.LinkByName(interfaceName)
	if err != nil {
		return fmt.Errorf("获取网络接口 %s 失败: %w", interfaceName, err)
	}

	// 处理DHCP配置
	if cfg.DHCP != nil {
		if *cfg.DHCP {
			// 启用DHCP - 在Linux中通常通过NetworkManager或systemd-networkd处理
			// 这里可以扩展实现DHCP配置
			return fmt.Errorf("DHCP配置暂未实现，请使用系统工具配置")
		} else {
			// 配置静态IP
			if err := m.configureStaticIP(ctx, link, cfg); err != nil {
				return fmt.Errorf("配置静态IP失败: %w", err)
			}
		}
	}

	// 配置网关
	if len(cfg.Gateway) > 0 {
		if err := m.configureGateway(ctx, link, cfg.Gateway[0]); err != nil {
			return fmt.Errorf("配置网关失败: %w", err)
		}
	}

	// 配置DNS
	if len(cfg.DNS) > 0 {
		if err := m.configureDNS(ctx, cfg.DNS); err != nil {
			return fmt.Errorf("配置DNS失败: %w", err)
		}
	}

	return nil
}

// configureStaticIP 配置静态IP
func (m *SManager) configureStaticIP(ctx context.Context, link netlink.Link, cfg public.InterfaceConfig) error {
	// 清除现有IP地址
	addresses, err := netlink.AddrList(link, netlink.FAMILY_V4)
	if err != nil {
		return fmt.Errorf("获取现有IP地址失败: %w", err)
	}

	for _, addr := range addresses {
		if err := netlink.AddrDel(link, &addr); err != nil {
			// 记录警告但不中断操作
			fmt.Printf("删除IP地址 %s 失败: %v\n", addr.IP.String(), err)
		}
	}

	// 添加新的IP地址
	for _, ipv4Config := range cfg.IPv4 {
		if ipv4Config == nil {
			continue
		}

		// 解析IP地址和子网掩码
		ip := net.ParseIP(ipv4Config.IPv4)
		if ip == nil {
			return fmt.Errorf("无效的IP地址: %s", ipv4Config.IPv4)
		}

		var mask net.IPMask
		if ipv4Config.SubnetMask != "" {
			maskIP := net.ParseIP(ipv4Config.SubnetMask)
			if maskIP == nil {
				return fmt.Errorf("无效的子网掩码: %s", ipv4Config.SubnetMask)
			}
			mask = net.IPMask(maskIP.To4())
		} else {
			// 默认使用/24子网掩码
			mask = net.CIDRMask(24, 32)
		}

		addr := &netlink.Addr{
			IPNet: &net.IPNet{
				IP:   ip,
				Mask: mask,
			},
		}

		if err := netlink.AddrAdd(link, addr); err != nil {
			return fmt.Errorf("添加IP地址 %s 失败: %w", ipv4Config.IPv4, err)
		}
	}

	return nil
}

// configureGateway 配置网关
func (m *SManager) configureGateway(ctx context.Context, link netlink.Link, gateway string) error {
	// 获取现有路由
	routes, err := netlink.RouteList(link, netlink.FAMILY_V4)
	if err != nil {
		return fmt.Errorf("获取现有路由失败: %w", err)
	}

	// 删除该接口的现有静态路由（非默认路由）
	for _, route := range routes {
		// 只删除非默认路由（有具体目标网络的路由）
		if route.Dst != nil && route.LinkIndex == link.Attrs().Index {
			if err := netlink.RouteDel(&route); err != nil {
				fmt.Printf("删除现有静态路由 %s 失败: %v\n", route.Dst.String(), err)
			}
		}
	}

	// 验证网关地址
	gatewayIP := net.ParseIP(gateway)
	if gatewayIP == nil {
		return fmt.Errorf("无效的网关地址: %s", gateway)
	}

	// 为接口添加静态路由（指向网关）
	// 这里添加一个指向网关的静态路由，而不是默认路由
	route := &netlink.Route{
		LinkIndex: link.Attrs().Index,
		Dst: &net.IPNet{
			IP:   gatewayIP,
			Mask: net.CIDRMask(32, 32), // 单主机路由
		},
		Scope: netlink.SCOPE_LINK, // 链路范围
	}

	if err := netlink.RouteAdd(route); err != nil {
		return fmt.Errorf("添加网关静态路由失败: %w", err)
	}

	return nil
}

// configureDNS 配置DNS服务器
func (m *SManager) configureDNS(ctx context.Context, dnsServers []string) error {
	if len(dnsServers) == 0 {
		return nil // 没有DNS服务器需要配置
	}

	// 验证DNS服务器地址格式
	for _, dns := range dnsServers {
		if net.ParseIP(dns) == nil {
			return fmt.Errorf("无效的DNS服务器地址: %s", dns)
		}
	}

	// 使用nmcli命令配置DNS服务器
	// 首先检查NetworkManager是否可用
	if !m.isNetworkManagerAvailable() {
		return fmt.Errorf("NetworkManager不可用，无法配置DNS")
	}

	// 获取当前连接名称
	connectionName, err := m.getCurrentConnectionName()
	if err != nil {
		return fmt.Errorf("获取当前连接名称失败: %w", err)
	}

	// 使用nmcli设置DNS服务器
	return m.setDNSWithNmcli(ctx, connectionName, dnsServers)
}

// isNetworkManagerAvailable 检查NetworkManager是否可用
func (m *SManager) isNetworkManagerAvailable() bool {
	cmd := exec.Command("which", "nmcli")
	err := cmd.Run()
	return err == nil
}

// getCurrentConnectionName 获取当前连接名称
func (m *SManager) getCurrentConnectionName() (string, error) {
	// 获取活动连接
	cmd := exec.Command("nmcli", "-t", "-f", "NAME,DEVICE,TYPE,STATE", "connection", "show", "--active")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("获取活动连接失败: %w", err)
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		parts := strings.Split(line, ":")
		if len(parts) >= 4 {
			connectionName := parts[0]
			_ = parts[1] // device (未使用)
			connectionType := parts[2]
			state := parts[3]

			// 查找以太网或WiFi连接
			if (connectionType == "802-3-ethernet" || connectionType == "wifi") && state == "activated" {
				return connectionName, nil
			}
		}
	}

	return "", fmt.Errorf("未找到活动的网络连接")
}

// setDNSWithNmcli 使用nmcli设置DNS服务器
func (m *SManager) setDNSWithNmcli(ctx context.Context, connectionName string, dnsServers []string) error {
	// 构建DNS服务器列表字符串
	dnsList := strings.Join(dnsServers, ",")

	// 使用nmcli修改连接的DNS设置
	cmd := exec.CommandContext(ctx, "nmcli", "connection", "modify", connectionName, "ipv4.dns", dnsList)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("设置DNS服务器失败: %w", err)
	}

	// 重新激活连接以应用DNS设置
	cmd = exec.CommandContext(ctx, "nmcli", "connection", "up", connectionName)
	if err := cmd.Run(); err != nil {
		// DNS设置可能已经生效，即使重新激活失败
		fmt.Printf("重新激活连接失败，但DNS设置可能已生效: %v\n", err)
	}

	return nil
}

// SetInterfaceState 设置接口状态（启用/禁用）
func (m *SManager) SetInterfaceState(ctx context.Context, id public.InterfaceID, up bool) error {
	interfaceName := string(id)

	// 获取网络接口
	link, err := netlink.LinkByName(interfaceName)
	if err != nil {
		return fmt.Errorf("获取网络接口 %s 失败: %w", interfaceName, err)
	}

	// 设置接口状态
	if up {
		if err := netlink.LinkSetUp(link); err != nil {
			return fmt.Errorf("启用网络接口 %s 失败: %w", interfaceName, err)
		}
	} else {
		if err := netlink.LinkSetDown(link); err != nil {
			return fmt.Errorf("禁用网络接口 %s 失败: %w", interfaceName, err)
		}
	}

	return nil
}

// Ping 执行ping测试
func (m *SManager) Ping(ctx context.Context, target string) (*public.PingResult, error) {
	// 验证目标地址
	if target == "" {
		return &public.PingResult{
			Target:    target,
			Success:   false,
			Error:     "目标地址不能为空",
			Timestamp: time.Now().Unix(),
		}, nil
	}

	// 使用系统的ping命令
	cmd := exec.CommandContext(ctx, "ping", "-c", "1", "-W", "3", target)
	output, err := cmd.Output()

	result := &public.PingResult{
		Target:    target,
		Timestamp: time.Now().Unix(),
	}

	if err != nil {
		result.Success = false
		result.Error = fmt.Sprintf("ping失败: %v", err)
		return result, nil
	}

	// 解析ping输出获取延迟时间
	latency, err := m.parsePingOutput(string(output))
	if err != nil {
		result.Success = false
		result.Error = fmt.Sprintf("解析ping结果失败: %v", err)
		return result, nil
	}

	result.Success = true
	result.Latency = latency
	return result, nil
}

// parsePingOutput 解析ping命令输出
func (m *SManager) parsePingOutput(output string) (float64, error) {
	// Linux ping输出格式示例：
	// PING 8.8.8.8 (8.8.8.8) 56(84) bytes of data.
	// 64 bytes from 8.8.8.8: icmp_seq=1 ttl=57 time=12.3 ms

	// 查找time=xxx ms模式
	re := regexp.MustCompile(`time=(\d+\.?\d*)\s*ms`)
	matches := re.FindStringSubmatch(output)

	if len(matches) < 2 {
		return 0, fmt.Errorf("无法从ping输出中解析延迟时间")
	}

	latency, err := strconv.ParseFloat(matches[1], 64)
	if err != nil {
		return 0, fmt.Errorf("解析延迟时间失败: %v", err)
	}

	return latency, nil
}

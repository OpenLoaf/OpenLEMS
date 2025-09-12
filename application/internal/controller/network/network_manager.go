package network

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	"application/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

// NetworkManager 网络管理器
type NetworkManager struct{}

// NewNetworkManager 创建网络管理器实例
func NewNetworkManager() *NetworkManager {
	return &NetworkManager{}
}

// GetInterfaceList 获取网络接口列表
func (nm *NetworkManager) GetInterfaceList(ctx context.Context, includeLoopback bool) ([]*entity.SNetworkInterface, error) {
	// 只支持Linux系统
	if runtime.GOOS != "linux" {
		return nm.getInterfaceListFallback(ctx, includeLoopback)
	}

	// 检查nmcli是否可用
	if !nm.isNmcliAvailable() {
		g.Log().Warningf(ctx, "nmcli 不可用，使用备用方法获取网络接口")
		return nm.getInterfaceListFallback(ctx, includeLoopback)
	}

	return nm.getInterfaceListWithNmcli(ctx, includeLoopback)
}

// getInterfaceListWithNmcli 使用nmcli获取网络接口列表
func (nm *NetworkManager) getInterfaceListWithNmcli(ctx context.Context, includeLoopback bool) ([]*entity.SNetworkInterface, error) {
	var interfaces []*entity.SNetworkInterface

	// 获取连接信息
	connections, err := nm.getNmcliConnections(ctx)
	if err != nil {
		g.Log().Warningf(ctx, "获取nmcli连接信息失败: %v，使用备用方法", err)
		return nm.getInterfaceListFallback(ctx, includeLoopback)
	}

	// 获取设备信息
	devices, err := nm.getNmcliDevices(ctx)
	if err != nil {
		g.Log().Warningf(ctx, "获取nmcli设备信息失败: %v，使用备用方法", err)
		return nm.getInterfaceListFallback(ctx, includeLoopback)
	}

	// 合并连接和设备信息
	for deviceName, device := range devices {
		// 过滤条件检查
		if !includeLoopback && deviceName == "lo" {
			continue
		}

		// 只处理有MAC地址的接口
		if device.MAC == "" || device.MAC == "--" {
			continue
		}

		iface := &entity.SNetworkInterface{
			Name:  deviceName,
			Type:  nm.normalizeInterfaceType(device.Type),
			MAC:   device.MAC,
			Up:    device.State == "connected" || device.State == "connecting",
			MTU:   device.MTU,
			Index: 0, // nmcli不提供index，使用系统接口获取
		}

		// 从系统接口获取Index
		if netIface, err := net.InterfaceByName(deviceName); err == nil {
			iface.Index = netIface.Index
		}

		// 获取连接配置信息
		if conn, exists := connections[deviceName]; exists {
			iface.DHCP = conn.Method == "auto"
			iface.IPAddresses = conn.IPAddresses
			iface.Netmask = conn.Netmask
			iface.Gateway = conn.Gateway
		}

		interfaces = append(interfaces, iface)
	}

	return interfaces, nil
}

// getInterfaceListFallback 备用方法获取网络接口列表
func (nm *NetworkManager) getInterfaceListFallback(ctx context.Context, includeLoopback bool) ([]*entity.SNetworkInterface, error) {
	// 获取所有网络接口
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("获取网络接口列表失败: %v", err)
	}

	var interfaces []*entity.SNetworkInterface
	for _, netIface := range netInterfaces {
		// 只返回有MAC地址的接口
		if len(netIface.HardwareAddr) == 0 {
			continue
		}

		// 过滤条件检查
		if !includeLoopback && netIface.Name == "lo" {
			continue
		}

		// 构建接口信息
		iface := &entity.SNetworkInterface{
			Name:  netIface.Name,
			Type:  nm.getInterfaceType(netIface),
			MAC:   netIface.HardwareAddr.String(),
			Up:    netIface.Flags&net.FlagUp != 0,
			MTU:   netIface.MTU,
			Index: netIface.Index,
		}

		// 获取IP地址信息
		if addrs, err := netIface.Addrs(); err == nil {
			var ipv4Addrs []string
			for _, addr := range addrs {
				if ipNet, ok := addr.(*net.IPNet); ok && ipNet.IP.To4() != nil {
					ipv4Addrs = append(ipv4Addrs, ipNet.IP.String())
					if iface.Netmask == "" {
						iface.Netmask = nm.getNetmaskFromIPNet(ipNet)
					}
				}
			}
			iface.IPAddresses = ipv4Addrs
		}

		// 获取网关和DHCP信息（使用原有方法）
		iface.Gateway = nm.getGatewayForInterface(ctx, netIface.Name)
		iface.DHCP = nm.isDHCPEnabled(ctx, netIface.Name)

		interfaces = append(interfaces, iface)
	}

	return interfaces, nil
}

// UpdateInterface 更新网络接口配置
func (nm *NetworkManager) UpdateInterface(ctx context.Context, req *UpdateInterfaceRequest) error {
	// 只支持Linux系统
	if runtime.GOOS != "linux" {
		return fmt.Errorf("网络接口更新功能仅支持Linux系统")
	}

	// 检查nmcli是否可用
	if !nm.isNmcliAvailable() {
		g.Log().Warningf(ctx, "nmcli 不可用，使用备用方法更新网络接口")
		return nm.updateInterfaceFallback(ctx, req)
	}

	return nm.updateInterfaceWithNmcli(ctx, req)
}

// updateInterfaceWithNmcli 使用nmcli更新网络接口配置
func (nm *NetworkManager) updateInterfaceWithNmcli(ctx context.Context, req *UpdateInterfaceRequest) error {
	g.Log().Infof(ctx, "使用nmcli更新网络接口配置: %s", req.Name)

	// 检查连接是否存在
	connectionName := nm.getConnectionName(ctx, req.Name)
	if connectionName == "" {
		// 创建新连接
		if err := nm.createConnection(ctx, req); err != nil {
			return fmt.Errorf("创建网络连接失败: %v", err)
		}
		connectionName = req.Name
	}

	// 更新连接配置
	if err := nm.modifyConnection(ctx, connectionName, req); err != nil {
		return fmt.Errorf("修改网络连接失败: %v", err)
	}

	// 激活连接
	if err := nm.activateConnection(ctx, connectionName, req.Name); err != nil {
		return fmt.Errorf("激活网络连接失败: %v", err)
	}

	g.Log().Infof(ctx, "网络接口配置更新完成: %s", req.Name)
	return nil
}

// updateInterfaceFallback 备用方法更新网络接口配置
func (nm *NetworkManager) updateInterfaceFallback(ctx context.Context, req *UpdateInterfaceRequest) error {
	if req.DHCP {
		return nm.applyDHCPConfigFallback(ctx, req)
	} else {
		return nm.applyStaticConfigFallback(ctx, req)
	}
}

// getNmcliConnections 获取nmcli连接信息
func (nm *NetworkManager) getNmcliConnections(ctx context.Context) (map[string]*ConnectionInfo, error) {
	cmd := exec.Command("nmcli", "-t", "-f", "NAME,DEVICE,TYPE,AUTOCONNECT,ACTIVE,STATE", "connection", "show")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("执行nmcli connection show失败: %v", err)
	}

	connections := make(map[string]*ConnectionInfo)
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")
		if len(fields) < 6 {
			continue
		}

		deviceName := fields[1]
		if deviceName == "" || deviceName == "--" {
			continue
		}

		conn := &ConnectionInfo{
			Name:        fields[0],
			Device:      deviceName,
			Type:        fields[2],
			AutoConnect: fields[3] == "yes",
			Active:      fields[4] == "yes",
			State:       fields[5],
		}

		// 获取详细配置信息
		if err := nm.getConnectionDetails(ctx, conn); err != nil {
			g.Log().Warningf(ctx, "获取连接 %s 详细信息失败: %v", conn.Name, err)
		}

		connections[deviceName] = conn
	}

	return connections, nil
}

// getNmcliDevices 获取nmcli设备信息
func (nm *NetworkManager) getNmcliDevices(ctx context.Context) (map[string]*DeviceInfo, error) {
	cmd := exec.Command("nmcli", "-t", "-f", "DEVICE,TYPE,STATE,CONNECTION", "device", "status")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("执行nmcli device status失败: %v", err)
	}

	devices := make(map[string]*DeviceInfo)
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}

		fields := strings.Split(line, ":")
		if len(fields) < 4 {
			continue
		}

		deviceName := fields[0]
		device := &DeviceInfo{
			Name:       deviceName,
			Type:       fields[1],
			State:      fields[2],
			Connection: fields[3],
		}

		// 获取设备详细信息
		if err := nm.getDeviceDetails(ctx, device); err != nil {
			g.Log().Warningf(ctx, "获取设备 %s 详细信息失败: %v", deviceName, err)
		}

		devices[deviceName] = device
	}

	return devices, nil
}

// getConnectionDetails 获取连接详细信息
func (nm *NetworkManager) getConnectionDetails(ctx context.Context, conn *ConnectionInfo) error {
	if !conn.Active {
		return nil
	}

	cmd := exec.Command("nmcli", "-t", "-f", "ipv4.method,ipv4.addresses,ipv4.gateway,ipv4.dns", "connection", "show", conn.Name)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "ipv4.method":
			conn.Method = value
		case "ipv4.addresses":
			if value != "" && value != "--" {
				// 解析IP地址和子网掩码
				addresses := strings.Split(value, ",")
				for _, addr := range addresses {
					addr = strings.TrimSpace(addr)
					if strings.Contains(addr, "/") {
						parts := strings.Split(addr, "/")
						if len(parts) == 2 {
							conn.IPAddresses = append(conn.IPAddresses, parts[0])
							if conn.Netmask == "" {
								if prefixLen, err := strconv.Atoi(parts[1]); err == nil {
									conn.Netmask = nm.prefixToHexMask(prefixLen)
								}
							}
						}
					}
				}
			}
		case "ipv4.gateway":
			if value != "" && value != "--" {
				conn.Gateway = value
			}
		case "ipv4.dns":
			if value != "" && value != "--" {
				// 解析DNS服务器地址
				dnsServers := strings.Split(value, ",")
				for _, dns := range dnsServers {
					dns = strings.TrimSpace(dns)
					if dns != "" && net.ParseIP(dns) != nil {
						conn.DNS = append(conn.DNS, dns)
					}
				}
			}
		}
	}

	return nil
}

// getDeviceDetails 获取设备详细信息
func (nm *NetworkManager) getDeviceDetails(ctx context.Context, device *DeviceInfo) error {
	cmd := exec.Command("nmcli", "-t", "-f", "GENERAL.HWADDR,GENERAL.MTU", "device", "show", device.Name)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		switch key {
		case "GENERAL.HWADDR":
			if value != "" && value != "--" {
				device.MAC = value
			}
		case "GENERAL.MTU":
			if mtu, err := strconv.Atoi(value); err == nil {
				device.MTU = mtu
			}
		}
	}

	return nil
}

// getConnectionName 获取设备对应的连接名称
func (nm *NetworkManager) getConnectionName(ctx context.Context, deviceName string) string {
	cmd := exec.Command("nmcli", "-t", "-f", "NAME", "connection", "show", "--active")
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	// 查找设备对应的活动连接
	cmd = exec.Command("nmcli", "-t", "-f", "NAME,DEVICE", "connection", "show", "--active")
	output, err = cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) >= 2 && fields[1] == deviceName {
			return fields[0]
		}
	}

	return ""
}

// createConnection 创建新的网络连接
func (nm *NetworkManager) createConnection(ctx context.Context, req *UpdateInterfaceRequest) error {
	var cmd *exec.Cmd

	if req.DHCP {
		// 创建DHCP连接
		args := []string{"connection", "add",
			"type", "ethernet",
			"con-name", req.Name,
			"ifname", req.Name,
			"ipv4.method", "auto"}

		// 添加DNS配置
		if len(req.DNS) > 0 {
			dnsList := strings.Join(req.DNS, ",")
			args = append(args, "ipv4.dns", dnsList)
		}

		cmd = exec.Command("nmcli", args...)
	} else {
		// 创建静态IP连接
		if len(req.IPAddresses) == 0 {
			return fmt.Errorf("静态配置需要提供IP地址")
		}

		// 将十六进制掩码转换为CIDR
		prefixLen, err := nm.hexMaskToPrefix(req.Netmask)
		if err != nil {
			return fmt.Errorf("无效的子网掩码: %v", err)
		}

		// 构建IP地址列表
		var addresses []string
		for _, ip := range req.IPAddresses {
			addresses = append(addresses, fmt.Sprintf("%s/%d", ip, prefixLen))
		}

		args := []string{"connection", "add",
			"type", "ethernet",
			"con-name", req.Name,
			"ifname", req.Name,
			"ipv4.method", "manual",
			"ipv4.addresses", strings.Join(addresses, ",")}

		if req.Gateway != "" {
			args = append(args, "ipv4.gateway", req.Gateway)
		}

		// 添加DNS配置
		if len(req.DNS) > 0 {
			dnsList := strings.Join(req.DNS, ",")
			args = append(args, "ipv4.dns", dnsList)
		}

		cmd = exec.Command("nmcli", args...)
	}

	g.Log().Debugf(ctx, "创建连接命令: %s", strings.Join(cmd.Args, " "))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("创建连接失败: %v, 输出: %s", err, string(output))
	}

	return nil
}

// modifyConnection 修改网络连接配置
func (nm *NetworkManager) modifyConnection(ctx context.Context, connectionName string, req *UpdateInterfaceRequest) error {
	var cmd *exec.Cmd

	if req.DHCP {
		// 修改为DHCP配置
		args := []string{"connection", "modify", connectionName,
			"ipv4.method", "auto",
			"ipv4.addresses", "",
			"ipv4.gateway", ""}

		// 设置DNS服务器
		if len(req.DNS) > 0 {
			dnsList := strings.Join(req.DNS, ",")
			args = append(args, "ipv4.dns", dnsList)
		} else {
			// 清除DNS配置
			args = append(args, "ipv4.dns", "")
		}

		cmd = exec.Command("nmcli", args...)
	} else {
		// 修改为静态IP配置
		if len(req.IPAddresses) == 0 {
			return fmt.Errorf("静态配置需要提供IP地址")
		}

		// 将十六进制掩码转换为CIDR
		prefixLen, err := nm.hexMaskToPrefix(req.Netmask)
		if err != nil {
			return fmt.Errorf("无效的子网掩码: %v", err)
		}

		// 构建IP地址列表
		var addresses []string
		for _, ip := range req.IPAddresses {
			addresses = append(addresses, fmt.Sprintf("%s/%d", ip, prefixLen))
		}

		args := []string{"connection", "modify", connectionName,
			"ipv4.method", "manual",
			"ipv4.addresses", strings.Join(addresses, ",")}

		if req.Gateway != "" {
			args = append(args, "ipv4.gateway", req.Gateway)
		} else {
			args = append(args, "ipv4.gateway", "")
		}

		// 设置DNS服务器
		if len(req.DNS) > 0 {
			dnsList := strings.Join(req.DNS, ",")
			args = append(args, "ipv4.dns", dnsList)
		} else {
			// 清除DNS配置
			args = append(args, "ipv4.dns", "")
		}

		cmd = exec.Command("nmcli", args...)
	}

	g.Log().Debugf(ctx, "修改连接命令: %s", strings.Join(cmd.Args, " "))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("修改连接失败: %v, 输出: %s", err, string(output))
	}

	return nil
}

// activateConnection 激活网络连接
func (nm *NetworkManager) activateConnection(ctx context.Context, connectionName, deviceName string) error {
	cmd := exec.Command("nmcli", "connection", "up", connectionName, "ifname", deviceName)
	g.Log().Debugf(ctx, "激活连接命令: %s", strings.Join(cmd.Args, " "))

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("激活连接失败: %v, 输出: %s", err, string(output))
	}

	// 等待连接建立
	time.Sleep(2 * time.Second)
	return nil
}

// GetDNSServers 获取DNS服务器列表
func (nm *NetworkManager) GetDNSServers(ctx context.Context) []string {
	if runtime.GOOS != "linux" {
		return []string{}
	}

	// 优先使用nmcli获取DNS信息
	if nm.isNmcliAvailable() {
		if dns := nm.getDNSFromNmcli(ctx); len(dns) > 0 {
			return dns
		}
	}

	// 备用方法：从resolv.conf获取
	return nm.getDNSFromResolvConf(ctx)
}

// getDNSFromNmcli 从nmcli获取DNS服务器
func (nm *NetworkManager) getDNSFromNmcli(ctx context.Context) []string {
	cmd := exec.Command("nmcli", "-t", "-f", "IP4.DNS", "device", "show")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	var dnsServers []string
	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "IP4.DNS[") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				dns := strings.TrimSpace(parts[1])
				if dns != "" && net.ParseIP(dns) != nil {
					dnsServers = append(dnsServers, dns)
				}
			}
		}
	}

	return dnsServers
}

// getDNSFromResolvConf 从resolv.conf获取DNS服务器
func (nm *NetworkManager) getDNSFromResolvConf(ctx context.Context) []string {
	// 使用systemd-resolved获取DNS
	if nm.isSystemdResolvedActive() {
		return nm.getDNSFromSystemdResolved(ctx)
	}

	// 直接读取resolv.conf
	return nm.getDNSFromResolvConfFile(ctx)
}

// 工具方法
func (nm *NetworkManager) isNmcliAvailable() bool {
	_, err := exec.LookPath("nmcli")
	return err == nil
}

func (nm *NetworkManager) normalizeInterfaceType(deviceType string) string {
	switch strings.ToLower(deviceType) {
	case "ethernet":
		return "ethernet"
	case "wifi", "wireless":
		return "wifi"
	case "loopback":
		return "loopback"
	default:
		return "ethernet"
	}
}

func (nm *NetworkManager) prefixToHexMask(prefixLen int) string {
	if prefixLen < 0 || prefixLen > 32 {
		return "ffffff00" // 默认 /24
	}

	mask := uint32(0xFFFFFFFF << (32 - prefixLen))
	return fmt.Sprintf("%08x", mask)
}

func (nm *NetworkManager) hexMaskToPrefix(hexMask string) (int, error) {
	maskValue, err := strconv.ParseUint(hexMask, 16, 32)
	if err != nil {
		return 0, fmt.Errorf("无效的十六进制掩码: %s", hexMask)
	}

	// 计算前缀长度
	ones := 0
	for i := 31; i >= 0; i-- {
		if (maskValue>>uint(i))&1 == 1 {
			ones++
		} else {
			break
		}
	}

	return ones, nil
}

// 保留原有的备用方法
func (nm *NetworkManager) getInterfaceType(iface net.Interface) string {
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

func (nm *NetworkManager) getNetmaskFromIPNet(ipNet *net.IPNet) string {
	return ipNet.Mask.String()
}

func (nm *NetworkManager) getGatewayForInterface(ctx context.Context, interfaceName string) string {
	if runtime.GOOS == "linux" {
		return nm.getGatewayLinux(interfaceName)
	}
	return ""
}

func (nm *NetworkManager) isDHCPEnabled(ctx context.Context, interfaceName string) bool {
	if runtime.GOOS == "linux" {
		return nm.isDHCPEnabledLinux(ctx, interfaceName)
	}
	return false
}

// 从原文件中保留的方法（简化版本）
func (nm *NetworkManager) getGatewayLinux(interfaceName string) string {
	cmd := exec.Command("ip", "route", "show", "dev", interfaceName)
	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "default") || strings.HasPrefix(line, "0.0.0.0/0") {
			fields := strings.Fields(line)
			for i, field := range fields {
				if field == "via" && i+1 < len(fields) {
					gateway := fields[i+1]
					if net.ParseIP(gateway) != nil {
						return gateway
					}
				}
			}
		}
	}
	return ""
}

func (nm *NetworkManager) isDHCPEnabledLinux(ctx context.Context, interfaceName string) bool {
	// 检查NetworkManager配置
	cmd := exec.Command("nmcli", "connection", "show", interfaceName)
	output, err := cmd.Output()
	if err == nil {
		lines := strings.Split(string(output), "\n")
		for _, line := range lines {
			if strings.Contains(line, "ipv4.method") && strings.Contains(line, "auto") {
				return true
			}
		}
	}
	return false
}

// getDNSFromSystemdResolved 从systemd-resolved获取DNS
func (nm *NetworkManager) getDNSFromSystemdResolved(ctx context.Context) []string {
	cmd := exec.Command("resolvectl", "status")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	var dnsServers []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, "DNS Servers:") {
			parts := strings.Split(line, "DNS Servers:")
			if len(parts) > 1 {
				dnsPart := strings.TrimSpace(parts[1])
				servers := strings.Fields(dnsPart)
				for _, server := range servers {
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

// getDNSFromResolvConfFile 直接从resolv.conf文件读取DNS
func (nm *NetworkManager) getDNSFromResolvConfFile(ctx context.Context) []string {
	cmd := exec.Command("cat", "/etc/resolv.conf")
	output, err := cmd.Output()
	if err != nil {
		return nil
	}

	var dnsServers []string
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "nameserver ") {
			dns := strings.TrimSpace(strings.TrimPrefix(line, "nameserver "))
			if dns != "" && net.ParseIP(dns) != nil {
				dnsServers = append(dnsServers, dns)
			}
		}
	}
	return dnsServers
}

// isSystemdResolvedActive 检查systemd-resolved是否活跃
func (nm *NetworkManager) isSystemdResolvedActive() bool {
	cmd := exec.Command("systemctl", "is-active", "systemd-resolved")
	output, err := cmd.Output()
	return err == nil && strings.TrimSpace(string(output)) == "active"
}

// 备用配置方法
func (nm *NetworkManager) applyDHCPConfigFallback(ctx context.Context, req *UpdateInterfaceRequest) error {
	cmds := [][]string{
		{"ip", "addr", "flush", "dev", req.Name},
		{"ip", "link", "set", req.Name, "up"},
		{"dhclient", "-r", req.Name},
		{"dhclient", req.Name},
	}

	for _, cmd := range cmds {
		g.Log().Debugf(ctx, "执行命令: %s", strings.Join(cmd, " "))
		output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		if err != nil && !strings.Contains(cmd[0], "dhclient") {
			return fmt.Errorf("命令执行失败 %s: %v, 输出: %s", strings.Join(cmd, " "), err, string(output))
		}
	}

	return nil
}

func (nm *NetworkManager) applyStaticConfigFallback(ctx context.Context, req *UpdateInterfaceRequest) error {
	prefixLen, err := nm.hexMaskToPrefix(req.Netmask)
	if err != nil {
		return err
	}

	cmds := [][]string{
		{"ip", "addr", "flush", "dev", req.Name},
	}

	for _, ip := range req.IPAddresses {
		cmds = append(cmds, []string{"ip", "addr", "add", fmt.Sprintf("%s/%d", ip, prefixLen), "dev", req.Name})
	}

	cmds = append(cmds, []string{"ip", "link", "set", req.Name, "up"})

	if req.Gateway != "" {
		networkAddr := nm.calculateNetworkAddress(req.IPAddresses[0], prefixLen)
		cmds = append(cmds, []string{"ip", "route", "del", networkAddr, "dev", req.Name})
		cmds = append(cmds, []string{"ip", "route", "add", networkAddr, "via", req.Gateway, "dev", req.Name})
	}

	for _, cmd := range cmds {
		g.Log().Debugf(ctx, "执行命令: %s", strings.Join(cmd, " "))
		output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput()
		if err != nil {
			if len(cmd) >= 3 && cmd[1] == "route" && cmd[2] == "del" {
				continue // 忽略删除路由的错误
			}
			return fmt.Errorf("命令执行失败 %s: %v, 输出: %s", strings.Join(cmd, " "), err, string(output))
		}
	}

	return nil
}

func (nm *NetworkManager) calculateNetworkAddress(ipAddress string, prefixLen int) string {
	ip := net.ParseIP(ipAddress)
	if ip == nil {
		return ""
	}

	ipNet := &net.IPNet{
		IP:   ip.To4(),
		Mask: net.CIDRMask(prefixLen, 32),
	}

	return ipNet.String()
}

//go:build windows

package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"os/exec"
	"strings"
	"t_network_manager/public"
)

// SManager 为 Windows 的 NetworkManager 实现
type SManager struct{}

// NewSManager 创建 Windows 网络管理器实例
func NewSManager() public.NetworkManager {
	return &SManager{}
}

// GetAllInterfaces 获取所有网络接口信息
func (m *SManager) GetAllInterfaces(ctx context.Context) ([]*public.InterfaceSummary, error) {
	// 使用 PowerShell 获取网络适配器信息
	cmd := exec.CommandContext(ctx, "powershell", "-Command", `
		Get-NetAdapter | Where-Object {$_.Status -eq "Up" -or $_.Status -eq "Down"} | 
		Select-Object Name, InterfaceDescription, MacAddress, Status, LinkSpeed | 
		ConvertTo-Json
	`)

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("获取网络适配器列表失败: %w", err)
	}

	// 解析 PowerShell 输出
	interfaces, err := m.parseNetAdapterOutput(string(output))
	if err != nil {
		return nil, fmt.Errorf("解析网络适配器信息失败: %w", err)
	}

	// 为每个接口获取详细信息
	var result []*public.InterfaceSummary
	for _, iface := range interfaces {
		summary, err := m.getInterfaceDetails(ctx, iface)
		if err != nil {
			continue // 跳过有问题的接口
		}
		result = append(result, summary)
	}

	return result, nil
}

// getInterfaceDetails 获取单个接口的详细信息
func (m *SManager) getInterfaceDetails(ctx context.Context, iface map[string]interface{}) (*public.InterfaceSummary, error) {
	name, ok := iface["Name"].(string)
	if !ok {
		return nil, fmt.Errorf("接口名称无效")
	}

	summary := &public.InterfaceSummary{
		ID:   public.InterfaceID(name),
		Name: name,
	}

	// 获取MAC地址
	if mac, ok := iface["MacAddress"].(string); ok {
		summary.MAC = mac
	}

	// 获取接口状态
	if status, ok := iface["Status"].(string); ok {
		summary.IsUp = status == "Up"
	}

	// 获取IP配置信息
	if err := m.getIPConfig(ctx, name, summary); err != nil {
		return nil, err
	}

	// 获取网关信息
	if err := m.getGatewayInfo(ctx, name, summary); err != nil {
		return nil, err
	}

	// 获取DNS配置
	if err := m.getDNSConfig(ctx, name, summary); err != nil {
		return nil, err
	}

	// 获取DHCP状态
	if err := m.getDHCPStatus(ctx, name, summary); err != nil {
		return nil, err
	}

	return summary, nil
}

// getIPConfig 获取IP配置信息
func (m *SManager) getIPConfig(ctx context.Context, interfaceName string, summary *public.InterfaceSummary) error {
	// 使用 PowerShell 获取IP配置
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		Get-NetIPAddress -InterfaceAlias "%s" -AddressFamily IPv4 | 
		Select-Object IPAddress, PrefixLength | 
		ConvertTo-Json
	`, interfaceName))

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("获取IP配置失败: %w", err)
	}

	// 解析IP配置
	ipConfigs, err := m.parseIPConfigOutput(string(output))
	if err != nil {
		return fmt.Errorf("解析IP配置失败: %w", err)
	}

	summary.IPv4 = ipConfigs
	return nil
}

// getGatewayInfo 获取网关信息
func (m *SManager) getGatewayInfo(ctx context.Context, interfaceName string, summary *public.InterfaceSummary) error {
	// 使用 PowerShell 获取网关信息
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		Get-NetRoute -InterfaceAlias "%s" -DestinationPrefix "0.0.0.0/0" | 
		Select-Object NextHop | 
		ConvertTo-Json
	`, interfaceName))

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("获取网关信息失败: %w", err)
	}

	// 解析网关信息
	gateways, err := m.parseGatewayOutput(string(output))
	if err != nil {
		return fmt.Errorf("解析网关信息失败: %w", err)
	}

	summary.Gateway = gateways
	return nil
}

// getDNSConfig 获取DNS配置
func (m *SManager) getDNSConfig(ctx context.Context, interfaceName string, summary *public.InterfaceSummary) error {
	// 使用 PowerShell 获取DNS配置
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		Get-DnsClientServerAddress -InterfaceAlias "%s" -AddressFamily IPv4 | 
		Select-Object ServerAddresses | 
		ConvertTo-Json
	`, interfaceName))

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("获取DNS配置失败: %w", err)
	}

	// 解析DNS配置
	dnsServers, err := m.parseDNSOutput(string(output))
	if err != nil {
		return fmt.Errorf("解析DNS配置失败: %w", err)
	}

	summary.DNS = dnsServers
	return nil
}

// getDHCPStatus 获取DHCP状态
func (m *SManager) getDHCPStatus(ctx context.Context, interfaceName string, summary *public.InterfaceSummary) error {
	// 使用 PowerShell 获取DHCP状态
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		Get-NetIPInterface -InterfaceAlias "%s" -AddressFamily IPv4 | 
		Select-Object Dhcp | 
		ConvertTo-Json
	`, interfaceName))

	output, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("获取DHCP状态失败: %w", err)
	}

	// 解析DHCP状态
	dhcpEnabled, err := m.parseDHCPStatus(string(output))
	if err != nil {
		return fmt.Errorf("解析DHCP状态失败: %w", err)
	}

	summary.DHCP = dhcpEnabled
	return nil
}

// UpdateInterfaceConfig 更新接口配置
func (m *SManager) UpdateInterfaceConfig(ctx context.Context, id public.InterfaceID, cfg public.InterfaceConfig) error {
	interfaceName := string(id)

	// 处理DHCP配置
	if cfg.DHCP != nil {
		if *cfg.DHCP {
			// 启用DHCP
			if err := m.enableDHCP(ctx, interfaceName); err != nil {
				return fmt.Errorf("启用DHCP失败: %w", err)
			}
		} else {
			// 配置静态IP
			if err := m.configureStaticIP(ctx, interfaceName, cfg); err != nil {
				return fmt.Errorf("配置静态IP失败: %w", err)
			}
		}
	}

	// 配置网关
	if len(cfg.Gateway) > 0 {
		if err := m.configureGateway(ctx, interfaceName, cfg.Gateway[0]); err != nil {
			return fmt.Errorf("配置网关失败: %w", err)
		}
	}

	// 配置DNS
	if len(cfg.DNS) > 0 {
		if err := m.configureDNS(ctx, interfaceName, cfg.DNS); err != nil {
			return fmt.Errorf("配置DNS失败: %w", err)
		}
	}

	return nil
}

// SetInterfaceState 设置接口状态（启用/禁用）
func (m *SManager) SetInterfaceState(ctx context.Context, id public.InterfaceID, up bool) error {
	interfaceName := string(id)

	// 使用 PowerShell 启用或禁用网络适配器
	var action string
	if up {
		action = "Enable"
	} else {
		action = "Disable"
	}

	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		%s-NetAdapter -Name "%s" -Confirm:$false
	`, action, interfaceName))

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("%s网络接口 %s 失败: %w", action, interfaceName, err)
	}

	return nil
}

// 辅助方法：解析网络适配器输出
func (m *SManager) parseNetAdapterOutput(output string) ([]map[string]interface{}, error) {
	// 清理输出，移除可能的 BOM 和空白字符
	output = strings.TrimSpace(output)
	if output == "" {
		return []map[string]interface{}{}, nil
	}

	// 尝试解析 JSON 数组
	var adapters []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &adapters); err != nil {
		// 如果解析失败，尝试解析单个对象
		var adapter map[string]interface{}
		if err := json.Unmarshal([]byte(output), &adapter); err != nil {
			return nil, fmt.Errorf("解析网络适配器JSON失败: %w", err)
		}
		adapters = []map[string]interface{}{adapter}
	}

	return adapters, nil
}

// 辅助方法：解析IP配置输出
func (m *SManager) parseIPConfigOutput(output string) ([]*public.Ipv4Config, error) {
	output = strings.TrimSpace(output)
	if output == "" {
		return []*public.Ipv4Config{}, nil
	}

	var ipConfigs []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &ipConfigs); err != nil {
		// 尝试解析单个对象
		var ipConfig map[string]interface{}
		if err := json.Unmarshal([]byte(output), &ipConfig); err != nil {
			return nil, fmt.Errorf("解析IP配置JSON失败: %w", err)
		}
		ipConfigs = []map[string]interface{}{ipConfig}
	}

	var result []*public.Ipv4Config
	for _, config := range ipConfigs {
		ipv4Config := &public.Ipv4Config{}

		if ipAddr, ok := config["IPAddress"].(string); ok {
			ipv4Config.IPv4 = ipAddr
		}

		if prefixLength, ok := config["PrefixLength"].(float64); ok {
			// 将前缀长度转换为子网掩码
			ipv4Config.SubnetMask = m.prefixLengthToSubnetMask(int(prefixLength))
		}

		if ipv4Config.IPv4 != "" {
			result = append(result, ipv4Config)
		}
	}

	return result, nil
}

// 辅助方法：解析网关输出
func (m *SManager) parseGatewayOutput(output string) ([]string, error) {
	output = strings.TrimSpace(output)
	if output == "" {
		return []string{}, nil
	}

	var routes []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &routes); err != nil {
		var route map[string]interface{}
		if err := json.Unmarshal([]byte(output), &route); err != nil {
			return nil, fmt.Errorf("解析网关JSON失败: %w", err)
		}
		routes = []map[string]interface{}{route}
	}

	var gateways []string
	for _, route := range routes {
		if nextHop, ok := route["NextHop"].(string); ok && nextHop != "" {
			gateways = append(gateways, nextHop)
		}
	}

	return gateways, nil
}

// 辅助方法：解析DNS输出
func (m *SManager) parseDNSOutput(output string) ([]string, error) {
	output = strings.TrimSpace(output)
	if output == "" {
		return []string{}, nil
	}

	var dnsConfigs []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &dnsConfigs); err != nil {
		var dnsConfig map[string]interface{}
		if err := json.Unmarshal([]byte(output), &dnsConfig); err != nil {
			return nil, fmt.Errorf("解析DNS JSON失败: %w", err)
		}
		dnsConfigs = []map[string]interface{}{dnsConfig}
	}

	var dnsServers []string
	for _, config := range dnsConfigs {
		if serverAddresses, ok := config["ServerAddresses"].([]interface{}); ok {
			for _, addr := range serverAddresses {
				if addrStr, ok := addr.(string); ok && addrStr != "" {
					dnsServers = append(dnsServers, addrStr)
				}
			}
		}
	}

	return dnsServers, nil
}

// 辅助方法：解析DHCP状态
func (m *SManager) parseDHCPStatus(output string) (bool, error) {
	output = strings.TrimSpace(output)
	if output == "" {
		return false, nil
	}

	var dhcpConfigs []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &dhcpConfigs); err != nil {
		var dhcpConfig map[string]interface{}
		if err := json.Unmarshal([]byte(output), &dhcpConfig); err != nil {
			return false, fmt.Errorf("解析DHCP JSON失败: %w", err)
		}
		dhcpConfigs = []map[string]interface{}{dhcpConfig}
	}

	for _, config := range dhcpConfigs {
		if dhcp, ok := config["Dhcp"].(string); ok {
			return strings.ToLower(dhcp) == "enabled", nil
		}
	}

	return false, nil
}

// 辅助方法：启用DHCP
func (m *SManager) enableDHCP(ctx context.Context, interfaceName string) error {
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		Set-NetIPInterface -InterfaceAlias "%s" -AddressFamily IPv4 -Dhcp Enabled
	`, interfaceName))

	return cmd.Run()
}

// 辅助方法：配置静态IP
func (m *SManager) configureStaticIP(ctx context.Context, interfaceName string, cfg public.InterfaceConfig) error {
	// 配置静态IP地址
	for _, ipv4Config := range cfg.IPv4 {
		if ipv4Config == nil {
			continue
		}

		// 计算前缀长度
		prefixLength := 24 // 默认前缀长度
		if ipv4Config.SubnetMask != "" {
			prefixLength = m.calculatePrefixLength(ipv4Config.SubnetMask)
		}

		cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
			New-NetIPAddress -InterfaceAlias "%s" -IPAddress %s -PrefixLength %d -AddressFamily IPv4
		`, interfaceName, ipv4Config.IPv4, prefixLength))

		if err := cmd.Run(); err != nil {
			return fmt.Errorf("配置IP地址 %s 失败: %w", ipv4Config.IPv4, err)
		}
	}

	return nil
}

// 辅助方法：配置网关
func (m *SManager) configureGateway(ctx context.Context, interfaceName string, gateway string) error {
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		New-NetRoute -InterfaceAlias "%s" -DestinationPrefix "0.0.0.0/0" -NextHop %s
	`, interfaceName, gateway))

	return cmd.Run()
}

// 辅助方法：配置DNS
func (m *SManager) configureDNS(ctx context.Context, interfaceName string, dnsServers []string) error {
	dnsList := strings.Join(dnsServers, ",")
	cmd := exec.CommandContext(ctx, "powershell", "-Command", fmt.Sprintf(`
		Set-DnsClientServerAddress -InterfaceAlias "%s" -ServerAddresses %s
	`, interfaceName, dnsList))

	return cmd.Run()
}

// 辅助方法：计算前缀长度
func (m *SManager) calculatePrefixLength(subnetMask string) int {
	// 将子网掩码转换为前缀长度
	// 例如：255.255.255.0 -> 24
	ip := net.ParseIP(subnetMask)
	if ip == nil {
		return 24 // 默认值
	}

	mask := net.IPMask(ip.To4())
	prefixLength, _ := mask.Size()
	return prefixLength
}

// 辅助方法：前缀长度转换为子网掩码
func (m *SManager) prefixLengthToSubnetMask(prefixLength int) string {
	// 将前缀长度转换为子网掩码
	// 例如：24 -> 255.255.255.0
	mask := net.CIDRMask(prefixLength, 32)
	return net.IP(mask).String()
}

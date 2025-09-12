//go:build darwin

package internal

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"t_network_manager/public"
)

// SManager 为 macOS 的 NetworkManager 实现
type SManager struct{}

// var SMangerInstance tnm.NetworkManager = {}
func NewSManager() public.NetworkManager {
	return &SManager{}
}

// GetAllInterfaces 获取所有网络接口信息
func (m *SManager) GetAllInterfaces(ctx context.Context) ([]*public.InterfaceSummary, error) {
	// 获取所有网络服务
	cmd := exec.CommandContext(ctx, "networksetup", "-listallnetworkservices")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("获取网络服务列表失败: %w", err)
	}

	services := strings.Split(string(output), "\n")
	var interfaces []*public.InterfaceSummary

	for _, service := range services {
		service = strings.TrimSpace(service)
		if service == "" || strings.Contains(service, "An asterisk") {
			continue
		}

		// 获取接口详细信息
		summary, err := m.getInterfaceDetails(ctx, service)
		if err != nil {
			continue // 跳过有问题的接口
		}

		interfaces = append(interfaces, summary)
	}

	return interfaces, nil
}

// getInterfaceDetails 获取单个接口的详细信息
func (m *SManager) getInterfaceDetails(ctx context.Context, serviceName string) (*public.InterfaceSummary, error) {
	summary := &public.InterfaceSummary{
		ID:   public.InterfaceID(serviceName),
		Name: serviceName,
	}

	// 获取硬件信息
	if err := m.getHardwareInfo(ctx, serviceName, summary); err != nil {
		return nil, err
	}

	// 获取IP配置
	if err := m.getIPConfig(ctx, serviceName, summary); err != nil {
		return nil, err
	}

	// 获取DNS配置
	if err := m.getDNSConfig(ctx, serviceName, summary); err != nil {
		return nil, err
	}

	// 获取网关信息
	if err := m.getGatewayInfo(ctx, serviceName, summary); err != nil {
		return nil, err
	}

	return summary, nil
}

// getHardwareInfo 获取硬件信息（MAC地址等）
func (m *SManager) getHardwareInfo(ctx context.Context, serviceName string, summary *public.InterfaceSummary) error {
	// 获取硬件端口信息
	cmd := exec.CommandContext(ctx, "networksetup", "-getmacaddress", serviceName)
	output, err := cmd.Output()
	if err == nil {
		// 解析MAC地址
		re := regexp.MustCompile(`([0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2}:[0-9a-fA-F]{2})`)
		matches := re.FindStringSubmatch(string(output))
		if len(matches) > 1 {
			summary.MAC = matches[1]
		}
	}

	// 检查接口状态
	cmd = exec.CommandContext(ctx, "networksetup", "-getinfo", serviceName)
	output, err = cmd.Output()
	if err == nil {
		// 简单检查是否有IP地址来判断接口是否启用
		summary.IsUp = strings.Contains(string(output), "IP address:")
	}

	return nil
}

// getIPConfig 获取IP配置信息
func (m *SManager) getIPConfig(ctx context.Context, serviceName string, summary *public.InterfaceSummary) error {
	cmd := exec.CommandContext(ctx, "networksetup", "-getinfo", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "MTU:") {
			// 解析MTU
			parts := strings.Fields(line)
			if len(parts) > 1 {
				if mtu, err := strconv.Atoi(parts[1]); err == nil {
					summary.MTU = mtu
				}
			}
		}
	}

	return nil
}

// getDNSConfig 获取DNS配置
func (m *SManager) getDNSConfig(ctx context.Context, serviceName string, summary *public.InterfaceSummary) error {
	cmd := exec.CommandContext(ctx, "networksetup", "-getdnsservers", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !strings.Contains(line, "There aren't any DNS Servers") {
			summary.DNS = append(summary.DNS, line)
		}
	}

	return nil
}

// getGatewayInfo 获取网关信息
func (m *SManager) getGatewayInfo(ctx context.Context, serviceName string, summary *public.InterfaceSummary) error {
	cmd := exec.CommandContext(ctx, "networksetup", "-getinfo", serviceName)
	output, err := cmd.Output()
	if err != nil {
		return err
	}

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Router:") {
			parts := strings.Fields(line)
			if len(parts) > 1 {
				// 判断是否是ip地址
				if net.ParseIP(parts[1]) != nil {
					summary.Gateway = append(summary.Gateway, parts[1])
				}
			}
		}
	}

	return nil
}

// UpdateInterfaceConfig 更新接口配置
func (m *SManager) UpdateInterfaceConfig(ctx context.Context, id public.InterfaceID, cfg public.InterfaceConfig) error {
	serviceName := string(id)

	// 设置DHCP或静态IP
	if cfg.DHCP != nil {
		if *cfg.DHCP {
			// 启用DHCP
			cmd := exec.CommandContext(ctx, "networksetup", "-setdhcp", serviceName)
			if err := cmd.Run(); err != nil {
				return fmt.Errorf("设置DHCP失败: %w", err)
			}
		} else {
			// 设置静态IP
			if len(cfg.IPv4) > 0 {
				ipv4 := cfg.IPv4[0]
				cmd := exec.CommandContext(ctx, "networksetup", "-setmanual", serviceName, ipv4.IPv4, ipv4.SubnetMask)
				if err := cmd.Run(); err != nil {
					return fmt.Errorf("设置静态IP失败: %w", err)
				}
			}
		}
	}

	// 设置网关
	if len(cfg.Gateway) > 0 {
		gateway := cfg.Gateway[0]
		cmd := exec.CommandContext(ctx, "networksetup", "-setadditionalroutes", serviceName, gateway, "0.0.0.0", "0.0.0.0")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("设置网关失败: %w", err)
		}
	}

	// 设置DNS
	if len(cfg.DNS) > 0 {
		args := []string{"-setdnsservers", serviceName}
		args = append(args, cfg.DNS...)

		cmd := exec.CommandContext(ctx, "networksetup", args...)
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("设置DNS失败: %w", err)
		}
	}

	return nil
}

// SetInterfaceState 设置接口状态（启用/禁用）
func (m *SManager) SetInterfaceState(ctx context.Context, id public.InterfaceID, up bool) error {
	serviceName := string(id)

	// 在macOS中，我们通过启用/禁用网络服务来控制接口状态
	if up {
		cmd := exec.CommandContext(ctx, "networksetup", "-setnetworkserviceenabled", serviceName, "on")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("启用网络服务失败: %w", err)
		}
	} else {
		cmd := exec.CommandContext(ctx, "networksetup", "-setnetworkserviceenabled", serviceName, "off")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("禁用网络服务失败: %w", err)
		}
	}

	return nil
}

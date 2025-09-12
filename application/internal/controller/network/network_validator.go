package network

import (
	"fmt"
	"net"
	"strconv"
)

// NetworkValidator 网络配置验证器
type NetworkValidator struct{}

// NewNetworkValidator 创建网络验证器实例
func NewNetworkValidator() *NetworkValidator {
	return &NetworkValidator{}
}

// ValidateUpdateRequest 验证更新请求
func (v *NetworkValidator) ValidateUpdateRequest(req *UpdateInterfaceRequest) error {
	// 验证接口名称
	if err := v.validateInterfaceName(req.Name); err != nil {
		return err
	}

	// 验证配置模式
	if req.DHCP {
		return v.validateDHCPConfig(req)
	} else {
		return v.validateStaticConfig(req)
	}
}

// validateInterfaceName 验证接口名称
func (v *NetworkValidator) validateInterfaceName(name string) error {
	if name == "" {
		return &ValidationError{Field: "name", Message: "网络接口名称不能为空"}
	}

	// 检查接口是否存在
	_, err := net.InterfaceByName(name)
	if err != nil {
		return &ValidationError{Field: "name", Message: fmt.Sprintf("网络接口 %s 不存在: %v", name, err)}
	}

	return nil
}

// validateDHCPConfig 验证DHCP配置
func (v *NetworkValidator) validateDHCPConfig(req *UpdateInterfaceRequest) error {
	// DHCP模式下不需要IP地址和子网掩码
	if len(req.IPAddresses) > 0 {
		return &ValidationError{Field: "ipAddresses", Message: "DHCP模式下不应设置IP地址"}
	}

	if req.Netmask != "" {
		return &ValidationError{Field: "netmask", Message: "DHCP模式下不应设置子网掩码"}
	}

	// 网关地址可选，但如果设置了需要验证格式
	if req.Gateway != "" {
		if err := v.validateIPAddress("gateway", req.Gateway); err != nil {
			return err
		}
	}

	// 验证DNS服务器地址
	for i, dns := range req.DNS {
		if err := v.validateIPAddress(fmt.Sprintf("dns[%d]", i), dns); err != nil {
			return err
		}
	}

	return nil
}

// validateStaticConfig 验证静态配置
func (v *NetworkValidator) validateStaticConfig(req *UpdateInterfaceRequest) error {
	// 静态配置模式：需要IP地址和子网掩码
	if len(req.IPAddresses) == 0 {
		return &ValidationError{Field: "ipAddresses", Message: "静态配置模式下IP地址列表不能为空"}
	}

	if req.Netmask == "" {
		return &ValidationError{Field: "netmask", Message: "静态配置模式下子网掩码不能为空"}
	}

	// 验证所有IP地址格式
	for i, ip := range req.IPAddresses {
		if err := v.validateIPAddress(fmt.Sprintf("ipAddresses[%d]", i), ip); err != nil {
			return err
		}
	}

	// 验证子网掩码格式
	if err := v.validateHexMask(req.Netmask); err != nil {
		return err
	}

	// 验证网关地址格式
	if req.Gateway != "" {
		if err := v.validateIPAddress("gateway", req.Gateway); err != nil {
			return err
		}

		// 验证网关地址与IP地址是否在同一网段
		if err := v.validateGatewayInSameSubnet(req.IPAddresses, req.Gateway, req.Netmask); err != nil {
			return err
		}
	}

	// 验证DNS服务器地址
	for i, dns := range req.DNS {
		if err := v.validateIPAddress(fmt.Sprintf("dns[%d]", i), dns); err != nil {
			return err
		}
	}

	return nil
}

// validateIPAddress 验证IP地址格式
func (v *NetworkValidator) validateIPAddress(field, ip string) error {
	if net.ParseIP(ip) == nil {
		return &ValidationError{Field: field, Message: fmt.Sprintf("IP地址格式不正确: %s", ip)}
	}

	// 检查是否为IPv4地址
	if parsedIP := net.ParseIP(ip); parsedIP != nil && parsedIP.To4() == nil {
		return &ValidationError{Field: field, Message: fmt.Sprintf("仅支持IPv4地址: %s", ip)}
	}

	return nil
}

// validateHexMask 验证十六进制子网掩码格式
func (v *NetworkValidator) validateHexMask(mask string) error {
	// 检查长度（应该是8个字符，如ffffff00）
	if len(mask) != 8 {
		return &ValidationError{Field: "netmask", Message: "子网掩码应为8个十六进制字符（如ffffff00）"}
	}

	// 检查是否都是十六进制字符
	for _, char := range mask {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return &ValidationError{Field: "netmask", Message: "子网掩码包含无效字符，应为十六进制格式"}
		}
	}

	// 尝试解析为32位整数
	maskValue, err := strconv.ParseUint(mask, 16, 32)
	if err != nil {
		return &ValidationError{Field: "netmask", Message: fmt.Sprintf("无效的十六进制掩码: %s", mask)}
	}

	// 验证掩码的连续性
	if err := v.validateMaskContinuity(maskValue); err != nil {
		return &ValidationError{Field: "netmask", Message: err.Error()}
	}

	return nil
}

// validateMaskContinuity 验证掩码的连续性
func (v *NetworkValidator) validateMaskContinuity(maskValue uint64) error {
	// 计算1的个数
	ones := 0
	for i := 31; i >= 0; i-- {
		if (maskValue>>uint(i))&1 == 1 {
			ones++
		} else {
			break
		}
	}

	// 检查剩余位是否都为0
	for i := 0; i < 32-ones; i++ {
		if (maskValue>>uint(i))&1 == 1 {
			return fmt.Errorf("子网掩码不连续，无效的掩码格式")
		}
	}

	return nil
}

// validateGatewayInSameSubnet 验证网关地址与IP地址是否在同一网段
func (v *NetworkValidator) validateGatewayInSameSubnet(ipAddresses []string, gateway, hexMask string) error {
	// 解析网关地址
	gatewayIP := net.ParseIP(gateway)
	if gatewayIP == nil {
		return &ValidationError{Field: "gateway", Message: fmt.Sprintf("网关地址格式不正确: %s", gateway)}
	}

	// 将十六进制掩码转换为net.IPMask
	maskValue, err := strconv.ParseUint(hexMask, 16, 32)
	if err != nil {
		return &ValidationError{Field: "netmask", Message: fmt.Sprintf("无效的十六进制掩码: %s", hexMask)}
	}

	// 创建子网掩码
	mask := net.IPv4Mask(
		byte(maskValue>>24),
		byte(maskValue>>16),
		byte(maskValue>>8),
		byte(maskValue),
	)

	// 检查每个IP地址是否与网关在同一网段
	for i, ipStr := range ipAddresses {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			return &ValidationError{Field: fmt.Sprintf("ipAddresses[%d]", i), Message: fmt.Sprintf("IP地址格式不正确: %s", ipStr)}
		}

		// 计算IP地址和网关的网络地址
		ipNetwork := &net.IPNet{IP: ip.To4(), Mask: mask}
		gatewayNetwork := &net.IPNet{IP: gatewayIP.To4(), Mask: mask}

		// 比较网络地址
		if !ipNetwork.Contains(gatewayIP) && !gatewayNetwork.Contains(ip) {
			return &ValidationError{
				Field:   "gateway",
				Message: fmt.Sprintf("网关地址 %s 与IP地址 %s 不在同一网段（掩码: %s）", gateway, ipStr, hexMask),
			}
		}
	}

	return nil
}

// ValidateInterfaceExists 验证网络接口是否存在且可用
func (v *NetworkValidator) ValidateInterfaceExists(name string) error {
	iface, err := net.InterfaceByName(name)
	if err != nil {
		return &ValidationError{Field: "name", Message: fmt.Sprintf("网络接口 %s 不存在: %v", name, err)}
	}

	// 检查接口是否有MAC地址（排除loopback等虚拟接口的某些情况）
	if len(iface.HardwareAddr) == 0 && name != "lo" {
		return &ValidationError{Field: "name", Message: fmt.Sprintf("网络接口 %s 没有MAC地址，可能不是物理接口", name)}
	}

	return nil
}

// ValidateConfiguration 验证网络配置是否合理
func (v *NetworkValidator) ValidateConfiguration(req *UpdateInterfaceRequest) []ValidationError {
	var errors []ValidationError

	// 基础验证
	if err := v.ValidateUpdateRequest(req); err != nil {
		if validationErr, ok := err.(*ValidationError); ok {
			errors = append(errors, *validationErr)
		} else {
			errors = append(errors, ValidationError{Field: "general", Message: err.Error()})
		}
		return errors
	}

	// 额外的配置合理性检查
	if !req.DHCP {
		// 检查IP地址范围
		for i, ip := range req.IPAddresses {
			if err := v.validateIPAddressRange(ip); err != nil {
				errors = append(errors, ValidationError{
					Field:   fmt.Sprintf("ipAddresses[%d]", i),
					Message: err.Error(),
				})
			}
		}

		// 检查网关可达性
		if req.Gateway != "" {
			if err := v.validateGatewayReachability(req.Gateway); err != nil {
				errors = append(errors, ValidationError{
					Field:   "gateway",
					Message: err.Error(),
				})
			}
		}
	}

	return errors
}

// validateIPAddressRange 验证IP地址范围
func (v *NetworkValidator) validateIPAddressRange(ip string) error {
	parsedIP := net.ParseIP(ip)
	if parsedIP == nil {
		return fmt.Errorf("无效的IP地址")
	}

	ipv4 := parsedIP.To4()
	if ipv4 == nil {
		return fmt.Errorf("仅支持IPv4地址")
	}

	// 检查是否为私有地址范围或公网地址
	if v.isPrivateIP(ipv4) {
		return nil // 私有地址范围是安全的
	}

	// 检查是否为特殊用途地址
	if v.isSpecialUseIP(ipv4) {
		return fmt.Errorf("不能使用特殊用途的IP地址")
	}

	// 公网地址需要特别注意
	return fmt.Errorf("使用公网IP地址请确保配置正确")
}

// isPrivateIP 检查是否为私有IP地址
func (v *NetworkValidator) isPrivateIP(ip net.IP) bool {
	// 10.0.0.0/8
	if ip[0] == 10 {
		return true
	}
	// 172.16.0.0/12
	if ip[0] == 172 && ip[1] >= 16 && ip[1] <= 31 {
		return true
	}
	// 192.168.0.0/16
	if ip[0] == 192 && ip[1] == 168 {
		return true
	}
	return false
}

// isSpecialUseIP 检查是否为特殊用途IP地址
func (v *NetworkValidator) isSpecialUseIP(ip net.IP) bool {
	// 127.0.0.0/8 (loopback)
	if ip[0] == 127 {
		return true
	}
	// 169.254.0.0/16 (link-local)
	if ip[0] == 169 && ip[1] == 254 {
		return true
	}
	// 224.0.0.0/4 (multicast)
	if ip[0] >= 224 && ip[0] <= 239 {
		return true
	}
	// 0.0.0.0/8
	if ip[0] == 0 {
		return true
	}
	return false
}

// validateGatewayReachability 验证网关可达性（基础检查）
func (v *NetworkValidator) validateGatewayReachability(gateway string) error {
	// 这里可以添加ping测试或其他可达性检查
	// 目前只做基础格式验证
	if net.ParseIP(gateway) == nil {
		return fmt.Errorf("网关地址格式不正确")
	}
	return nil
}

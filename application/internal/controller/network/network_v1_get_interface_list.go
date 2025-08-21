package network

import (
	"bufio"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"time"

	v1 "application/api/network/v1"
	"application/internal/model/entity"
)

// GetNetworkInterfaceList 获取本机网络接口列表
func (c *ControllerV1) GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error) {
	start := time.Now()
	// macOS 走 system_profiler，以获取更完整的网卡信息
	if runtime.GOOS == "darwin" {
		dnsServers := readSystemDNSServers()
		list, derr := getInterfacesFromSystemProfiler(req, dnsServers)
		if derr != nil {
			return nil, derr
		}
		// 按名称排序
		for i := 0; i < len(list)-1; i++ {
			for j := i + 1; j < len(list); j++ {
				if list[i].Name > list[j].Name {
					list[i], list[j] = list[j], list[i]
				}
			}
		}
		res = &v1.GetNetworkInterfaceListRes{Interfaces: list, Total: len(list), DNS: dnsServers}
		_ = start
		return res, nil
	}

	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	gatewayMap := getDefaultGatewayPerInterface()
	dnsServers := readSystemDNSServers()

	var list []*entity.SNetworkInterface
	for _, nif := range interfaces {
		// 可选过滤：仅以太网
		if req.OnlyEthernet && !isEthernetInterface(nif.Name) {
			continue
		}
		item := &entity.SNetworkInterface{
			Name:      nif.Name,
			MAC:       nif.HardwareAddr.String(),
			Connected: (nif.Flags & net.FlagUp) != 0,
			Gateway:   gatewayMap[nif.Name],
		}

		if (nif.Flags & net.FlagLoopback) != 0 {
			item.Type = "loopback"
		} else if (nif.Flags & net.FlagPointToPoint) != 0 {
			item.Type = "p2p"
		} else {
			item.Type = "ethernet"
		}

		addrs, _ := nif.Addrs()
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok {
				ip := ipNet.IP
				if ip == nil || ip.IsLoopback() {
					continue
				}
				if ip.To4() != nil {
					item.IPv4 = ip.String()
					mask := ipNet.Mask
					if mask != nil {
						item.Netmask = net.IP(mask).String()
					}
					break
				}
			}
		}

		// 若接口完全无 IPv4 且非 up，跳过（避免无效虚拟接口）
		if item.IPv4 == "" && !item.Connected {
			continue
		}
		list = append(list, item)
	}

	// 按名称排序，稳定输出
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i].Name > list[j].Name {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	res = &v1.GetNetworkInterfaceListRes{Interfaces: list, Total: len(list), DNS: dnsServers}
	_ = start
	return res, nil
}

// readSystemDNSServers 从 /etc/resolv.conf 读取 DNS 列表
func readSystemDNSServers() []string {
	file, err := os.Open("/etc/resolv.conf")
	if err != nil {
		return nil
	}
	defer file.Close()

	var servers []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 2 && strings.ToLower(fields[0]) == "nameserver" {
			ip := fields[1]
			servers = append(servers, ip)
		}
	}
	return servers
}

// getDefaultGatewayPerInterface 获取默认网关(按接口)
func getDefaultGatewayPerInterface() map[string]string {
	result := make(map[string]string)
	switch runtime.GOOS {
	case "linux":
		data, err := os.ReadFile("/proc/net/route")
		if err != nil {
			return result
		}
		lines := strings.Split(string(data), "\n")
		for _, line := range lines[1:] { // 跳过表头
			fields := strings.Fields(line)
			if len(fields) < 3 {
				continue
			}
			iface := fields[0]
			destination := fields[1]
			gatewayHex := fields[2]
			if destination == "00000000" && len(gatewayHex) == 8 {
				b, err := hex.DecodeString(gatewayHex)
				if err != nil || len(b) != 4 {
					continue
				}
				// 小端序
				ip := net.IPv4(b[3], b[2], b[1], b[0]).String()
				result[iface] = ip
			}
		}
	case "darwin":
		// macOS 通过 route 命令获取默认网关及接口
		cmd := exec.Command("route", "-n", "get", "default")
		out, err := cmd.Output()
		if err != nil {
			return result
		}
		var gw, iface string
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "gateway:") {
				gw = strings.TrimSpace(strings.TrimPrefix(line, "gateway:"))
			}
			if strings.HasPrefix(line, "interface:") {
				iface = strings.TrimSpace(strings.TrimPrefix(line, "interface:"))
			}
		}
		if gw != "" && iface != "" {
			result[iface] = gw
		}
	}
	return result
}

// macOS：使用 system_profiler 获取网卡信息
func getInterfacesFromSystemProfiler(req *v1.GetNetworkInterfaceListReq, dnsServers []string) ([]*entity.SNetworkInterface, error) {
	cmd := exec.Command("system_profiler", "SPNetworkDataType", "-json")
	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("system_profiler failed: %v", err)
	}

	// 只解析我们关心的字段
	var sp struct {
		SPNetworkDataType []struct {
			Name      string                 `json:"_name"`
			Hardware  string                 `json:"hardware"`
			Iface     string                 `json:"interface"`
			Type      string                 `json:"type"`
			Ethernet  map[string]interface{} `json:"Ethernet"`
			IPv4      map[string]interface{} `json:"IPv4"`
			DNS       map[string]interface{} `json:"DNS"`
			IPAddress []string               `json:"ip_address"`
		} `json:"SPNetworkDataType"`
	}
	if err := json.Unmarshal(out, &sp); err != nil {
		return nil, fmt.Errorf("parse system_profiler json failed: %v", err)
	}

	var list []*entity.SNetworkInterface
	for _, it := range sp.SPNetworkDataType {
		iface := it.Iface
		if iface == "" {
			continue
		}

		// 仅以太网过滤
		if req.OnlyEthernet {
			hw := strings.ToLower(it.Hardware)
			if !(strings.Contains(hw, "ethernet") || strings.HasPrefix(strings.ToLower(iface), "en")) {
				continue
			}
		}

		item := &entity.SNetworkInterface{
			Name:      it.Name,
			Type:      strings.ToLower(it.Hardware),
			Connected: false,
		}

		// MAC
		if it.Ethernet != nil {
			if macVal, ok := it.Ethernet["MAC SourceAddress"].(string); ok {
				item.MAC = macVal
			}
		}

		// IPv4 地址与子网掩码
		if it.IPv4 != nil {
			if addrs, ok := it.IPv4["Addresses"].([]interface{}); ok && len(addrs) > 0 {
				if s, ok2 := addrs[0].(string); ok2 {
					item.IPv4 = s
				}
			}
			if masks, ok := it.IPv4["SubnetMasks"].([]interface{}); ok && len(masks) > 0 {
				if s, ok2 := masks[0].(string); ok2 {
					item.Netmask = s
				}
			}
			if r, ok := it.IPv4["Router"].(string); ok && r != "" {
				item.Gateway = r
			}
		}
		if item.IPv4 == "" && len(it.IPAddress) > 0 {
			item.IPv4 = it.IPAddress[0]
		}

		// 连接状态简单判定：有 IP 即视为连接
		item.Connected = item.IPv4 != ""

		// 若仅以太网且未连接且无 IP，可以保留由前端选择是否显示
		list = append(list, item)
	}
	return list, nil
}

// isEthernetInterface 判断 iface 是否为以太网(过滤掉 Wi‑Fi/回环等)
func isEthernetInterface(iface string) bool {
	switch runtime.GOOS {
	case "linux":
		// 通过 /sys/class/net/<iface>/type 判断：1=ARPHRD_ETHER(以太网)、772=loopback、280=CAN、65534=none
		if iface == "lo" {
			return false
		}
		if b, err := os.ReadFile(fmt.Sprintf("/sys/class/net/%s/type", iface)); err == nil {
			t := strings.TrimSpace(string(b))
			if t == "1" {
				return true
			}
			return false
		}
		// 回退：过滤无线与明显的非以太网前缀
		wirelessPath := filepath.Join("/sys/class/net", iface, "wireless")
		if _, err := os.Stat(wirelessPath); err == nil {
			return false
		}
		lname := strings.ToLower(iface)
		if strings.HasPrefix(lname, "can") || strings.HasPrefix(lname, "vcan") || strings.HasPrefix(lname, "tun") || strings.HasPrefix(lname, "tap") {
			return false
		}
		return true
	case "darwin":
		// 解析 networksetup 输出判断是否为有线以太网
		cmd := exec.Command("networksetup", "-listallhardwareports")
		out, err := cmd.Output()
		if err != nil {
			// 未能获取映射，回退：保留所有 en*，但排除 en0（多为 Wi‑Fi）与桥接名
			return strings.HasPrefix(iface, "en") && iface != "en0"
		}
		var currentPort, currentDev string
		lowerHas := func(s string, keys []string) bool {
			s = strings.ToLower(s)
			for _, k := range keys {
				if strings.Contains(s, k) {
					return true
				}
			}
			return false
		}
		noKeys := []string{"wi-fi", "wifi", "bluetooth", "bridge", "firewire"}
		for _, line := range strings.Split(string(out), "\n") {
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "Hardware Port:") {
				currentPort = strings.TrimSpace(strings.TrimPrefix(line, "Hardware Port:"))
			} else if strings.HasPrefix(line, "Device:") {
				currentDev = strings.TrimSpace(strings.TrimPrefix(line, "Device:"))
				if currentDev == iface {
					if lowerHas(currentPort, noKeys) {
						return false
					}
					// 只要不是 Wi‑Fi/Bridge/BT 等，且设备为 en*，一律视为以太网（含 USB/雷电）
					return strings.HasPrefix(currentDev, "en")
				}
			}
		}
		// 未找到映射：回退规则，保留 en*（常见以太网命名），排除 awdl/llw/utun/bridge
		if strings.HasPrefix(iface, "en") {
			return true
		}
		lname := strings.ToLower(iface)
		if strings.HasPrefix(lname, "awdl") || strings.HasPrefix(lname, "llw") || strings.HasPrefix(lname, "utun") || strings.HasPrefix(lname, "bridge") || lname == "lo0" {
			return false
		}
		return true
	default:
		return iface != "lo"
	}
}

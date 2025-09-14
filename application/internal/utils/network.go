package utils

import "net"

// GetLocalIPv4Addrs 获取所有本地IPv4地址
func GetLocalIPv4Addrs() ([]string, error) {
	var ipv4Addrs []string

	// 获取所有网络接口
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	for _, iface := range interfaces {
		// 跳过down状态的接口
		if iface.Flags&net.FlagUp == 0 {
			continue
		}

		// 跳过loopback接口
		if iface.Flags&net.FlagLoopback != 0 {
			continue
		}

		// 获取接口的地址
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}

		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}

			// 只获取IPv4地址，且不是回环地址
			if ip != nil && ip.To4() != nil && !ip.IsLoopback() {
				ipv4Addrs = append(ipv4Addrs, ip.String())
			}
		}
	}

	return ipv4Addrs, nil
}

package network

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strings"

	v1 "application/api/network/v1"

	"github.com/gogf/gf/v2/frame/g"
)

// UpdateNetworkInterface 仅更新 IP/掩码/网关/DNS
func (c *ControllerV1) UpdateNetworkInterface(ctx context.Context, req *v1.UpdateNetworkInterfaceReq) (res *v1.UpdateNetworkInterfaceRes, err error) {
	// 1) 基础校验（GoFrame 已做格式校验，这里做互斥/逻辑校验）
	if net.ParseIP(req.IP) == nil {
		return nil, fmt.Errorf("IP地址格式不正确")
	}
	if _, ipnet, perr := net.ParseCIDR(req.IP + "/24"); ipnet == nil || perr != nil { // 仅用于快速触发解析
		// 忽略
	}
	if req.Gateway != "" && net.ParseIP(req.Gateway) == nil {
		return nil, fmt.Errorf("网关地址格式不正确")
	}
	// DNS 已移出本接口

	// 2) 应用配置
	g.Log().Infof(ctx, "开始应用网络配置: name=%s ip=%s mask=%s gw=%s", req.Name, req.IP, req.Netmask, req.Gateway)
	switch runtime.GOOS {
	case "linux":
		if err = applyLinux(req); err != nil {
			g.Log().Errorf(ctx, "应用Linux网络配置失败: %v", err)
			return nil, err
		}
	case "darwin":
		if err = applyDarwin(req); err != nil {
			g.Log().Errorf(ctx, "应用Darwin网络配置失败: %v", err)
			return nil, err
		}
	default:
		return nil, errors.New("unsupported OS")
	}

	// 3) 验证
	ok, verifyMsg := verifyInterface(req)
	g.Log().Infof(ctx, "应用完成，验证结果: ok=%v msg=%s", ok, verifyMsg)
	if !ok {
		return nil, fmt.Errorf(verifyMsg)
	}
	return &v1.UpdateNetworkInterfaceRes{}, nil
}

func applyLinux(req *v1.UpdateNetworkInterfaceReq) error {
	// ip addr flush dev <name> ; ip addr add <ip>/<mask> dev <name>
	mask, err := maskToPrefix(req.Netmask)
	if err != nil {
		return err
	}
	cmds := [][]string{
		{"ip", "addr", "flush", "dev", req.Name},
		{"ip", "addr", "add", fmt.Sprintf("%s/%d", req.IP, mask), "dev", req.Name},
	}
	if req.Gateway != "" {
		cmds = append(cmds, []string{"ip", "route", "replace", "default", "via", req.Gateway, "dev", req.Name})
	}
	for _, c := range cmds {
		if out, err := exec.Command(c[0], c[1:]...).CombinedOutput(); err != nil {
			return fmt.Errorf("%s failed: %v, %s", strings.Join(c, " "), err, string(out))
		}
	}
	// DNS 不在此接口更新
	return nil
}

func applyDarwin(req *v1.UpdateNetworkInterfaceReq) error {
	// networksetup -setmanual <device> <ip> <subnet> <router>
	if out, err := exec.Command("networksetup", "-setmanual", req.Name, req.IP, req.Netmask, req.Gateway).CombinedOutput(); err != nil {
		return fmt.Errorf("setmanual failed: %v, %s", err, string(out))
	}
	// DNS 不在此接口更新
	return nil
}

func verifyInterface(req *v1.UpdateNetworkInterfaceReq) (bool, string) {
	// 读取系统当前状态进行比对
	ifi, err := net.InterfaceByName(req.Name)
	if err != nil {
		return false, err.Error()
	}
	addrs, _ := ifi.Addrs()
	var ipOK bool
	for _, a := range addrs {
		if ipNet, ok := a.(*net.IPNet); ok && ipNet.IP.To4() != nil {
			ipOK = ipNet.IP.String() == req.IP
			if ipOK {
				break
			}
		}
	}
	if !ipOK {
		return false, "IP未生效"
	}
	// 网关与 DNS 的强校验在不同系统复杂，这里报告基本成功
	return true, "配置已应用"
}

func maskToPrefix(mask string) (int, error) {
	ip := net.ParseIP(mask)
	if ip == nil {
		return 0, fmt.Errorf("子网掩码格式不正确")
	}
	ip = ip.To4()
	if ip == nil {
		return 0, fmt.Errorf("子网掩码格式不正确")
	}
	ones := 0
	for _, b := range []byte{ip[0], ip[1], ip[2], ip[3]} {
		for i := 7; i >= 0; i-- {
			if (b>>uint(i))&1 == 1 {
				ones++
			} else {
				break
			}
		}
	}
	return ones, nil
}

func emptyIf(v string, def string) string {
	if v == "" {
		return def
	}
	return v
}

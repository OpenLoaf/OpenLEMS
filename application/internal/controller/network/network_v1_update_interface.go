package network

import (
	"context"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	v1 "application/api/network/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// UpdateNetworkInterface 更新网络接口配置（仅支持Linux）
func (c *ControllerV1) UpdateNetworkInterface(ctx context.Context, req *v1.UpdateNetworkInterfaceReq) (res *v1.UpdateNetworkInterfaceRes, err error) {
	// 只支持Linux系统
	if runtime.GOOS != "linux" {
		return nil, gerror.NewCode(gcode.CodeNotSupported, "网络接口更新功能仅支持Linux系统")
	}

	// 基础校验
	if net.ParseIP(req.IP) == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "IP地址格式不正确")
	}

	// 验证子网掩码格式（十六进制格式，如ffffff00）
	if !isValidHexMask(req.Netmask) {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "子网掩码格式不正确，应为十六进制格式（如ffffff00）")
	}

	if req.Gateway != "" && net.ParseIP(req.Gateway) == nil {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "网关地址格式不正确")
	}

	g.Log().Infof(ctx, "开始应用网络配置: name=%s ip=%s mask=%s gw=%s", req.Name, req.IP, req.Netmask, req.Gateway)

	// 应用配置
	if err = applyLinuxConfig(ctx, req); err != nil {
		g.Log().Errorf(ctx, "应用Linux网络配置失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "应用网络配置失败")
	}

	// 验证配置
	ok, verifyMsg := verifyInterface(req)
	g.Log().Infof(ctx, "应用完成，验证结果: ok=%v msg=%s", ok, verifyMsg)
	if !ok {
		return nil, gerror.NewCode(gcode.CodeInternalError, verifyMsg)
	}

	return &v1.UpdateNetworkInterfaceRes{}, nil
}

func applyLinuxConfig(ctx context.Context, req *v1.UpdateNetworkInterfaceReq) error {
	// 将十六进制子网掩码转换为CIDR前缀长度
	prefixLen, err := hexMaskToPrefix(req.Netmask)
	if err != nil {
		return err
	}

	// 构建命令序列
	cmds := [][]string{
		{"ip", "addr", "flush", "dev", req.Name},
		{"ip", "addr", "add", fmt.Sprintf("%s/%d", req.IP, prefixLen), "dev", req.Name},
		{"ip", "link", "set", req.Name, "up"},
	}

	// 如果指定了网关，添加默认路由
	if req.Gateway != "" {
		cmds = append(cmds, []string{"ip", "route", "replace", "default", "via", req.Gateway, "dev", req.Name})
	}

	// 执行命令
	for _, cmd := range cmds {
		g.Log().Debugf(ctx, "执行命令: %s", strings.Join(cmd, " "))
		if output, err := exec.Command(cmd[0], cmd[1:]...).CombinedOutput(); err != nil {
			return fmt.Errorf("命令执行失败 %s: %v, 输出: %s", strings.Join(cmd, " "), err, string(output))
		}
	}

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

// isValidHexMask 验证十六进制子网掩码格式
func isValidHexMask(mask string) bool {
	// 检查长度（应该是8个字符，如ffffff00）
	if len(mask) != 8 {
		return false
	}

	// 检查是否都是十六进制字符
	for _, char := range mask {
		if !((char >= '0' && char <= '9') || (char >= 'a' && char <= 'f') || (char >= 'A' && char <= 'F')) {
			return false
		}
	}

	// 尝试解析为32位整数
	_, err := strconv.ParseUint(mask, 16, 32)
	return err == nil
}

// hexMaskToPrefix 将十六进制子网掩码转换为CIDR前缀长度
func hexMaskToPrefix(hexMask string) (int, error) {
	// 解析十六进制掩码
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

	// 验证掩码的连续性（确保没有0在1之后）
	temp := maskValue
	for i := 0; i < 32-ones; i++ {
		if (temp>>uint(i))&1 == 1 {
			return 0, fmt.Errorf("无效的子网掩码: %s（掩码不连续）", hexMask)
		}
	}

	return ones, nil
}

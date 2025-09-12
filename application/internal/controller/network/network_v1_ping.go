package network

import (
	"context"
	"net"

	v1 "application/api/network/v1"
	"t_network_manager"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// Ping 执行ping测试
func (c *ControllerV1) Ping(ctx context.Context, req *v1.PingReq) (res *v1.PingRes, err error) {
	// 1. 参数验证
	if err := c.validatePingRequest(req); err != nil {
		return nil, err
	}

	// 2. 获取网络管理器实例
	networkManager := t_network_manager.GetInstance()
	if networkManager == nil {
		g.Log().Errorf(ctx, "获取网络管理器实例失败")
		return nil, gerror.NewCode(gcode.CodeInternalError, "网络管理器初始化失败")
	}

	// 3. 执行ping测试
	result, err := networkManager.Ping(ctx, req.Target)
	if err != nil {
		g.Log().Errorf(ctx, "执行ping测试失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "ping测试失败")
	}

	// 4. 记录测试结果
	if result.Success {
		g.Log().Infof(ctx, "ping %s 成功，延迟: %.2fms", req.Target, result.Latency)
	} else {
		g.Log().Warningf(ctx, "ping %s 失败: %s", req.Target, result.Error)
	}

	return &v1.PingRes{
		Result: result,
	}, nil
}

// validatePingRequest 验证ping请求参数
func (c *ControllerV1) validatePingRequest(req *v1.PingReq) error {
	if req.Target == "" {
		return gerror.NewCode(gcode.CodeInvalidParameter, "目标地址不能为空")
	}

	// 验证目标地址格式（IP地址或域名）
	if net.ParseIP(req.Target) == nil {
		// 如果不是IP地址，检查是否为有效的域名格式
		if !c.isValidDomain(req.Target) {
			return gerror.NewCode(gcode.CodeInvalidParameter, "无效的目标地址格式")
		}
	}

	return nil
}

// isValidDomain 检查是否为有效的域名格式
func (c *ControllerV1) isValidDomain(domain string) bool {
	// 简单的域名格式验证
	if len(domain) == 0 || len(domain) > 253 {
		return false
	}

	// 检查是否包含非法字符
	for _, char := range domain {
		if !((char >= 'a' && char <= 'z') ||
			(char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') ||
			char == '.' || char == '-') {
			return false
		}
	}

	// 检查是否以点开头或结尾
	if domain[0] == '.' || domain[len(domain)-1] == '.' {
		return false
	}

	// 检查是否包含连续的点
	for i := 0; i < len(domain)-1; i++ {
		if domain[i] == '.' && domain[i+1] == '.' {
			return false
		}
	}

	return true
}

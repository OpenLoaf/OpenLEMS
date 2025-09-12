package network

import (
	"context"
	"time"

	v1 "application/api/network/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// GetNetworkInterfaceList 获取本机网络接口列表
func (c *ControllerV1) GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error) {
	start := time.Now()
	g.Log().Infof(ctx, "开始获取网络接口列表，参数: includeLoopback=%v", req.IncludeLoopback)

	// 创建网络管理器
	networkManager := NewNetworkManager()

	// 获取网络接口列表
	interfaces, err := networkManager.GetInterfaceList(ctx, req.IncludeLoopback)
	if err != nil {
		g.Log().Errorf(ctx, "获取网络接口列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取网络接口列表失败")
	}

	// 获取DNS服务器地址列表
	dnsServers := networkManager.GetDNSServers(ctx)

	g.Log().Infof(ctx, "获取网络接口列表完成，共 %d 个接口，耗时: %v", len(interfaces), time.Since(start))

	return &v1.GetNetworkInterfaceListRes{
		Interfaces: interfaces,
		Total:      len(interfaces),
		DNS:        dnsServers,
	}, nil
}

package network

import (
	"context"
	"net"
	"time"

	v1 "application/api/network/v1"
	"application/internal/model/entity"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/vishvananda/netlink"
)

// GetNetworkInterfaceList 获取本机网络接口列表
func (c *ControllerV1) GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error) {
	start := time.Now()
	g.Log().Infof(ctx, "开始获取网络接口列表，参数: includeLoopback=%v", req.IncludeLoopback)

	// 获取所有网络接口
	links, err := netlink.LinkList()
	if err != nil {
		g.Log().Errorf(ctx, "获取网络接口列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取网络接口列表失败")
	}

	var interfaces []*entity.SNetworkInterface
	for _, link := range links {
		attrs := link.Attrs()

		// 只返回有MAC地址的接口
		if attrs.HardwareAddr == nil || attrs.HardwareAddr.String() == "" {
			continue
		}

		// 过滤条件检查
		if !req.IncludeLoopback && attrs.Name == "lo" {
			continue
		}

		// 构建接口信息
		iface := &entity.SNetworkInterface{
			Name:  attrs.Name,
			Type:  link.Type(),
			MAC:   attrs.HardwareAddr.String(),
			Up:    attrs.Flags&net.FlagUp != 0,
			MTU:   attrs.MTU,
			Index: attrs.Index,
		}

		// 获取IP地址信息
		addrs, err := netlink.AddrList(link, 0)
		if err != nil {
			g.Log().Warningf(ctx, "获取接口 %s 的IP地址失败: %+v", attrs.Name, err)
		} else {
			var ipv4Addrs, ipv6Addrs []string
			for _, addr := range addrs {
				ipStr := addr.IP.String()
				if addr.IP.To4() != nil {
					ipv4Addrs = append(ipv4Addrs, ipStr)
					if iface.IPv4 == "" {
						iface.IPv4 = ipStr
						iface.Netmask = getNetmaskFromCIDR(addr.IPNet.String())
					}
				} else {
					ipv6Addrs = append(ipv6Addrs, ipStr)
					if iface.IPv6 == "" {
						iface.IPv6 = ipStr
					}
				}
			}
			iface.IPAddresses = append(ipv4Addrs, ipv6Addrs...)
		}

		// 判断连接状态
		iface.Connected = iface.Up && len(iface.IPAddresses) > 0

		interfaces = append(interfaces, iface)
	}

	g.Log().Infof(ctx, "获取网络接口列表完成，共 %d 个接口，耗时: %v", len(interfaces), time.Since(start))
	return &v1.GetNetworkInterfaceListRes{
		Interfaces: interfaces,
		Total:      len(interfaces),
	}, nil
}

// getNetmaskFromCIDR 从CIDR格式中提取子网掩码
func getNetmaskFromCIDR(cidr string) string {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		return ""
	}
	return ipNet.Mask.String()
}

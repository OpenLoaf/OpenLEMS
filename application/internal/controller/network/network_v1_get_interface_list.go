package network

import (
	"common/c_log"
	"context"
	"strings"
	"sync"
	"t_network_manager/public"
	"time"

	v1 "application/api/network/v1"
	"t_network_manager"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
)

// 缓存相关常量
const (
	NetworkInterfaceCacheKey = "network:interfaces" // 网络接口缓存键
	DefaultCacheDuration     = 5 * time.Minute      // 默认缓存过期时间
)

// 全局缓存实例
var (
	networkInterfaceCache *gcache.Cache
	networkCacheOnce      sync.Once
)

// initNetworkCache 初始化网络接口缓存
func initNetworkCache() {
	networkCacheOnce.Do(func() {
		networkInterfaceCache = gcache.New()
	})
}

// GetNetworkInterfaceList 获取本机网络接口列表
func (c *ControllerV1) GetNetworkInterfaceList(ctx context.Context, req *v1.GetNetworkInterfaceListReq) (res *v1.GetNetworkInterfaceListRes, err error) {
	// 1. 初始化缓存
	initNetworkCache()

	// 2. 检查是否需要强制刷新缓存
	if !req.ForceRefresh {
		// 尝试从缓存获取数据
		if cachedData, err := c.getCachedInterfaces(ctx); err == nil && cachedData != nil {
			c_log.Debugf(ctx, "从缓存获取网络接口列表，共 %d 个接口", len(cachedData))
			return &v1.GetNetworkInterfaceListRes{
				Interfaces: c.filterInterfaces(cachedData, req.IncludeLoopback),
			}, nil
		}
	}

	// 3. 缓存未命中或强制刷新，从网络管理器获取数据
	networkManager := t_network_manager.GetInstance()
	if networkManager == nil {
		c_log.Errorf(ctx, "获取网络管理器实例失败")
		return nil, gerror.NewCode(gcode.CodeInternalError, "网络管理器初始化失败")
	}

	// 4. 获取所有网络接口
	interfaces, err := networkManager.GetAllInterfaces(ctx)
	if err != nil {
		c_log.Errorf(ctx, "获取网络接口列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取网络接口列表失败")
	}

	// 5. 缓存数据
	if err := c.setCachedInterfaces(ctx, interfaces); err != nil {
		c_log.Warningf(ctx, "缓存网络接口数据失败: %+v", err)
	}

	// 6. 数据过滤和构建响应
	result := c.filterInterfaces(interfaces, req.IncludeLoopback)
	c_log.Infof(ctx, "成功获取 %d 个网络接口（强制刷新: %v）", len(result), req.ForceRefresh)

	return &v1.GetNetworkInterfaceListRes{
		Interfaces: result,
	}, nil
}

// isLoopbackInterface 判断是否为回环接口
func (c *ControllerV1) isLoopbackInterface(name string) bool {
	// 常见的回环接口名称
	loopbackNames := []string{
		"lo", "lo0", "Loopback", "Loopback Pseudo-Interface 1",
	}

	name = strings.ToLower(name)
	for _, loopbackName := range loopbackNames {
		if strings.Contains(name, strings.ToLower(loopbackName)) {
			return true
		}
	}

	return false
}

// getCachedInterfaces 从缓存获取网络接口数据
func (c *ControllerV1) getCachedInterfaces(ctx context.Context) ([]*public.InterfaceSummary, error) {
	value, err := networkInterfaceCache.Get(ctx, NetworkInterfaceCacheKey)
	if err != nil {
		return nil, err
	}

	if value == nil {
		return nil, nil
	}

	// 从 gvar.Var 中获取实际值
	actualValue := value.Val()
	if interfaces, ok := actualValue.([]*public.InterfaceSummary); ok {
		return interfaces, nil
	}

	return nil, nil
}

// setCachedInterfaces 设置网络接口数据到缓存
func (c *ControllerV1) setCachedInterfaces(ctx context.Context, interfaces []*public.InterfaceSummary) error {
	return networkInterfaceCache.Set(ctx, NetworkInterfaceCacheKey, interfaces, DefaultCacheDuration)
}

// filterInterfaces 过滤网络接口
func (c *ControllerV1) filterInterfaces(interfaces []*public.InterfaceSummary, includeLoopback bool) []*public.InterfaceSummary {
	var result []*public.InterfaceSummary

	for _, iface := range interfaces {
		// 根据参数决定是否包含回环接口
		if !includeLoopback && c.isLoopbackInterface(iface.Name) {
			continue
		}
		if iface.MAC == "" {
			continue
		}
		result = append(result, iface)
	}

	return result
}

// clearNetworkInterfaceCache 清除网络接口缓存
func (c *ControllerV1) clearNetworkInterfaceCache(ctx context.Context) {
	initNetworkCache()
	networkInterfaceCache.Remove(ctx, NetworkInterfaceCacheKey)
	c_log.Debugf(ctx, "已清除网络接口缓存")
}

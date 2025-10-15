package utils

import (
	"context"
	"os"
	"time"

	"t_machine_id"

	"github.com/gogf/gf/v2/frame/g"
)

// PrintWebServerInfo 打印Web服务器信息
func PrintWebServerInfo(ctx context.Context, serverAddress string) {
	// 延迟一点时间，确保GoFrame先打印完路由信息
	time.Sleep(100 * time.Millisecond)

	// 获取所有IPv4地址
	ipv4Addrs, err := GetLocalIPv4Addrs()
	if err != nil {
		g.Log().Warningf(ctx, "获取IPv4地址失败: %v", err)
		ipv4Addrs = []string{"localhost"}
	}

	// 获取系统序列号
	systemNumber := t_machine_id.GetMachineId()

	// 打印服务器访问地址
	g.Log().Infof(ctx, "==========================================")
	g.Log().Infof(ctx, "🚀 LEMS Web服务已启动！PID: %d", os.Getpid())
	g.Log().Infof(ctx, "🔢 系统序列号: %s", systemNumber)

	// 打印所有IPv4地址的服务器访问地址
	g.Log().Infof(ctx, "📡 服务器地址:")
	g.Log().Infof(ctx, "   http://localhost%s", serverAddress)
	for _, ip := range ipv4Addrs {
		g.Log().Infof(ctx, "   http://%s%s", ip, serverAddress)
	}

	// 获取当前配置的 profile，如果不是 prod 模式则显示 API 文档和调试工具地址
	profile, _ := g.Cfg().GetWithCmd(context.Background(), "profile", "prod")
	if profile.String() != "prod" {
		// 打印API文档和调试工具地址（使用第一个IPv4地址或localhost）
		primaryIP := "localhost"
		if len(ipv4Addrs) > 0 {
			primaryIP = ipv4Addrs[0]
		}
		g.Log().Infof(ctx, "📖 API文档: http://%s%s/api.json", primaryIP, serverAddress)
		g.Log().Infof(ctx, "🔧 调试工具: http://%s%s/debug", primaryIP, serverAddress)
	}
	g.Log().Infof(ctx, "==========================================")
}

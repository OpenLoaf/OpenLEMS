package common

import (
	"common/c_base"
	"context"
)

type IManager interface {
	Start(ctx context.Context)   // 启动服务
	Shutdown()                   // 停止管理器（释放资源、退出 goroutine）
	Cleanup() error              // 清理过期/无效资源（定时调用）
	Status() c_base.EServerState // 返回运行状态
}

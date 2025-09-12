package public

import (
	"context"
)

// NetworkManager 通用网络管理器接口（统一定义）
type NetworkManager interface {
	// 基础查询
	GetAllInterfaces(ctx context.Context) ([]*InterfaceSummary, error)

	// 配置与状态
	UpdateInterfaceConfig(ctx context.Context, id InterfaceID, cfg InterfaceConfig) error
	SetInterfaceState(ctx context.Context, id InterfaceID, up bool) error
}

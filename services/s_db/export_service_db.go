package s_db

import (
	"common"
	"context"
	"s_db/internal/impl"
)

// NewDbDriverConfigService 创建基础包中获取驱动配置服务的实现
func NewDbDriverConfigService(ctx context.Context) common.IDriverConfigService {
	return impl.NewDriverConfigService(ctx)
}

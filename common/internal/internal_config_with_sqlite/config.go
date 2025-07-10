package internal_config_with_sqlite

import (
	"context"
	"sqlite/service"
	"sync"

	database "sqlite"

	"github.com/gogf/gf/v2/os/gcmd"
)

type SConfig struct {
	instance service.IConfigManage
}

var (
	configInitOnce = sync.Once{}
	ConfigInstance *SConfig
)

func init() {
	gcmd.GetOptWithEnv("device")
}

func NewConfigInstance(ctx context.Context) *SConfig {
	configInitOnce.Do(func() {
		ConfigInstance = &SConfig{
			instance: database.NewConfigManage(ctx),
		}
	})
	return ConfigInstance
}

// func (c *SConfig) OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, error) {
// 	functionName := c_base.ConstNewPluginFunctionName
// 	if len(symName) != 0 {
// 		functionName = symName[0]
// 	}
// 	return openPlugin(ctx, gfile.Join(c.driverPath, path), functionName)
// }

// func openPlugin(ctx context.Context, path string, symName string) (plugin.Symbol, error) {
// 	if !gfile.Exists(path) {
// 		panic(gerror.Newf("插件%s不存在", path))
// 	}
// 	g.Log().Infof(ctx, "加载插件：%s", path)
// 	// 打开插件
// 	p, err := plugin.Open(path)
// 	if err != nil {
// 		panic(gerror.Newf("打开插件%s失败,symName%s。失败原因：%v", path, symName, err))
// 	}

// 	// 查找并使用结构体和函数
// 	return p.Lookup(symName)
// }

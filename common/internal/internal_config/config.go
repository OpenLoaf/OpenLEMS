package internal_config

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"plugin"
	"sync"
)

type SConfig struct {
	deviceCfgName string // 设备配置文件名
	driverPath    string // 驱动存放路径

}

var (
	configInitOnce = sync.Once{}
	ConfigInstance *SConfig
)

func NewConfigInstance(deviceCfgName, driverPath string) *SConfig {
	if deviceCfgName == "" {
		deviceCfgName = "device"
	}
	if driverPath == "" {
		driverPath = "driver"
	}
	configInitOnce.Do(func() {
		ConfigInstance = &SConfig{
			deviceCfgName: deviceCfgName,
			driverPath:    driverPath,
		}
	})
	return ConfigInstance
}

func (c *SConfig) OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, error) {
	functionName := c_base.ConstNewPluginFunctionName
	if len(symName) != 0 {
		functionName = symName[0]
	}
	return openPlugin(ctx, gfile.Join(c.driverPath, path), functionName)
}

func openPlugin(ctx context.Context, path string, symName string) (plugin.Symbol, error) {
	if !gfile.Exists(path) {
		panic(gerror.Newf("插件%s不存在", path))
	}
	g.Log().Infof(ctx, "加载插件：%s", path)
	// 打开插件
	p, err := plugin.Open(path)
	if err != nil {
		panic(gerror.Newf("打开插件%s失败原因：%v", path, err))
	}

	// 查找并使用结构体和函数
	return p.Lookup(symName)
}

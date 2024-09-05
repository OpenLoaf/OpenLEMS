package internal_config

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"github.com/gogf/gf/v2/util/gconv"
)

func (c *SConfig) GetDriverConfig(ctx context.Context) *c_base.SDriverConfig {
	config, configPath, err := getConfig[c_base.SDriverConfig](ctx, c.deviceCfgName, "device")
	if err != nil {
		panic(gerror.Newf("解析设备配置文件中的[drivers]失败！配置文件路径: %s 错误原因：%v", configPath, err))
	}
	g.Log().Infof(ctx, "加载设备配置文件：%s", configPath)
	return config
}

func (c *SConfig) GetProtocolsConfig(ctx context.Context) []*c_base.SProtocolConfig {
	list, configPath, err := getConfigList[c_base.SProtocolConfig](ctx, c.deviceCfgName, "protocols")
	if err != nil {
		panic(gerror.Newf("解析设备配置文件中的[protocols]失败！配置文件路径: %s 错误原因：%v", configPath, err))
	}
	return list
}

func (c *SConfig) GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig {
	list := c.GetProtocolsConfig(ctx)
	for _, protocol := range list {
		if protocol.Id == id {
			return protocol
		}
	}
	return nil
}

func getConfig[T any](ctx context.Context, cfgName string, key string) (*T, string, error) {
	data := g.Cfg(cfgName).MustData(ctx)[key]
	if data == nil {
		return nil, cfgName, gerror.Newf("配置文件:%s中没有devices字段", cfgName)
	}
	// 解析
	config := new(T)
	err := gconv.Scan(data, config)
	if err != nil {
		return nil, cfgName, err
	}

	adapter := g.Cfg(cfgName).GetAdapter().(*gcfg.AdapterFile)
	path, err := adapter.GetFilePath()
	if err != nil {
		return nil, "", err
	}
	return config, path, err
}

func getConfigList[T any](ctx context.Context, cfgName string, key string) ([]*T, string, error) {
	data := g.Cfg(cfgName).MustData(ctx)[key]
	if data == nil {
		return nil, cfgName, gerror.Newf("配置文件:%s中没有devices字段", cfgName)
	}
	// 解析列表

	clientConfigList := make([]*T, len(data.([]interface{})))
	err := gconv.Scan(data, &clientConfigList)
	if err != nil {
		return nil, cfgName, err
	}

	adapter := g.Cfg(cfgName).GetAdapter().(*gcfg.AdapterFile)
	path, err := adapter.GetFilePath()
	if err != nil {
		return nil, "", err
	}
	return clientConfigList, path, err
}

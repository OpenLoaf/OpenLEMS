//go:build !dev && !windows
// +build !dev,!windows

package internal_cmd_device

import (
	"common"
	"common/c_base"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"strings"
)

func init() {
	//g.Log().Noticef(context.Background(), "当前环境为生产环境，从driver文件中获取驱动！")
}

// 生产环境下使用驱动加载的方式进行加载
func (d *SDeviceCmd) getDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) c_base.IDriver {
	if deviceConfig.Driver == "" {
		panic(gerror.Newf("设备[%s]%s驱动名称为空！", deviceConfig.Id, deviceConfig.Name))
	}

	// 获取最新的驱动文件路径
	latestDriverPath := fmt.Sprintf("%s.driver", deviceConfig.Driver)

	// 获取驱动的类型
	driverGroups := strings.Split(deviceConfig.Driver, "_")
	if driverGroups == nil || len(driverGroups) == 0 {
		panic(gerror.Newf("驱动名称错误！%s", deviceConfig.Driver))
	}

	ctx = context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, deviceConfig.Id)

	pluginSymbol, err := common.OpenPlugin(ctx, latestDriverPath)
	if err != nil {
		panic(gerror.Newf("加载插件[%s]失败！原因：%v", latestDriverPath, err))
	}

	var dv c_base.IDriver
	if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDriver); ok {
		dv = driverNewFunction(ctx)

		if dv.GetDriverType() != c_base.EDeviceType(driverGroups[0]) {
			panic(gerror.Newf("%s 中驱动类型不匹配！期望类型：%s, 实际类型：%s", deviceConfig.Driver, dv.GetDriverType(), driverGroups[0]))
		}
	} else {
		panic(gerror.Newf("加载插件[%s]失败！原因：未找到函数：%s", latestDriverPath, c_base.ConstNewPluginFunctionName))
	}

	return dv
}

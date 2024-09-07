//go:build dev || windows
// +build dev windows

package driver

import (
	"basic_v1/gpio_basic_v1"
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"pylonTechUs108_v1/bms_pylon_tech_us108_v1"
	"pylon_checkwatt_v1/ess_pylon_checkwatt_v1"
	"reflect"
	"starCharge100E_v1/pcs_star_charge_100E_v1"
	"station_energy_store/sess_basic_v1"
	"strings"
)

func init() {
	g.Log().Warningf(context.Background(), "当前环境为开发环境，直接加载驱动插件！而非从driver文件中获取！")
}

var pluginNewMethodCache = map[string]any{
	"bms_pylon_tech_us108_v1": bms_pylon_tech_us108_v1.NewPlugin,
	"ess_pylon_checkwatt_v1":  ess_pylon_checkwatt_v1.NewPlugin,
	"gpio_basic_v1":           gpio_basic_v1.NewPlugin,
	"pcs_star_charge_100E_v1": pcs_star_charge_100E_v1.NewPlugin,
	"sess_basic_v1":           sess_basic_v1.NewPlugin,
}

func (d *DeviceCmd) getDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) c_base.IDriver {
	if deviceConfig.Driver == "" {
		panic(gerror.Newf("设备[%s]%s驱动名称为空！", deviceConfig.Id, deviceConfig.Name))
	}

	ctx = context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, deviceConfig.Id)

	// 获取驱动的类型
	driverGroups := strings.Split(deviceConfig.Driver, "_")
	if driverGroups == nil || len(driverGroups) == 0 {
		panic(gerror.Newf("驱动名称错误！%s", deviceConfig.Driver))
	}

	pluginNewMethod := pluginNewMethodCache[deviceConfig.Driver]
	if pluginNewMethod == nil {
		panic(gerror.Newf("未找到驱动插件[%s]的NewPlugin方法！请检查pluginNewMethodCache或配置文件", deviceConfig.Driver))
	}

	// 准备参数
	args := []reflect.Value{reflect.ValueOf(ctx)}
	// 调用函数并获取结果
	results := reflect.ValueOf(pluginNewMethod).Call(args)

	if dv, ok := results[0].Interface().(c_base.IDriver); ok {
		if dv.GetDriverType() != c_base.EDeviceType(driverGroups[0]) {
			panic(gerror.Newf("%s 中驱动类型不匹配！期望类型：%s, 实际类型：%s", deviceConfig.Driver, dv.GetDriverType(), driverGroups[0]))
		}
		return dv
	} else {
		panic(gerror.Newf("加载插件[%s]失败！原因：未找到函数：%s", deviceConfig.Driver, c_base.ConstNewPluginFunctionName))
	}
}

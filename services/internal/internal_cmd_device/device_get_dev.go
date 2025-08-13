//go:build dev || windows
// +build dev windows

package internal_cmd_device

import (
	"ammeter_acrel_10r_v1/ammeter_acrel_10r_v1"
	"basic_v1/gpio_basic_v1"
	"bms_lnxall/bms_lnxall_v1"
	"common/c_base"
	"context"
	"ess_boost_gold/ess_boost_gold_v1"
	"ess_lnxall/ess_boost_lnxall_v1"
	"fire_control/fire_control_v1"
	"pcs_elecod/pcs_elecod_mac_v1"
	"pcs_enjoy/pcs_enjoy_basic_v1"
	"pcs_lnxall/pcs_lnxall_v1"
	"pylonTechUs108_v1/bms_pylon_tech_us108_v1"
	"pylon_checkwatt_v1/ess_pylon_checkwatt_v1"
	"reflect"
	"starCharge100E_v1/pcs_star_charge_100E_v1"
	"station_energy_store/sess_basic_v1"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	g.Log().Warningf(context.Background(), "当前环境为开发环境，直接加载驱动插件！而非从driver文件中获取！")
}

var pluginNewMethodCache = map[string]any{
	"bms_pylon_tech_us108_v1.0.0": bms_pylon_tech_us108_v1.NewPlugin,
	"ess_pylon_checkwatt_v1.0.0":  ess_pylon_checkwatt_v1.NewPlugin,
	"ess_boost_gold_v1.0.0":       ess_boost_gold_v1.NewPlugin,
	"ess_boost_lnxall_v1.0.0":     ess_boost_lnxall_v1.NewPlugin,
	"gpio_basic_v1.0.0":           gpio_basic_v1.NewPlugin,
	"pcs_star_charge_100E_v1.0.0": pcs_star_charge_100E_v1.NewPlugin,
	"pcs_enjoy_basic_v1.0.0":      pcs_enjoy_basic_v1.NewPlugin,
	"sess_basic_v1.0.0":           sess_basic_v1.NewPlugin,
	"ammeter_acrel_10r_v1.0.0":    ammeter_acrel_10r_v1.NewPlugin,
	"pcs_lnxall_v1.0.0":           pcs_lnxall_v1.NewPlugin,
	"bms_lnxall_v1.0.0":           bms_lnxall_v1.NewPlugin,
	"fire_control_v1.0.0":         fire_control_v1.NewPlugin,
	"pcs_elecod_mac_v1":           pcs_elecod_mac_v1.NewPlugin,
}

func GetAllDriverNames() []string {
	var driverNames []string
	for driverName := range pluginNewMethodCache {
		driverNames = append(driverNames, driverName)
	}
	return driverNames
}

// GetAllDriversInfo 获取所有驱动的详细信息
func GetAllDriversInfo(ctx context.Context) []c_base.DriverInfo {
	var driversInfo []c_base.DriverInfo

	for driverName, newMethod := range pluginNewMethodCache {
		driverInfo := c_base.DriverInfo{
			Name:      driverName,
			Available: newMethod != nil,
		}

		// 尝试创建驱动实例获取详细信息
		if newMethod != nil {
			func() {
				defer func() {
					if r := recover(); r != nil {
						// 如果创建驱动失败，标记为不可用
						driverInfo.Available = false
						g.Log().Warningf(ctx, "创建驱动[%s]失败: %v", driverName, r)
					}
				}()

				// 调用NewPlugin方法创建驱动实例
				args := []reflect.Value{reflect.ValueOf(ctx)}
				results := reflect.ValueOf(newMethod).Call(args)

				if len(results) > 0 {
					if driver, ok := results[0].Interface().(c_base.IDriver); ok {
						driverInfo.Type = driver.GetDriverType()
						driverInfo.Description = driver.GetDescription()
					}
				}
			}()
		}

		driversInfo = append(driversInfo, driverInfo)
	}

	return driversInfo
}

// GetDriverInfo 获取指定驱动的详细信息
func GetDriverInfo(ctx context.Context, driverName string) (*c_base.DriverInfo, error) {
	newMethod, exists := pluginNewMethodCache[driverName]
	if !exists {
		return nil, gerror.Newf("未找到驱动[%s]", driverName)
	}

	driverInfo := &c_base.DriverInfo{
		Name:      driverName,
		Available: newMethod != nil,
	}

	if newMethod != nil {
		defer func() {
			if r := recover(); r != nil {
				driverInfo.Available = false
				g.Log().Warningf(ctx, "创建驱动[%s]失败: %v", driverName, r)
			}
		}()

		// 调用NewPlugin方法创建驱动实例
		args := []reflect.Value{reflect.ValueOf(ctx)}
		results := reflect.ValueOf(newMethod).Call(args)

		if len(results) > 0 {
			if driver, ok := results[0].Interface().(c_base.IDriver); ok {
				driverInfo.Type = driver.GetDriverType()
				driverInfo.Description = driver.GetDescription()
			}
		}
	}

	return driverInfo, nil
}

// GetDriversByType 根据设备类型获取驱动信息
func GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.DriverInfo {
	allDrivers := GetAllDriversInfo(ctx)
	var filteredDrivers []c_base.DriverInfo

	for _, driver := range allDrivers {
		if driver.Type == deviceType {
			filteredDrivers = append(filteredDrivers, driver)
		}
	}

	return filteredDrivers
}

func (d *SDeviceCmd) getDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) c_base.IDriver {
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

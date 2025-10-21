package internal

import (
	"ammeter_demo/ammeter_demo_v1"
	"basic_gpio_in/gpio_in_basic_v1"
	"basic_gpio_out/gpio_out_basic_v1"
	"bms_pylon_tech_us108/bms_pylon_tech_us108_v1"
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"ess_demo/ess_demo_v1"
	"load_demo/load_demo_v1"
	"pcs_elecod/pcs_elecod_mac_v1"
	"pv_demo/pv_demo_v1"
	"pylon_checkwatt_v1/ess_pylon_checkwatt_v1"
	"reflect"
	"starCharge100E_v1/pcs_star_charge_100E_v1"
	"strings"

	"github.com/pkg/errors"
)

func init() {
	c_log.Warningf(context.Background(), "直接加载驱动插件！而非从driver文件中获取！")
}

var pluginNewMethodCache = map[string]any{
	"bms_pylon_tech_us108_v1.0.0": bms_pylon_tech_us108_v1.NewPlugin,
	"ess_pylon_checkwatt_v1.0.0":  ess_pylon_checkwatt_v1.NewPlugin,
	"pcs_star_charge_100E_v1.0.0": pcs_star_charge_100E_v1.NewPlugin,
	"ess_demo_v1.0.0":             ess_demo_v1.NewPlugin,
	"pv_demo_v1.0.0":              pv_demo_v1.NewPlugin,
	"load_demo_v1.0.0":            load_demo_v1.NewPlugin,
	"ammeter_demo_v1.0.0":         ammeter_demo_v1.NewPlugin,
	"pcs_elecod_mac_v1.0.0":       pcs_elecod_mac_v1.NewPlugin,
	"gpio_out_basic_v1.0.0":       gpio_out_basic_v1.NewPlugin,
	"gpio_in_basic_v1.0.0":        gpio_in_basic_v1.NewPlugin,
}
var pluginDriverInfo = map[string]*c_base.SDriverInfo{
	"bms_pylon_tech_us108_v1.0.0": bms_pylon_tech_us108_v1.GetDriverInfo(),
	"ess_pylon_checkwatt_v1.0.0":  ess_pylon_checkwatt_v1.GetDriverInfo(),
	"pcs_star_charge_100E_v1.0.0": pcs_star_charge_100E_v1.GetDriverInfo(),
	"ess_demo_v1.0.0":             ess_demo_v1.GetDriverInfo(),
	"pv_demo_v1.0.0":              pv_demo_v1.GetDriverInfo(),
	"load_demo_v1.0.0":            load_demo_v1.GetDriverInfo(),
	"ammeter_demo_v1.0.0":         ammeter_demo_v1.GetDriverInfo(),
	"pcs_elecod_mac_v1.0.0":       pcs_elecod_mac_v1.GetDriverInfo(),
	"gpio_out_basic_v1.0.0":       gpio_out_basic_v1.GetDriverInfo(),
	"gpio_in_basic_v1.0.0":        gpio_in_basic_v1.GetDriverInfo(),
}

func GetAllDriversInfo() map[string]*c_base.SDriverInfo {
	return pluginDriverInfo
}

// GetDriverInfo 获取指定驱动的详细信息
func GetDriverInfo(driverName string) (*c_base.SDriverInfo, error) {
	if driverInfo, ok := pluginDriverInfo[driverName]; ok {
		driverInfo.Name = driverName
		return driverInfo, nil
	}
	return nil, errors.Errorf("未找到驱动[%s]", driverName)
}

// GetDriversByType 根据设备类型获取驱动信息
func GetDriversByType(ctx context.Context, deviceType c_enum.EDeviceType) []*c_base.SDriverInfo {
	allDrivers := GetAllDriversInfo()
	var filteredDrivers []*c_base.SDriverInfo

	for _, driver := range allDrivers {
		if driver.Type == deviceType {
			filteredDrivers = append(filteredDrivers, driver)
		}
	}

	return filteredDrivers
}

func getDriver(driverName string, device c_base.IDevice) (d c_base.IDriver, err error) {

	if driverName == "" {
		return nil, errors.Errorf("驱动未设置")
	}
	// 获取驱动的类型
	driverGroups := strings.Split(driverName, "_")
	if driverGroups == nil || len(driverGroups) == 0 {
		return nil, errors.Errorf("驱动名称错误！%s", driverName)
	}

	// 检查GPIO协议是否在Linux系统上运行
	//deviceConfig := device.GetConfig()
	//if deviceConfig != nil && deviceConfig.ProtocolConfig != nil {
	//	protocolType := deviceConfig.ProtocolConfig.GetProtocol()
	//	if (protocolType == c_enum.EGpiod || protocolType == c_enum.EGpioSfs) && runtime.GOOS != "linux" {
	//		return nil, errors.Errorf("GPIO协议只能在Linux系统上使用，当前系统：%s，协议类型：%s", runtime.GOOS, protocolType)
	//	}
	//}

	pluginNewMethod := pluginNewMethodCache[driverName]
	if pluginNewMethod == nil {
		return nil, errors.Errorf("未找到驱动插件[%s]的NewPlugin方法！请检查pluginNewMethodCache或配置文件", driverName)
	}

	defer func() {
		if r := recover(); r != nil {
			err = errors.Errorf("执行[%s]驱动New方法失败! 原因请查看控制台日志", driverName)
			c_log.Errorf(context.Background(), "执行[%s]驱动New方法失败！原因：%s", driverName, r.(error).Error())
		}
	}()

	// 准备参数
	args := []reflect.Value{reflect.ValueOf(device)}
	// 调用函数并获取结果
	results := reflect.ValueOf(pluginNewMethod).Call(args)

	if dv, ok := results[0].Interface().(c_base.IDriver); ok {
		//if dv.GetDriverType() != c_enum.EDeviceType(driverGroups[0]) {
		//	panic(errors.Errorf("%s 中驱动类型不匹配！期望类型：%s, 实际类型：%s", driverName, dv.GetDriverType(), driverGroups[0]))
		//}
		return dv, nil
	} else {
		return nil, errors.Errorf("加载插件[%s]失败！原因：未找到函数：%s", driverName, c_base.ConstNewPluginFunctionName)
	}
}

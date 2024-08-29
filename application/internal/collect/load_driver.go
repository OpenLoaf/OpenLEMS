package collect

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/util"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"plug_protocol_modbus/p_modbus"
	"pylonTechUs108_v1/pylon_tech_us108"
	"starCharge100E_v1/star_charge_100e"
	"strings"
)

func loadDriver(ctx context.Context, deviceConfig *p_modbus.SModbusDeviceConfig) c_base.IDriver {
	// 先加载驱动
	latestDriverPath, i, err := c_base.GetLatestDriver(ctx, deviceConfig.Driver)

	if err != nil {
		//return nil, fmt.Errorf("get latest driver err: %v", err)
		panic(fmt.Sprintf("获取驱动[%s]失败: %v", deviceConfig.Name, err))
	}
	g.Log().Infof(ctx, "设备加载驱动：%v 最新版本为v%v", deviceConfig.Driver, i)

	// 获取驱动的类型
	driverGroups := strings.Split(deviceConfig.Driver, "_")
	if driverGroups == nil || len(driverGroups) == 0 {
		panic(fmt.Sprintf("驱动名称错误！%s", deviceConfig.Driver))
	}
	// 用驱动的名称中的组类型来创建

	var dv c_base.IDriver

	var _group = c_base.EDeviceType(driverGroups[0])
	switch _group {

	case c_base.EDeviceAmmeter:
		dv = getDriver[c_device.IAmmeter](ctx, latestDriverPath)

		if deviceConfig.CabinetId != 0 {
			groupType := c_base.EGroupType(_group)

			_tempInstanceCache.Ammeters[groupType] = append(_tempInstanceCache.Ammeters[groupType], dv.(c_device.IAmmeter))
		} else {
			_tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId).Ammeter = dv.(c_device.IAmmeter)
		}

	case c_base.EDevicePcs:
		//dv = getDriver[c_device.IPcs](ctx, latestDriverPath)
		dv = &star_charge_100e.StarCharge100EPcs{}
		ess := _tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId)
		ess.Pcs = append(ess.Pcs, dv.(c_device.IPcs))
	case c_base.EDeviceBms:

		// TODO: 改成插件加载
		dv = &pylon_tech_us108.PylonTechUs108Bms{}
		//dv = getDriver[c_device.IBms](ctx, latestDriverPath)
		_tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId).Bms = dv.(c_device.IBms)
	case c_base.EDeviceFire:
		dv = getDriver[c_device.IFire](ctx, latestDriverPath)
		_tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId).Fire = dv.(c_device.IFire)
	case c_base.EDeviceEnergyStore:
		dv = getDriver[c_device.IEnergyStore](ctx, latestDriverPath)
		_tempInstanceCache.Ess = append(_tempInstanceCache.Ess, dv.(c_device.IEnergyStore))
	case c_base.EDeviceHumiture:
		dv = getDriver[c_device.IHumiture](ctx, latestDriverPath)
		_tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId).Humiture = dv.(c_device.IHumiture)
	case c_base.EDevicePv:
		dv = getDriver[c_device.IPv](ctx, latestDriverPath)
		_tempInstanceCache.Pv = append(_tempInstanceCache.Pv, dv.(c_device.IPv))
	case c_base.EDeviceCoolingAc:
		dv = getDriver[c_device.ICoolingAc](ctx, latestDriverPath)
		_tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId).Cooling = dv.(c_device.ICoolingBasic)
	case c_base.EDeviceCoolingLiquid:
		dv = getDriver[c_device.ICoolingLiquid](ctx, latestDriverPath)
		_tempInstanceCache.GetCabinetEss(deviceConfig.CabinetId).Cooling = dv.(c_device.ICoolingBasic)
	case c_base.EDeviceLoad:
		dv = getDriver[c_device.ILoad](ctx, latestDriverPath)
		_tempInstanceCache.Load = append(_tempInstanceCache.Load, dv.(c_device.ILoad))
	case c_base.EChargePile:
		dv = getDriver[c_device.ICharge](ctx, latestDriverPath)
		_tempInstanceCache.ChargePile = append(_tempInstanceCache.ChargePile, dv.(c_device.ILoad))
	case c_base.EGenerator:
		dv = getDriver[c_device.IGenerator](ctx, latestDriverPath)
		_tempInstanceCache.Generator = append(_tempInstanceCache.Generator, dv.(c_device.IGenerator))
	}

	if dv == nil {
		panic(fmt.Sprintf("加载驱动失败！配置为:%+v", deviceConfig))
	}
	return dv
}

func getDriver[T c_base.IDriver](ctx context.Context, path string) T {

	dv, err := util.OpenPlugin(ctx, c_base.GetSystemConfig().DriverPath+"/"+path)
	if err != nil {
		panic(fmt.Sprintf("加载插件[%s]失败！原因：%v", path, err))
	}

	driverNewFunction := dv.(func() (T, error))
	instance, err := driverNewFunction()
	if err != nil {
		panic(fmt.Sprintf("加载插件[%s]失败！原因：%v", path, err))
	}
	return instance
}

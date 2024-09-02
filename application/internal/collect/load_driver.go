package collect

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/util"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"pylonTechUs108_v1/pylon_tech_us108"
	"starCharge100E_v1/star_charge_100e"
	"strings"
)

func loadDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) c_base.IDriver {
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

		//if deviceConfig.CabinetId != 0 {
		//	groupType := c_base.EStationType(_group)
		//
		//	_tempInstanceCache.Ammeters[groupType] = append(_tempInstanceCache.Ammeters[groupType], dv.(c_device.IAmmeter))
		//}

	case c_base.EDevicePcs:
		//dv = getDriver[c_device.IPcs](ctx, latestDriverPath)
		dv = &star_charge_100e.StarCharge100EPcs{
			Ctx: ctx,
		}
	case c_base.EDeviceBms:

		// TODO: 改成插件加载
		dv = pylon_tech_us108.NewPlugin(ctx)
	case c_base.EDeviceFire:
		dv = getDriver[c_device.IFire](ctx, latestDriverPath)
	case c_base.EDeviceEnergyStore:
		dv = getDriver[c_device.IEnergyStore](ctx, latestDriverPath)
	case c_base.EDeviceHumiture:
		dv = getDriver[c_device.IHumiture](ctx, latestDriverPath)
	case c_base.EDevicePv:
		dv = getDriver[c_device.IPv](ctx, latestDriverPath)
	case c_base.EDeviceCoolingAc:
		dv = getDriver[c_device.ICoolingAc](ctx, latestDriverPath)
	case c_base.EDeviceCoolingLiquid:
		dv = getDriver[c_device.ICoolingLiquid](ctx, latestDriverPath)
	case c_base.EDeviceLoad:
		dv = getDriver[c_device.ILoad](ctx, latestDriverPath)
	case c_base.EDeviceChargePile:
		dv = getDriver[c_device.ICharge](ctx, latestDriverPath)
	case c_base.EDeviceGenerator:
		dv = getDriver[c_device.IGenerator](ctx, latestDriverPath)
	}

	if dv == nil {
		panic(fmt.Sprintf("加载驱动失败！配置为:%+v", deviceConfig))
	}

	//if deviceConfig.CabinetId != 0 {
	//	_tempInstanceCache.AddCabinetDevice(deviceConfig.CabinetId, dv)
	//}

	//deviceConfig.CheckTypeIs(dv.GetDriverType())
	deviceConfig.Type = dv.GetDriverType()

	return dv
}

func getDriver[T c_base.IDriver](ctx context.Context, path string) T {

	dv, err := util.OpenPlugin(ctx, c_base.GetSystemConfig().DriverPath+"/"+path)
	if err != nil {
		panic(fmt.Sprintf("加载插件[%s]失败！原因：%v", path, err))
	}

	driverNewFunction := dv.(func(ctx context.Context) (T, error))
	instance, err := driverNewFunction(ctx)
	if err != nil {
		panic(fmt.Sprintf("加载插件[%s]失败！原因：%v", path, err))
	}
	return instance
}

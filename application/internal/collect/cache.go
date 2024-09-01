package collect

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"pylon_checkwatt_v1/pylon_checkwatt"
)

type tmpStation struct {
	Ammeters map[c_base.EStationType][]c_device.IAmmeter

	Pv          []c_device.IPv
	Load        []c_device.ILoad
	Ess         []c_device.IEnergyStore
	EssGpioList []c_device.IGpio

	ChargePile []c_device.ICharge
	Generator  []c_device.IGenerator

	cabinetEss map[uint8][]c_base.IDriver
}

var _tempInstanceCache = &tmpStation{
	cabinetEss: make(map[uint8][]c_base.IDriver),
}

func (t *tmpStation) AddCabinetDevice(cabinetId uint8, driver c_base.IDriver) {
	if cabinetId == 0 {
		panic("cabinetId 不能为0")
	}
	cabinetDrivers := t.cabinetEss[cabinetId]
	if cabinetDrivers == nil {
		cabinetDrivers = []c_base.IDriver{driver}
		t.cabinetEss[cabinetId] = cabinetDrivers
		return
	}
	t.cabinetEss[cabinetId] = append(cabinetDrivers, driver)
}

func (t *tmpStation) Init(ctx context.Context) {

	for cabinetId, drivers := range t.cabinetEss {

		ess, err := pylon_checkwatt.NewEss(ctx, cabinetId, drivers, listToMap[c_device.IGpio](drivers))
		if err != nil {
			panic(err)
		}
		//ess.Init(nil, nil)

		t.Ess = append(t.Ess, ess)
	}

	/*	if ammeters, exist := t.Ammeters[c_base.EStationEnergyStore]; exist || len(t.Ess) != 0 {
			// 场站储能
			if len(ammeters) == 0 {
				common_station.RegisterGroupEnergyStore(ctx, nil, t.Ess, t.EssGpioList)
			} else if len(ammeters) == 1 {
				common_station.RegisterGroupEnergyStore(ctx, ammeters[0], t.Ess, t.EssGpioList)
			} else {
				panic("场站储能只能有一个根电表")
			}
		}

		if ammeters, exist := t.Ammeters[c_base.EStationEntrance]; exist {
			// 场站总入口
			if len(ammeters) == 0 {
				panic("场站总入口必须有一个根电表")
			}
			root, slaves := getMasterAndList[c_device.IAmmeter](ammeters)

			common_station.RegisterEntrance(ctx, root, slaves)
		}
	*/
}

func listToMap[T c_base.IDriver](drivers []c_base.IDriver) map[string]T {
	_map := make(map[string]T)
	for _, driver := range drivers {
		if input, ok := driver.(T); ok {
			_map[input.GetDeviceConfig().Id] = input
		}
	}
	return _map
}

func getMasterAndList[T c_base.IDriver](list []T) (master T, slaves []T) {
	for _, v := range list {
		if v.GetDeviceConfig().IsMaster {
			master = v
		} else {
			slaves = append(slaves, v)
		}
	}
	return master, slaves

}

package common_station

import (
	"common_station/internal/internal_energy_store"
	"common_station/internal/internal_entrance"
	"context"
	"ems-plan/c_device"
)

// RegisterGroupEnergyStore 创建储能组
func RegisterGroupEnergyStore(ctx context.Context, rootAmmeter c_device.IAmmeter, energyStores []c_device.IEnergyStore, gpios []c_device.IGpio) {
	StationInstance.RegisterInstance(internal_energy_store.NewGroupEnergyStore(ctx, rootAmmeter, energyStores, gpios))
}

// RegisterEntrance 创建场站总能量使用情况
func RegisterEntrance(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter) {
	StationInstance.RegisterInstance(internal_entrance.NewEntrance(ctx, rootAmmeter, ammeters))
}

//
//// RegisterLoad 创建负荷组
//func RegisterLoad(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootLoad c_base.ILoad, loads []c_base.ILoad) {
//	StationInstance.RegisterInstance(internal_group.NewLoad(ctx, rootAmmeter, ammeters, rootLoad, loads))
//}
//
//// RegisterPv 创建光伏组
//func RegisterPv(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootPv c_base.IPv, pvs []c_base.IPv) {
//	StationInstance.RegisterInstance(internal_group.NewPv(ctx, rootAmmeter, ammeters, rootPv, pvs))
//}
//
//// RegisterGenerator 创建发电机组
//func RegisterGenerator(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootGenerator c_base.IGenerator, generators []c_base.IGenerator) {
//	StationInstance.RegisterInstance(internal_group.NewGenerator(ctx, rootAmmeter, ammeters, rootGenerator, generators))
//}

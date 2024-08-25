package common_group

import (
	"common_group/internal/internal_group"
	"context"
	"ems-plan/c_device"
)

// RegisterGroupEnergyStore 创建储能组
func RegisterGroupEnergyStore(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, energyStores []c_device.IEnergyStore) {
	GroupInstance.RegisterInstance(internal_group.NewGroupEnergyStore(ctx, rootAmmeter, ammeters, energyStores))
}

// RegisterEntrance 创建场站总能量使用情况
func RegisterEntrance(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter) {
	GroupInstance.RegisterInstance(internal_group.NewEntrance(ctx, rootAmmeter, ammeters))
}

// RegisterLoad 创建负荷组
func RegisterLoad(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootLoad c_device.ILoad, loads []c_device.ILoad) {
	GroupInstance.RegisterInstance(internal_group.NewLoad(ctx, rootAmmeter, ammeters, rootLoad, loads))
}

// RegisterPv 创建光伏组
func RegisterPv(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootPv c_device.IPv, pvs []c_device.IPv) {
	GroupInstance.RegisterInstance(internal_group.NewPv(ctx, rootAmmeter, ammeters, rootPv, pvs))
}

// RegisterGenerator 创建发电机组
func RegisterGenerator(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootGenerator c_device.IGenerator, generators []c_device.IGenerator) {
	GroupInstance.RegisterInstance(internal_group.NewGenerator(ctx, rootAmmeter, ammeters, rootGenerator, generators))
}

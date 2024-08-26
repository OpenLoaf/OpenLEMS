package common_station

import (
	"common_group/internal/internal_group"
	"context"
	"ems-plan/c_device"
)

// RegisterGroupEnergyStore 创建储能组
func RegisterGroupEnergyStore(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter, energyStores []c_base.IEnergyStore) {
	StationInstance.RegisterInstance(internal_group.NewGroupEnergyStore(ctx, rootAmmeter, ammeters, energyStores))
}

// RegisterEntrance 创建场站总能量使用情况
func RegisterEntrance(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter) {
	StationInstance.RegisterInstance(internal_group.NewEntrance(ctx, rootAmmeter, ammeters))
}

// RegisterLoad 创建负荷组
func RegisterLoad(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter, rootLoad c_base.ILoad, loads []c_base.ILoad) {
	StationInstance.RegisterInstance(internal_group.NewLoad(ctx, rootAmmeter, ammeters, rootLoad, loads))
}

// RegisterPv 创建光伏组
func RegisterPv(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter, rootPv c_base.IPv, pvs []c_base.IPv) {
	StationInstance.RegisterInstance(internal_group.NewPv(ctx, rootAmmeter, ammeters, rootPv, pvs))
}

// RegisterGenerator 创建发电机组
func RegisterGenerator(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter, rootGenerator c_base.IGenerator, generators []c_base.IGenerator) {
	StationInstance.RegisterInstance(internal_group.NewGenerator(ctx, rootAmmeter, ammeters, rootGenerator, generators))
}

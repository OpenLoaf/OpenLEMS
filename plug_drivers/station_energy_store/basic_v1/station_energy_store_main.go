package main

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"station_energy_store/station_energy_store"
)

func NewPlugin(ctx context.Context, deviceConfig *c_base.SDriverConfig, drivers []c_base.IDriver) c_device.IStationEnergyStore {
	return station_energy_store.NewGroupEnergyStore(ctx, deviceConfig, drivers)
}

func main() {

}

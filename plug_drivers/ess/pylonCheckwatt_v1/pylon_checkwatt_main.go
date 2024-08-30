//go:generate ../build.sh
package main

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"pylon_checkwatt_v1/pylon_checkwatt"
)

// NewCabinetPlugin 必须的方法，不能取消
func NewCabinetPlugin(ctx context.Context, cabinetId uint8, drivers []c_base.IDriver) (c_device.IEnergyStore, error) {
	return pylon_checkwatt.NewEss(ctx, cabinetId, drivers)
}

func main() {

}

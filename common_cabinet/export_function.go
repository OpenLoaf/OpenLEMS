package common_cabinet

import (
	"common_cabinet/internal/internal_bms"
	"common_cabinet/internal/internal_pcs"
	"context"
	"ems-plan/c_device"
)

func NewPcs(ctx context.Context, cabinetId uint8, master c_device.IPcs, pcsList []c_device.IPcs) c_device.IPcsBasic {
	return internal_pcs.NewPcs(ctx, cabinetId, master, pcsList)
}

func NewBms(ctx context.Context, cabinetId uint8, bms c_device.IBms) c_device.IBmsBasic {
	return internal_bms.NewBms(ctx, cabinetId, bms)
}

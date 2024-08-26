package common_cabinet

import (
	"common_cabinet/internal/internal_bms"
	"common_cabinet/internal/internal_pcs"
	"context"
	"ems-plan/c_device"
)

func NewPcs(ctx context.Context, cabinetId uint8, master c_base.IPcs, pcsList []c_base.IPcs) c_base.IPcsBasic {
	return internal_pcs.NewPcs(ctx, cabinetId, master, pcsList)
}

func NewBms(ctx context.Context, cabinetId uint8, bms c_base.IBms) c_base.IBmsBasic {
	return internal_bms.NewBms(ctx, cabinetId, bms)
}

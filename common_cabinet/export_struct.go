package common_cabinet

import (
	"common_cabinet/internal/cb_bms"
	"common_cabinet/internal/cb_cooling"
	"common_cabinet/internal/cb_fire"
	"common_cabinet/internal/cb_humiture"
	"common_cabinet/internal/cb_pcs"
)

type (
	CabinetBms      = cb_bms.CabinetBms
	CabinetCooling  = cb_cooling.CabinetCooling
	CabinetFire     = cb_fire.CabinetFire
	CabinetPcs      = cb_pcs.CabinetPcs
	CabinetHumidity = cb_humiture.CabinetHumiture
)

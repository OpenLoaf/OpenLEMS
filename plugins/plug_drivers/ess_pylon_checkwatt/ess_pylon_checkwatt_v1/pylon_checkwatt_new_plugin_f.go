package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_none"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sEssPylonCheckwatt{
		SAlarmHandler:      &c_base.SAlarmHandler{},
		ctx:                ctx,
		pcs:                c_none.NonePcs,
		bms:                c_none.NoneBms,
		SDriverDescription: c_base.BuildDescriptionFromYaml(buildYaml),
	}
}

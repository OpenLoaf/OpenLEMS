package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sPylonCheckwattEss{
		SAlarmHandler:      &c_base.SAlarmHandler{},
		ctx:                ctx,
		SDriverDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

package ess_pylon_checkwatt_v1

import (
	"context"
	_ "embed"
	"ems-plan/c_base"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &PylonCheckwattEss{
		SAlarmHandler: &c_base.SAlarmHandler{},
		ctx:           ctx,
		SDescription:  c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

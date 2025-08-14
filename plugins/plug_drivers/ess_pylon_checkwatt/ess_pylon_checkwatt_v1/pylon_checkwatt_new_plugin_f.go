package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &sPylonCheckwattEss{
		SAlarmHandler: &c_base.SAlarmHandler{},
		ctx:           ctx,
		SDescription:  c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

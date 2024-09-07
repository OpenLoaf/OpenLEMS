package bms_pylon_tech_us108_v1

import (
	"context"
	_ "embed"
	"ems-plan/c_base"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &PylonTechUs108Bms{
		ctx:          ctx,
		SDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

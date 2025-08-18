package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sBmsPylonTechUs108{
		ctx:                ctx,
		SDriverDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

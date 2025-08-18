package ammeter_acrel_10r_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sAmmeterAcrel10r{
		ctx:                ctx,
		SDriverDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

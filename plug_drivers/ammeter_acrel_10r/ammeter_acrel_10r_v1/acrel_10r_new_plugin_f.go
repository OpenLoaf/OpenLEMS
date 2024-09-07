package ammeter_acrel_10r_v1

import (
	"context"
	_ "embed"
	"ems-plan/c_base"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &sAmmeterAcrel10r{
		ctx:          ctx,
		SDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

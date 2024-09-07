package pcs_star_charge_100E_v1

import (
	"context"
	_ "embed"
	"ems-plan/c_base"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &sStarCharge100EPcs{
		ctx:          ctx,
		SDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

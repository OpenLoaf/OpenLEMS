package pcs_star_charge_100E_v1

import (
	"common/c_base"
	"common/c_none"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sPcsStarCharge100E{
		ctx:                ctx,
		IModbusProtocol:    c_none.NoneProtocol,
		SDriverDescription: c_base.BuildDescriptionFromYaml(buildYaml),
	}
}

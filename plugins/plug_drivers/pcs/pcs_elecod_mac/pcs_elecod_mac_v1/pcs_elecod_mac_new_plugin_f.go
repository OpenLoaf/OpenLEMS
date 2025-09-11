package pcs_elecod_mac_v1

import (
	"common/c_base"
	"common/c_none"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sPcsElecodMac{
		ctx:                ctx,
		ICanbusProtocol:    c_none.NoneProtocol,
		SDriverDescription: c_base.BuildDescriptionFromYaml(buildYaml),
	}
}

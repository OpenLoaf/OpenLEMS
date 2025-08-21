package pcs_elecod_mac_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sPcsElecodMac{
		ctx:                ctx,
		SDriverDescription: c_base.BuildDescriptionFromYaml(buildYaml),
	}
}

package pcs_elecod_mac_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &sPcsElecodBasic{
		ctx:          ctx,
		SDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

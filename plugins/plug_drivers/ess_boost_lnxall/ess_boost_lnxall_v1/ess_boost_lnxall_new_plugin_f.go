package ess_boost_lnxall_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDevice {
	return &sEssBoostLnxallEss{
		//SAlarmHandler: &c_base.SAlarmHandler{},
		ctx:                ctx,
		SDriverDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

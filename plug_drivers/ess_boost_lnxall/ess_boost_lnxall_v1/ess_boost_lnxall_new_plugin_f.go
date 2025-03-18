package ess_boost_lnxall_v1

import (
	"common/c_base"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_base.IDriver {
	return &sEssBoostLnxallEss{
		//SAlarmHandler: &c_base.SAlarmHandler{},
		ctx:          ctx,
		SDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}
}

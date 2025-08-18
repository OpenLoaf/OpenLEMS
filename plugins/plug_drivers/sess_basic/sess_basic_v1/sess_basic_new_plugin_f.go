package sess_basic_v1

import (
	"common/c_base"
	"common/c_device"
	"context"
	_ "embed"
)

//go:embed build.yaml
var buildYaml []byte

func NewPlugin(ctx context.Context) c_device.IStationEnergyStore {
	instance := &sStationEnergyStore{
		SAlarmHandler:      &c_base.SAlarmHandler{},
		ctx:                ctx,
		SDriverDescription: c_base.BuildDescriptionFromYaml(ctx, buildYaml),
	}

	return instance
}

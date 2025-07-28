package p_canbus

import (
	"common/c_base"
	"context"
)

type ICanbusProtocol interface {
	c_base.IProtocol

	RegisterRead(ctx context.Context, group *SCanbusTask, gs ...*SCanbusTask)

	GetBool(meta *c_base.Meta) (bool, error)
	GetIntValue(meta *c_base.Meta) (int, error)
	GetInt32Value(meta *c_base.Meta) (int32, error)
	GetUintValue(meta *c_base.Meta) (uint, error)
	GetUint32Value(meta *c_base.Meta) (uint32, error)
	GetFloat32Value(meta *c_base.Meta) (float32, error)
	GetFloat32Values(metas ...*c_base.Meta) ([]float32, error)
	GetFloat64Value(meta *c_base.Meta) (float64, error)
	GetFloat64Values(meta ...*c_base.Meta) ([]float64, error)
	GetCanbusDeviceConfig() *SCanbusDeviceConfig
}

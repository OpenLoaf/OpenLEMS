package c_device

import (
	"common/c_base"
	"common/c_proto"
	"context"
)

type SRealGpio struct {
	c_proto.IGpiodProtocol

	DeviceCtx context.Context
	cancel    context.CancelFunc
}

var _ c_base.IDevice = (*SRealGpio)(nil)

func NewRealGpio(ctx context.Context, protocol c_proto.IGpiodProtocol) (*SRealGpio, error) {
	deviceCtx, cancel := context.WithCancel(ctx)

	device := &SRealGpio{
		DeviceCtx:      deviceCtx,
		cancel:         cancel,
		IGpiodProtocol: protocol,
	}

	return device, nil
}

func (s *SRealGpio) IsVirtualDevice() bool {
	return false
}

// 实现新的IDevice接口方法 - GPIO设备默认实现
func (s *SRealGpio) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{}
}

package c_device

import (
	"common/c_base"
	"common/c_proto"
	"context"

	"github.com/pkg/errors"
)

type SRealGpio struct {
	c_proto.IGpiodProtocol

	DeviceCtx    context.Context
	cancel       context.CancelFunc
	deviceConfig *c_base.SDeviceConfig // 配置
}

var _ c_base.IDevice = (*SRealGpio)(nil)

func NewRealGpio(ctx context.Context, deviceConfig *c_base.SDeviceConfig) (*SRealGpio, error) {
	if deviceConfig == nil {
		// 必须有设备配置
		panic(errors.New("deviceConfig is nil"))
	}
	deviceCtx, cancel := context.WithCancel(ctx)

	device := &SRealGpio{
		DeviceCtx:    deviceCtx,
		cancel:       cancel,
		deviceConfig: deviceConfig,
	}

	return device, nil
}

func (s *SRealGpio) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *SRealGpio) IsVirtualDevice() bool {
	return false
}

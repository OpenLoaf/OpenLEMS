package canbus

import (
	"canbus/internal"
	"canbus/p_canbus"
	"common/c_base"
	"context"
	"go.einride.tech/can"
)

func NewCanbusChan(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (<-chan can.Frame, chan<- can.Frame, error) {
	return internal.NewCanbusConnect(ctx, protocolConfig)
}

func NewCanbusProvider(ctx context.Context, deviceType c_base.EDeviceType, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, receiverChan <-chan can.Frame, transmitterChan chan<- can.Frame) (p_canbus.ICanbusProtocol, error) {
	return internal.NewCanbusProvider(ctx, deviceType, clientConfig, deviceConfig, receiverChan, transmitterChan)
}

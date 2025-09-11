package canbus

import (
	"common/c_base"
	"common/c_enum"
	"common/c_proto"
	"context"
	"p_canbus/internal"

	"go.einride.tech/can"
)

func NewCanbusChan(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (<-chan can.Frame, chan<- can.Frame, error) {
	return internal.NewCanbusConnect(ctx, protocolConfig)
}

func NewCanbusProvider(ctx context.Context, deviceType c_enum.EDeviceType, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig,
	receiverChan <-chan can.Frame, transmitterChan chan<- can.Frame) (c_proto.ICanbusProtocol, error) {
	return internal.NewCanbusProvider(ctx, clientConfig, deviceConfig, receiverChan, transmitterChan)
}

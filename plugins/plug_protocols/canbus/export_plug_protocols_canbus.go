package canbus

import (
	"canbus/internal"
	"canbus/p_canbus"
	"common/c_base"
	"context"
	"go.einride.tech/can"
)

func NewCanbusChan(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (<-chan can.Frame, chan<- can.Frame) {
	return internal.NewCanbusConnect(ctx, protocolConfig)
}

func NewCanbusProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig, receiverChan <-chan can.Frame, transmitterChan chan<- can.Frame) (p_canbus.ICanbusProtocol, error) {
	return internal.NewCanbusProvider(ctx, clientConfig, deviceConfig, receiverChan, transmitterChan)
}

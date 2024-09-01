package plug_protocol_modbus

import (
	"context"
	"ems-plan/c_base"
	"github.com/torykit/go-modbus"
	"plug_protocol_modbus/internal"
	"plug_protocol_modbus/p_modbus"
)

func NewModbusClient(ctx context.Context, protocolConfig *c_base.SProtocolConfig) modbus.Client {
	return internal.NewModbusClient(ctx, protocolConfig)
}

func NewModbusProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig, client any) (p_modbus.IModbusProtocol, error) {
	return internal.NewModbusProvider(ctx, clientConfig, deviceConfig, client)
}

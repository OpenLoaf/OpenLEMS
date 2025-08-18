package modbus

import (
	"common/c_base"
	"common/c_modbus"
	"context"
	"github.com/torykit/go-modbus"
	"modbus/internal"
)

// NewModbusClient 一个协议一个client
func NewModbusClient(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (modbus.Client, error) {
	return internal.NewModbusClient(ctx, protocolConfig)
}

// NewModbusProvider 一个设备一个provider
func NewModbusProvider(ctx context.Context, deviceType c_base.EDeviceType, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, client any) (c_modbus.IModbusProtocol, error) {
	return internal.NewModbusProvider(ctx, deviceType, clientConfig, deviceConfig, client)
}

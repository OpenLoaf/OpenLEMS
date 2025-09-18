package p_modbus

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"context"
	"p_modbus/internal"

	"github.com/torykit/go-modbus"
)

// NewModbusClient 一个协议一个client
func NewModbusClient(ctx context.Context, protocolConfig *c_base.SProtocolConfig) (modbus.Client, error) {
	return internal.NewModbusClient(ctx, protocolConfig)
}

// NewModbusProvider 一个设备一个provider
func NewModbusProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, client any) (c_proto.IModbusProtocol, error) {
	return internal.NewModbusProvider(ctx, clientConfig, deviceConfig, client)
}

// GetModbusDeviceConfigFields 获取modbus的设备配置
func GetModbusDeviceConfigFields() []*c_base.SFieldDefinition {
	modbusDeviceConfig := &c_proto.SModbusDeviceConfig{}
	fields, err := c_base.BuildConfigStructFields(modbusDeviceConfig)
	if err != nil {
		c_log.BizErrorf(context.Background(), "解析Modbus配置信息结构失败！")
		return nil
	}
	return fields
}

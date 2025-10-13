// =================================================================================
// Key generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package remote

import (
	"context"

	v1 "application/api/remote/v1"
)

type IRemoteV1 interface {
	GetMqttStatus(ctx context.Context, req *v1.GetMqttStatusReq) (res *v1.GetMqttStatusRes, err error)
	ReloadMqtt(ctx context.Context, req *v1.ReloadMqttReq) (res *v1.ReloadMqttRes, err error)
	MqttSwitch(ctx context.Context, req *v1.MqttSwitchReq) (res *v1.MqttSwitchRes, err error)
	GetModbusStatus(ctx context.Context, req *v1.GetModbusStatusReq) (res *v1.GetModbusStatusRes, err error)
	ReloadModbus(ctx context.Context, req *v1.ReloadModbusReq) (res *v1.ReloadModbusRes, err error)
	ModbusSwitch(ctx context.Context, req *v1.ModbusSwitchReq) (res *v1.ModbusSwitchRes, err error)
}

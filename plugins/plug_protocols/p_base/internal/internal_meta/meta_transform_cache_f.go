package internal_meta

import (
	"common/c_base"
)

func MetaProcess(meta *c_base.Meta, value any) any {
	return SystemTypeTransform(meta.SystemType, value, meta.ReadType, meta.BitLength, meta.Factor, meta.Offset)
}

func processAlarm(protocol c_base.IProtocol, deviceId string, deviceType c_base.EDeviceType, meta *c_base.Meta, IsTrigger bool, value any) {
	// TODO 触发告警
	//now := time.Now()
	//protocol.TriggerAlarm(&c_base.SAlarmDetail{
	//	DeviceId:   deviceId,
	//	DeviceType: deviceType,
	//	Level:      meta.Level,
	//	Meta:       meta,
	//	IsTrigger:  IsTrigger,
	//	HappenTime: &now,
	//	Value:      value,
	//})
}

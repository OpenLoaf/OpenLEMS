package internal_meta

import (
	"common/c_base"
	"time"
)

func MetaProcess(meta *c_base.Meta, value any) any {
	return SystemTypeTransform(meta.SystemType, value, meta.ReadType, meta.BitLength, meta.Factor, meta.Offset)
}

func processAlarm(protocol c_base.IProtocol, deviceId string, deviceType c_base.EDeviceType, meta *c_base.Meta, IsTrigger bool, value any) {
	protocol.TriggerAlarm(&c_base.SAlarmDetail{
		DeviceId:   deviceId,
		DeviceType: deviceType,
		Level:      meta.Level,
		Meta:       meta,
		IsTrigger:  IsTrigger,
		HappenTime: time.Now(),
		Value:      value,
	})
}

package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetDeviceTelemetry(ctx context.Context, req *v1.GetDeviceTelemetryReq) (res *v1.GetDeviceTelemetryRes, err error) {
	telemetry := make(map[string]*v1.DeviceTelemetryData)
	alarmLevelMap := make(map[string]string)

	common.GetDeviceManager().IteratorAssAllDevicesWrapper(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		deviceId := config.Id
		if deviceId == "" || config.DriverInfo == nil {
			return true
		}

		if device == nil {
			return true
		}

		telemetry[deviceId] = &v1.DeviceTelemetryData{
			LastUpdateTime:  device.GetLastUpdateTime(),
			TelemetryValues: config.DriverInfo.GetAllTelemetry(device),
		}

		alarmLevelMap[deviceId] = device.GetAlarmLevel().String()

		return true
	})

	return &v1.GetDeviceTelemetryRes{
		AlarmLevelMap: alarmLevelMap,
		Telemetry:     telemetry,
	}, nil
}

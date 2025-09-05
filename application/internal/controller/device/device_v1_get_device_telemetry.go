package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"context"
)

func (c *ControllerV1) GetDeviceTelemetry(ctx context.Context, req *v1.GetDeviceTelemetryReq) (res *v1.GetDeviceTelemetryRes, err error) {

	if common.GetDeviceManager().Status() == c_base.EStateInit {
		// 系统还在初始化中
		return nil, nil
	}

	telemetry := make(map[string]*v1.DeviceTelemetryData)
	alarmLevelMap := make(map[string]string)
	protocolStatus := make(map[string]string)

	common.GetDeviceManager().IteratorChildDevicesById(req.DeviceId, func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		deviceId := config.Id
		if device == nil {
			return true
		}

		if values := config.DriverInfo.GetAllTelemetry(device); len(values) > 0 {
			telemetry[deviceId] = &v1.DeviceTelemetryData{
				LastUpdateTime:  device.GetLastUpdateTime(),
				TelemetryValues: values,
			}
		}

		alarmLevelMap[deviceId] = device.GetAlarmLevel().String()
		protocolStatus[deviceId] = device.GetProtocolStatus().String()

		return true
	})

	return &v1.GetDeviceTelemetryRes{
		ProtocolStatus: protocolStatus,
		AlarmLevelMap:  alarmLevelMap,
		Telemetry:      telemetry,
	}, nil
}

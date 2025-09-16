package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
	"sync"
)

func (c *ControllerV1) GetDeviceTelemetry(ctx context.Context, req *v1.GetDeviceTelemetryReq) (res *v1.GetDeviceTelemetryRes, err error) {

	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		// 系统还在初始化中
		return nil, nil
	}

	// 使用线程安全的 map 来存储结果
	var (
		telemetry      = make(map[string]*v1.DeviceTelemetryData)
		alarmLevelMap  = make(map[string]string)
		protocolStatus = make(map[string]string)
		mu             sync.RWMutex
		wg             sync.WaitGroup
	)

	// 收集所有设备配置，用于并行处理
	var deviceConfigs []*c_base.SDeviceConfig
	common.GetDeviceManager().IteratorChildDevicesById(req.DeviceId, func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		deviceConfigs = append(deviceConfigs, config)
		return true
	})

	// 使用 goroutine 并行处理每个设备
	for _, config := range deviceConfigs {
		wg.Add(1)
		go func(deviceConfig *c_base.SDeviceConfig) {
			defer wg.Done()

			deviceId := deviceConfig.Id
			device := common.GetDeviceManager().GetDeviceById(deviceId)
			if device == nil {
				return
			}

			// 获取遥测数据
			var telemetryData *v1.DeviceTelemetryData
			if values := c_base.GetAllTelemetry(device); len(values) > 0 {
				telemetryData = &v1.DeviceTelemetryData{
					LastUpdateTime:  device.GetLastUpdateTime(),
					TelemetryValues: values,
				}
			}

			// 获取告警级别和协议状态
			alarmLevel := device.GetAlarmLevel().String()
			protocolStat := device.GetProtocolStatus().String()

			// 线程安全地写入结果
			mu.Lock()
			if telemetryData != nil {
				telemetry[deviceId] = telemetryData
			}
			alarmLevelMap[deviceId] = alarmLevel
			protocolStatus[deviceId] = protocolStat
			mu.Unlock()
		}(config)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	return &v1.GetDeviceTelemetryRes{
		ProtocolStatus: protocolStatus,
		AlarmLevelMap:  alarmLevelMap,
		Telemetry:      telemetry,
	}, nil
}

package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
	"sync"
)

func (c *ControllerV1) GetDeviceStatus(ctx context.Context, req *v1.GetDeviceStatusReq) (res *v1.GetDeviceStatusRes, err error) {
	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		// 系统还在初始化中
		return &v1.GetDeviceStatusRes{
			DeviceStatus: make(map[string]*v1.DeviceStatusData),
		}, nil
	}

	// 使用线程安全的 map 来存储结果
	var (
		deviceStatus = make(map[string]*v1.DeviceStatusData)
		mu           sync.RWMutex
		wg           sync.WaitGroup
	)

	// 收集设备配置，用于并行处理
	var deviceConfigs []*c_base.SDeviceConfig

	if req.DeviceId != "" {
		// 如果指定了设备ID，只处理该设备及其子设备
		common.GetDeviceManager().IteratorChildDevicesById(req.DeviceId, func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
			deviceConfigs = append(deviceConfigs, config)
			return true
		})
	} else {
		// 如果没有指定设备ID，处理所有设备
		common.GetDeviceManager().IteratorAllDevices(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
			deviceConfigs = append(deviceConfigs, config)
			return true
		})
	}

	// 使用 goroutine 并行处理每个设备
	for _, config := range deviceConfigs {
		wg.Add(1)
		go func(deviceConfig *c_base.SDeviceConfig) {
			defer wg.Done()

			deviceId := deviceConfig.Id
			device := common.GetDeviceManager().GetDeviceById(deviceId)
			if device == nil {
				// 设备不在运行中，仍然返回状态信息
				mu.Lock()
				deviceStatus[deviceId] = &v1.DeviceStatusData{
					ProtocolStatus: "Disconnected",
					AlarmLevel:     "Unknown",
					LastUpdateTime: nil,
				}
				mu.Unlock()
				return
			}

			// 获取设备状态信息
			alarmLevel := device.GetAlarmLevel().String()
			protocolStat := device.GetProtocolStatus().String()
			lastUpdateTime := device.GetLastUpdateTime()

			// 线程安全地写入结果
			mu.Lock()
			deviceStatus[deviceId] = &v1.DeviceStatusData{
				ProtocolStatus: protocolStat,
				AlarmLevel:     alarmLevel,
				LastUpdateTime: lastUpdateTime,
			}
			mu.Unlock()
		}(config)
	}

	// 等待所有 goroutine 完成
	wg.Wait()

	return &v1.GetDeviceStatusRes{
		DeviceStatus: deviceStatus,
	}, nil
}

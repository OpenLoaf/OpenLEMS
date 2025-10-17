package internal

import (
	"context"
	"sort"
	"time"

	"common"
	"common/c_base"
	"common/c_log"
	"common/c_type"
)

// executeStrategy 执行单个策略（简化版）
func executeStrategy(ctx context.Context, strategy *SEnergyManageStrategy) error {
	// 1. 获取绑定的设备列表
	devices := getDevicesByIds(strategy.EssDeviceIdsList)

	// 2. 遍历每个设备
	for _, device := range devices {
		// 2.1 类型断言为 IEnergyStore
		essDevice, ok := device.(c_type.IEnergyStore)
		if !ok {
			c_log.Warningf(ctx, "设备 %s 不是储能设备类型，跳过", device.GetConfig().Name)
			continue
		}

		// 2.2 获取当前 SOC
		socPtr, err := essDevice.GetSoc()
		if err != nil {
			c_log.Errorf(ctx, "获取设备 %s SOC 失败: %v", device.GetConfig().Name, err)
			continue
		}

		var soc float64
		if socPtr != nil {
			soc = float64(*socPtr)
		}

		// 2.3 计算目标功率
		targetPower := calculateTargetPower(soc, strategy.ConfigParsed, time.Now())

		// 2.4 获取当前状态（简化版，暂时不获取状态）
		currentStatus := "unknown"

		// 2.5 日志输出
		c_log.Infof(ctx, "策略执行 - 设备: %s, 当前SOC: %.2f%%, 目标功率: %dW, 当前状态: %s",
			device.GetConfig().Name, soc, targetPower, currentStatus)
	}

	return nil
}

// calculateTargetPower 计算目标功率
func calculateTargetPower(soc float64, config *SStrategyConfig, now time.Time) int32 {
	if config == nil {
		return 0 // 默认待机
	}

	// 1. SOC 边界检查
	if soc <= config.SocMinRatio {
		return -5000 // 充电（负值）
	}
	if soc >= config.SocMaxRatio {
		return 5000 // 放电（正值）
	}

	// 2. 按小时目标点位
	if len(config.Points) > 0 {
		targetSoc := getTargetSocByHour(now.Hour(), config.Points)
		diff := targetSoc - soc
		if diff > 1 {
			return -3000 // 充电
		} else if diff < -1 {
			return 3000 // 放电
		}
	}

	// 3. 默认待机
	return 0
}

// getTargetSocByHour 根据小时获取目标SOC（插值）
func getTargetSocByHour(hour int, points [][2]int) float64 {
	if len(points) == 0 {
		return 50.0 // 默认50%
	}

	// 按小时排序
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	// 查找当前小时在哪个区间
	for i := 0; i < len(points)-1; i++ {
		currHour := points[i][0]
		nextHour := points[i+1][0]
		currPercent := points[i][1]
		nextPercent := points[i+1][1]

		if hour >= currHour && hour <= nextHour {
			// 线性插值
			if nextHour == currHour {
				return float64(currPercent)
			}
			ratio := float64(hour-currHour) / float64(nextHour-currHour)
			return float64(currPercent) + ratio*float64(nextPercent-currPercent)
		}
	}

	// 如果超出范围，返回最后一个点
	return float64(points[len(points)-1][1])
}

// getDevicesByIds 根据设备ID列表获取设备实例
func getDevicesByIds(deviceIds []string) []c_base.IDevice {
	var devices []c_base.IDevice
	deviceManager := common.GetDeviceManager()

	for _, deviceId := range deviceIds {
		device := deviceManager.GetDeviceById(deviceId)
		if device != nil {
			devices = append(devices, device)
		}
	}

	return devices
}

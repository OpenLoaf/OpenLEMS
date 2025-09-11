package internal

import (
	"context"
	"fmt"
	"p_gpio_sysfs/p_gpio_sysfs"
	"strconv"

	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

// ScanGpioSysfs 使用 periph.io 扫描可用的 GPIO 引脚
// root 参数保留以保持接口兼容性，但实际使用 periph.io 的 GPIO 发现功能
func ScanGpioSysfs(ctx context.Context, root string) (*p_gpio_sysfs.SGpioScanResult, error) {
	result := &p_gpio_sysfs.SGpioScanResult{Chips: make([]*p_gpio_sysfs.SGpioChipInfo, 0), Gpios: make([]*p_gpio_sysfs.SGpioInfo, 0)}

	// 初始化 periph.io
	_, err := host.Init()
	if err != nil {
		return result, fmt.Errorf("初始化 periph.io 失败：%v", err)
	}

	// 扫描可用的 GPIO 引脚
	// 这里我们尝试扫描常见的 GPIO 引脚名称
	// 实际实现可能需要根据具体硬件平台调整
	for i := 0; i < MaxGpioScanCount; i++ {
		pinName := fmt.Sprintf("GPIO%d", i)
		pin := gpioreg.ByName(pinName)
		if pin != nil {
			gpioInfo := &p_gpio_sysfs.SGpioInfo{
				Name:      pinName,
				Path:      fmt.Sprintf("/sys/class/gpio/gpio%d", i),
				Number:    i,
				Direction: "unknown", // periph.io 不直接提供方向信息
				Value:     "unknown", // 需要实际读取才能获取值
			}
			result.Gpios = append(result.Gpios, gpioInfo)
		}
	}

	// 添加一个通用的芯片信息
	chip := &p_gpio_sysfs.SGpioChipInfo{
		Name:  DefaultGpioChipName,
		Path:  fmt.Sprintf("/sys/class/gpio/%s", DefaultGpioChipName),
		Label: DefaultGpioChipLabel,
		Base:  0,
		Ngpio: len(result.Gpios),
	}
	result.Chips = append(result.Chips, chip)

	return result, nil
}

func atoiSafe(s string) int {
	if s == "" {
		return 0
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return v
}

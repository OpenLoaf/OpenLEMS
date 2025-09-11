package internal

import (
	"context"
	"os"
	"p_gpio_sysfs/p_gpio_sysfs"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gogf/gf/v2/os/gfile"
)

// ScanGpioSysfs 扫描指定 sysfs gpio 路径，汇总 gpiochip 与已导出的 gpioN 信息
// 常见 root 为 /sys/class/gpio
func ScanGpioSysfs(ctx context.Context, root string) (*p_gpio_sysfs.SGpioScanResult, error) {
	result := &p_gpio_sysfs.SGpioScanResult{Chips: make([]*p_gpio_sysfs.SGpioChipInfo, 0), Gpios: make([]*p_gpio_sysfs.SGpioInfo, 0)}

	entries, err := os.ReadDir(root)
	if err != nil {
		return result, err
	}

	for _, e := range entries {
		if !e.IsDir() {
			continue
		}
		name := e.Name()
		full := filepath.Join(root, name)

		if strings.HasPrefix(name, "gpiochip") {
			chip := &p_gpio_sysfs.SGpioChipInfo{Name: name, Path: full}
			chip.Label = strings.TrimSpace(gfile.GetContents(filepath.Join(full, "label")))
			chip.Base = atoiSafe(strings.TrimSpace(gfile.GetContents(filepath.Join(full, "base"))))
			chip.Ngpio = atoiSafe(strings.TrimSpace(gfile.GetContents(filepath.Join(full, "ngpio"))))
			result.Chips = append(result.Chips, chip)
			continue
		}

		if strings.HasPrefix(name, "gpio") {
			num := atoiSafe(strings.TrimPrefix(name, "gpio"))
			gpio := &p_gpio_sysfs.SGpioInfo{
				Name:   name,
				Path:   full,
				Number: num,
			}
			// direction 与 value 可能不存在（未导出或权限限制），忽略错误
			gpio.Direction = strings.TrimSpace(gfile.GetContents(filepath.Join(full, GpioPathDirection)))
			gpio.Value = strings.TrimSpace(gfile.GetContents(filepath.Join(full, GpioPathValue)))
			result.Gpios = append(result.Gpios, gpio)
		}
	}

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

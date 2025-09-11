package internal

// 保留这些常量以保持向后兼容性，但实际使用 periph.io 时不再需要
const (
	GpioPathValue     = "value"
	GpioPathDirection = "direction"
)

// periph.io 相关的常量
const (
	// 默认扫描的 GPIO 引脚数量上限
	MaxGpioScanCount = 100

	// 默认的 GPIO 芯片名称
	DefaultGpioChipName = "gpiochip0"

	// 默认的 GPIO 芯片标签
	DefaultGpioChipLabel = "periph.io GPIO"
)

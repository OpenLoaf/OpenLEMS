package p_gpio_sysfs

type EGpioDirection string

const (
	EGpioDirectionNone EGpioDirection = ""
	EGpioDirectionIn   EGpioDirection = "in"
	EGpioDirectionOut  EGpioDirection = "out"
)

//go:generate stringer -type=EGpioDirection -output=c_gpio_direction_e.go
package c_proto

type EGpioDirection string

const (
	EGpioDirectionNone EGpioDirection = ""
	EGpioDirectionIn   EGpioDirection = "in"
	EGpioDirectionOut  EGpioDirection = "out"
)

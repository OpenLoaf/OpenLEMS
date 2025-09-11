package c_proto

import (
	"common/c_enum"
)

type SDeviceGpioConfig struct {
	Io        int                `json:"io" name:"GPIO编号" min:"0" max:"999" default:"0" required:"true" step:"1"` // GPIO编号
	Direction EGpioDirection     `json:"direction" name:"方向" required:"true"`                                     // 方向：输入/输出
	Level     c_enum.EAlarmLevel `json:"level" name:"告警级别"`                                                       // 告警级别
	Reverse   bool               `json:"reverse" name:"反向" default:"false"`                                       // 是否反向
}

type EGpioDirection string

const (
	EGpioDirectionNone EGpioDirection = ""
	EGpioDirectionIn   EGpioDirection = "in"
	EGpioDirectionOut  EGpioDirection = "out"
)

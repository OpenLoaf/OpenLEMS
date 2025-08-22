package c_proto

import (
	"common/c_base"
)

type SDeviceGpioConfig struct {
	Direction EGpioDirection `json:"direction" dc:"GPIO方向IN/OUT" name:"direction" brief:"GPIO方向"`
	Io        uint           `json:"io" dc:"GPIO编号" name:"io" brief:"GPIO编号"`
	//Path      string             `json:"path" dc:"GPIO路径"`
	Level   c_base.EAlarmLevel `json:"level" dc:"告警级别" name:"level" brief:"告警级别"`
	Reverse bool               `json:"reverse" dc:"是否反转" name:"reverse" brief:"是否反转"`
}

package c_proto

import (
	"common/c_enum"
)

// SGpioProtocolConfig gpio协议的配置
type SGpioProtocolConfig struct {
	Io        int                   `json:"io" name:"GPIO编号" required:"true" ct:"number" vt:"int" min:"1" max:"99" default:"0" required:"true" step:"1" dc:"GPIO引脚编号，范围0-99"`
	Direction c_enum.EGpioDirection `json:"direction" required:"true" name:"方向" ct:"singleSelect" vt:"string" selectOptions:"in:输入,out:输出" required:"true" dc:"GPIO引脚方向：输入用于读取状态，输出用于控制"`
}

// SGpioDeviceConfig gpio设备的配置
type SGpioDeviceConfig struct {
	Level       c_enum.EAlarmLevel `json:"level" name:"告警级别" ct:"singleSelect" vt:"int" selectOptions:"0:无告警,1:警告,2:警报,3:故障" default:"0" dc:"GPIO状态变化时的告警级别"`
	HighTrigger bool               `json:"highTrigger" name:"高电平触发" ct:"switch" vt:"bool" default:"true" dc:"触发模式"`
}

// SGpiodProtocolConfig gpiod协议的配置
type SGpiodProtocolConfig struct {
	Name      string                `json:"name" dc:"GPIO名称"`
	Direction c_enum.EGpioDirection `json:"direction" required:"true" name:"方向" ct:"singleSelect" vt:"string" selectOptions:"in:输入,out:输出" required:"true" dc:"GPIO引脚方向：输入用于读取状态，输出用于控制"`
	ChipIndex uint8                 `json:"chipIndex" dc:"GPIO芯片名称，如gpiochip0"`
	Pin       uint8                 `json:"pin" dc:"GPIO引脚编号，范围0-99"`
	LowActive bool                  `json:"lowActive" dc:"低电平有效"`
}

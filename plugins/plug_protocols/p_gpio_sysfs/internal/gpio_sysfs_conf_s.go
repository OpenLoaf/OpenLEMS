package internal

import "common/c_proto"

// SGpioSysfsConfig 描述一个 GPIO 芯片
type sGpioSysfsConfig struct {
	c_proto.SGpioProtocolConfig
	Input  []uint16 `json:"input"`
	Output []uint16 `json:"output"`
}

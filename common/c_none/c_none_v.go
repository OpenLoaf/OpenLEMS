package c_none

import "fmt"

var (
	noneErr      = fmt.Errorf("this device is not real")
	NoneProtocol = &sNoneProtocol{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NonePcs      = &sNonePcs{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
	NoneBms      = &sNoneBms{sNoneAlarm: sNoneAlarm{}, sNoneDeviceRuntimeInfo: sNoneDeviceRuntimeInfo{}}
)

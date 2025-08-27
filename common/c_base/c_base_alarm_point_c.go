package c_base

import "github.com/shockerli/cvt"

var alarmTrigger = func(a any) bool {
	return cvt.Bool(a)
}

var (
	ProtocolClientDisconnectedPoint = &Meta{Name: "ProtocolClientDisconnected", Cn: "通讯连接断开", SystemType: SBool, Level: EAlarmLevelError, Trigger: alarmTrigger}
)

package c_base

import (
	"common/c_enum"

	"github.com/shockerli/cvt"
)

var alarmTrigger = func(a any) bool {
	return cvt.Bool(a)
}

var (
	ProtocolClientDisconnectedPoint = &SModbusPoint{Name: "ProtocolClientDisconnected", Cn: "通讯连接断开", SystemType: SBool, Level: c_enum.EAlarmLevelError, Trigger: alarmTrigger}
)

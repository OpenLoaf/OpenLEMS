package c_default

import (
	"common/c_enum"

	"github.com/shockerli/cvt"
)

func fAlarmTriggerBool(value interface{}, l c_enum.EAlarmLevel) (trigger bool, level c_enum.EAlarmLevel, err error) {
	level = l
	trigger, err = cvt.BoolE(value)
	return
}
func fAlarmTriggerNotZero(value interface{}, l c_enum.EAlarmLevel) (trigger bool, level c_enum.EAlarmLevel, err error) {
	level = l
	if v, err := cvt.IntE(value); err == nil {
		trigger = v == 0
	}
	return
}

func FAlarmTriggerWarnBool(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
	return fAlarmTriggerBool(value, c_enum.EAlarmLevelWarn)
}

func FAlarmTriggerAlertBool(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
	return fAlarmTriggerBool(value, c_enum.EAlarmLevelAlert)
}

func FAlarmTriggerErrorBool(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
	return fAlarmTriggerBool(value, c_enum.EAlarmLevelError)
}

func FAlarmTriggerWarnNotZero(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
	return fAlarmTriggerNotZero(value, c_enum.EAlarmLevelWarn)
}

func FAlarmTriggerAlertNotZero(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
	return fAlarmTriggerNotZero(value, c_enum.EAlarmLevelAlert)
}

func FAlarmTriggerErrorNotZero(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
	return fAlarmTriggerNotZero(value, c_enum.EAlarmLevelError)
}

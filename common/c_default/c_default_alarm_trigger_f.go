package c_default

import (
	"common/c_base"
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
		trigger = v != 0
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

// FAlarmTriggerRangeThan 根据数值范围触发不同级别的告警，优先级：错误 > 警报 > 警告
func FAlarmTriggerRangeThan(value interface{}, rangeTrigger *c_base.SAlarmRangeTrigger) (trigger bool, level c_enum.EAlarmLevel, err error) {
	// 将输入值转换为float64类型
	v, err := cvt.Float64E(value)
	if err != nil {
		return false, c_enum.EAlarmLevelNone, err
	}

	// 检查是否超出错误级别范围（最高优先级）
	if trigger, level := rangeTrigger.Error.CheckAlarm(v, c_enum.EAlarmLevelError); trigger {
		return true, level, nil
	}

	// 检查是否超出警报级别范围
	if trigger, level := rangeTrigger.Alert.CheckAlarm(v, c_enum.EAlarmLevelAlert); trigger {
		return true, level, nil
	}

	// 检查是否超出警告级别范围
	if trigger, level := rangeTrigger.Warn.CheckAlarm(v, c_enum.EAlarmLevelWarn); trigger {
		return true, level, nil
	}

	// 值在正常范围内，不触发告警
	return false, c_enum.EAlarmLevelNone, nil
}

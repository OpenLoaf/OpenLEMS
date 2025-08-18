//go:generate stringer -type=EAlarmLevel -trimprefix=E -output=c_base_alarm_level_e_string.go
package c_base

type EAlarmLevel int

const (
	ENone  EAlarmLevel = iota // 默认非告警
	EWarn                     // 警告，系统正常工作
	EAlarm                    // 警报，系统降低功率
	EError                    // 故障，一旦有一个系统全部停机
)

//
//func (l EAlarmLevel) Name() string {
//	switch l {
//	case ENone:
//		return "-"
//	case EWarn:
//		return "预警"
//	case EAlarm:
//		return "警报"
//	case EError:
//		return "故障"
//	}
//	return "未知"
//}
//
//func (l EAlarmLevel) FullName() string {
//	switch l {
//	case ENone:
//		return "正常"
//	case EWarn:
//		return "预警"
//	case EAlarm:
//		return "警报，系统将降低功率"
//	case EError:
//		return "故障，系统将全部停机"
//	}
//	return "未知的告警级别"
//}

package c_default

import (
	"common/c_base"
	"common/c_enum"
)

// 电力系统常用点位配置

// 三相电流相关点位
var (
	VPointIa   = &c_base.SPoint{Key: "Ia", Name: "A相电流", Unit: "A", Desc: "A相电流值", Sort: 1, Min: 0, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic}    // A相电流
	VPointIb   = &c_base.SPoint{Key: "Ib", Name: "B相电流", Unit: "A", Desc: "B相电流值", Sort: 2, Min: 0, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic}    // B相电流
	VPointIc   = &c_base.SPoint{Key: "Ic", Name: "C相电流", Unit: "A", Desc: "C相电流值", Sort: 3, Min: 0, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic}    // C相电流
	VPointIavg = &c_base.SPoint{Key: "Iavg", Name: "平均电流", Unit: "A", Desc: "三相平均电流", Sort: 4, Min: 0, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic} // 平均电流
	VPointImax = &c_base.SPoint{Key: "Imax", Name: "最大电流", Unit: "A", Desc: "三相最大电流", Sort: 5, Min: 0, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic} // 最大电流
)

// 三相电压相关点位
var (
	VPointUa = &c_base.SPoint{Key: "Ua", Name: "A相电压", Unit: "V", Desc: "A相电压值", Sort: 10, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 150, After: 280}, // 错误：< 150V 或 > 280V
			Alert: &c_base.SAlarmOvertop{Before: 180, After: 250}, // 警报：< 180V 或 > 250V
			Warn:  &c_base.SAlarmOvertop{Before: 200, After: 240}, // 警告：< 200V 或 > 240V
		})
	}} // A相电压
	VPointUb = &c_base.SPoint{Key: "Ub", Name: "B相电压", Unit: "V", Desc: "B相电压值", Sort: 11, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 150, After: 280}, // 错误：< 150V 或 > 280V
			Alert: &c_base.SAlarmOvertop{Before: 180, After: 250}, // 警报：< 180V 或 > 250V
			Warn:  &c_base.SAlarmOvertop{Before: 200, After: 240}, // 警告：< 200V 或 > 240V
		})
	}} // B相电压
	VPointUc = &c_base.SPoint{Key: "Uc", Name: "C相电压", Unit: "V", Desc: "C相电压值", Sort: 12, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 150, After: 280}, // 错误：< 150V 或 > 280V
			Alert: &c_base.SAlarmOvertop{Before: 180, After: 250}, // 警报：< 180V 或 > 250V
			Warn:  &c_base.SAlarmOvertop{Before: 200, After: 240}, // 警告：< 200V 或 > 240V
		})
	}} // C相电压
	VPointUavg = &c_base.SPoint{Key: "Uavg", Name: "平均电压", Unit: "V", Desc: "三相平均电压", Sort: 13, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 150, After: 280}, // 错误：< 150V 或 > 280V
			Alert: &c_base.SAlarmOvertop{Before: 180, After: 250}, // 警报：< 180V 或 > 250V
			Warn:  &c_base.SAlarmOvertop{Before: 200, After: 240}, // 警告：< 200V 或 > 240V
		})
	}} // 平均电压
	VPointUmax = &c_base.SPoint{Key: "Umax", Name: "最大电压", Unit: "V", Desc: "三相最大电压", Sort: 14, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 150, After: 280}, // 错误：< 150V 或 > 280V
			Alert: &c_base.SAlarmOvertop{Before: 180, After: 250}, // 警报：< 180V 或 > 250V
			Warn:  &c_base.SAlarmOvertop{Before: 200, After: 240}, // 警告：< 200V 或 > 240V
		})
	}} // 最大电压
	VPointUab = &c_base.SPoint{Key: "Uab", Name: "AB线电压", Unit: "V", Desc: "AB线电压值", Sort: 15, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 250, After: 500}, // 错误：< 250V 或 > 500V
			Alert: &c_base.SAlarmOvertop{Before: 300, After: 450}, // 警报：< 300V 或 > 450V
			Warn:  &c_base.SAlarmOvertop{Before: 350, After: 420}, // 警告：< 350V 或 > 420V
		})
	}} // AB线电压
	VPointUbc = &c_base.SPoint{Key: "Ubc", Name: "BC线电压", Unit: "V", Desc: "BC线电压值", Sort: 16, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 250, After: 500}, // 错误：< 250V 或 > 500V
			Alert: &c_base.SAlarmOvertop{Before: 300, After: 450}, // 警报：< 300V 或 > 450V
			Warn:  &c_base.SAlarmOvertop{Before: 350, After: 420}, // 警告：< 350V 或 > 420V
		})
	}} // BC线电压
	VPointUca = &c_base.SPoint{Key: "Uca", Name: "CA线电压", Unit: "V", Desc: "CA线电压值", Sort: 17, Min: 0, Max: 1000, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 250, After: 500}, // 错误：< 250V 或 > 500V
			Alert: &c_base.SAlarmOvertop{Before: 300, After: 450}, // 警报：< 300V 或 > 450V
			Warn:  &c_base.SAlarmOvertop{Before: 350, After: 420}, // 警告：< 350V 或 > 420V
		})
	}} // CA线电压
)

// 功率相关点位
var (
	VPointPa = &c_base.SPoint{Key: "Pa", Name: "A相有功功率", Unit: "kW", Desc: "A相有功功率", Sort: 20, Precise: 2, Group: VPointGroupSystemBasic}   // A相有功功率
	VPointPb = &c_base.SPoint{Key: "Pb", Name: "B相有功功率", Unit: "kW", Desc: "B相有功功率", Sort: 21, Precise: 2, Group: VPointGroupSystemBasic}   // B相有功功率
	VPointPc = &c_base.SPoint{Key: "Pc", Name: "C相有功功率", Unit: "kW", Desc: "C相有功功率", Sort: 22, Precise: 2, Group: VPointGroupSystemBasic}   // C相有功功率
	VPointP  = &c_base.SPoint{Key: "P", Name: "总有功功率", Unit: "kW", Desc: "三相总有功功率", Sort: 23, Precise: 2, Group: VPointGroupSystemBasic}    // 总有功功率
	VPointQa = &c_base.SPoint{Key: "Qa", Name: "A相无功功率", Unit: "kVar", Desc: "A相无功功率", Sort: 24, Precise: 2, Group: VPointGroupSystemBasic} // A相无功功率
	VPointQb = &c_base.SPoint{Key: "Qb", Name: "B相无功功率", Unit: "kVar", Desc: "B相无功功率", Sort: 25, Precise: 2, Group: VPointGroupSystemBasic} // B相无功功率
	VPointQc = &c_base.SPoint{Key: "Qc", Name: "C相无功功率", Unit: "kVar", Desc: "C相无功功率", Sort: 26, Precise: 2, Group: VPointGroupSystemBasic} // C相无功功率
	VPointQ  = &c_base.SPoint{Key: "Q", Name: "总无功功率", Unit: "kVar", Desc: "三相总无功功率", Sort: 27, Precise: 2, Group: VPointGroupSystemBasic}  // 总无功功率
	VPointSa = &c_base.SPoint{Key: "Sa", Name: "A相视在功率", Unit: "kVA", Desc: "A相视在功率", Sort: 28, Precise: 2, Group: VPointGroupSystemBasic}  // A相视在功率
	VPointSb = &c_base.SPoint{Key: "Sb", Name: "B相视在功率", Unit: "kVA", Desc: "B相视在功率", Sort: 29, Precise: 2, Group: VPointGroupSystemBasic}  // B相视在功率
	VPointSc = &c_base.SPoint{Key: "Sc", Name: "C相视在功率", Unit: "kVA", Desc: "C相视在功率", Sort: 30, Precise: 2, Group: VPointGroupSystemBasic}  // C相视在功率
	VPointS  = &c_base.SPoint{Key: "S", Name: "总视在功率", Unit: "kVA", Desc: "三相总视在功率", Sort: 31, Precise: 2, Group: VPointGroupSystemBasic}   // 总视在功率
)

// 功率因数相关点位
var (
	VPointPFa = &c_base.SPoint{Key: "PFa", Name: "A相功率因数", Unit: "", Desc: "A相功率因数", Sort: 35, Min: -1, Max: 1, Precise: 3, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -1.0, After: -0.5}, // 错误：< -0.5
			Alert: &c_base.SAlarmOvertop{Before: -1.0, After: -0.7}, // 警报：< -0.7
			Warn:  &c_base.SAlarmOvertop{Before: -1.0, After: -0.8}, // 警告：< -0.8
		})
	}} // A相功率因数
	VPointPFb = &c_base.SPoint{Key: "PFb", Name: "B相功率因数", Unit: "", Desc: "B相功率因数", Sort: 36, Min: -1, Max: 1, Precise: 3, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -1.0, After: -0.5}, // 错误：< -0.5
			Alert: &c_base.SAlarmOvertop{Before: -1.0, After: -0.7}, // 警报：< -0.7
			Warn:  &c_base.SAlarmOvertop{Before: -1.0, After: -0.8}, // 警告：< -0.8
		})
	}} // B相功率因数
	VPointPFc = &c_base.SPoint{Key: "PFc", Name: "C相功率因数", Unit: "", Desc: "C相功率因数", Sort: 37, Min: -1, Max: 1, Precise: 3, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -1.0, After: -0.5}, // 错误：< -0.5
			Alert: &c_base.SAlarmOvertop{Before: -1.0, After: -0.7}, // 警报：< -0.7
			Warn:  &c_base.SAlarmOvertop{Before: -1.0, After: -0.8}, // 警告：< -0.8
		})
	}} // C相功率因数
	VPointPF = &c_base.SPoint{Key: "PF", Name: "总功率因数", Unit: "", Desc: "三相总功率因数", Sort: 38, Min: -1, Max: 1, Precise: 3, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -1.0, After: -0.5}, // 错误：< -0.5
			Alert: &c_base.SAlarmOvertop{Before: -1.0, After: -0.7}, // 警报：< -0.7
			Warn:  &c_base.SAlarmOvertop{Before: -1.0, After: -0.8}, // 警告：< -0.8
		})
	}} // 总功率因数
)

// 目标功率相关点位
var (
	VPointTargetP  = &c_base.SPoint{Key: "TargetP", Name: "目标有功功率", Unit: "kW", Desc: "目标有功功率设定值", Sort: 32, Precise: 2, Group: VPointGroupSystemBasic}                 // 目标有功功率
	VPointTargetQ  = &c_base.SPoint{Key: "TargetQ", Name: "目标无功功率", Unit: "kVar", Desc: "目标无功功率设定值", Sort: 33, Precise: 2, Group: VPointGroupSystemBasic}               // 目标无功功率
	VPointTargetPF = &c_base.SPoint{Key: "TargetPF", Name: "目标功率因数", Unit: "", Desc: "目标功率因数设定值", Sort: 34, Min: -1, Max: 1, Precise: 3, Group: VPointGroupSystemBasic} // 目标功率因数
)

// 频率相关点位
var (
	VPointFreq = &c_base.SPoint{Key: "Freq", Name: "频率", Unit: "Hz", Desc: "电网频率", Sort: 40, Min: 45, Max: 55, Precise: 2, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 45, After: 55}, // 错误：< 45Hz 或 > 55Hz
			Alert: &c_base.SAlarmOvertop{Before: 47, After: 53}, // 警报：< 47Hz 或 > 53Hz
			Warn:  &c_base.SAlarmOvertop{Before: 48, After: 52}, // 警告：< 48Hz 或 > 52Hz
		})
	}} // 频率
)

// 电能相关点位
var (
	VPointEa = &c_base.SPoint{Key: "Ea", Name: "A相电能", Unit: "kWh", Desc: "A相累计电能", Sort: 50, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic} // A相电能
	VPointEb = &c_base.SPoint{Key: "Eb", Name: "B相电能", Unit: "kWh", Desc: "B相累计电能", Sort: 51, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic} // B相电能
	VPointEc = &c_base.SPoint{Key: "Ec", Name: "C相电能", Unit: "kWh", Desc: "C相累计电能", Sort: 52, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic} // C相电能
	VPointE  = &c_base.SPoint{Key: "E", Name: "总电能", Unit: "kWh", Desc: "三相总累计电能", Sort: 53, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic}  // 总电能
)

// 充放电量相关点位
var (
	VPointTodayCharge    = &c_base.SPoint{Key: "TodayCharge", Name: "今日充电量", Unit: "kWh", Desc: "今日累计充电量", Sort: 54, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic}      // 今日充电量
	VPointTodayDischarge = &c_base.SPoint{Key: "TodayDischarge", Name: "今日放电量", Unit: "kWh", Desc: "今日累计放电量", Sort: 55, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic}   // 今日放电量
	VPointTotalCharge    = &c_base.SPoint{Key: "TotalCharge", Name: "历史总充电量", Unit: "kWh", Desc: "历史累计总充电量", Sort: 56, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic}    // 历史总充电量
	VPointTotalDischarge = &c_base.SPoint{Key: "TotalDischarge", Name: "历史总放电量", Unit: "kWh", Desc: "历史累计总放电量", Sort: 57, Min: 0, Max: 0, Precise: 2, Group: VPointGroupSystemBasic} // 历史总放电量
)

// 温度相关点位
var (
	VPointTemp = &c_base.SPoint{Key: "Temp", Name: "温度", Unit: "℃", Desc: "设备温度", Sort: 60, Min: -40, Max: 85, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -20, After: 80}, // 错误：< -20℃ 或 > 80℃
			Alert: &c_base.SAlarmOvertop{Before: -10, After: 70}, // 警报：< -10℃ 或 > 70℃
			Warn:  &c_base.SAlarmOvertop{Before: 0, After: 60},   // 警告：< 0℃ 或 > 60℃
		})
	}} // 温度
	VPointTempMax = &c_base.SPoint{Key: "TempMax", Name: "最高温度", Unit: "℃", Desc: "设备最高温度", Sort: 61, Min: -40, Max: 85, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -20, After: 80}, // 错误：< -20℃ 或 > 80℃
			Alert: &c_base.SAlarmOvertop{Before: -10, After: 70}, // 警报：< -10℃ 或 > 70℃
			Warn:  &c_base.SAlarmOvertop{Before: 0, After: 60},   // 警告：< 0℃ 或 > 60℃
		})
	}} // 最高温度
	VPointTempMin = &c_base.SPoint{Key: "TempMin", Name: "最低温度", Unit: "℃", Desc: "设备最低温度", Sort: 62, Min: -40, Max: 85, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -20, After: 80}, // 错误：< -20℃ 或 > 80℃
			Alert: &c_base.SAlarmOvertop{Before: -10, After: 70}, // 警报：< -10℃ 或 > 70℃
			Warn:  &c_base.SAlarmOvertop{Before: 0, After: 60},   // 警告：< 0℃ 或 > 60℃
		})
	}} // 最低温度
	VPointIGBTTemp = &c_base.SPoint{Key: "IGBTTemp", Name: "IGBT温度", Unit: "℃", Desc: "IGBT模块温度", Sort: 63, Min: -40, Max: 200, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: -20, After: 130}, // 错误：< -20℃ 或 > 130℃
			Alert: &c_base.SAlarmOvertop{Before: -10, After: 110}, // 警报：< -10℃ 或 > 110℃
			Warn:  &c_base.SAlarmOvertop{Before: 0, After: 95},    // 警告：< 0℃ 或 > 95℃
		})
	}} // IGBT温度
)

// 湿度相关点位
var (
	VPointHumidity = &c_base.SPoint{Key: "Humidity", Name: "湿度", Unit: "%RH", Desc: "环境湿度", Sort: 70, Min: 0, Max: 100, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 0, After: 95}, // 错误：> 95%
			Alert: &c_base.SAlarmOvertop{Before: 0, After: 90}, // 警报：> 90%
			Warn:  &c_base.SAlarmOvertop{Before: 0, After: 80}, // 警告：> 80%
		})
	}} // 湿度
)

// 状态相关点位
var (
	VPointStatus    = &c_base.SPoint{Key: "Status", Name: "状态", Unit: "", Desc: "设备运行状态", Sort: 80, Min: 0, Max: 255, Precise: 0, Group: VPointGroupSystemBasic}        // 状态
	VPointErrorCode = &c_base.SPoint{Key: "ErrorCode", Name: "错误代码", Unit: "", Desc: "设备错误代码", Sort: 82, Min: 0, Max: 65535, Precise: 0, Group: VPointGroupSystemBasic} // 错误代码
)

// 电池相关点位 (BMS)
var (
	VPointSOC = &c_base.SPoint{Key: "SOC", Name: "电池电量", Unit: "%", Desc: "电池剩余电量百分比", Sort: 90, Min: 0, Max: 100, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 0, After: 100}, // 错误：> 100%
		})
	}} // 电池电量
	VPointSOH = &c_base.SPoint{Key: "SOH", Name: "电池健康度", Unit: "%", Desc: "电池健康度百分比", Sort: 91, Min: 0, Max: 100, Precise: 1, Group: VPointGroupSystemBasic, Trigger: func(value interface{}) (bool, c_enum.EAlarmLevel, error) {
		return FAlarmTriggerRangeThan(value, &c_base.SAlarmRangeTrigger{
			Error: &c_base.SAlarmOvertop{Before: 0, After: 100}, // 错误：> 100%
		})
	}} // 电池健康度
	VPointDcVoltage = &c_base.SPoint{Key: "Voltage", Name: "电池电压", Unit: "V", Desc: "电池总电压", Sort: 92, Min: 0, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic}       // 电池电压
	VPointDcCurrent = &c_base.SPoint{Key: "Current", Name: "电池电流", Unit: "A", Desc: "电池充放电电流", Sort: 93, Min: -1000, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic} // 电池电流
	VPointDcPower   = &c_base.SPoint{Key: "Power", Name: "电池直流功率", Unit: "kW", Desc: "电池直流功率", Sort: 94, Min: -1000, Max: 1000, Precise: 2, Group: VPointGroupSystemBasic} // 电池直流功率
)

// 最大功率和电流相关点位
var (
	VPointMaxChargePower      = &c_base.SPoint{Key: "MaxChargePower", Name: "最大充电功率", Unit: "kW", Desc: "设备最大充电功率", Sort: 95, Precise: 2, Group: VPointGroupSystemBasic}     // 最大充电功率
	VPointMaxDischargePower   = &c_base.SPoint{Key: "MaxDischargePower", Name: "最大放电功率", Unit: "kW", Desc: "设备最大放电功率", Sort: 96, Precise: 2, Group: VPointGroupSystemBasic}  // 最大放电功率
	VPointMaxChargeCurrent    = &c_base.SPoint{Key: "MaxChargeCurrent", Name: "最大充电电流", Unit: "A", Desc: "设备最大充电电流", Sort: 97, Precise: 2, Group: VPointGroupSystemBasic}    // 最大充电电流
	VPointMaxDischargeCurrent = &c_base.SPoint{Key: "MaxDischargeCurrent", Name: "最大放电电流", Unit: "A", Desc: "设备最大放电电流", Sort: 98, Precise: 2, Group: VPointGroupSystemBasic} // 最大放电电流
	VPointMaxPower            = &c_base.SPoint{Key: "MaxPower", Name: "最大功率", Unit: "kW", Desc: "设备最大功率", Sort: 99, Precise: 2, Group: VPointGroupSystemBasic}               // 最大功率
	VPointMaxCurrent          = &c_base.SPoint{Key: "MaxCurrent", Name: "最大电流", Unit: "A", Desc: "设备最大电流", Sort: 100, Precise: 2, Group: VPointGroupSystemBasic}             // 最大电流
)

package c_default

import (
	"common/c_base"
)

// 电力系统常用点位配置

// 三相电流相关点位
var (
	VPointIa   = &c_base.SPoint{Key: "Ia", Name: "A相电流", Unit: "A", Desc: "A相电流值", Sort: 1, Min: 0, Max: 1000, Precise: 2}    // A相电流
	VPointIb   = &c_base.SPoint{Key: "Ib", Name: "B相电流", Unit: "A", Desc: "B相电流值", Sort: 2, Min: 0, Max: 1000, Precise: 2}    // B相电流
	VPointIc   = &c_base.SPoint{Key: "Ic", Name: "C相电流", Unit: "A", Desc: "C相电流值", Sort: 3, Min: 0, Max: 1000, Precise: 2}    // C相电流
	VPointIavg = &c_base.SPoint{Key: "Iavg", Name: "平均电流", Unit: "A", Desc: "三相平均电流", Sort: 4, Min: 0, Max: 1000, Precise: 2} // 平均电流
	VPointImax = &c_base.SPoint{Key: "Imax", Name: "最大电流", Unit: "A", Desc: "三相最大电流", Sort: 5, Min: 0, Max: 1000, Precise: 2} // 最大电流
)

// 三相电压相关点位
var (
	VPointUa   = &c_base.SPoint{Key: "Ua", Name: "A相电压", Unit: "V", Desc: "A相电压值", Sort: 10, Min: 0, Max: 1000, Precise: 1}    // A相电压
	VPointUb   = &c_base.SPoint{Key: "Ub", Name: "B相电压", Unit: "V", Desc: "B相电压值", Sort: 11, Min: 0, Max: 1000, Precise: 1}    // B相电压
	VPointUc   = &c_base.SPoint{Key: "Uc", Name: "C相电压", Unit: "V", Desc: "C相电压值", Sort: 12, Min: 0, Max: 1000, Precise: 1}    // C相电压
	VPointUavg = &c_base.SPoint{Key: "Uavg", Name: "平均电压", Unit: "V", Desc: "三相平均电压", Sort: 13, Min: 0, Max: 1000, Precise: 1} // 平均电压
	VPointUmax = &c_base.SPoint{Key: "Umax", Name: "最大电压", Unit: "V", Desc: "三相最大电压", Sort: 14, Min: 0, Max: 1000, Precise: 1} // 最大电压
	VPointUab  = &c_base.SPoint{Key: "Uab", Name: "AB线电压", Unit: "V", Desc: "AB线电压值", Sort: 15, Min: 0, Max: 1000, Precise: 1} // AB线电压
	VPointUbc  = &c_base.SPoint{Key: "Ubc", Name: "BC线电压", Unit: "V", Desc: "BC线电压值", Sort: 16, Min: 0, Max: 1000, Precise: 1} // BC线电压
	VPointUca  = &c_base.SPoint{Key: "Uca", Name: "CA线电压", Unit: "V", Desc: "CA线电压值", Sort: 17, Min: 0, Max: 1000, Precise: 1} // CA线电压
)

// 功率相关点位
var (
	VPointPa = &c_base.SPoint{Key: "Pa", Name: "A相有功功率", Unit: "kW", Desc: "A相有功功率", Sort: 20, Precise: 2}   // A相有功功率
	VPointPb = &c_base.SPoint{Key: "Pb", Name: "B相有功功率", Unit: "kW", Desc: "B相有功功率", Sort: 21, Precise: 2}   // B相有功功率
	VPointPc = &c_base.SPoint{Key: "Pc", Name: "C相有功功率", Unit: "kW", Desc: "C相有功功率", Sort: 22, Precise: 2}   // C相有功功率
	VPointP  = &c_base.SPoint{Key: "P", Name: "总有功功率", Unit: "kW", Desc: "三相总有功功率", Sort: 23, Precise: 2}    // 总有功功率
	VPointQa = &c_base.SPoint{Key: "Qa", Name: "A相无功功率", Unit: "kVar", Desc: "A相无功功率", Sort: 24, Precise: 2} // A相无功功率
	VPointQb = &c_base.SPoint{Key: "Qb", Name: "B相无功功率", Unit: "kVar", Desc: "B相无功功率", Sort: 25, Precise: 2} // B相无功功率
	VPointQc = &c_base.SPoint{Key: "Qc", Name: "C相无功功率", Unit: "kVar", Desc: "C相无功功率", Sort: 26, Precise: 2} // C相无功功率
	VPointQ  = &c_base.SPoint{Key: "Q", Name: "总无功功率", Unit: "kVar", Desc: "三相总无功功率", Sort: 27, Precise: 2}  // 总无功功率
	VPointSa = &c_base.SPoint{Key: "Sa", Name: "A相视在功率", Unit: "kVA", Desc: "A相视在功率", Sort: 28, Precise: 2}  // A相视在功率
	VPointSb = &c_base.SPoint{Key: "Sb", Name: "B相视在功率", Unit: "kVA", Desc: "B相视在功率", Sort: 29, Precise: 2}  // B相视在功率
	VPointSc = &c_base.SPoint{Key: "Sc", Name: "C相视在功率", Unit: "kVA", Desc: "C相视在功率", Sort: 30, Precise: 2}  // C相视在功率
	VPointS  = &c_base.SPoint{Key: "S", Name: "总视在功率", Unit: "kVA", Desc: "三相总视在功率", Sort: 31, Precise: 2}   // 总视在功率
)

// 功率因数相关点位
var (
	VPointPFa = &c_base.SPoint{Key: "PFa", Name: "A相功率因数", Unit: "", Desc: "A相功率因数", Sort: 35, Min: -1, Max: 1, Precise: 3} // A相功率因数
	VPointPFb = &c_base.SPoint{Key: "PFb", Name: "B相功率因数", Unit: "", Desc: "B相功率因数", Sort: 36, Min: -1, Max: 1, Precise: 3} // B相功率因数
	VPointPFc = &c_base.SPoint{Key: "PFc", Name: "C相功率因数", Unit: "", Desc: "C相功率因数", Sort: 37, Min: -1, Max: 1, Precise: 3} // C相功率因数
	VPointPF  = &c_base.SPoint{Key: "PF", Name: "总功率因数", Unit: "", Desc: "三相总功率因数", Sort: 38, Min: -1, Max: 1, Precise: 3}  // 总功率因数
)

// 频率相关点位
var (
	VPointFreq = &c_base.SPoint{Key: "Freq", Name: "频率", Unit: "Hz", Desc: "电网频率", Sort: 40, Min: 45, Max: 55, Precise: 2} // 频率
)

// 电能相关点位
var (
	VPointEa = &c_base.SPoint{Key: "Ea", Name: "A相电能", Unit: "kWh", Desc: "A相累计电能", Sort: 50, Min: 0, Max: 0, Precise: 2} // A相电能
	VPointEb = &c_base.SPoint{Key: "Eb", Name: "B相电能", Unit: "kWh", Desc: "B相累计电能", Sort: 51, Min: 0, Max: 0, Precise: 2} // B相电能
	VPointEc = &c_base.SPoint{Key: "Ec", Name: "C相电能", Unit: "kWh", Desc: "C相累计电能", Sort: 52, Min: 0, Max: 0, Precise: 2} // C相电能
	VPointE  = &c_base.SPoint{Key: "E", Name: "总电能", Unit: "kWh", Desc: "三相总累计电能", Sort: 53, Min: 0, Max: 0, Precise: 2}  // 总电能
)

// 温度相关点位
var (
	VPointTemp    = &c_base.SPoint{Key: "Temp", Name: "温度", Unit: "℃", Desc: "设备温度", Sort: 60, Min: -40, Max: 85, Precise: 1}        // 温度
	VPointTempMax = &c_base.SPoint{Key: "TempMax", Name: "最高温度", Unit: "℃", Desc: "设备最高温度", Sort: 61, Min: -40, Max: 85, Precise: 1} // 最高温度
	VPointTempMin = &c_base.SPoint{Key: "TempMin", Name: "最低温度", Unit: "℃", Desc: "设备最低温度", Sort: 62, Min: -40, Max: 85, Precise: 1} // 最低温度
)

// 湿度相关点位
var (
	VPointHumidity = &c_base.SPoint{Key: "Humidity", Name: "湿度", Unit: "%RH", Desc: "环境湿度", Sort: 70, Min: 0, Max: 100, Precise: 1} // 湿度
)

// 状态相关点位
var (
	VPointStatus    = &c_base.SPoint{Key: "Status", Name: "状态", Unit: "", Desc: "设备运行状态", Sort: 80, Min: 0, Max: 255, Precise: 0}        // 状态
	VPointErrorCode = &c_base.SPoint{Key: "ErrorCode", Name: "错误代码", Unit: "", Desc: "设备错误代码", Sort: 82, Min: 0, Max: 65535, Precise: 0} // 错误代码
)

// 电池相关点位 (BMS)
var (
	VPointSOC       = &c_base.SPoint{Key: "SOC", Name: "电池电量", Unit: "%", Desc: "电池剩余电量百分比", Sort: 90, Min: 0, Max: 100, Precise: 1}        // 电池电量
	VPointSOH       = &c_base.SPoint{Key: "SOH", Name: "电池健康度", Unit: "%", Desc: "电池健康度百分比", Sort: 91, Min: 0, Max: 100, Precise: 1}        // 电池健康度
	VPointDcVoltage = &c_base.SPoint{Key: "Voltage", Name: "电池电压", Unit: "V", Desc: "电池总电压", Sort: 92, Min: 0, Max: 1000, Precise: 2}       // 电池电压
	VPointDcCurrent = &c_base.SPoint{Key: "Current", Name: "电池电流", Unit: "A", Desc: "电池充放电电流", Sort: 93, Min: -1000, Max: 1000, Precise: 2} // 电池电流
	VPointDcPower   = &c_base.SPoint{Key: "Power", Name: "电池直流功率", Unit: "kW", Desc: "电池直流功率", Sort: 94, Min: -1000, Max: 1000, Precise: 2} // 电池直流功率
)

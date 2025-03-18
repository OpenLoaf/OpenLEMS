package pcs_lnxall_v1

import "common/c_base"

var (
	Ua     = &c_base.Meta{Name: "Ua", Cn: "A相电压", Unit: "V", Addr: 40001, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Ub     = &c_base.Meta{Name: "Ub", Cn: "B相电压", Unit: "V", Addr: 40003, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Uc     = &c_base.Meta{Name: "Uc", Cn: "C相电压", Unit: "V", Addr: 40005, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Ia     = &c_base.Meta{Name: "Ia", Cn: "A相电流", Unit: "A", Addr: 40007, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Ib     = &c_base.Meta{Name: "Ib", Cn: "B相电流", Unit: "A", Addr: 40009, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Ic     = &c_base.Meta{Name: "Ic", Cn: "C相电流", Unit: "A", Addr: 40011, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Freq   = &c_base.Meta{Name: "Freq", Cn: "电网频率", Unit: "Hz", Addr: 40013, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01}
	Pa     = &c_base.Meta{Name: "Pa", Cn: "A相有功功率", Unit: "W", Addr: 40015, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Pb     = &c_base.Meta{Name: "Pb", Cn: "B相有功功率", Unit: "W", Addr: 40017, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Pc     = &c_base.Meta{Name: "Pc", Cn: "C相有功功率", Unit: "W", Addr: 40019, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Sa     = &c_base.Meta{Name: "Sa", Cn: "A相视在功率", Unit: "VA", Addr: 40021, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Sb     = &c_base.Meta{Name: "Sb", Cn: "B相视在功率", Unit: "VA", Addr: 40023, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Sc     = &c_base.Meta{Name: "Sc", Cn: "C相视在功率", Unit: "VA", Addr: 40025, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Qa     = &c_base.Meta{Name: "Qa", Cn: "A相无功功率", Unit: "var", Addr: 40027, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Qb     = &c_base.Meta{Name: "Qb", Cn: "B相无功功率", Unit: "var", Addr: 40029, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Qc     = &c_base.Meta{Name: "Qc", Cn: "C相无功功率", Unit: "var", Addr: 40031, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	PTotal = &c_base.Meta{Name: "PTotal", Cn: "总有功功率", Unit: "W", Addr: 40033, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	QTotal = &c_base.Meta{Name: "QTotal", Cn: "总无功功率", Unit: "var", Addr: 40035, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	STotal = &c_base.Meta{Name: "STotal", Cn: "总视在功率", Unit: "VA", Addr: 40037, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	PF     = &c_base.Meta{Name: "PF", Cn: "功率因数", Unit: "cosφ", Addr: 40039, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01}
)

var (
	Vbatt = &c_base.Meta{Name: "Vbatt", Cn: "电池电压", Unit: "V", Addr: 40041, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Ibatt = &c_base.Meta{Name: "Ibatt", Cn: "电池电流", Unit: "A", Addr: 40043, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Pdc   = &c_base.Meta{Name: "Pdc", Cn: "直流功率", Unit: "W", Addr: 40045, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Idc   = &c_base.Meta{Name: "Idc", Cn: "直流总电流", Unit: "A", Addr: 40047, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Vbus  = &c_base.Meta{Name: "Vbus", Cn: "总母线电压", Unit: "V", Addr: 40049, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Vp    = &c_base.Meta{Name: "Vp", Cn: "正母线电压", Unit: "V", Addr: 40051, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	Vn    = &c_base.Meta{Name: "Vn", Cn: "负母线电压", Unit: "V", Addr: 40053, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
)

var (
	WorkState          = &c_base.Meta{Name: "WorkState", Cn: "工作状态", Unit: "", Addr: 40055, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1, Desc: "0:停机;1:开机中;5:待机;9:恒压运行;32:故障;257:恒流运行;"}
	IGBTTemp           = &c_base.Meta{Name: "IGBTTemp", Cn: "IGBT温度", Unit: "℃", Addr: 40057, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	PaPF               = &c_base.Meta{Name: "PaPF", Cn: "A相功率因数", Unit: "cosφ", Addr: 40059, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01}
	PbPF               = &c_base.Meta{Name: "PbPF", Cn: "B相功率因数", Unit: "cosφ", Addr: 40061, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01}
	PcPF               = &c_base.Meta{Name: "PcPF", Cn: "C相功率因数", Unit: "cosφ", Addr: 40063, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.01}
	DcHistoryCharge    = &c_base.Meta{Name: "DcHistoryCharge", Cn: "直流历史充电量", Unit: "Ah", Addr: 40065, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	DcDayCharge        = &c_base.Meta{Name: "DcDayCharge", Cn: "直流日充电量", Unit: "Ah", Addr: 40067, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	DcHistoryDischarge = &c_base.Meta{Name: "DcHistoryDischarge", Cn: "直流历史放电量", Unit: "Ah", Addr: 40069, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	DcDayDischarge     = &c_base.Meta{Name: "DcDayDischarge", Cn: "直流日放电量", Unit: "Ah", Addr: 40071, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	AcHistoryCharge    = &c_base.Meta{Name: "AcHistoryCharge", Cn: "交流历史充电量", Unit: "kWh", Addr: 40073, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	AcDayCharge        = &c_base.Meta{Name: "AcDayCharge", Cn: "交流日充电量", Unit: "kWh", Addr: 40075, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	AcHistoryDischarge = &c_base.Meta{Name: "AcHistoryDischarge", Cn: "交流历史放电量", Unit: "kWh", Addr: 40077, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	AcDayDischarge     = &c_base.Meta{Name: "AcDayDischarge", Cn: "交流日放电量", Unit: "kWh", Addr: 40079, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	OnlineState        = &c_base.Meta{Name: "OnlineState", Cn: "在线状态", Unit: "", Addr: 40081, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1, Desc: "0:离线;1:在线;"}
)

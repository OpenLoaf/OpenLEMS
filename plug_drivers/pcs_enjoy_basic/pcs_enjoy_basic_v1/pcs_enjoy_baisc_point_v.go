package pcs_enjoy_basic_v1

import "common/c_base"

var (
	Ua               = &c_base.Meta{Addr: 0x6023, Name: "Ua", Cn: "A相电压", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "V", Desc: "Phase A voltage"}
	Ub               = &c_base.Meta{Addr: 0x6024, Name: "Ub", Cn: "B相电压", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "V", Desc: "Phase B voltage"}
	Uc               = &c_base.Meta{Addr: 0x6025, Name: "Uc", Cn: "C相电压", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "V", Desc: "Phase C voltage"}
	Ia               = &c_base.Meta{Addr: 0x6026, Name: "Ia", Cn: "A相电流", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "A", Desc: "Phase A current"}
	Ib               = &c_base.Meta{Addr: 0x6027, Name: "Ib", Cn: "B相电流", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "A", Desc: "Phase B current"}
	Ic               = &c_base.Meta{Addr: 0x6028, Name: "Ic", Cn: "C相电流", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "A", Desc: "Phase C current"}
	Freq             = &c_base.Meta{Addr: 0x602C, Name: "Freq", Cn: "电网频率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "Hz", Desc: "Grid frequency"}
	Grid_Seq         = &c_base.Meta{Addr: 0x602D, Name: "Grid_Seq", Cn: "电网相序", ReadType: c_base.RInt16, SystemType: c_base.SUint16, Factor: 1, Unit: "", Desc: "0: 正序；1: 负序"}
	Pcs_Degrade      = &c_base.Meta{Addr: 0x602E, Name: "Pcs_Degrade", Cn: "PCS 降额系数", ReadType: c_base.RInt16, SystemType: c_base.SUint16, Factor: 1, Unit: "", Desc: "正常：4096；（3686 表示降额 0.9,降额 3686/4096=0.9）--适用于 5.13.0 及以上软件版本"}
	Pcs_Degrade_Flag = &c_base.Meta{Addr: 0x602F, Name: "Pcs_Degrade_Flag", Cn: "PCS 降额标志", ReadType: c_base.RInt16, SystemType: c_base.SUint16, Factor: 1, Unit: "", Desc: "0：正常；1：IGBT 过温降额，2，环境温度降额，3：IGBT、环境温度都降额；--适用于 5.13.0 及以上软件版本"}
	Pa               = &c_base.Meta{Addr: 0x6030, Name: "Pa", Cn: "A相有功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "W", Desc: "Phase A active power"}
	Pb               = &c_base.Meta{Addr: 0x6031, Name: "Pb", Cn: "B相有功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "W", Desc: "Phase B active power"}
	Pc               = &c_base.Meta{Addr: 0x6032, Name: "Pc", Cn: "C相有功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "W", Desc: "Phase C active power"}
	Sa               = &c_base.Meta{Addr: 0x6033, Name: "Sa", Cn: "A相视在功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "VA", Desc: "Phase A apparent power"}
	Sb               = &c_base.Meta{Addr: 0x6034, Name: "Sb", Cn: "B相视在功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "VA", Desc: "Phase B apparent power"}
	Sc               = &c_base.Meta{Addr: 0x6035, Name: "Sc", Cn: "C相视在功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "VA", Desc: "Phase C apparent power"}
	Qa               = &c_base.Meta{Addr: 0x6036, Name: "Qa", Cn: "A相无功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "Var", Desc: "Phase A reactive power"}
	Qb               = &c_base.Meta{Addr: 0x6037, Name: "Qb", Cn: "B相无功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "Var", Desc: "Phase B reactive power"}
	Qc               = &c_base.Meta{Addr: 0x6038, Name: "Qc", Cn: "C相无功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "Var", Desc: "Phase C reactive power"}
)

var (
	Pt = &c_base.Meta{Addr: 0x6039, Name: "Pt", Cn: "总有功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "W", Desc: "Total active power"}
	Qt = &c_base.Meta{Addr: 0x603A, Name: "Qt", Cn: "总无功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "Var", Desc: "Total reactive power"}
	St = &c_base.Meta{Addr: 0x603B, Name: "St", Cn: "总视在功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "VA", Desc: "Total apparent power"}
	Pf = &c_base.Meta{Addr: 0x603C, Name: "Pf", Cn: "功率因数", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.01, Unit: "", Desc: "Power factor"}
)

var (
	Pcs_Status = &c_base.Meta{Addr: 0x6057, Name: "Pcs_Status", Cn: "工作状态", ReadType: c_base.RInt16, SystemType: c_base.SUint16, Factor: 1, Unit: "", Desc: "0: 停机；1：待机；2：运行；3：故障"}
	IGBT_Temp  = &c_base.Meta{Addr: 0x6058, Name: "IGBT_Temp", Cn: "IGBT 温度", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "℃", Desc: "IGBT temperature"}
	Env_Temp   = &c_base.Meta{Addr: 0x6059, Name: "Env_Temp", Cn: "环境温度", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "℃", Desc: "Environment temperature"}
)

//  控制点位
/**

下设有功充电/放电功率 0x0D57 I16，读写，单位 KW，放大 10 倍，正表示放电功率；负表示充电功率；

无功功率补偿功率设置 0x0D58 I16，读写，单位 KVar，放大 10 倍

离网模式设置 0x5066 U16, 读写， 1:离网使能； 0：并网使能
*/

var (
	Set_Ap = &c_base.Meta{Addr: 0x0D57, Name: "Set_Ap", Cn: "设置功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "KW", Desc: "Set power"}
	Set_Qp = &c_base.Meta{Addr: 0x0D58, Name: "Set_Qp", Cn: "设置无功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "KVar", Desc: "Set reactive power"}
)

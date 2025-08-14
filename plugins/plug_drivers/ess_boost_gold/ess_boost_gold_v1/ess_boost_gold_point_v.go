package ess_boost_gold_v1

import "common/c_base"

var (
	Pt     = &c_base.Meta{Addr: 100, Name: "Pt", Cn: "总有功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW", Desc: "Total active power"}
	Qt     = &c_base.Meta{Addr: 101, Name: "Qt", Cn: "总无功功率", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar", Desc: "Total reactive power"}
	Vt     = &c_base.Meta{Addr: 102, Name: "Vt", Cn: "总电压", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "V", Desc: "Total voltage"}
	It     = &c_base.Meta{Addr: 103, Name: "It", Cn: "总电流", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "A", Desc: "Total current"}
	Soc    = &c_base.Meta{Addr: 104, Name: "Soc", Cn: "电池soc", ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 1, Unit: "%", Desc: "Battery soc"}
	Temp   = &c_base.Meta{Addr: 105, Name: "Temp", Cn: "电池平均温度", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 1, Unit: "℃", Desc: "Average battery temperature"}
	Diff   = &c_base.Meta{Addr: 106, Name: "Diff", Cn: "电池最大温差", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 1, Unit: "℃", Desc: "Maximum temperature difference of battery"}
	Vmax   = &c_base.Meta{Addr: 107, Name: "Vmax", Cn: "最高单体电压值", ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Unit: "V", Desc: "Maximum voltage of single cell"}
	Vmin   = &c_base.Meta{Addr: 108, Name: "Vmin", Cn: "最低单体电压值", ReadType: c_base.RUint16, SystemType: c_base.SFloat32, Factor: 0.001, Unit: "V", Desc: "Minimum voltage of single cell"}
	Status = &c_base.Meta{Addr: 111, Name: "Status", Cn: "状态", ReadType: c_base.RUint16}
)

var CONTROL_ON_OFF = &c_base.Meta{Addr: 2000, Name: "CONTROL_ON_OFF", Cn: "总起停指令", ReadType: c_base.RUint16, Desc: "0:关机;1:开机;2:复位;3:测试"}

var (
	Set_Ap_Power = &c_base.Meta{Addr: 105, Name: "Set_Ap_Power", Cn: "总有功功率指令", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kW", Desc: "Total active power command"}
	Set_Rp_Power = &c_base.Meta{Addr: 105, Name: "Set_Rp_Power", Cn: "总无功功率指令", ReadType: c_base.RInt16, SystemType: c_base.SFloat32, Factor: 0.1, Unit: "kVar", Desc: "Total reactive power command"}
)

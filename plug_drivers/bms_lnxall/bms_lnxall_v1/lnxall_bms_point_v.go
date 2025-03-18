package bms_lnxall_v1

import "common/c_base"

var (
	BATT_VOLTAGE    = &c_base.Meta{Name: "BATT_VOLTAGE", Cn: "电池簇电压", Unit: "V", Addr: 40001, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	BATT_CURRENT    = &c_base.Meta{Name: "BATT_CURRENT", Cn: "电池簇电流值", Unit: "A", Addr: 40003, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	RACK_SOC        = &c_base.Meta{Name: "RACK_SOC", Cn: "RackSOC", Unit: "%", Addr: 40005, ReadType: c_base.RFloat32, SystemType: c_base.SFloat32, Factor: 1}
	RACK_SOH        = &c_base.Meta{Name: "RACK_SOH", Cn: "RackSOH", Unit: "%", Addr: 40007, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	RACK_R_PLUS     = &c_base.Meta{Name: "RACK_R_PLUS", Cn: "电池簇绝缘电阻R+", Unit: "KΩ", Addr: 40009, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	RACK_R_MINUS    = &c_base.Meta{Name: "RACK_R_MINUS", Cn: "电池簇绝缘电阻R-", Unit: "KΩ", Addr: 40011, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	RACK_STATUS     = &c_base.Meta{Name: "RACK_STATUS", Cn: "电池簇电池状态", Unit: "", Addr: 40013, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1, Desc: "1:充电;2:放电;3:开路"}
	DI_STATUS       = &c_base.Meta{Name: "DI_STATUS", Cn: "DI检测状态", Unit: "", Addr: 40015, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	DO_STATUS       = &c_base.Meta{Name: "DO_STATUS", Cn: "DO输出状态", Unit: "", Addr: 40017, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	TEMP_NUM        = &c_base.Meta{Name: "TEMP_NUM", Cn: "实际温度采集点数", Unit: "", Addr: 40019, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	TEMP_MAX        = &c_base.Meta{Name: "TEMP_MAX", Cn: "电池最高温度", Unit: "℃", Addr: 40021, ReadType: c_base.RFloat32, SystemType: c_base.SFloat32, Factor: 1}
	TEMP_MAX_RACK   = &c_base.Meta{Name: "TEMP_MAX_RACK", Cn: "电池最高温度所在模块号", Unit: "", Addr: 40023, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	TEMP_MAX_MODULE = &c_base.Meta{Name: "TEMP_MAX_MODULE", Cn: "电池最高温度模块内序号", Unit: "", Addr: 40025, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	TEMP_MIN        = &c_base.Meta{Name: "TEMP_MIN", Cn: "电池最低温度", Unit: "℃", Addr: 40027, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 1}
	TEMP_MIN_RACK   = &c_base.Meta{Name: "TEMP_MIN_RACK", Cn: "电池最低温度所在模块号", Unit: "", Addr: 40029, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	TEMP_MIN_MODULE = &c_base.Meta{Name: "TEMP_MIN_MODULE", Cn: "电池最低温度模块内序号", Unit: "", Addr: 40031, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	TEMP_AVG        = &c_base.Meta{Name: "TEMP_AVG", Cn: "电池平均温度", Unit: "℃", Addr: 40033, ReadType: c_base.RFloat32, SystemType: c_base.SFloat32, Factor: 1}
	BATT_NUM        = &c_base.Meta{Name: "BATT_NUM", Cn: "电池组电池总节数", Unit: "", Addr: 40035, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	BATT_AVG        = &c_base.Meta{Name: "BATT_AVG", Cn: "单体平均电压", Unit: "V", Addr: 40037, ReadType: c_base.RFloat32, SystemType: c_base.SFloat32, Factor: 1}

	BATT_MAX        = &c_base.Meta{Name: "BATT_MAX", Cn: "最高单体电压", Unit: "V", Addr: 40039, ReadType: c_base.RFloat32, SystemType: c_base.SFloat32, Factor: 1}
	BATT_MAX_RACK   = &c_base.Meta{Name: "BATT_MAX_RACK", Cn: "最高单体电压所在模块号", Unit: "", Addr: 40041, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	BATT_MAX_MODULE = &c_base.Meta{Name: "BATT_MAX_MODULE", Cn: "最高单体电压模块内序号", Unit: "", Addr: 40043, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1}
	BATT_MIN        = &c_base.Meta{Name: "BATT_MIN", Cn: "最低单体电压", Unit: "V", Addr: 40045, ReadType: c_base.RFloat32, SystemType: c_base.SFloat32, Factor: 1}
)

var (
	TOTAL_CHARGE_ENERGY     = &c_base.Meta{Name: "TOTAL_CHARGE_ENERGY", Cn: "累计充电电量", Unit: "kWh", Addr: 40087, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	TOTAL_DISCHARGE_ENERGY  = &c_base.Meta{Name: "TOTAL_DISCHARGE_ENERGY", Cn: "累计放电电量", Unit: "kWh", Addr: 40089, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	SINGLE_CHARGE_ENERGY    = &c_base.Meta{Name: "SINGLE_CHARGE_ENERGY", Cn: "单次充电电量", Unit: "kWh", Addr: 40091, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	SINGLE_DISCHARGE_ENERGY = &c_base.Meta{Name: "SINGLE_DISCHARGE_ENERGY", Cn: "单次放电电量", Unit: "kWh", Addr: 40093, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Factor: 0.1}
	LINKED_STATUS           = &c_base.Meta{Name: "LINKED_STATUS", Cn: "在线状态", Unit: "", Addr: 40095, ReadType: c_base.RUint32, SystemType: c_base.SUint16, Factor: 1, Desc: "0:离线;1:在线;"}
)

package ess_boost_lnxall_v1

import "common/c_base"

var (
	ESS_ON_OFF       = &c_base.Meta{Name: "ESS_ON_OFF", Cn: "开关机", ReadType: c_base.RUint16, Desc: "0:关机;1:开机"}
	ESS_SET_AP_POWER = &c_base.Meta{Name: "ESS_SET_AP_POWER", Cn: "设置有功功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: "设置有功功率"}
	ESS_SET_RP_POWER = &c_base.Meta{Name: "ESS_SET_RP_POWER", Cn: "设置无功功率", Unit: "kVAR", ReadType: c_base.RFloat32, Desc: "设置无功功率"}
)

var (
	LIMIT_ESS_POWER_ENABLE          = &c_base.Meta{Name: "LIMIT_ESS_POWER_ENABLE", Cn: "EMS限功率运行", ReadType: c_base.RUint32, Desc: "0:未限制;1:已限制"}
	ESS_INVERSE_FLOW_ENABLE         = &c_base.Meta{Name: "ESS_INVERSE_FLOW_ENABLE", Cn: "EMS逆流运行", ReadType: c_base.RUint32, Desc: "0:未逆流;1:已逆流"}
	ESS_MAX_DEMAND                  = &c_base.Meta{Name: "ESS_MAX_DEMAND", Cn: "最大需量", ReadType: c_base.RUint32, Desc: ""}
	ESS_SYSTEM_STATUS               = &c_base.Meta{Name: "ESS_SYSTEM_STATUS", Cn: "系统运行状态", ReadType: c_base.RUint32, Desc: "0:正常;1:异常"}
	ESS_INVERSE_WINDOW_POWER        = &c_base.Meta{Name: "ESS_INVERSE_WINDOW_POWER", Cn: "逆流窗口功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_LOAD_POWER                  = &c_base.Meta{Name: "ESS_LOAD_POWER", Cn: "负载功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_EXPECTED_POWER              = &c_base.Meta{Name: "ESS_EXPECTED_POWER", Cn: "期望有功功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_LC_INPUT_POWER              = &c_base.Meta{Name: "ESS_LC_INPUT_POWER", Cn: "LC输入功率", ReadType: c_base.RFloat32, Desc: ""}
	ESS_RATED_POWER                 = &c_base.Meta{Name: "ESS_RATED_POWER", Cn: "额定功率", Unit: "kW", ReadType: c_base.RUint32, Desc: ""}
	ESS_RATED_CAPACITY              = &c_base.Meta{Name: "ESS_RATED_CAPACITY", Cn: "额定容量", ReadType: c_base.RUint32, Desc: ""}
	ESS_ALARM_NUM                   = &c_base.Meta{Name: "ESS_ALARM_NUM", Cn: "告警数量", ReadType: c_base.RUint32, Desc: ""}
	ESS_PROTECT_DISABLE             = &c_base.Meta{Name: "ESS_PROTECT_DISABLE", Cn: "禁充保护", ReadType: c_base.RUint32, Desc: "0:未保护;1:已保护"}
	ESS_PROTECT_ENABLE              = &c_base.Meta{Name: "ESS_PROTECT_ENABLE", Cn: "禁放保护", ReadType: c_base.RUint32, Desc: "0:未保护;1:已保护"}
	ESS_LAST_POWER                  = &c_base.Meta{Name: "ESS_LAST_POWER", Cn: "上次充放功率", ReadType: c_base.RUint32, Desc: ""}
	ESS_MAX_CHARGE_POWER            = &c_base.Meta{Name: "ESS_MAX_CHARGE_POWER", Cn: "最大可充功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_MAX_DISCHARGE_POWER         = &c_base.Meta{Name: "ESS_MAX_DISCHARGE_POWER", Cn: "最大可放功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_PV_PERCENT                  = &c_base.Meta{Name: "ESS_PV_PERCENT", Cn: "光伏设置百分比", Unit: "%", ReadType: c_base.RFloat32, Desc: ""}
	ESS_CHARGE_LIMIT_ENABLE         = &c_base.Meta{Name: "ESS_CHARGE_LIMIT_ENABLE", Cn: "充电功率限制", ReadType: c_base.RUint32, Desc: "0:正常;1:已限制"}
	ESS_DISCHARGE_LIMIT_ENABLE      = &c_base.Meta{Name: "ESS_DISCHARGE_LIMIT_ENABLE", Cn: "放电功率限制", ReadType: c_base.RUint32, Desc: "0:正常;1:已限制"}
	ESS_TARGET_HEAT_TEMP            = &c_base.Meta{Name: "ESS_TARGET_HEAT_TEMP", Cn: "温控目标制热温度", Unit: "℃", ReadType: c_base.RFloat32, Desc: ""}
	ESS_TARGET_COOL_TEMP            = &c_base.Meta{Name: "ESS_TARGET_COOL_TEMP", Cn: "温控目标制冷温度", Unit: "℃", ReadType: c_base.RFloat32, Desc: ""}
	ESS_CABINET_STATUS              = &c_base.Meta{Name: "ESS_CABINET_STATUS", Cn: "柜1运行状态", ReadType: c_base.RUint32, Desc: "0:正常;1:异常"}
	ESS_CABINET_PROTECT_DISABLE     = &c_base.Meta{Name: "ESS_CABINET_PROTECT_DISABLE", Cn: "柜1禁充保护", ReadType: c_base.RUint32, Desc: "0:未保护;1:已保护"}
	ESS_CABINET_PROTECT_ENABLE      = &c_base.Meta{Name: "ESS_CABINET_PROTECT_ENABLE", Cn: "柜1禁放保护", ReadType: c_base.RUint32, Desc: "0:未保护;1:已保护"}
	ESS_CABINET_MAX_CHARGE_POWER    = &c_base.Meta{Name: "ESS_CABINET_MAX_CHARGE_POWER", Cn: "柜1最大可充功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_CABINET_MAX_DISCHARGE_POWER = &c_base.Meta{Name: "ESS_CABINET_MAX_DISCHARGE_POWER", Cn: "柜1最大可放功率", Unit: "kW", ReadType: c_base.RFloat32, Desc: ""}
	ESS_ONLINE_STATUS               = &c_base.Meta{Name: "ESS_ONLINE_STATUS", Cn: "在线状态", ReadType: c_base.RUint32, Desc: "0:离线;1:在线"}
)

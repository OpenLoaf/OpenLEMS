package ess_pylon_checkwatt_v1

import (
	"common/c_base"
	"common/c_enum"
)

var (
	// 遥测点位定义 - 直接创建，启动时验证SPoint字段
	telemetrySocPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "soc",
			Name:      "SOC",
			Unit:      "%",
			ValueType: c_enum.EFloat32,
			Desc:      "电池电量",
			Min:       0,
			Max:       100,
		},
		MethodName: "GetSoc",
	}

	telemetrySohPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "soh",
			Name:      "SOH",
			Unit:      "%",
			ValueType: c_enum.EFloat32,
			Desc:      "电池健康度",
			Min:       0,
			Max:       100,
		},
		MethodName: "GetSoh",
	}

	telemetryPowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "power",
			Name:      "功率",
			Unit:      "kW",
			ValueType: c_enum.EFloat64,
			Desc:      "当前功率",
		},
		MethodName: "GetPower",
	}

	telemetryCapacityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "capacity",
			Name:      "容量",
			Unit:      "kWh",
			ValueType: c_enum.EUint32,
			Desc:      "电池容量",
		},
		MethodName: "GetCapacity",
	}

	telemetryTargetPowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "targetPower",
			Name:      "目标功率",
			Unit:      "kW",
			ValueType: c_enum.EInt32,
			Desc:      "目标功率",
		},
		MethodName: "GetTargetPower",
	}

	telemetryApparentPowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "apparentPower",
			Name:      "视在功率",
			Unit:      "kVA",
			ValueType: c_enum.EFloat64,
			Desc:      "视在功率",
		},
		MethodName: "GetApparentPower",
	}

	telemetryReactivePowerPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "reactivePower",
			Name:      "无功功率",
			Unit:      "kVar",
			ValueType: c_enum.EFloat64,
			Desc:      "无功功率",
		},
		MethodName: "GetReactivePower",
	}

	telemetryTodayIncomingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "todayIncomingQuantity",
			Name:      "当日充电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "当日充电量",
		},
		MethodName: "GetTodayIncomingQuantity",
	}

	telemetryTodayOutgoingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "todayOutgoingQuantity",
			Name:      "当日放电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "当日放电量",
		},
		MethodName: "GetTodayOutgoingQuantity",
	}

	telemetryHistoryIncomingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyIncomingQuantity",
			Name:      "历史充电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "历史充电量",
		},
		MethodName: "GetHistoryIncomingQuantity",
	}

	telemetryHistoryOutgoingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyOutgoingQuantity",
			Name:      "历史放电量",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "历史放电量",
		},
		MethodName: "GetHistoryOutgoingQuantity",
	}
)

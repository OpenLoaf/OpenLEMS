package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"time"
)

var (
	// 遥测点位定义 - 直接创建，启动时验证SPoint字段
	telemetryPTotalPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "pTotal",
			Name:      "功率",
			Unit:      "kW",
			ValueType: c_enum.EFloat32,
			Desc:      "总有功功率",
		},
		MethodName: "GetPTotal",
	}

	telemetryFrequencyPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "frequency",
			Name:      "频率",
			Unit:      "Hz",
			ValueType: c_enum.EFloat32,
			Desc:      "系统频率",
		},
		MethodName: "GetFrequency",
	}

	telemetryPfTotalPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "pfTotal",
			Name:      "功率因素",
			ValueType: c_enum.EFloat32,
			Desc:      "总功率因数",
		},
		MethodName: "GetPfTotal",
	}

	telemetryHistoryIncomingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyIncomingQuantity",
			Name:      "正向总有功电能",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "正向总有功电能",
		},
		MethodName: "GetHistoryIncomingQuantity",
	}

	telemetryHistoryOutgoingQuantityPoint = &c_base.SReflectPoint{
		SPoint: &c_base.SPoint{
			Key:       "historyOutgoingQuantity",
			Name:      "反向总有功电能",
			Unit:      "kWh",
			ValueType: c_enum.EFloat64,
			Desc:      "反向总有功电能",
		},
		MethodName: "GetHistoryOutgoingQuantity",
	}

	// 协议点位定义 - 直接创建，启动时验证SPoint字段
	Status = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Status",
				Name:      "设备状态字",
				ValueType: c_enum.EInt16,
				Desc:      "设备状态字",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0190,
	}

	PhaseAVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseAVoltage",
				Name:      "A相电压",
				Unit:      "V",
				ValueType: c_enum.EInt16,
				Desc:      "A相电压",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0191,
	}

	PhaseBVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseBVoltage",
				Name:      "B相电压",
				Unit:      "V",
				ValueType: c_enum.EInt16,
				Desc:      "B相电压",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0192,
	}

	PhaseCVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseCVoltage",
				Name:      "C相电压",
				Unit:      "V",
				ValueType: c_enum.EInt16,
				Desc:      "C相电压",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0193,
	}

	PhaseACurrent = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseACurrent",
				Name:      "A相电流",
				Unit:      "A",
				ValueType: c_enum.EInt16,
				Desc:      "A相电流",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0194,
	}

	PhaseBCurrent = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseBCurrent",
				Name:      "B相电流",
				Unit:      "A",
				ValueType: c_enum.EInt16,
				Desc:      "B相电流",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0195,
	}

	PhaseCCurrent = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseCCurrent",
				Name:      "C相电流",
				Unit:      "A",
				ValueType: c_enum.EInt16,
				Desc:      "C相电流",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x0196,
	}

	PhaseAActivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseAActivePower",
				Name:      "A相有功功率",
				Unit:      "kW",
				ValueType: c_enum.EInt16,
				Desc:      "A相有功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0197,
	}

	PhaseBActivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseBActivePower",
				Name:      "B相有功功率",
				Unit:      "kW",
				ValueType: c_enum.EInt16,
				Desc:      "B相有功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0198,
	}

	PhaseCActivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseCActivePower",
				Name:      "C相有功功率",
				Unit:      "kW",
				ValueType: c_enum.EInt16,
				Desc:      "C相有功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x0199,
	}

	PhaseAReactivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseAReactivePower",
				Name:      "A相无功功率",
				Unit:      "kVar",
				ValueType: c_enum.EInt16,
				Desc:      "A相无功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x019A,
	}

	PhaseBReactivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseBReactivePower",
				Name:      "B相无功功率",
				Unit:      "kVar",
				ValueType: c_enum.EInt16,
				Desc:      "B相无功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x019B,
	}

	PhaseCReactivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseCReactivePower",
				Name:      "C相无功功率",
				Unit:      "kVar",
				ValueType: c_enum.EInt16,
				Desc:      "C相无功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x019C,
	}

	PhaseAApparentPower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseAApparentPower",
				Name:      "A相视在功率",
				Unit:      "kVA",
				ValueType: c_enum.EInt16,
				Desc:      "A相视在功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x019D,
	}

	PhaseBApparentPower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseBApparentPower",
				Name:      "B相视在功率",
				Unit:      "kVA",
				ValueType: c_enum.EInt16,
				Desc:      "B相视在功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x019E,
	}

	PhaseCApparentPower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PhaseCApparentPower",
				Name:      "C相视在功率",
				Unit:      "kVA",
				ValueType: c_enum.EInt16,
				Desc:      "C相视在功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x019F,
	}

	TotalActivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TotalActivePower",
				Name:      "总有功功率",
				Unit:      "kW",
				ValueType: c_enum.EInt16,
				Desc:      "总有功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x01A0,
	}

	TotalReactivePower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TotalReactivePower",
				Name:      "总无功功率",
				Unit:      "kVar",
				ValueType: c_enum.EInt16,
				Desc:      "总无功功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x01A1,
	}

	TotalApparentPower = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "TotalApparentPower",
				Name:      "总视在功率",
				Unit:      "kVA",
				ValueType: c_enum.EInt16,
				Desc:      "总视在功率",
			},
			DataAccess: c_default.VDataAccessInt16Scale01,
		},
		Addr: 0x01A2,
	}

	ForwardActiveEnergy = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "ForwardActiveEnergy",
				Name:      "正向有功电量",
				Unit:      "kWh",
				ValueType: c_enum.EUint16,
				Desc:      "正向有功电量",
			},
			DataAccess: c_default.VDataAccessUInt16Scale001,
		},
		Addr: 0x01A3,
	}

	ReverseActiveEnergy = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "ReverseActiveEnergy",
				Name:      "反向有功电量",
				Unit:      "kWh",
				ValueType: c_enum.EUint16,
				Desc:      "反向有功电量",
			},
			DataAccess: c_default.VDataAccessUInt16Scale001,
		},
		Addr: 0x01A4,
	}

	Frequency = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "Frequency",
				Name:      "频率",
				Unit:      "Hz",
				ValueType: c_enum.EInt16,
				Desc:      "频率",
			},
			DataAccess: c_default.VDataAccessInt16Scale001,
		},
		Addr: 0x01A5,
	}

	PowerFactor = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "PowerFactor",
				Name:      "功率因数",
				ValueType: c_enum.EInt16,
				Desc:      "功率因数",
			},
			DataAccess: c_default.VDataAccessInt16Scale001,
		},
		Addr: 0x01A6,
	}

	RatedLineVoltage = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "RatedLineVoltage",
				Name:      "额定线电压",
				Unit:      "V",
				ValueType: c_enum.EInt16,
				Desc:      "额定线电压",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x01A7,
	}

	RatedFrequency = &c_proto.SModbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       "RatedFrequency",
				Name:      "额定频率",
				Unit:      "Hz",
				ValueType: c_enum.EInt16,
				Desc:      "额定频率",
			},
			DataAccess: c_default.VDataAccessInt16,
		},
		Addr: 0x01A8,
	}
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:      "ReadTask",
	Addr:      Status.Addr,
	Quantity:  RatedFrequency.Addr - Status.Addr + 1,
	Function:  c_enum.EMqHoldingRegisters,
	CycleMill: 500,
	Lifetime:  3 * time.Second,
	Points: []*c_proto.SModbusPoint{
		Status, PhaseAVoltage, PhaseBVoltage, PhaseCVoltage,
		PhaseACurrent, PhaseBCurrent, PhaseCCurrent,
		PhaseAActivePower, PhaseBActivePower, PhaseCActivePower,
		PhaseAReactivePower, PhaseBReactivePower, PhaseCReactivePower,
		PhaseAApparentPower, PhaseBApparentPower, PhaseCApparentPower,
		TotalActivePower, TotalReactivePower, TotalApparentPower,
		ForwardActiveEnergy, ReverseActiveEnergy,
		Frequency, PowerFactor, RatedLineVoltage, RatedFrequency,
	},
}

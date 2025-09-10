package pv_demo_v1

import (
	"common/c_base"
	"common/c_proto"
	"fmt"
	"github.com/shockerli/cvt"
	"time"
)

var (
	// 根据规格表定义光伏设备点表
	Status          = &c_base.SModbusPoint{Name: "Status", Cn: "设备状态字", Addr: 0x0064, ReadType: c_base.RInt16}
	Power           = &c_base.SModbusPoint{Name: "Power", Cn: "当前功率", Addr: 0x0065, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	GeneratedEnergy = &c_base.SModbusPoint{Name: "GeneratedEnergy", Cn: "累计发电量", Addr: 0x0066, ReadType: c_base.RUint16, Factor: 0.1, Unit: "kWh"}
	OnOffState      = &c_base.SModbusPoint{Name: "OnOffState", Cn: "开机状态", Addr: 0x0067, ReadType: c_base.RInt16, Factor: 0.1, StatusExplain: func(value any) string {
		if v, err := cvt.Uint8E(value); err == nil {
			switch v {
			case 0:
				return "关机"
			case 1:
				return "开机"
			}
		}
		return fmt.Sprintf("未知值: %v", value)
	}}
	PowerLimit        = &c_base.SModbusPoint{Name: "PowerLimit", Cn: "功率限制值", Addr: 0x0068, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	InstalledCapacity = &c_base.SModbusPoint{Name: "InstalledCapacity", Cn: "装机容量", Addr: 0x0069, ReadType: c_base.RInt16, Factor: 0.1, Unit: "kW"}
	Irradiance        = &c_base.SModbusPoint{Name: "Irradiance", Cn: "当前辐照度", Addr: 0x006A, ReadType: c_base.RInt16, Factor: 0.1, Unit: "W/m²"}
	Temperature       = &c_base.SModbusPoint{Name: "Temperature", Cn: "当前温度", Addr: 0x006B, ReadType: c_base.RInt16, Factor: 0.1, Unit: "°C"}
	Efficiency        = &c_base.SModbusPoint{Name: "Efficiency", Cn: "转换效率", Addr: 0x006C, ReadType: c_base.RInt16, Factor: 0.01}
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:        "ReadTask",
	DisplayName: "查询数据",
	Addr:        Status.Addr,
	Quantity:    Efficiency.Addr - Status.Addr + 1,
	Function:    c_enum.EMqHoldingRegisters,
	CycleMill:   500,
	Lifetime:    3 * time.Second,
	Metas: []*c_base.SModbusPoint{
		Status, Power, GeneratedEnergy, OnOffState, PowerLimit, InstalledCapacity, Irradiance, Temperature, Efficiency,
	},
}

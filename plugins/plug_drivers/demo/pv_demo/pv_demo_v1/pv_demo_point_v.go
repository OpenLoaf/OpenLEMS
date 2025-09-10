package pv_demo_v1

import (
	"common/c_base"
	"common/c_default"
	"common/c_enum"
	"common/c_proto"
	"fmt"
	"time"

	"github.com/shockerli/cvt"
)

var (
	// 根据规格表定义光伏设备点表
	Status          = &c_proto.SModbusPoint{Addr: 0x0064, SPoint: &c_base.SPoint{Key: "Status", Name: "设备状态字", Desc: "设备状态字"}, DataAccess: c_default.VDataAccessInt16}
	Power           = &c_proto.SModbusPoint{Addr: 0x0065, SPoint: &c_base.SPoint{Key: "Power", Name: "当前功率", Unit: "kW", Desc: "当前功率"}, DataAccess: c_default.VDataAccessInt16Scale01}
	GeneratedEnergy = &c_proto.SModbusPoint{Addr: 0x0066, SPoint: &c_base.SPoint{Key: "GeneratedEnergy", Name: "累计发电量", Unit: "kWh", Desc: "累计发电量"}, DataAccess: c_default.VDataAccessUInt16Scale01}
	OnOffState      = &c_proto.SModbusPoint{
		Addr:       0x0067,
		SPoint:     &c_base.SPoint{Key: "OnOffState", Name: "开机状态", Desc: "开机状态"},
		DataAccess: c_default.VDataAccessInt16Scale01,
		StatusExplain: func(value any) (string, error) {
			if v, err := cvt.Uint8E(value); err == nil {
				switch v {
				case 0:
					return "关机", nil
				case 1:
					return "开机", nil
				}
			}
			return fmt.Sprintf("未知值: %v", value), nil
		},
	}
	PowerLimit        = &c_proto.SModbusPoint{Addr: 0x0068, SPoint: &c_base.SPoint{Key: "PowerLimit", Name: "功率限制值", Unit: "kW", Desc: "功率限制值"}, DataAccess: c_default.VDataAccessInt16Scale01}
	InstalledCapacity = &c_proto.SModbusPoint{Addr: 0x0069, SPoint: &c_base.SPoint{Key: "InstalledCapacity", Name: "装机容量", Unit: "kW", Desc: "装机容量"}, DataAccess: c_default.VDataAccessInt16Scale01}
	Irradiance        = &c_proto.SModbusPoint{Addr: 0x006A, SPoint: &c_base.SPoint{Key: "Irradiance", Name: "当前辐照度", Unit: "W/m²", Desc: "当前辐照度"}, DataAccess: c_default.VDataAccessInt16Scale01}
	Temperature       = &c_proto.SModbusPoint{Addr: 0x006B, SPoint: &c_base.SPoint{Key: "Temperature", Name: "当前温度", Unit: "°C", Desc: "当前温度"}, DataAccess: c_default.VDataAccessInt16Scale01}
	Efficiency        = &c_proto.SModbusPoint{Addr: 0x006C, SPoint: &c_base.SPoint{Key: "Efficiency", Name: "转换效率", Desc: "转换效率"}, DataAccess: c_default.VDataAccessInt16Scale001}
)

var ReadTask = &c_proto.SModbusPointTask{
	Name:      "ReadTask",
	Addr:      Status.Addr,
	Quantity:  Efficiency.Addr - Status.Addr + 1,
	Function:  c_enum.EMqHoldingRegisters,
	CycleMill: 500,
	Lifetime:  3 * time.Second,
	Points: []*c_proto.SModbusPoint{
		Status, Power, GeneratedEnergy, OnOffState, PowerLimit, InstalledCapacity, Irradiance, Temperature, Efficiency,
	},
}

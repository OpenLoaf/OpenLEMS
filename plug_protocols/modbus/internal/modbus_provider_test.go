package internal

/*
import (
	"context"
	"common/c_base"
	common_analysis "example/ems/common-analysis"
	"example/ems/common/config"
	"modbus/p_modbus"
	"testing"
	"time"
)

var (
	Ua = &c_base.Meta{Id: "Ua", Addr: 0, ReadType: c_base.RInt16, Factor: 0.1, Min: 0, Max: 400, Unit: "V", Desc: "A相电压", Trigger: nil}
	Ub = &c_base.Meta{Id: "Ub", Addr: 1, ReadType: c_base.RInt16, Factor: 0.1, Min: 0, Max: 400, Unit: "V", Desc: "B相电压", Trigger: nil}
	Uc = &c_base.Meta{Id: "Uc", Addr: 2, ReadType: c_base.RInt16, Factor: 0.1, Min: 0, Max: 400, Unit: "V", Desc: "C相电压", Trigger: nil}
	P  = &c_base.Meta{Id: "P", Addr: 3, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Min: -1000, Max: 1000, Unit: "KWh", Desc: "有功功率", Trigger: nil}
	Q  = &c_base.Meta{Id: "Q", Addr: 5, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Min: -1000, Max: 1000, Unit: "KVar", Desc: "无功功率", Trigger: nil}

	TargetPower = &c_base.Meta{Id: "TargetPower", Addr: 7, ReadType: c_base.RUint32, SystemType: c_base.SFloat32, Endianness: c_base.EMiddleEndian, Factor: 0.0001, Min: -1000, Max: 1000, Unit: "KWh", Desc: "目标功率", Trigger: nil}
)

var (
	Vol = &p_modbus.ModbusGroup{Addr: Ua.Addr, Quantity: 3, Function: p_modbus.MqHoldingRegisters, CycleMill: 1000, Lifetime: 0, Id: "Vol",
		Metas: []*c_base.Meta{Ua, Ub, Uc}}
	Power = &p_modbus.ModbusGroup{Addr: P.Addr, Quantity: 4, Function: p_modbus.MqHoldingRegisters, CycleMill: 200, Lifetime: 10 * time.Second, Id: "Power",
		Metas: []*c_base.Meta{P, Q}}
)

func NewTcpClient() (context.Context, p_modbus.IModbusProtocol, error) {
	ctx := context.TODO()
	protocolConfig := &c_base.SProtocolConfig{
		Id:     "测试Modbus",
		Protocol: "modbus_tcp",
		Address:  "127.0.0.1:1504",
		Enable:   true,
	}
	deviceConfig := &p_modbus.SModbusDeviceConfig{
		UnitId: 1,
		DeviceConfig: config.DeviceConfig{
			Id: "测试设备",
		},
	}

	client, broadcaster := common_analysis.NewModbusClient(ctx, protocolConfig)

	modbusProvider, err := NewModbusProvider(ctx, protocolConfig, deviceConfig, client, broadcaster) // 进行直接掉用

	return ctx, modbusProvider, err
}

func TestModbusProvider_RegisterRead(t *testing.T) {
	_, client, err := NewTcpClient()
	if err != nil {
		t.Error(err)
	}
	/*	vars, err := client.ReadGroupSync(Vol, false, Ua)
		if err != nil {
			t.Error(err)
		}
		fmt.Printf("Ua: %v\n", vars[0].Float32())

		//client.PrintCacheValues()

		// 读2位以上的寄存器值
		fmt.Println("读2位以上的寄存器值")
		pVar, err := client.ReadSingleSync(P, p_modbus.MqHoldingRegisters, -1, false)
		fmt.Println(P.ValueToString(pVar))

		client.PrintCacheValues()
*/
// 设置TargetPower
//
//	setPower := 10002
//	err = client.WriteSingleRegister(TargetPower, int64(setPower))
//	if err != nil {
//		t.Error(err)
//	}
//
//	v, err := client.ReadSingleSync(TargetPower, p_modbus.MqHoldingRegisters, -1, false)
//	if err != nil {
//		return
//	}
//	if v.Int() != setPower {
//		t.Errorf("set power failed, %v", v.Int())
//	}
//}
//
//func BenchmarkTestWriteAndRead(b *testing.B) {
//	_, client, err := NewTcpClient()
//	if err != nil {
//		b.Error(err)
//	}
//	for i := 0; i < b.N; i++ {
//		setPower := i
//		err = client.WriteSingleRegister(TargetPower, int64(setPower))
//		if err != nil {
//			b.Error(err)
//		}
//
//		v, err := client.ReadSingleSync(TargetPower, p_modbus.MqHoldingRegisters, -1, false)
//		if err != nil {
//			return
//		}
//		if v.Int() != setPower {
//			b.Errorf("set power failed, %v", v.Int())
//		}
//	}
//}*/

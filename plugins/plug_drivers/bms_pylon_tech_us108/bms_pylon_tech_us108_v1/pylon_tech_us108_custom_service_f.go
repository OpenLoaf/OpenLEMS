package bms_pylon_tech_us108_v1

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"fmt"
)

func (p *sBmsPylonTechUs108) CustomDcOffService() error {
	fmt.Println("CustomDcOffService")
	c_log.BizInfo(p.DeviceCtx, "触发直流下电指令")
	return nil
}

func (p *sBmsPylonTechUs108) CustomDcOnService() error {
	fmt.Println("CustomDcOnService")
	c_log.BizInfo(p.DeviceCtx, "触发直流上电指令")
	return nil
}

func (p *sBmsPylonTechUs108) CustomSyncTimeService() error {
	c_log.BizInfo(p.DeviceCtx, "触发时间同步指令")
	return nil
}

func (p *sBmsPylonTechUs108) CustomTestTrigger() error {
	c_log.BizInfo(p.DeviceCtx, "触发测试温度告警触发指令")
	return p.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		return protocol.WriteSingleRegister(&c_base.Meta{
			Addr:     0x1100,
			ReadType: c_base.RUint16,
		}, 0x1202)
	})
}

func (p *sBmsPylonTechUs108) CustomTestClear() error {
	c_log.BizInfo(p.DeviceCtx, "触发测试温度告警清除指令")
	return p.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		return protocol.WriteSingleRegister(&c_base.Meta{
			Addr:     0x1100,
			ReadType: c_base.RUint16,
		}, 0x1002)
	})
}

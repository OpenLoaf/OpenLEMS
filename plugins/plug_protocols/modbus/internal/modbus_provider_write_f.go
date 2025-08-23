package internal

import (
	"c_protocol"
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (p *ModbusProtocolProvider) WriteSingleRegister(meta *c_base.Meta, value int32) error {
	if !p.IsActivate() {
		return fmt.Errorf("device %s connect is not activated", p.deviceConfig.Id)
	}
	//result := meta.ReadType.Encoder(int64(value), meta.Factor, meta.Offset, meta.Endianness)
	result := c_protocol.ReadTypeEncoder(meta.ReadType, int64(value), meta.Factor, meta.Offset, meta.Endianness)
	// 通关result来计算需要多少个寄存器位置
	registerLength := len(result) / 2
	if registerLength == 1 {
		//uint16Value := meta.Endianness.DecodeToUint16(result)
		uint16Value := c_protocol.ECharSequenceDecodeToUint16(meta.Endianness, result)
		p.log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceConfig.Id, meta.Name, meta.Addr, uint16Value)
		err := p.client.WriteSingleRegister(p.modbusDeviceConfig.UnitId, meta.Addr, uint16Value)
		if err != nil {
			return err
		}
	} else {
		p.log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceConfig.Id, meta.Name, meta.Addr, result)
		err := p.client.WriteMultipleRegistersBytes(p.modbusDeviceConfig.UnitId, meta.Addr, uint16(registerLength), result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ModbusProtocolProvider) WriteMultipleRegisters(group *c_proto.SModbusTask, values []int64) error {
	if !p.IsActivate() {
		return fmt.Errorf("device %s connect is not activated", p.deviceConfig.Id)
	}
	dataLength := group.Quantity
	if len(group.Metas) != len(values) {
		panic(gerror.Newf("点位数量与值数量不一致！点位数量：%d, 值数量：%d", len(group.Metas), dataLength))
	}
	bytes := make([]byte, dataLength*2)

	// 需要验证一下meta的顺序是否正确
	metaIndex := uint16(0)
	for i, meta := range group.Metas {
		if metaIndex == 0 {
			metaIndex = meta.Addr
		} else {

			if meta.Addr != (metaIndex + c_protocol.ReadTypeRegisterSize(meta.ReadType)) {
				panic(gerror.Newf("点位的顺序不正确！点位：%s, 地址：%d，实际地址应该为: %d", meta.Name, meta.Addr, metaIndex+c_protocol.ReadTypeRegisterSize(meta.ReadType)))
			}
			metaIndex = meta.Addr
		}
		valueBytes := c_protocol.ReadTypeEncoder(meta.ReadType, values[i], meta.Factor, meta.Offset, meta.Endianness)
		copy(bytes[i*2:], valueBytes)
		p.log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceConfig.Id, meta.Name, meta.Addr, valueBytes)
	}

	err := p.client.WriteMultipleRegistersBytes(p.modbusDeviceConfig.UnitId, group.Addr, dataLength, bytes)
	if err != nil {
		c_log.BizWarningf(p.ctx, "WriteMultipleRegisters失败！StationType: [%s] Error: [%v]", group.Name, err)
		return err
	}
	return nil
}

func (p *ModbusProtocolProvider) WriteSingleCoil(meta *c_base.Meta, isOn bool) error {
	err := p.client.WriteSingleCoil(p.modbusDeviceConfig.UnitId, meta.Addr, isOn)

	return err
}

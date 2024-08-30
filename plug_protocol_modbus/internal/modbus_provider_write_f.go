package internal

import (
	"ems-plan/c_base"
	"fmt"
	"plug_protocol_modbus/p_modbus"
)

func (p *ModbusProvider) WriteSingleRegister(meta *c_base.Meta, value int32) error {
	result := meta.ReadType.Encoder(int64(value), meta.Factor, meta.Offset, meta.Endianness)
	// 通关result来计算需要多少个寄存器位置
	registerLength := len(result) / 2
	if registerLength == 1 {
		uint16Value := meta.Endianness.DecodeToUint16(result)
		p.log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, uint16Value)
		err := p.client.WriteSingleRegister(p.unitId, meta.Addr, uint16Value)
		if err != nil {
			return err
		}
	} else {
		p.log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, result)
		err := p.client.WriteMultipleRegistersBytes(p.unitId, meta.Addr, uint16(registerLength), result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ModbusProvider) WriteMultipleRegisters(group *p_modbus.ModbusGroup, values []int64) error {
	dataLength := group.Quantity
	if len(group.Metas) != len(values) {
		panic(fmt.Sprintf("点位数量与值数量不一致！点位数量：%d, 值数量：%d", len(group.Metas), dataLength))
	}
	bytes := make([]byte, dataLength*2)

	// 需要验证一下meta的顺序是否正确
	metaIndex := uint16(0)
	for i, meta := range group.Metas {
		if metaIndex == 0 {
			metaIndex = meta.Addr
		} else {
			if meta.Addr != (metaIndex + meta.ReadType.RegisterSize()) {
				panic(fmt.Sprintf("点位的顺序不正确！点位：%s, 地址：%d，实际地址应该为: %d", meta.Name, meta.Addr, metaIndex+meta.ReadType.RegisterSize()))
			}
			metaIndex = meta.Addr
		}
		valueBytes := meta.ReadType.Encoder(values[i], meta.Factor, meta.Offset, meta.Endianness)
		copy(bytes[i*2:], valueBytes)
		p.log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, valueBytes)
	}

	err := p.client.WriteMultipleRegistersBytes(p.unitId, group.Addr, dataLength, bytes)
	if err != nil {
		p.log.Warningf(p.ctx, "WriteMultipleRegisters失败！Group: [%s] Error: [%v]", group.Name, err)
		return err
	}
	return nil
}

func (p *ModbusProvider) WriteSingleCoil(meta *c_base.Meta, isOn bool) error {
	err := p.client.WriteSingleCoil(p.unitId, meta.Addr, isOn)

	return err
}

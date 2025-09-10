package internal

import (
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"p_base"

	"github.com/pkg/errors"
)

func (p *ModbusProtocolProvider) WriteSingleRegister(meta *c_proto.SModbusPoint, value int32) error {
	if p.GetProtocolStatus() != c_enum.EProtocolConnected {
		return errors.Errorf("device %s connect is not activated", p.deviceId)
	}

	// 使用新的 EncoderBytes 函数编码数据
	result, err := p_base.EncoderBytes(value, meta.DataAccess.DataFormat, meta.DataAccess.ByteEndian, meta.DataAccess.WordOrder, meta.DataAccess.Offset, meta.DataAccess.Factor)
	if err != nil {
		return errors.Wrapf(err, "编码点位 %s 的值失败", meta.Name)
	}

	// 处理长度冲突 - 使用新的结构
	expectedLength := int(p_base.GetEffectiveByteLength(meta.DataAccess.ByteLength, meta.DataAccess.DataFormat))
	if expectedLength > 0 && len(result) != expectedLength {
		c_log.BizWarningf(p.ctx, "点位 %s 编码长度不匹配：期望 %d 字节，实际 %d 字节，DataFormat: %v",
			meta.Name, expectedLength, len(result), meta.DataAccess.DataFormat)

		// 根据期望长度调整结果
		result = p.adjustDataLength(result, expectedLength)
	}

	// 计算需要多少个寄存器位置（每个寄存器2字节）
	registerLength := len(result) / 2
	if registerLength == 1 {
		// 单寄存器写入
		uint16Value := uint16(result[0])<<8 | uint16(result[1])
		c_log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, uint16Value)
		err := p.client.WriteSingleRegister(p.modbusDeviceConfig.UnitId, meta.Addr, uint16Value)
		if err != nil {
			return err
		}
	} else {
		// 多寄存器写入
		c_log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, result)
		err := p.client.WriteMultipleRegistersBytes(p.modbusDeviceConfig.UnitId, meta.Addr, uint16(registerLength), result)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *ModbusProtocolProvider) WriteMultipleRegisters(group *c_proto.SModbusPointTask, values []int64) error {
	if p.GetProtocolStatus() != c_enum.EProtocolConnected {
		return errors.Errorf("device %s connect is not activated", p.deviceId)
	}

	dataLength := group.Quantity
	if len(group.Points) != len(values) {
		return errors.Errorf("点位数量与值数量不一致！点位数量：%d, 值数量：%d", len(group.Points), len(values))
	}

	bytes := make([]byte, dataLength*2)

	// 需要验证一下meta的顺序是否正确
	metaIndex := uint16(0)
	for i, meta := range group.Points {
		if metaIndex == 0 {
			metaIndex = meta.Addr
		} else {
			// 计算当前点位应该占用的寄存器数量
			expectedAddr := metaIndex + p_base.GetQuantityFromDataAccess(meta.DataAccess)
			if meta.Addr != expectedAddr {
				return errors.Errorf("点位的顺序不正确！点位：%s, 地址：%d，实际地址应该为: %d", meta.Name, meta.Addr, expectedAddr)
			}
			metaIndex = meta.Addr
		}

		// 使用新的 EncoderBytes 函数编码数据
		valueBytes, err := p_base.EncoderBytes(values[i], meta.DataAccess.DataFormat, meta.DataAccess.ByteEndian, meta.DataAccess.WordOrder, meta.DataAccess.Offset, meta.DataAccess.Factor)
		if err != nil {
			return errors.Wrapf(err, "编码点位 %s 的值失败", meta.Name)
		}

		// 将编码后的字节复制到总字节数组中
		copy(bytes[i*2:], valueBytes)
		c_log.Debugf(p.ctx, "%s 写入点位：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, valueBytes)
	}

	err := p.client.WriteMultipleRegistersBytes(p.modbusDeviceConfig.UnitId, group.Addr, dataLength, bytes)
	if err != nil {
		c_log.BizWarningf(p.ctx, "WriteMultipleRegisters失败！StationType: [%s] Error: [%v]", group.Name, err)
		return err
	}
	return nil
}

func (p *ModbusProtocolProvider) WriteSingleCoil(meta *c_proto.SModbusPoint, isOn bool) error {
	if p.GetProtocolStatus() != c_enum.EProtocolConnected {
		return errors.Errorf("device %s connect is not activated", p.deviceId)
	}

	c_log.Debugf(p.ctx, "%s 写入线圈：%s 地址：0x%x 值：%v", p.deviceId, meta.Name, meta.Addr, isOn)
	err := p.client.WriteSingleCoil(p.modbusDeviceConfig.UnitId, meta.Addr, isOn)
	if err != nil {
		return errors.Wrapf(err, "写入线圈 %s 失败", meta.Name)
	}

	return nil
}

// adjustDataLength 调整数据长度以匹配期望长度
func (p *ModbusProtocolProvider) adjustDataLength(data []byte, expectedLength int) []byte {
	if len(data) == expectedLength {
		return data
	}

	if len(data) < expectedLength {
		// 数据不足，用0填充
		result := make([]byte, expectedLength)
		copy(result, data)
		c_log.Debugf(p.ctx, "数据长度不足，用0填充：原始长度 %d，期望长度 %d", len(data), expectedLength)
		return result
	} else {
		// 数据过长，截断
		result := make([]byte, expectedLength)
		copy(result, data[:expectedLength])
		c_log.Debugf(p.ctx, "数据长度过长，截断：原始长度 %d，期望长度 %d", len(data), expectedLength)
		return result
	}
}

package internal

import (
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"fmt"
	"p_base"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

func (p *ModbusProtocolProvider) ReadSingleSync(point *c_proto.SModbusPoint, function c_enum.EModbusReadFunction, lifetime time.Duration, readCache bool) (any, error) {

	var (
		vr  any
		err error
	)

	if readCache {
		vr, err = p.GetValue(point)
	}

	if err != nil {
		return nil, err
	}
	if vr != nil {
		return vr, nil
	}

	if p.GetProtocolStatus() != c_enum.EProtocolConnected {
		return nil, errors.New("当前连接未连接，无法查询数据")
	}

	name := fmt.Sprintf("SingleRead:%s", point.Name)

	result, err := p.readValues(name, point.Addr, point.GetQuantity(), function)
	if err != nil {
		return nil, err
	}
	if result == nil || len(result) == 0 {
		return nil, errors.Errorf("读取到的数据为空！")
	}
	values, err := p.analysisModbus(name, point.Addr, point.GetQuantity(), lifetime, result, point)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, errors.Errorf("获取的值为空！")
	}
	return values[0], nil
}

// ReadGroupSync 同步读取
func (p *ModbusProtocolProvider) ReadGroupSync(group *c_proto.SModbusPointTask, readCache bool, metas ...*c_proto.SModbusPoint) ([]any, error) {
	returnMetasLength := len(metas)
	if readCache && metas != nil && returnMetasLength != 0 {
		vars := make([]any, returnMetasLength)
		var needRead bool
		for i, meta := range metas {
			value, err := p.GetValue(meta)
			if err != nil || value == nil {
				// 如果有错误或者无数据，就直接退出循环，执行后面的数据读取指令
				needRead = true
				break
			}
			vars[i] = value
		}
		if !needRead {
			// 如果不需要读，直接返回
			return vars, nil
		}
	}

	result, err := p.readValues(group.Name, group.Addr, group.Quantity, group.Function)
	if err != nil {
		return nil, err
	}
	allGroupVars, err := p.analysisModbus(group.Name, group.Addr, group.Quantity, group.Lifetime, result, group.Points...)
	if err != nil {
		return nil, err
	}
	if metas == nil || returnMetasLength == 0 {
		// 如果没有指定metas，就返回空值
		return nil, nil
	}

	// 从allGroupVars中取出metas对应的值
	vars := make([]any, returnMetasLength)

	for i, meta := range metas {
		for j := 0; j < len(allGroupVars); j++ {
			if group.Points[j] == meta {
				// 一样的点位
				vars[i] = allGroupVars[j]
			}
		}
	}

	if len(vars) != returnMetasLength {
		return nil, errors.Errorf("metas数量和结果数量不一样！")
	}

	return vars, nil
}

func (p *ModbusProtocolProvider) read(name string, addr uint16, quantity uint16, function c_enum.EModbusReadFunction) ([]byte, error) {
	var (
		result []byte
		err    error
	)

	// 累计分钟请求次数
	p.metricProtocol.AddMinuteReadCount()

	queryTime := time.Now()
	switch function {
	case c_enum.EMqReadCoils:
		result, err = p.client.ReadCoils(p.modbusDeviceConfig.UnitId, addr, quantity)
	case c_enum.EMqDiscreteInputs:
		result, err = p.client.ReadDiscreteInputs(p.modbusDeviceConfig.UnitId, addr, quantity)
	case c_enum.EMqHoldingRegisters:
		result, err = p.client.ReadHoldingRegistersBytes(p.modbusDeviceConfig.UnitId, addr, quantity)
	case c_enum.EMqInputRegisters:
		result, err = p.client.ReadInputRegistersBytes(p.modbusDeviceConfig.UnitId, addr, quantity)
	}
	// 累计请求时间
	p.metricProtocol.CalcReadTime(name, time.Since(queryTime))
	// 累计请求返回的数据量
	p.metricProtocol.AddMinuteResultSize(len(result))

	// 累计失败次数
	if err != nil {
		p.metricProtocol.AddMinuteFailedCount()
	}

	return result, err
}

func (p *ModbusProtocolProvider) readValues(name string, addr, quantity uint16, function c_enum.EModbusReadFunction) ([]byte, error) {
	p.modbusRwMutex.Lock()
	defer p.modbusRwMutex.Unlock()
	result, err := p.read(name, addr, quantity, function)
	if err != nil {
		if err.Error() == "EOF" {
			_ = p.client.Close()
		} else {
			_ = p.client.Close()
			c_log.BizWarningf(p.ctx, "[%s] Modbus Failed！unitId:%d Add: 0x%X Len: %d  Fun: %s ;Cause：%+v", name, p.modbusDeviceConfig.UnitId, addr, quantity, function.String(), err)
		}
		return nil, err
	}
	if result == nil || len(result) == 0 {
		_ = p.client.Close()
		return nil, errors.Errorf("[%s] Modbus Task Return Empty！", name)
	}

	c_log.Debugf(p.ctx, "[%v] Modbus Task Return：[% x]", name, result)
	// 更新最后更新时间
	now := time.Now()
	p.lastUpdateTime = &now
	return result, nil
}

func (p *ModbusProtocolProvider) analysisModbus(groupName string, addr, quantity uint16, lifetime time.Duration, result []byte, points ...*c_proto.SModbusPoint) ([]any, error) {
	if points == nil || len(points) == 0 || result == nil {
		return nil, errors.Errorf("[%s] Analysis的查询方法 value或points参数为空！", groupName)
	}

	var (
		results    = make([]any, len(points))
		errMessage string
		err        error
	)
	for i := 0; i < len(points); i++ {
		point := points[i]
		if point == nil {
			continue
		}

		value, err := decoder(p.deviceId, p.deviceType, result, addr, quantity, point)

		if p.cache != nil && err == nil {
			err = p.IProtocolCacheValue.CacheValue(point, value, lifetime)
		}

		if err != nil {
			message := fmt.Sprintf("[%s-%s] %v;", groupName, point.Name, err)
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}

		p.UpdateAlarm(p.deviceId, point, value)
		results[i] = value
	}
	if errMessage != "" {
		err = errors.Errorf(errMessage)
	}

	return results, err
}

func decoder(deviceId string, deviceType c_enum.EDeviceType, bytes []byte, addr, quantity uint16, point *c_proto.SModbusPoint) (*c_base.SPointValue, error) {
	var value any
	var err error

	//if task.CustomDecoder != nil {
	//	value, err = task.CustomDecoder(bytes, task, point)
	//	if err != nil {
	//		return nil, errors.Wrap(err, "custom decoder failed")
	//	}
	//} else {
	// 使用通用的DecoderBytes函数进行解析
	// 根据BitLength决定如何传递参数
	var byteIndex, byteLength, bitIndex, bitLength uint16

	if point.BitLength != 0 {
		// 位级别：使用BitAddr和BitLength
		// 边界检查：确保点位地址在任务范围内
		// 每个寄存器占16位，所以任务的位地址范围是 [s.Addr*16, (s.Addr+s.Quantity)*16-1]
		taskStartBit := addr * 16
		taskEndBit := (addr+quantity)*16 - 1
		// 检查起始地址和结束地址都在范围内
		bitEndAddress := point.BitAddr + point.BitLength - 1
		if point.BitAddr < taskStartBit || bitEndAddress > taskEndBit {
			return nil, errors.Errorf("bit address range [%d:%d] is out of task range [%d:%d]",
				point.BitAddr, bitEndAddress, taskStartBit, taskEndBit)
		}

		// 纯位模式：byteLength=0，使用bitIndex和bitLength
		byteIndex = 0
		byteLength = 0
		bitIndex = point.BitAddr - taskStartBit // 相对位索引
		bitLength = point.BitLength
	} else {
		// 字节级别：Address和Length都是寄存器级别的
		// 边界检查：确保点位地址在任务范围内
		if point.Addr < addr || point.Addr+point.Length > addr+quantity {
			return nil, errors.Errorf("register address range [%d:%d] is out of task range [%d:%d]",
				point.Addr, point.Addr+point.Length-1, addr, addr+quantity-1)
		}

		// 纯字节模式：bitLength=0，使用byteIndex和byteLength
		registerOffset := point.Addr - addr // 寄存器偏移量
		byteIndex = registerOffset * 2      // 转换为字节偏移量
		byteLength = point.Length * 2       // 寄存器数量转换为字节数量
		bitIndex = 0
		bitLength = 0
	}

	value, err = p_base.DecoderBytes(
		bytes,
		byteIndex,
		byteLength,
		bitIndex,
		bitLength,
		point.ByteEndian,
		point.WordOrder,
		point.DataFormat,
		point.ValueType,
		point.Offset,
		point.Factor,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode point %s", point.GetKey())
	}

	err = p_base.ValidateValueRange(value, point.Min, point.Max)
	if err != nil {
		return nil, errors.Wrapf(err, "value %v out of range", value)
	}

	return c_base.NewPointValue(deviceId, deviceType, point, value), nil
}

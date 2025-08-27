package internal

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/pkg/errors"
	"p_base"
	"time"
)

func (p *ModbusProtocolProvider) ReadSingleSync(meta *c_base.Meta, function c_proto.EModbusReadFunction, lifetime time.Duration, readCache bool) (any, error) {
	var (
		vr  any
		err error
	)
	if readCache {
		vr, err = p.GetValue(meta)
	}

	if err != nil {
		return nil, err
	}
	if vr != nil {
		return vr, nil
	}

	if p.GetStatus() != c_base.EProtocolConnected {
		return nil, errors.New("当前连接未连接，无法查询数据")
	}

	name := fmt.Sprintf("SingleRead:%s", meta.Name)
	result, err := p.readValues(name, meta.Addr, p_base.ReadTypeRegisterSize(meta.ReadType), function)
	if err != nil {
		return nil, err
	}
	if result == nil || len(result) == 0 {
		return nil, gerror.Newf("读取到的数据为空！")
	}
	values, err := p.analysisModbus(name, meta.Addr, lifetime, result, meta)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, gerror.Newf("获取的值为空！")
	}
	return values[0], nil
}

// ReadGroupSync 同步读取
func (p *ModbusProtocolProvider) ReadGroupSync(group *c_proto.SModbusTask, readCache bool, metas ...*c_base.Meta) ([]any, error) {
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
	allGroupVars, err := p.analysisModbus(group.Name, group.Addr, group.Lifetime, result, group.Metas...)
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
			if group.Metas[j] == meta {
				// 一样的点位
				vars[i] = allGroupVars[j]
			}
		}
	}

	if len(vars) != returnMetasLength {
		return nil, gerror.Newf("metas数量和结果数量不一样！")
	}

	return vars, nil
}

func (p *ModbusProtocolProvider) read(name string, addr uint16, quantity uint16, function c_proto.EModbusReadFunction) ([]byte, error) {
	var (
		result []byte
		err    error
	)

	// 累计分钟请求次数
	p.metricProtocol.AddMinuteReadCount()

	queryTime := time.Now()
	switch function {
	case c_proto.EMqReadCoils:
		result, err = p.client.ReadCoils(p.modbusDeviceConfig.UnitId, addr, quantity)
	case c_proto.EMqDiscreteInputs:
		result, err = p.client.ReadDiscreteInputs(p.modbusDeviceConfig.UnitId, addr, quantity)
	case c_proto.EMqHoldingRegisters:
		result, err = p.client.ReadHoldingRegistersBytes(p.modbusDeviceConfig.UnitId, addr, quantity)
	case c_proto.EMqInputRegisters:
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

func (p *ModbusProtocolProvider) readValues(name string, addr, quantity uint16, function c_proto.EModbusReadFunction) ([]byte, error) {
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
		return nil, gerror.Newf("[%s] Modbus Task Return Empty！", name)
	}

	c_log.Debugf(p.ctx, "[%v] Modbus Task Return：[% x]", name, result)
	// 更新最后更新时间
	now := time.Now()
	p.lastUpdateTime = &now
	return result, nil
}

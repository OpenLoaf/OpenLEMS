package protocol

import (
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"plug_protocol_modbus/p_modbus"
	"time"
)

func (p *ModbusProvider) ReadSingleSync(meta *c_base.Meta, function p_modbus.ModbusReadFunction, lifetime time.Duration, readCache bool) (*gvar.Var, error) {
	var (
		vr  *gvar.Var
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
	name := fmt.Sprintf("SingleRead:%s", meta.Name)
	result, err := p.readValues(name, meta.Addr, meta.ReadType.RegisterSize(), function)
	if err != nil {
		return nil, err
	}
	if result == nil || len(result) == 0 {
		return nil, fmt.Errorf("读取到的数据为空！")
	}
	values, err := common_analysis.AnalysisModbus(p.ctx, p.cache, p.alarmProvider, name, meta.Addr, lifetime, result, meta)
	if err != nil {
		return nil, err
	}
	if len(values) == 0 {
		return nil, fmt.Errorf("获取的值为空！")
	}
	return values[0], nil
}

// ReadGroupSync 同步读取
func (p *ModbusProvider) ReadGroupSync(group *p_modbus.ModbusGroup, readCache bool, metas ...*c_base.Meta) ([]*gvar.Var, error) {
	returnMetasLength := len(metas)
	if readCache && metas != nil && returnMetasLength != 0 {
		vars := make([]*gvar.Var, returnMetasLength)
		for i, meta := range metas {
			value, err := p.GetValue(meta)
			if err != nil || value == nil {
				// 如果有错误或者无数据，就直接退出循环，执行后面的数据读取指令
				break
			}
			vars[i] = value
		}
		return vars, nil
	}

	result, err := p.readValues(group.Name, group.Addr, group.Quantity, group.Function)
	if err != nil {
		return nil, err
	}
	allGroupVars, err := common_analysis.AnalysisModbus(p.ctx, p.cache, p.alarmProvider, group.Name, group.Addr, group.Lifetime, result, group.Metas...)
	if err != nil {
		return nil, err
	}
	if metas == nil || returnMetasLength == 0 {
		// 如果没有指定metas，就返回空值
		return nil, nil
	}

	// 从allGroupVars中取出metas对应的值
	vars := make([]*gvar.Var, returnMetasLength)

	for i, meta := range metas {
		for j := 0; j < len(allGroupVars); j++ {
			if group.Metas[j] == meta {
				// 一样的点位
				vars[i] = allGroupVars[j]
			}
		}
	}

	if len(vars) != returnMetasLength {
		return nil, fmt.Errorf("metas数量和结果数量不一样！")
	}

	return vars, nil
}

func (p *ModbusProvider) read(addr uint16, quantity uint16, function p_modbus.ModbusReadFunction) ([]byte, error) {
	var (
		result []byte
		err    error
	)

	switch function {
	case p_modbus.MqReadCoils:
		result, err = p.client.ReadCoils(p.unitId, addr, quantity)
	case p_modbus.MqDiscreteInputs:
		result, err = p.client.ReadDiscreteInputs(p.unitId, addr, quantity)
	case p_modbus.MqHoldingRegisters:
		result, err = p.client.ReadHoldingRegistersBytes(p.unitId, addr, quantity)
	case p_modbus.MqInputRegisters:
		result, err = p.client.ReadInputRegistersBytes(p.unitId, addr, quantity)
	}
	return result, err
}

func (p *ModbusProvider) readValues(name string, addr, quantity uint16, function p_modbus.ModbusReadFunction) ([]byte, error) {
	p.modbusRwMutex.Lock()
	defer p.modbusRwMutex.Unlock()
	result, err := p.read(addr, quantity, function)
	if err != nil {
		if err.Error() == "EOF" {
			_ = p.client.Close()
		} else {
			_ = p.client.Close()
			p.log.Warningf(p.ctx, "[%v-%v] Modbus任务获取数据失败！失败原因：%+v", p.DeviceId, name, err)
			if p.notifyChannel != nil {
				p.notifyChannel <- &protocol.Notify{
					Type:    protocol.ReadFailed,
					Message: fmt.Sprintf("[%v-%v] Modbus任务获取数据失败！失败原因：%+v", p.DeviceId, name, err),
				}
			}
		}
		return nil, err
	}
	if result == nil || len(result) == 0 {
		_ = p.client.Close()
		return nil, fmt.Errorf("[%v-%v] Modbus任务获取数据为空！", p.DeviceId, name)
	}

	p.log.Debugf(p.ctx, "[%v-%v] Modbus任务获取到数据：[% x]", p.DeviceId, name, result)
	// 更新最后更新时间
	now := time.Now()
	p.lastUpdateTime = &now
	return result, nil
}

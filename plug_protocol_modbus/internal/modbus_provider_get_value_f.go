package internal

import (
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/util/gconv"
)

func (p *ModbusProtocolProvider) GetValue(meta *c_base.Meta) (*gvar.Var, error) {
	get, err := p.cache.Get(p.ctx, meta)
	if err != nil {
		return &gvar.Var{}, err
	}
	metaValue := &c_base.MetaValue{}
	err = get.Structs(metaValue)
	if err != nil {
		return nil, err
	}
	if metaValue == nil {
		return nil, fmt.Errorf("[%v-%s] 获取的值为空！", p.deviceConfig.Id, meta.Name)
	}
	return metaValue.Value, err
}
func (p *ModbusProtocolProvider) GetIntValue(meta *c_base.Meta) (int, error) {
	get, err := p.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, fmt.Errorf("[%v-%s] 获取的值为空！", p.deviceConfig.Id, meta.Name)
	}
	return get.Int(), err
}

func (p *ModbusProtocolProvider) GetUintValue(meta *c_base.Meta) (uint, error) {
	get, err := p.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, fmt.Errorf("[%v-%s] 获取的值为空！", p.deviceConfig.Id, meta.Name)
	}
	return get.Uint(), err
}

func (p *ModbusProtocolProvider) GetFloat32Value(meta *c_base.Meta) (float32, error) {
	// TODO 判断一下调用的参数类型是否是float32，使用断言

	get, err := p.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, fmt.Errorf("[%v-%s] 获取的值为空！", p.deviceConfig.Id, meta.Name)
	}
	return get.Float32(), err
}

func (p *ModbusProtocolProvider) GetFloat32Values(metas ...*c_base.Meta) ([]float32, error) {
	list := make([]float32, len(metas))
	for i, poi := range metas {
		get, err := p.GetFloat32Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

func (p *ModbusProtocolProvider) GetFloat64Value(meta *c_base.Meta) (float64, error) {
	get, err := p.GetValue(meta)
	if err != nil {
		return 0, err
	}
	return gconv.Float64(get), nil
}

func (p *ModbusProtocolProvider) GetFloat64Values(metas ...*c_base.Meta) ([]float64, error) {
	list := make([]float64, len(metas))
	for i, poi := range metas {
		get, err := p.GetFloat64Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

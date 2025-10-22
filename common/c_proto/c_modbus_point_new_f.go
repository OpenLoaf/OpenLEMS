package c_proto

import (
	"common/c_base"
	"common/c_enum"
)

// NewModbusPoint 创建基础 Modbus 点位
func NewModbusPoint(addr uint16, key, name string, valueType c_enum.EValueType, dataAccess *c_base.SDataAccess) *SModbusPoint {
	point := &SModbusPoint{
		Addr: addr,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
			},
			DataAccess: dataAccess,
		},
	}

	// 自动从 DataAccess 推断小数位数
	if dataAccess != nil {
		point.SProtocolPoint.SPoint.Precise = uint8(dataAccess.GetDecimalPlaces())
	}

	return point
}

// NewModbusPointWithUnit 创建带单位的 Modbus 点位
func NewModbusPointWithUnit(addr uint16, key, name string, valueType c_enum.EValueType, unit string, dataAccess *c_base.SDataAccess) *SModbusPoint {
	point := &SModbusPoint{
		Addr: addr,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
				Unit:      unit,
			},
			DataAccess: dataAccess,
		},
	}

	// 自动从 DataAccess 推断小数位数
	if dataAccess != nil {
		point.SProtocolPoint.SPoint.Precise = uint8(dataAccess.GetDecimalPlaces())
	}

	return point
}

// NewModbusPointWithDesc 创建带描述的 Modbus 点位
func NewModbusPointWithDesc(addr uint16, key, name string, valueType c_enum.EValueType, unit, desc string, dataAccess *c_base.SDataAccess) *SModbusPoint {
	point := &SModbusPoint{
		Addr: addr,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
				Unit:      unit,
				Desc:      desc,
			},
			DataAccess: dataAccess,
		},
	}

	// 自动从 DataAccess 推断小数位数
	if dataAccess != nil {
		point.SProtocolPoint.SPoint.Precise = uint8(dataAccess.GetDecimalPlaces())
	}

	return point
}

// NewModbusPointFromPreset 使用预定义 SPoint 创建点位
func NewModbusPointFromPreset(addr uint16, preset *c_base.SPoint, dataAccess *c_base.SDataAccess) *SModbusPoint {
	point := &SModbusPoint{
		Addr: addr,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     preset,
			DataAccess: dataAccess,
		},
	}

	// 如果预定义点位没有设置小数位数，从 DataAccess 推断
	if dataAccess != nil && preset != nil && preset.Precise == 0 {
		point.SProtocolPoint.SPoint.Precise = uint8(dataAccess.GetDecimalPlaces())
	}

	return point
}

// ModbusPointOption 点位选项函数类型
type ModbusPointOption func(*SModbusPoint)

// NewModbusPointExt 使用选项模式创建 Modbus 点位（复杂场景）
func NewModbusPointExt(addr uint16, opts ...ModbusPointOption) *SModbusPoint {
	point := &SModbusPoint{
		Addr: addr,
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{},
		},
	}

	// 应用所有选项
	for _, opt := range opts {
		opt(point)
	}

	// 如果小数位数为0且 DataAccess 存在，从 DataAccess 推断
	if point.SProtocolPoint.SPoint.Precise == 0 && point.SProtocolPoint.DataAccess != nil {
		point.SProtocolPoint.SPoint.Precise = uint8(point.SProtocolPoint.DataAccess.GetDecimalPlaces())
	}

	return point
}

// 选项函数定义
func WithKey(key string) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Key = key
	}
}

func WithName(name string) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Name = name
	}
}

func WithValueType(valueType c_enum.EValueType) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.ValueType = valueType
	}
}

func WithUnit(unit string) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Unit = unit
	}
}

func WithDesc(desc string) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Desc = desc
	}
}

func WithDataAccess(dataAccess *c_base.SDataAccess) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.DataAccess = dataAccess
	}
}

func WithValueExplain(explains []*c_base.SFieldExplain) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.ValueExplain = explains
	}
}

func WithTrigger(trigger func(value interface{}) (bool, c_enum.EAlarmLevel, error)) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.Trigger = trigger
	}
}

func WithGroup(group *c_base.SPointGroup) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.Group = group
	}
}

func WithSort(sort int) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.Sort = sort
	}
}

func WithMinMax(min, max int64) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Min = min
		p.SProtocolPoint.SPoint.Max = max
	}
}

func WithPrecise(precise uint8) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Precise = precise
	}
}

func WithHidden(hidden bool) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint.Hidden = hidden
	}
}

// 复合选项
func WithPresetPoint(preset *c_base.SPoint) ModbusPointOption {
	return func(p *SModbusPoint) {
		p.SProtocolPoint.SPoint = preset
	}
}

package c_proto

import (
	"common/c_base"
	"common/c_enum"
)

// NewCanbusPoint 创建基础 CANbus 点位
func NewCanbusPoint(key, name string, valueType c_enum.EValueType, dataAccess *c_base.SDataAccess) *SCanbusPoint {
	return &SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{
				Key:       key,
				Name:      name,
				ValueType: valueType,
			},
			DataAccess: dataAccess,
		},
	}
}

// NewCanbusPointWithUnit 创建带单位的 CANbus 点位
func NewCanbusPointWithUnit(key, name string, valueType c_enum.EValueType, unit string, dataAccess *c_base.SDataAccess) *SCanbusPoint {
	return &SCanbusPoint{
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
}

// NewCanbusPointWithDesc 创建带描述的 CANbus 点位
func NewCanbusPointWithDesc(key, name string, valueType c_enum.EValueType, unit, desc string, dataAccess *c_base.SDataAccess) *SCanbusPoint {
	return &SCanbusPoint{
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
}

// NewCanbusPointFromPreset 使用预定义 SPoint 创建点位
func NewCanbusPointFromPreset(preset *c_base.SPoint, dataAccess *c_base.SDataAccess) *SCanbusPoint {
	return &SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint:     preset,
			DataAccess: dataAccess,
		},
	}
}

// CanbusPointOption CANbus 点位选项函数类型
type CanbusPointOption func(*SCanbusPoint)

// NewCanbusPointExt 使用选项模式创建 CANbus 点位（复杂场景）
func NewCanbusPointExt(opts ...CanbusPointOption) *SCanbusPoint {
	point := &SCanbusPoint{
		SProtocolPoint: &c_base.SProtocolPoint{
			SPoint: &c_base.SPoint{},
		},
	}
	for _, opt := range opts {
		opt(point)
	}
	return point
}

// CANbus 选项函数定义
func WithCanbusKey(key string) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Key = key
	}
}

func WithCanbusName(name string) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Name = name
	}
}

func WithCanbusValueType(valueType c_enum.EValueType) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.ValueType = valueType
	}
}

func WithCanbusUnit(unit string) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Unit = unit
	}
}

func WithCanbusDesc(desc string) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Desc = desc
	}
}

func WithCanbusDataAccess(dataAccess *c_base.SDataAccess) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.DataAccess = dataAccess
	}
}

func WithCanbusValueExplain(explains []*c_base.SFieldExplain) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.ValueExplain = explains
	}
}

func WithCanbusTrigger(trigger func(value interface{}) (bool, c_enum.EAlarmLevel, error)) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.Trigger = trigger
	}
}

func WithCanbusGroup(group *c_base.SPointGroup) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.Group = group
	}
}

func WithCanbusSort(sort int) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.Sort = sort
	}
}

func WithCanbusMinMax(min, max int64) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Min = min
		p.SProtocolPoint.SPoint.Max = max
	}
}

func WithCanbusPrecise(precise uint8) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Precise = precise
	}
}

func WithCanbusHidden(hidden bool) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint.Hidden = hidden
	}
}

// 复合选项
func WithCanbusPresetPoint(preset *c_base.SPoint) CanbusPointOption {
	return func(p *SCanbusPoint) {
		p.SProtocolPoint.SPoint = preset
	}
}

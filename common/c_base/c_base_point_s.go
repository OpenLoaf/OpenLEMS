package c_base

import (
	"common/c_enum"
	"strconv"

	"github.com/shockerli/cvt"
)

// SPoint 点位元数据
type SPoint struct {
	Key       string                                                                      `json:"key" v:"required" ` // 名称
	Name      string                                                                      `json:"name" v:"required"` // 名称
	ValueType c_enum.EValueType                                                           `json:"value_type,omitempty" v:"required"`
	Group     *SPointGroup                                                                `json:"group" dc:"分组"`
	Unit      string                                                                      `json:"unit,omitempty"`             // 单位
	Desc      string                                                                      `json:"desc,omitempty"`             // 备注
	Sort      int                                                                         `json:"sort"`                       // 排序
	Min       int64                                                                       `json:"min,omitempty" dc:"正常范围最小值"` // 范围最小值,  默认为0
	Max       int64                                                                       `json:"max,omitempty" dc:"正常范围最大值"` // 范围最大值，默认为0
	Precise   uint8                                                                       `json:"precise,omitempty"`          // 设置浮点数精度（只是显示用）
	Trigger   func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) `json:"-" dc:"告警触发函数"`
	Hidden    bool                                                                        `json:"hidden" dc:"是否显示"`
}

func (s *SPoint) IsHidden() bool {
	return s.Hidden
}

func (s *SPoint) IsAlarmPoint() bool {
	return s.Trigger != nil
}

func (s *SPoint) GetValueType() c_enum.EValueType {
	return s.ValueType
}

func (s *SPoint) TriggerAlarm(value any) (trigger bool, level c_enum.EAlarmLevel, err error) {
	if s.Trigger != nil {
		return s.Trigger(value)
	}
	return false, c_enum.EAlarmLevelNone, nil
}

func (s *SPoint) GetValueExplain(value any) (string, error) {
	// 检查值是否为数值类型（整数或浮点数）
	switch value.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
	case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
	default:
		// 非数值类型使用默认的字符串转换
		return cvt.StringE(value)
	}

	// 尝试将值转换为浮点数
	if floatVal, err := cvt.Float64E(value); err == nil {
		// 使用 strconv.FormatFloat 进行精确格式化
		// 'f' 表示固定小数点格式，s.Precise 表示精度
		formatted := strconv.FormatFloat(floatVal, 'f', int(s.Precise), 64)
		return formatted, nil
	}

	// 如果无法转换为浮点数，使用默认的字符串转换
	return cvt.StringE(value)
}

func (s *SPoint) GetKey() string {
	return s.Key
}

func (s *SPoint) GetName() string {
	return s.Name
}

func (s *SPoint) GetGroup() *SPointGroup {
	return s.Group
}

func (s *SPoint) GetUnit() string {
	return s.Unit
}

func (s *SPoint) GetDesc() string {
	return s.Desc
}

func (s *SPoint) GetSort() int {
	return s.Sort
}

func (s *SPoint) GetMin() int64 {
	return s.Min
}

func (s *SPoint) GetMax() int64 {
	return s.Max
}

func (s *SPoint) GetPrecise() uint8 {
	return s.Precise
}

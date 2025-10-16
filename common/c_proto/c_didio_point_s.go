package c_proto

import (
	"common/c_base"
	"common/c_enum"
	"fmt"
)

type SDidioPoint struct {
	// 继承协议点位基础结构
	*c_base.SProtocolPoint

	// DIDO特定配置
	Pin       uint8 `json:"pin" v:"required" dc:"DIDO引脚编号"`
	ChipIndex uint8 `json:"chipIndex" v:"required" dc:"DIDO芯片索引"`

	// 覆盖字段（如果需要与基础结构不同的行为）
	Sort  int                 `json:"sort"`          // 覆盖SPoint的Sort
	Group *c_base.SPointGroup `json:"group" dc:"分组"` // 覆盖SPoint的Group

	// 覆盖字段（如果需要与基础结构不同的行为）
	ValueExplain []*c_base.SFieldExplain `json:"valueExplain,omitempty" yaml:"valueExplain"` // 值解释

	// 功能函数
	Trigger func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) `json:"-" dc:"告警触发函数"` // 可覆盖SPoint的触发函数
}

func (s *SDidioPoint) GetDataAccess() *c_base.SDataAccess {
	if s.SProtocolPoint != nil {
		return s.SProtocolPoint.DataAccess
	}
	return nil
}

func (s *SDidioPoint) String() string {
	return fmt.Sprintf("%s[chip%d-pin%d]", s.GetName(), s.ChipIndex, s.Pin)
}

func (s *SDidioPoint) IsAlarmPoint() bool {
	if s.Trigger != nil {
		return true
	}
	if s.SProtocolPoint != nil && s.SProtocolPoint.SPoint != nil {
		return s.SProtocolPoint.SPoint.IsAlarmPoint()
	}
	return false
}

func (s *SDidioPoint) TriggerAlarm(value any) (trigger bool, level c_enum.EAlarmLevel, err error) {
	if s.Trigger != nil {
		return s.Trigger(value)
	}
	if s.SProtocolPoint != nil && s.SProtocolPoint.SPoint != nil && s.SProtocolPoint.SPoint.Trigger != nil {
		return s.SProtocolPoint.SPoint.Trigger(value)
	}
	return false, level, nil
}

func (s *SDidioPoint) GetGroup() *c_base.SPointGroup {
	if s.Group != nil {
		return s.Group
	}
	if s.SProtocolPoint != nil && s.SProtocolPoint.SPoint != nil {
		return s.SProtocolPoint.SPoint.Group
	}
	return nil
}

func (s *SDidioPoint) GetSort() int {
	if s.Sort != 0 {
		return s.Sort
	}
	if s.SProtocolPoint != nil && s.SProtocolPoint.SPoint != nil {
		return s.SProtocolPoint.SPoint.Sort
	}
	return 0
}

// GetValueExplain 获取值解释，实现层层查找逻辑
func (s *SDidioPoint) GetValueExplain() []*c_base.SFieldExplain {
	// 1. 优先使用自己的 ValueExplain
	if len(s.ValueExplain) > 0 {
		return s.ValueExplain
	}
	// 2. 然后使用 SProtocolPoint 的 ValueExplain
	if s.SProtocolPoint != nil {
		return s.SProtocolPoint.GetValueExplain()
	}
	// 3. 最后返回空切片
	return []*c_base.SFieldExplain{}
}

// IsProtocolPoint 判断是否为协议点位，SDidioPoint 是协议点位，返回 true
func (s *SDidioPoint) IsProtocolPoint() bool {
	return true
}

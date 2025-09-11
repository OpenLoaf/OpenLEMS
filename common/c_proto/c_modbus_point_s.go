package c_proto

import (
	"common/c_base"
	"common/c_enum"
	"fmt"
)

type SModbusPoint struct {
	// 地址配置
	Addr uint16 `json:"addr" v:"required" dc:"Modbus起始地址"`

	// 点位信息
	*c_base.SPoint

	Sort  int                 `json:"sort"`          // 覆盖SPoint的Sort
	Group *c_base.SPointGroup `json:"group" dc:"分组"` // 覆盖SPoint的Group

	// 数据访问配置
	DataAccess *c_base.SDataAccess `json:"dataAccess" v:"required" dc:"数据访问配置"`

	// 功能函数
	StatusExplain func(value any) (text string, err error)                                    `json:"-" dc:"状态解释函数"`
	Trigger       func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) `json:"-" dc:"告警触发函数"` // 可覆盖SPoint的触发函数
}

func (s *SModbusPoint) GetDataAccess() *c_base.SDataAccess {
	return s.DataAccess
}

func (s *SModbusPoint) String() string {
	return fmt.Sprintf("%s[0x%x]", s.GetName(), s.Addr)
}

func (s *SModbusPoint) ValueExplain(value any) (string, error) {
	if s.StatusExplain == nil {
		return s.SPoint.ValueExplain(value)
	}
	return s.StatusExplain(value)
}
func (s *SModbusPoint) IsNotAlarm() bool {
	return s.Trigger == nil || s.SPoint.IsNotAlarm()
}

func (s *SModbusPoint) AlarmTrigger(value any) (trigger bool, level c_enum.EAlarmLevel, err error) {
	if s.Trigger != nil {
		return s.Trigger(value)
	}
	if s.SPoint.Trigger != nil {
		return s.SPoint.Trigger(value)
	}
	return false, level, nil
}
func (s *SModbusPoint) GetGroup() *c_base.SPointGroup {
	if s.Group != nil {
		return s.Group
	}
	return s.SPoint.Group
}
func (s *SModbusPoint) GetSort() int {
	if s.Sort != 0 {
		return s.Sort
	}
	return s.SPoint.Sort
}

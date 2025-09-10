package c_proto

import (
	"common/c_base"
	"fmt"
)

type SModbusPoint struct {
	// 地址配置
	Addr uint16 `json:"addr" v:"required" dc:"起始地址"`

	// 点位信息
	*c_base.SPoint

	// 数据访问配置
	DataAccess *c_base.SDataAccess `json:"dataAccess" v:"required" dc:"数据访问配置"`

	// 功能函数
	StatusExplain func(value any) (string, error)       `json:"-" dc:"状态解释函数"`
	Trigger       func(value interface{}) (bool, error) `json:"-" dc:"告警触发函数"`
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

func (s *SModbusPoint) AlarmTrigger(value any) (bool, error) {
	if s.Trigger != nil {
		return s.Trigger(value)
	}
	if s.SPoint.Trigger != nil {
		return s.SPoint.Trigger(value)
	}
	return false, nil
}

//func (s *SModbusPoint) GetPrecise() uint8 {
//
//}

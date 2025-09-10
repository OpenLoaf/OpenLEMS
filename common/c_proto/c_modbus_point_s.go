package c_proto

import (
	"common/c_base"

	"github.com/shockerli/cvt"
)

type SModbusPoint struct {
	*c_base.SPoint

	// 地址配置
	Addr uint16 `json:"addr" v:"required" dc:"起始地址"`

	// 数据访问配置
	DataAccess *c_base.SDataAccess `json:"dataAccess" v:"required" dc:"数据访问配置"`

	// 功能函数
	StatusExplain func(value any) (string, error)       `json:"-" dc:"状态解释函数"`
	Trigger       func(value interface{}) (bool, error) `json:"-" dc:"告警触发函数"`
}

func (s *SModbusPoint) ValueExplain(value any) (string, error) {
	if s.StatusExplain == nil {
		return cvt.String(value), nil
	}
	return s.StatusExplain(value)
}

func (s *SModbusPoint) AlarmTrigger(value any) (bool, error) {
	if s.Trigger == nil {
		return false, nil
	}

	return s.Trigger(value)
}

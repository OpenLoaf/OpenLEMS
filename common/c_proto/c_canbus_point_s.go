package c_proto

import (
	"common/c_base"
	"common/c_enum"
	"fmt"
)

type SCanbusPoint struct {

	// 点位信息
	*c_base.SPoint

	// 数据访问配置
	DataAccess *c_base.SDataAccess `json:"dataAccess" v:"required" dc:"数据访问配置"`

	// 功能函数
	StatusExplain func(value any) (text string, err error)                                    `json:"-" dc:"状态解释函数"`
	Trigger       func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) `json:"-" dc:"告警触发函数"` // 可覆盖SPoint的触发函数
}

func (s *SCanbusPoint) GetDataAccess() *c_base.SDataAccess {
	return s.DataAccess
}

func (s *SCanbusPoint) String() string {
	if s.DataAccess == nil {
		return s.GetName()
	}
	return fmt.Sprintf("%s-%s", s.GetName(), s.DataAccess)
}

func (s *SCanbusPoint) ValueExplain(value any) (string, error) {
	if s.StatusExplain == nil {
		return s.SPoint.ValueExplain(value)
	}
	return s.StatusExplain(value)
}
func (s *SCanbusPoint) IsNotAlarm() bool {
	return s.Trigger == nil || s.SPoint.IsNotAlarm()
}

func (s *SCanbusPoint) AlarmTrigger(value any) (trigger bool, level c_enum.EAlarmLevel, err error) {
	if s.Trigger != nil {
		return s.Trigger(value)
	}
	if s.SPoint.Trigger != nil {
		return s.SPoint.Trigger(value)
	}
	return false, level, nil
}

package c_base

import (
	"common/c_enum"
)

type IPoint interface {
	GetKey() string
	GetName() string
	GetGroup() *SPointGroup
	GetUnit() string
	GetDesc() string
	GetSort() int
	GetMin() int64
	GetMax() int64
	GetPrecise() uint8
	IsNotAlarm() bool                                                           //  是否有告警触发函数
	AlarmTrigger(value any) (trigger bool, level c_enum.EAlarmLevel, err error) // 判断触发或者消除告警
	ValueExplain(value any) (string, error)                                     // 获取Value解释，一般为状态类型的解释
	GetDataAccess() *SDataAccess                                                //  获取数据访问配置
}

//type IPointFull interface {
//	IPoint
//}

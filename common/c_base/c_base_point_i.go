package c_base

import (
	"common/c_enum"
)

type IPoint interface {
	GetKey() string
	GetName() string
	GetGroup() *SPointGroup
	GetLevel() c_enum.EAlarmLevel
	GetUnit() string
	GetDesc() string
	GetSort() int
	GetMin() int64
	GetMax() int64
	GetPrecise() uint8
	AlarmTrigger(value any) (bool, error)   // 判断触发或者消除告警
	ValueExplain(value any) (string, error) // 获取Value解释，一般为状态类型的解释
	GetDataAccess() *SDataAccess            //  获取数据访问配置
}

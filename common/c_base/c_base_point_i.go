package c_base

import (
	"common/c_enum"
)

type IPoint interface {
	GetKey() string                            // 点位Key 一个设备点位key不能重复
	GetName() string                           // 名称
	GetGroup() *SPointGroup                    // 分组
	GetUnit() string                           // 单位
	GetDesc() string                           // 备注
	GetSort() int                              // 排序
	GetMin() int64                             // 点位理论最小值
	GetMax() int64                             // 点位理论最大值
	GetPrecise() uint8                         // 小数点
	GetValueType() c_enum.EValueType           // 值类型
	GetValueExplain(value any) (string, error) // 获取Value解释，一般为状态类型的解释

	IsAlarmPoint() bool                                                         //  是否是告警点位
	TriggerAlarm(value any) (trigger bool, level c_enum.EAlarmLevel, err error) // 判断触发或者消除告警
}

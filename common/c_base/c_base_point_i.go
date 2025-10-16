package c_base

import (
	"common/c_enum"
)

// IPoint 简化的点位接口，只包含核心方法
type IPoint interface {
	GetKey() string                                                             // 点位Key
	GetName() string                                                            // 名称
	GetValueType() c_enum.EValueType                                            // 值类型
	GetValueExplainByValue(value any) (string, error)                           // 获取Value解释
	IsAlarmPoint() bool                                                         // 是否是告警点位
	TriggerAlarm(value any) (trigger bool, level c_enum.EAlarmLevel, err error) // 告警触发
	GetGroup() *SPointGroup                                                     // 获取分组信息
	GetSort() int                                                               // 获取排序
	GetMin() int64                                                              // 获取最小值
	GetMax() int64                                                              // 获取最大值
	GetPrecise() uint8                                                          // 获取精度
	GetUnit() string                                                            // 获取单位
	GetDesc() string                                                            // 获取描述
	IsHidden() bool                                                             // 是否隐藏
	AsProtocolPoint() *SProtocolPoint                                           // 转换为协议点位（如果是协议点位则返回，否则返回nil）
}

// 注意：子结构体不需要重复实现IPoint接口方法
// 通过结构体嵌套，自动继承SPoint的方法实现
// 只有在需要覆盖方法时才重新实现

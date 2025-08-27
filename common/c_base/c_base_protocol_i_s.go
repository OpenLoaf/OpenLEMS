package c_base

import "time"

type EProtocolStatus int // 连接状态
const (
	EProtocolDisconnected EProtocolStatus = iota // 连接断开
	EProtocolConnecting                          // 正在连接中
	EProtocolConnected                           // 连接成功
)

type IProtocol interface {
	GetStatus() EProtocolStatus            // 获取协议连接状态
	GetLastUpdateTime() *time.Time         // 获取最后更新时间
	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetValue(meta *Meta) (any, error)
	GetAlarmList() []*SAlarmDetail // 获取当前告警列表

	RegisterTask(task ITask, tasks ...ITask)                                             // 注册任务
	RegisterAlarmTriggerFunc(handler func(maxAlarm EAlarmLevel, nowAlarm *SAlarmDetail)) // 注册告警激活函数
	RegisterAlarmRemoveFunc(handler func(maxAlarm EAlarmLevel, nowAlarm *SAlarmDetail))  // 注册告警
	ProtocolListen()                                                                     // 启动协议监听
}

type IGetProtocolCacheValue interface {
	GetBool(meta *Meta) (bool, error)
	GetIntValue(meta *Meta) (int, error)
	GetInt32Value(meta *Meta) (int32, error)
	GetUintValue(meta *Meta) (uint, error)
	GetUint32Value(meta *Meta) (uint32, error)
	GetFloat32Value(meta *Meta) (float32, error)
	GetFloat32Values(metas ...*Meta) ([]float32, error)
	GetFloat64Value(meta *Meta) (float64, error)
	GetFloat64Values(meta ...*Meta) ([]float64, error)
}

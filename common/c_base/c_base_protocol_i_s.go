package c_base

import (
	"common/c_enum"
	"time"
)

type IProtocol interface {
	IAlarm
	GetProtocolStatus() c_enum.EProtocolStatus // 获取协议连接状态
	GetLastUpdateTime() *time.Time             // 获取最后更新时间
	GetPointValueList() []*SPointValue         // 获取所有缓存的数据列表

	GetValue(point IPoint) (any, error)

	RegisterTask(task IPointTask, tasks ...IPointTask) // 注册任务
	ProtocolListen()                                   // 启动协议监听

}

type IProtocolCacheValue interface {
	GetValue(point IPoint) (any, error)
	GetBool(meta IPoint) (bool, error)
	GetIntValue(meta IPoint) (int, error)
	GetInt32Value(meta IPoint) (int32, error)
	GetUintValue(meta IPoint) (uint, error)
	GetUint32Value(meta IPoint) (uint32, error)
	GetFloat32Value(meta IPoint) (float32, error)
	GetFloat32Values(metas ...IPoint) ([]float32, error)
	GetFloat64Value(meta IPoint) (float64, error)
	GetFloat64Values(meta ...IPoint) ([]float64, error)
	CacheValue(point IPoint, value any, lifetime time.Duration) error
}

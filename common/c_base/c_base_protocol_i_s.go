package c_base

import (
	"common/c_enum"
	"time"
)

type IProtocol interface {
	IAlarm
	GetProtocolStatus() c_enum.EProtocolStatus       // 获取协议连接状态
	GetLastUpdateTime() *time.Time                   // 获取最后更新时间
	GetProtocolPointValue(point IPoint) *SPointValue // 获取协议点位缓存值

	GetValue(point IPoint) (any, error)

	//RegisterTask(task IPointTask, tasks ...IPointTask) // 注册任务
	ProtocolListen()
	GetConfig() *SDeviceConfig // 启动协议监听

}

type IProtocolCacheValue interface {
	Clear()
	CacheValue(value *SPointValue, lifetime time.Duration) error

	GetValue(point IPoint) (any, error)
	GetBool(point IPoint) (*bool, error)
	GetIntValue(point IPoint) (*int, error)
	GetInt32Value(point IPoint) (*int32, error)
	GetUintValue(point IPoint) (*uint, error)
	GetUint32Value(point IPoint) (*uint32, error)
	GetFloat32Value(point IPoint) (*float32, error)
	GetFloat32Values(points ...IPoint) ([]*float32, error)
	GetFloat64Value(point IPoint) (*float64, error)
	GetFloat64Values(points ...IPoint) ([]*float64, error)
	GetProtocolPointValue(point IPoint) *SPointValue // 获取协议点位缓存值
}

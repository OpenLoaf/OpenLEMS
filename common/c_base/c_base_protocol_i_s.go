package c_base

import "time"

type IProtocol interface {
	IAlarm
	GetStatus() EProtocolStatus            // 获取协议连接状态
	GetLastUpdateTime() *time.Time         // 获取最后更新时间
	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetValue(meta *Meta) (any, error)

	RegisterTask(task ITask, tasks ...ITask) // 注册任务
	ProtocolListen()                         // 启动协议监听
}

type IGetProtocolCacheValue interface {
	GetValue(meta *Meta) (any, error)
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

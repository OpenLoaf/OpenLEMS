package c_proto

import (
	"common/c_base"
	"common/c_enum"
	"time"
)

type IModbusProtocol interface {
	c_base.IProtocol
	c_base.IProtocolCacheValue

	RegisterTask(task *SModbusPointTask, tasks ...*SModbusPointTask) // 注册任务

	ReadSingleSync(meta *SModbusPoint, function c_enum.EModbusReadFunction, lifetime time.Duration, readCache bool) (any, error) // lifetime 为0时候永不过期，为负数时候不缓存并删除缓存的值
	ReadGroupSync(group *SModbusPointTask, readCache bool, points ...*SModbusPoint) ([]any, error)                               // 同步读取,第二个参数为幻读,先从缓存中取，如果有值直接返回，无值再去执行查询方法

	WriteSingleRegister(meta *SModbusPoint, value int32) error
	WriteMultipleRegisters(group *SModbusPointTask, values []int64) error
}

type ICanbusProtocol interface {
	c_base.IProtocol
	c_base.IProtocolCacheValue

	RegisterTask(task *SCanbusTask, tasks ...*SCanbusTask) // 注册任务
	SendMessage(task *SCanbusTask, values []float64) error
}

type IGpiodProtocol interface {
	c_base.IProtocol

	InitGpioPoint(point c_base.IPoint) // 初始化点位

	RegisterHandler(handler func(status bool)) // 状态变化处理
	GetGpioStatus() *bool                      // 是否是高电平
	SetHigh() error                            // 设置为高电平
	SetLow() error                             // 设置为低电平
}

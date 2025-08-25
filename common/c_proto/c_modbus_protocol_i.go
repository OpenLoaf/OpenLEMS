package c_proto

import (
	"common/c_base"
	"time"
)

type IModbusProtocol interface {
	c_base.IProtocol
	c_base.IGetProtocolCacheValue

	// todo 删除lifetime
	ReadSingleSync(meta *c_base.Meta, function EModbusReadFunction, lifetime time.Duration, readCache bool) (any, error) // lifetime 为0时候永不过期，为负数时候不缓存并删除缓存的值
	ReadGroupSync(group *SModbusTask, readCache bool, metas ...*c_base.Meta) ([]any, error)                              // 同步读取,第二个参数为幻读,先从缓存中取，如果有值直接返回，无值再去执行查询方法

	WriteSingleRegister(meta *c_base.Meta, value int32) error
	WriteMultipleRegisters(group *SModbusTask, values []int64) error
}

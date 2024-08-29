package p_modbus

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

type IModbusProtocol interface {
	c_base.IProtocol

	RegisterRead(ctx context.Context, group *ModbusGroup, gs ...*ModbusGroup) // 注册group后，系统将自动解析group的查询后，读取数据并缓存

	ReadSingleSync(meta *c_base.Meta, function ModbusReadFunction, lifetime time.Duration, readCache bool) (*gvar.Var, error) // lifetime 为0时候永不过期，为负数时候不缓存并删除缓存的值
	ReadGroupSync(group *ModbusGroup, readCache bool, metas ...*c_base.Meta) ([]*gvar.Var, error)                             // 同步读取,第二个参数为幻读,先从缓存中取，如果有值直接返回，无值再去执行查询方法

	GetCache() *gcache.Cache
	GetValue(meta *c_base.Meta) (*gvar.Var, error)
	GetIntValue(meta *c_base.Meta) (int, error)
	GetUintValue(meta *c_base.Meta) (uint, error)
	GetFloat32Value(meta *c_base.Meta) (float32, error)
	GetFloat32Values(metas ...*c_base.Meta) ([]float32, error)
	GetFloat64Value(meta *c_base.Meta) (float64, error)
	GetFloat64Values(meta ...*c_base.Meta) ([]float64, error)

	WriteSingleRegister(meta *c_base.Meta, value int32) error
	WriteMultipleRegisters(group *ModbusGroup, values []int64) error
}

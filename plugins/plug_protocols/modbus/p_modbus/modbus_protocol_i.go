package p_modbus

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"time"
)

type IModbusProtocol interface {
	c_base.IProtocol

	RegisterRead(ctx context.Context, group *SModbusTask, gs ...*SModbusTask) // 注册group后，系统将自动解析group的查询后，读取数据并缓存

	ReadSingleSync(meta *c_base.Meta, function ModbusReadFunction, lifetime time.Duration, readCache bool) (*gvar.Var, error) // lifetime 为0时候永不过期，为负数时候不缓存并删除缓存的值
	ReadGroupSync(group *SModbusTask, readCache bool, metas ...*c_base.Meta) ([]*gvar.Var, error)                             // 同步读取,第二个参数为幻读,先从缓存中取，如果有值直接返回，无值再去执行查询方法

	GetValue(meta *c_base.Meta) (*gvar.Var, error)
	GetBool(meta *c_base.Meta) (bool, error)
	GetIntValue(meta *c_base.Meta) (int, error)
	GetInt32Value(meta *c_base.Meta) (int32, error)
	GetUintValue(meta *c_base.Meta) (uint, error)
	GetUint32Value(meta *c_base.Meta) (uint32, error)
	GetFloat32Value(meta *c_base.Meta) (float32, error)
	GetFloat32Values(metas ...*c_base.Meta) ([]float32, error)
	GetFloat64Value(meta *c_base.Meta) (float64, error)
	GetFloat64Values(meta ...*c_base.Meta) ([]float64, error)
	GetModbusDeviceConfig() *SModbusDeviceConfig

	WriteSingleRegister(meta *c_base.Meta, value int32) error
	WriteMultipleRegisters(group *SModbusTask, values []int64) error
}

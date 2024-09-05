package p_modbus

import (
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

// TODO 改名字加S
type ModbusGroup struct {
	Name           string
	Desc           string
	Addr           uint16
	Quantity       uint16
	Function       ModbusReadFunction
	CycleMill      int64
	Lifetime       time.Duration // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Transitory     bool          // 是短暂的，查询一次后，需要再次调用查询才能查询，而且不会一直轮询。默认是永久查询
	TransitoryTime time.Duration
	Metas          []*c_base.Meta
}

func (m *ModbusGroup) Check() {
	var pointNameSet gset.StrSet

	for _, p := range m.Metas {
		if !pointNameSet.AddIfNotExist(p.Name) {
			panic(gerror.Newf("ModbusGroup[%s] has duplicate point name: %s", m.Name, p.Name))
		}
	}
}

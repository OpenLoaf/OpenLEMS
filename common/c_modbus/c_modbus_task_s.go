package c_modbus

import (
	"common/c_base"
	"fmt"
	"time"
)

type SModbusTask struct {
	Name           string
	Desc           string
	Addr           uint16
	Quantity       uint16
	Function       EModbusReadFunction
	CycleMill      int64
	Lifetime       time.Duration // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Transitory     bool          // 是短暂的，查询一次后，需要再次调用查询才能查询，而且不会一直轮询。默认是永久查询
	TransitoryTime time.Duration
	Metas          []*c_base.Meta
}

func (m *SModbusTask) Check() {
	var pointNameMap = make(map[string]struct{})
	for _, p := range m.Metas {
		if _, exist := pointNameMap[p.Name]; exist {
			panic(fmt.Errorf("SModbusTask[%s] has duplicate point name: %s", m.Name, p.Name))
		}
		pointNameMap[p.Name] = struct{}{}
	}
}

package c_proto

import (
	"common/c_base"
	"common/c_log"
	"context"
	"github.com/pkg/errors"
	"time"
)

const DefaultCacheLifeTime = 3 * time.Second

type SModbusTask struct {
	Name           string
	DisplayName    string
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

func (m *SModbusTask) GetName() string {
	return m.Name
}

func (m *SModbusTask) GetDescription() string {
	return m.Desc
}

func (m *SModbusTask) GetMetas() []*c_base.Meta {
	return m.Metas
}

func (m *SModbusTask) GetLifeTime() time.Duration {
	return m.Lifetime
}

func (m *SModbusTask) Check(ctx context.Context) error {
	var pointNameMap = make(map[string]any)
	var err error
	for _, p := range m.Metas {
		if _, exist := pointNameMap[p.Name]; exist {
			c_log.BizErrorf(ctx, "SModbusTask[%s] has duplicate point name: %s", m.Name, p.Name)
			err = errors.New("duplicate point name")
		}
		pointNameMap[p.Name] = p
	}
	return err
}

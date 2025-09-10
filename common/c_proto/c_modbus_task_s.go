package c_proto

import (
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"time"

	"github.com/pkg/errors"
)

const DefaultCacheLifeTime = 3 * time.Second

type SModbusPointTask struct {
	Name           string                     `v:"required"`
	Addr           uint16                     `v:"required"`
	Quantity       uint16                     `v:"required"`
	Function       c_enum.EModbusReadFunction `v:"required"`
	CycleMill      int64                      `v:"required"`
	Points         []*SModbusPoint            `v:"required"`
	Lifetime       time.Duration              // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Transitory     bool                       // 是短暂的，查询一次后，需要再次调用查询才能查询，而且不会一直轮询。默认是永久查询
	TransitoryTime time.Duration
	Desc           string
	CustomDecoder  func(bytes []byte, task *SModbusPointTask, point c_base.IPoint) (any, error) // 手动解析，空代表使用默认的协议解析器
}

func (s *SModbusPointTask) GetName() string {
	return s.Name
}

func (s *SModbusPointTask) GetDescription() string {
	return s.Desc
}

func (s *SModbusPointTask) GetPoints() []c_base.IPoint {
	points := make([]c_base.IPoint, len(s.Points))
	for i, point := range s.Points {
		points[i] = point
	}
	return points
}

func (s *SModbusPointTask) GetLifeTime() time.Duration {
	return s.Lifetime
}

func (s *SModbusPointTask) Check(ctx context.Context) error {
	var pointNameMap = make(map[string]any)
	var err error
	for _, p := range s.Points {
		if _, exist := pointNameMap[p.GetName()]; exist {
			c_log.BizErrorf(ctx, "SModbusPointTask[%s] has duplicate point name: %s", s.Name, p.GetName())
			err = errors.New("duplicate point name")
		}
		pointNameMap[p.GetName()] = p
	}
	return err
}

package c_proto

import (
	"common/c_base"
	"common/c_log"
	"context"
	"time"

	"github.com/pkg/errors"
)

type SCanbusTask struct {
	Name           string
	Desc           string
	GetCanbusID    func(params map[string]any) uint32
	Lifetime       time.Duration                                                           // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Points         []*SCanbusPoint                                                         // 点位列表
	IsRemote       bool                                                                    // 是否是远程帧（写重要）
	IsExtended     bool                                                                    // 是否是扩展帧（写重要）
	SendMaxRetries int                                                                     // 发送最大重试次数，默认为3
	CustomDecoder  func(task *SCanbusTask, bytes []byte, point c_base.IPoint) (any, error) // 手动解析，空代表使用默认的协议解析器
	CustomEncoder  func(task *SCanbusTask, values []any) ([]byte, error)                   //  手动编码，空代表使用默认的协议编码器
}

func (s *SCanbusTask) GetName() string {
	return s.Name
}

func (s *SCanbusTask) GetDescription() string {
	return s.Desc
}

func (s *SCanbusTask) GetPoints() []c_base.IPoint {
	points := make([]c_base.IPoint, len(s.Points))
	for i, point := range s.Points {
		points[i] = point
	}
	return points
}

func (s *SCanbusTask) GetLifeTime() time.Duration {
	return s.Lifetime
}

func (s *SCanbusTask) Check(ctx context.Context) error {
	var pointNameMap = make(map[string]any)
	var err error
	for _, p := range s.Points {
		if _, exist := pointNameMap[p.GetName()]; exist {
			c_log.BizErrorf(ctx, "SCanbusTask[%s] has duplicate point name: %s", s.Name, p.GetName())
			err = errors.New("duplicate point name")
		}
		pointNameMap[p.GetName()] = p
	}
	return err
}

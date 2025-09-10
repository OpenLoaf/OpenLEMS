package c_proto

import (
	"common/c_base"
	"time"
)

type SCanbusTask struct {
	Name          string
	Desc          string
	GetCanbusID   func(params map[string]any) *uint32
	IDMatch       func(canId uint32) bool                              // 判断ID是否匹配，如果为空，直接判断是否和CanbusID相等
	Lifetime      time.Duration                                        // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Points        []c_base.IPoint                                      // 点位列表
	IsRemote      bool                                                 // 是否是远程帧（写重要）
	IsExtended    bool                                                 // 是否是扩展帧（写重要）
	CustomDecoder func(bytes []byte, point c_base.IPoint) (any, error) // 手动解析，空代表使用默认的协议解析器
}

func (s *SCanbusTask) GetName() string {
	return s.Name
}

func (s *SCanbusTask) GetDescription() string {
	return s.Desc
}

func (s *SCanbusTask) GetPoints() []c_base.IPoint {
	return s.Points
}

func (s *SCanbusTask) GetLifeTime() time.Duration {
	return s.Lifetime
}

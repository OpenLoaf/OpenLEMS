package p_canbus

import (
	"common/c_base"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
	"time"
)

type SCanbusTask struct {
	Name        string
	Desc        string
	GetCanbusID func(params map[string]any) *uint32
	IDMatch     func(canId uint32) bool // 判断ID是否匹配，如果为空，直接判断是否和CanbusID相等
	Lifetime    time.Duration           // lifetime 为0时候缓存永不过期，为负数时候不缓存并删除缓存的值
	Metas       []*c_base.Meta          // 点位列表

	IsRemote   bool // 是否是远程帧（写重要）
	IsExtended bool // 是否是扩展帧（写重要）
}

func (t *SCanbusTask) Check() {
	var pointNameSet gset.StrSet

	for _, p := range t.Metas {
		if !pointNameSet.AddIfNotExist(p.Name) {
			panic(gerror.Newf("SCanbusTask[%s] has duplicate point name: %s", t.Name, p.Name))
		}
	}
}

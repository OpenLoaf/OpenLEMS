package p_canbus

import (
	"common/c_base"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/errors/gerror"
)

type SCanbusTask struct {
	Name  string
	Desc  string
	CanID uint32
	Metas []*c_base.Meta
}

func (t *SCanbusTask) Check() {
	var pointNameSet gset.StrSet

	for _, p := range t.Metas {
		if !pointNameSet.AddIfNotExist(p.Name) {
			panic(gerror.Newf("SCanbusTask[%s] has duplicate point name: %s", t.Name, p.Name))
		}
	}
}

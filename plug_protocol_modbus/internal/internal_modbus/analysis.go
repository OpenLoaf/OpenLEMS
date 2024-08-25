package protocol

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/os/gcache"
	"math"
	"plug_protocol_analysis/p_analysis"
	"reflect"
	"time"
)

func analysisModbus(ctx context.Context, cache *gcache.Cache, alarmProvider alarm.IProvider, groupName string, addr uint16, lifetime time.Duration, result []byte, metas ...*c_base.Meta) ([]*gvar.Var, error) {
	if metas == nil || len(metas) == 0 || result == nil {
		return nil, fmt.Errorf("[%s] Analysis的查询方法 value或points参数为空！", groupName)
	}

	var (
		results    = make([]*gvar.Var, len(metas))
		errMessage string
		err        error
	)
	for i := 0; i < len(metas); i++ {
		meta := metas[i]
		if meta == nil {
			continue
		}

		if meta.Addr < addr {
			errMessage += fmt.Sprintf("[%s-%s] 点位地址:0x%x超出数据长度:%v;", groupName, meta.Name, meta.Addr, addr)
			continue
		}
		value, err := meta.ReadType.ReadValue(result[(meta.Addr-addr)*2:], meta.BitLength, meta.Endianness)
		if err != nil {
			errMessage += fmt.Sprintf("[%s-%s] %v;", groupName, meta.Name, err)
			continue
		}
		kind := meta.ReadType.GetReflectKind(meta.BitLength)
		if kind == reflect.Float64 && math.IsNaN(value.(float64)) {
			panic(fmt.Sprintf("[%s-%s] 读取到的float64位的值为NaN！请检查字段是否配置正确！\n%+v", groupName, meta.Name, meta))
		}
		vars, err := p_analysis.Process(ctx, value, cache, alarmProvider, lifetime, meta)
		if err != nil {
			errMessage += fmt.Sprintf("[%s-%s] %v;", groupName, meta.Name, err)
			continue
		}
		results[i] = vars
	}
	if errMessage != "" {
		err = fmt.Errorf(errMessage)
	}

	return results, err
}

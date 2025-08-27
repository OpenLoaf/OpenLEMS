package internal

import (
	"common/c_base"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
	"math"
	"p_base"
	"reflect"
	"time"
)

func (p *ModbusProtocolProvider) analysisModbus(groupName string, addr uint16, lifetime time.Duration, result []byte, metas ...*c_base.Meta) ([]any, error) {
	if metas == nil || len(metas) == 0 || result == nil {
		return nil, errors.Errorf("[%s] Analysis的查询方法 value或points参数为空！", groupName)
	}

	var (
		results    = make([]any, len(metas))
		errMessage string
		err        error
	)
	for i := 0; i < len(metas); i++ {
		meta := metas[i]
		if meta == nil {
			continue
		}

		if meta.Addr < addr {
			message := fmt.Sprintf("[%s-%s] 点位地址:0x%x超出数据长度:%v;", groupName, meta.Name, meta.Addr, addr)
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}
		index := (meta.Addr - addr) * 2
		if len(result) < int(index) {
			message := fmt.Sprintf("[%s-%s] 点位地址:0x%x超出数据长度:%v;返回的长度:%v,点位%v", groupName, meta.Name, meta.Addr, addr, len(result), index)
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}

		value, err := p_base.ReadTypeReadValue(meta.ReadType, result[index:], meta.BitLength, meta.Endianness)
		if err != nil {
			message := fmt.Sprintf("[%s-%s] %v;", groupName, meta.Name, err.Error())
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}
		//kind := meta.ReadType.GetReflectKind(meta.BitLength)
		kind := p_base.ReadTypeGetReflectKind(meta.ReadType, meta.BitLength)
		if kind == reflect.Float64 && math.IsNaN(value.(float64)) {
			panic(errors.Errorf("[%s-%s] 读取到的float64位的值为NaN！请检查字段是否配置正确！\n%+v", groupName, meta.Name, meta))
		}
		vars := p_base.MetaTransformModbus(meta, value)
		vars, err = p_base.CacheValue(p.ctx, p.deviceId, p.deviceType, p, meta, vars, p.cache, lifetime)

		p.UpdateAlarm(p.deviceId, p.deviceType, meta, vars)

		if err != nil {
			message := fmt.Sprintf("[%s-%s] %v;", groupName, meta.Name, err)
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}
		results[i] = vars
	}
	if errMessage != "" {
		err = errors.Errorf(errMessage)
	}

	return results, err
}

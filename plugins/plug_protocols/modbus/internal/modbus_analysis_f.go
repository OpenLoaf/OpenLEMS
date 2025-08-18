package internal

import (
	"c_protocol"
	"common/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"math"
	"reflect"
	"time"
)

func (p *ModbusProtocolProvider) analysisModbus(groupName string, addr uint16, lifetime time.Duration, result []byte, metas ...*c_base.Meta) ([]*gvar.Var, error) {
	if metas == nil || len(metas) == 0 || result == nil {
		return nil, gerror.Newf("[%s] Analysis的查询方法 value或points参数为空！", groupName)
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

		value, err := c_protocol.ReadTypeReadValue(meta.ReadType, result[index:], meta.BitLength, meta.Endianness)
		if err != nil {
			message := fmt.Sprintf("[%s-%s] %v;", groupName, meta.Name, err)
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}
		//kind := meta.ReadType.GetReflectKind(meta.BitLength)
		kind := c_protocol.ReadTypeGetReflectKind(meta.ReadType, meta.BitLength)
		if kind == reflect.Float64 && math.IsNaN(value.(float64)) {
			panic(gerror.Newf("[%s-%s] 读取到的float64位的值为NaN！请检查字段是否配置正确！\n%+v", groupName, meta.Name, meta))
		}
		vars, err := c_protocol.MetaTransformAndCache(p.ctx, p.deviceConfig.Id, p.deviceType, p, meta, value, p.cache, lifetime)
		if err != nil {
			message := fmt.Sprintf("[%s-%s] %v;", groupName, meta.Name, err)
			g.Log().Errorf(p.ctx, message)
			errMessage += message
			continue
		}
		results[i] = vars
	}
	if errMessage != "" {
		err = gerror.Newf(errMessage)
		//g.Log().Errorf(ctx, errMessage)
	}

	return results, err
}

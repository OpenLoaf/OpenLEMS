package internal_meta

import (
	"common/c_base"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

func CacheValue(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (*gvar.Var, error) {
	valueInt64 := gconv.Int64(value)
	// 范围验证
	if meta.SystemType != c_base.SBool && meta.Min != meta.Max && (valueInt64 < meta.Min || valueInt64 > meta.Max) {
		//log.Errorf(ctx, "[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Id, value, meta.Min, meta.Max)
		// TODO 此处触发Error级别的告警
		return nil, fmt.Errorf("[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Name, value, meta.Min, meta.Max)
	}

	// 判断是否是非信息类型，用于触发告警
	if meta.Level != 0 && protocol != nil {
		if meta.Trigger != nil {
			if meta.Trigger(value) {
				//alarmProvider.TriggerAlarm(meta, value)
				g.Log().Debugf(ctx, "[%s-%s] 触发[%s]", deviceId, meta.Name, meta.Level.String())

				processAlarm(protocol, deviceId, deviceType, meta, true, value)
			} else {
				processAlarm(protocol, deviceId, deviceType, meta, false, value)
			}
		} else if gconv.Bool(value) == false {
			processAlarm(protocol, deviceId, deviceType, meta, false, value)
		} else {
			processAlarm(protocol, deviceId, deviceType, meta, true, value)
		}
	}
	now := time.Now()

	// 缓存
	if cache != nil {
		err := cache.Set(ctx, meta, &c_base.MetaValue{
			Value:      gvar.New(value),
			HappenTime: &now,
		}, lifetime)
		if err != nil {
			return nil, err
		}
	}

	if meta.Debug {
		g.Log().Debugf(ctx, "[%s-%s] 值: %v", deviceId, meta.Cn, value)
	}

	return gvar.New(value), nil
}

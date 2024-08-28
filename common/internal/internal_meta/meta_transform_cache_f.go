package internal_meta

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"time"
)

func MetaProcess(ctx context.Context, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (*gvar.Var, error) {
	var deviceId = protocol.GetId()
	if meta == nil {
		return nil, fmt.Errorf("[%s] Analysis的查询方法获取到point为nil", deviceId)
	}
	/*	if lifetime == 0 {
		lifetime = common.DefaultCacheLifeTime
	}*/

	//originValue := value
	value = meta.SystemType.Transform(value, meta.ReadType, meta.BitLength, meta.Factor, meta.Offset)
	/*	if g.Log().GetLevel()&glog.LEVEL_DEBU > 0 {
		if meta.SystemType == point.SBool {
			g.Log().Debugf(ctx, "[%s-%s] 值: %v", deviceId, meta.Name, value)
		} else {
			g.Log().Debugf(ctx, "[%s-%s] 原始值:%5v, 乘以：%v, 再偏移：%v, 后的值:%v", meta.Cn, meta.Name, originValue, meta.Factor, meta.Offset, value)
		}
	}*/

	valueInt64 := gconv.Int64(value)

	// 范围验证
	if meta.SystemType != c_base.SBool && meta.Min != meta.Max && (valueInt64 < meta.Min || valueInt64 > meta.Max) {
		//log.Errorf(ctx, "[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Name, value, meta.Min, meta.Max)
		// TODO 此处触发Error级别的告警
		return nil, fmt.Errorf("[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Name, value, meta.Min, meta.Max)
	}

	// 判断是否是非信息类型，用于触发告警
	if meta.Level != 0 {
		if meta.Trigger != nil {
			if meta.Trigger(value) {
				//alarmProvider.TriggerAlarm(meta, value)
				g.Log().Debugf(ctx, "[%s-%s] 触发[%s]", deviceId, meta.Name, meta.Level.Name())

				processAlarm(ctx, protocol, deviceId, meta, true, value)
			} else {
				processAlarm(ctx, protocol, deviceId, meta, false, value)
			}
		} else if gconv.Bool(value) == false {
			processAlarm(ctx, protocol, deviceId, meta, false, value)
		} else {
			processAlarm(ctx, protocol, deviceId, meta, true, value)
		}
	}

	// 缓存
	err := cache.Set(ctx, meta, value, lifetime)
	if err != nil {
		return nil, err
	}

	if meta.Debug {
		g.Log().Debugf(ctx, "[%s-%s] 值: %v", deviceId, meta.Cn, value)
	}

	return gvar.New(value), nil
}

func processAlarm(ctx context.Context, protocol c_base.IProtocol, deviceId string, meta *c_base.Meta, IsTrigger bool, value any) {
	protocol.HandlerAlarmDetail(ctx, &c_base.SAlarmDetail{
		DeviceId:   deviceId,
		DeviceType: protocol.GetType(),
		Level:      meta.Level,
		Meta:       meta,
		IsTrigger:  IsTrigger,
		HappenTime: time.Now(),
		Value:      value,
	})
}

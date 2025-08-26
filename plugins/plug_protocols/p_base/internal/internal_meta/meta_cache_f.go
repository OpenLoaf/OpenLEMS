package internal_meta

import (
	"common/c_base"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"gopkg.in/errgo.v2/fmt/errors"
)

func CacheValue(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (any, error) {
	valueInt64 := gconv.Int64(value)
	// 范围验证
	if meta.SystemType != c_base.SBool && meta.Min != meta.Max && (valueInt64 < meta.Min || valueInt64 > meta.Max) {
		//log.Errorf(ctx, "[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Id, value, meta.Min, meta.Max)
		// TODO 此处触发Error级别的告警
		return nil, errors.Newf("[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Name, value, meta.Min, meta.Max)
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
			Value:      value,
			HappenTime: &now,
		}, lifetime)
		if err != nil {
			fmt.Println("cache set error:", err)
			return nil, err
		}
	}

	//g.Log().Debugf(ctx, "[%s-%s] 值: %v cache is null:%v", deviceId, meta.Cn, value, cache == nil)

	if meta.Debug {
		g.Log().Infof(ctx, "[deviceId:%s-%s] 值: %v", deviceId, meta.Cn, value)
	}

	return value, nil
}

func ProcessAlarm(ctx context.Context, meta *c_base.Meta, value any, cache *gcache.Cache, tiggerAlarm func(meta *c_base.Meta), removeAlarm func(meta *c_base.Meta)) {
	// 判断是否是非信息类型，用于触发告警
	if meta.Level == 0 {
		return
	}
	isTigger := false

	//triggerMessage := fmt.Sprintf("[%s] 触发[%s]告警, 触发值为: %v", meta.Name, meta.Level.String(), value)
	//removeMessage := fmt.Sprintf("[%s] 消除[%s]告警, 触发值为: %v", meta.Name, meta.Level.String(), value)
	if meta.Trigger != nil {
		if meta.Trigger(value) {
			isTigger = true
		}
	} else if gconv.Bool(value) == true {
		isTigger = true
	}

	// 先从缓存中判断一下当前点位值是否相等。如果相等说明是重复触发，如果不相等等说明发生变化
	oldValue, err := cache.Get(ctx, meta)
	if err != nil {
		// 缓存获取失败，直接根据当前状态处理告警
		if isTigger {
			tiggerAlarm(meta)
		} else {
			removeAlarm(meta)
		}
		return
	}

	// 获取缓存中的旧值
	var oldMetaValue *c_base.MetaValue
	if oldValue != nil {
		oldMetaValue = &c_base.MetaValue{}
		err = oldValue.Structs(oldMetaValue)
		if err != nil {
			// 解析缓存值失败，直接根据当前状态处理告警
			if isTigger {
				tiggerAlarm(meta)
			} else {
				removeAlarm(meta)
			}
			return
		}
	}

	// 判断旧值是否也触发了告警
	oldIsTrigger := false
	if oldMetaValue != nil {
		if meta.Trigger != nil {
			oldIsTrigger = meta.Trigger(oldMetaValue.Value)
		} else {
			oldIsTrigger = gconv.Bool(oldMetaValue.Value) == true
		}
	}

	// 判断告警状态是否发生变化
	if isTigger != oldIsTrigger {
		// 告警状态发生变化，执行相应的告警处理
		if isTigger {
			// 触发告警
			tiggerAlarm(meta)
		} else {
			// 消除告警
			removeAlarm(meta)
		}
	}
	// 如果告警状态没有变化，则不执行任何操作（避免重复触发）
}

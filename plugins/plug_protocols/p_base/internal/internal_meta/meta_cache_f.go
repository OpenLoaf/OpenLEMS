package internal_meta

import (
	"common/c_base"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
)

func CacheValue(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (any, error) {
	valueInt64 := gconv.Int64(value)
	// 范围验证
	if meta.SystemType != c_base.SBool && meta.Min != meta.Max && (valueInt64 < meta.Min || valueInt64 > meta.Max) {
		//log.Errorf(ctx, "[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Id, value, meta.Min, meta.Max)
		// TODO 此处触发Error级别的告警
		return nil, errors.Errorf("[%s-%s] 数据不在正常范围内!当前值:%v,理论上最小值：%v,最大值：%v", deviceId, meta.Name, value, meta.Min, meta.Max)
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

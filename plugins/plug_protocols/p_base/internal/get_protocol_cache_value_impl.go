package internal

import (
	"common/c_base"
	"common/c_enum"
	"context"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

type SGetProtocolCacheValueImpl struct {
	deviceId   string
	deviceType c_enum.EDeviceType
	ctx        context.Context
	cache      *gcache.Cache // 点位缓存
}

func NewGetProtocolCacheValue(ctx context.Context, deviceId string, deviceType c_enum.EDeviceType, cache *gcache.Cache) c_base.IProtocolCacheValue {
	return &SGetProtocolCacheValueImpl{
		deviceId:   deviceId,
		deviceType: deviceType,
		ctx:        ctx,
		cache:      cache,
	}
}

func (s *SGetProtocolCacheValueImpl) CacheValue(point c_base.IPoint, value any, lifetime time.Duration) error {
	if point == nil {
		return errors.Errorf("[%v-%s] 缓存点位不能为空！", s.deviceId, point.GetName())
	}
	if value == nil {
		return errors.Errorf("[%v-%s] 缓存值不能为空！", s.deviceId, point.GetName())
	}

	return s.cache.Set(s.ctx, point, c_base.NewPointValue(s.deviceId, s.deviceType, point, value), lifetime)
}

func (s *SGetProtocolCacheValueImpl) GetValue(point c_base.IPoint) (any, error) {
	cacheValue, err := s.cache.Get(s.ctx, point)
	if err != nil {
		return nil, err
	}
	if cacheValue == nil {
		return nil, errors.Errorf("cache value is empty")
	}
	metaValue := &c_base.SPointValue{}
	err = cacheValue.Structs(metaValue)
	if err != nil {
		return nil, err
	}
	if metaValue.GetValue() == nil {
		return nil, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}
	// todo 添加数据过期逻辑，比如超过3秒，数据过期，返回数据过期

	return metaValue.GetValue(), err
}

func (s *SGetProtocolCacheValueImpl) GetBool(point c_base.IPoint) (bool, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return false, err
	}
	if get == nil {
		return false, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}
	return cvt.BoolE(get)
}

func (s *SGetProtocolCacheValueImpl) GetIntValue(point c_base.IPoint) (int, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}
	return cvt.IntE(get)
}

func (s *SGetProtocolCacheValueImpl) GetInt32Value(point c_base.IPoint) (int32, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}
	return cvt.Int32E(get)
}

func (s *SGetProtocolCacheValueImpl) GetUintValue(point c_base.IPoint) (uint, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}
	return cvt.UintE(get)
}

func (s *SGetProtocolCacheValueImpl) GetUint32Value(point c_base.IPoint) (uint32, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}
	return cvt.Uint32E(get)
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Value(point c_base.IPoint) (float32, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.deviceId, point.GetName())
	}

	return cvt.Float32E(get)
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Values(points ...c_base.IPoint) ([]float32, error) {
	list := make([]float32, len(points))
	for i, poi := range points {
		get, err := s.GetFloat32Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat64Value(point c_base.IPoint) (float64, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return 0, err
	}
	return gconv.Float64(get), nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat64Values(points ...c_base.IPoint) ([]float64, error) {
	list := make([]float64, len(points))
	for i, poi := range points {
		get, err := s.GetFloat64Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

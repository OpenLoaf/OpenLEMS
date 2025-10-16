package internal

import (
	"common/c_base"
	"common/c_log"
	"context"
	"sync"
	"time"

	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

// 全局缓存
var pointCache = gcache.New()

type SGetProtocolCacheValueImpl struct {
	ctx      context.Context
	deviceId string
	pointMap map[string]struct{}
	mu       sync.RWMutex // 读写锁，保护 pointMap 的并发访问
}

func NewGetProtocolCacheValue(ctx context.Context, deviceId string) c_base.IProtocolCacheValue {
	return &SGetProtocolCacheValueImpl{
		ctx:      ctx,
		deviceId: deviceId,
		pointMap: make(map[string]struct{}),
	}
}

func (s *SGetProtocolCacheValueImpl) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	for key, _ := range s.pointMap {
		_, _ = pointCache.Remove(s.ctx, key)
	}
	s.pointMap = make(map[string]struct{})
}

func (s *SGetProtocolCacheValueImpl) getCacheKeyWithGroup(point c_base.IPoint) string {
	if point.GetGroup() == nil {
		return s.getCacheKey(point.GetKey(), "")
	}
	return s.getCacheKey(point.GetKey(), point.GetGroup().GroupKey)
}

// getCacheKey 获取缓存key
func (s *SGetProtocolCacheValueImpl) getCacheKey(pointKey, groupKey string) string {
	key := s.deviceId + ":" + pointKey
	if groupKey != "" {
		key += ":" + groupKey
	}
	return key
}

func (s *SGetProtocolCacheValueImpl) CacheValue(value *c_base.SPointValue, lifetime time.Duration) error {

	if value == nil {
		return errors.Errorf("[%v] 缓存值不能为空！", s.deviceId)
	}
	c_log.Debugf(s.ctx, "缓存设备: %s 值 %s[%s]：%v 过期时间: %v", s.deviceId, value.GetName(), value.GetKey(), value.GetValue(), lifetime)
	key := s.getCacheKeyWithGroup(value)
	err := pointCache.Set(s.ctx, key, value, lifetime)
	if err != nil {
		return errors.Wrapf(err, "[%v] 缓存值失败！", s.deviceId)
	}

	s.mu.Lock()
	s.pointMap[key] = struct{}{} // 缓存key
	s.mu.Unlock()

	return err
}

func (s *SGetProtocolCacheValueImpl) GetProtocolPointValue(point c_base.IPoint) *c_base.SPointValue {
	if point == nil {
		return nil
	}

	// 使用GetValue方法获取缓存值
	value, err := s.GetValue(point)
	if err != nil || value == nil {
		return nil
	}

	// 创建SPointValue并返回
	pointValue := c_base.NewPointValue(s.deviceId, point, value)
	return pointValue
}

func (s *SGetProtocolCacheValueImpl) GetValue(point c_base.IPoint) (any, error) {
	cacheValue, err := pointCache.Get(s.ctx, s.getCacheKeyWithGroup(point))
	if err != nil {
		return nil, err
	}
	if cacheValue == nil {
		return nil, nil
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

func (s *SGetProtocolCacheValueImpl) GetBool(point c_base.IPoint) (*bool, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value, err := cvt.BoolE(get)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetIntValue(point c_base.IPoint) (*int, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value, err := cvt.IntE(get)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetInt32Value(point c_base.IPoint) (*int32, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value, err := cvt.Int32E(get)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetUintValue(point c_base.IPoint) (*uint, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value, err := cvt.UintE(get)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetUint32Value(point c_base.IPoint) (*uint32, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value, err := cvt.Uint32E(get)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Value(point c_base.IPoint) (*float32, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value, err := cvt.Float32E(get)
	if err != nil {
		return nil, err
	}
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Values(points ...c_base.IPoint) ([]*float32, error) {
	list := make([]*float32, len(points))
	for i, poi := range points {
		get, err := s.GetFloat32Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat64Value(point c_base.IPoint) (*float64, error) {
	get, err := s.GetValue(point)
	if err != nil {
		return nil, err
	}
	if get == nil {
		return nil, nil // 没有采集到数据
	}
	value := gconv.Float64(get)
	return &value, nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat64Values(points ...c_base.IPoint) ([]*float64, error) {
	list := make([]*float64, len(points))
	for i, poi := range points {
		get, err := s.GetFloat64Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

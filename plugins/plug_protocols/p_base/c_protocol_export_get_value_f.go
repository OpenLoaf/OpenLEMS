package p_base

import (
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

type SGetProtocolCacheValueImpl struct {
	id    string
	ctx   context.Context
	cache *gcache.Cache // 点位缓存
}

func NewGetProtocolCacheValue(ctx context.Context, id string, cache *gcache.Cache) c_base.IGetProtocolCacheValue {
	return &SGetProtocolCacheValueImpl{
		id:    id,
		ctx:   ctx,
		cache: cache,
	}
}

func (s *SGetProtocolCacheValueImpl) GetValue(meta *c_base.Meta) (any, error) {
	cacheValue, err := s.cache.Get(s.ctx, meta)
	if err != nil {
		return nil, err
	}
	if cacheValue == nil {
		return nil, nil
	}
	metaValue := &c_base.MetaValue{}
	err = cacheValue.Structs(metaValue)
	if err != nil {
		return nil, err
	}
	if metaValue.Value == nil {
		return nil, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	// todo 添加数据过期逻辑，比如超过3秒，数据过期，返回数据过期

	return metaValue.Value, err
}

func (s *SGetProtocolCacheValueImpl) GetBool(meta *c_base.Meta) (bool, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return false, err
	}
	if get == nil {
		return false, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return cvt.BoolE(get)
}

func (s *SGetProtocolCacheValueImpl) GetIntValue(meta *c_base.Meta) (int, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return cvt.IntE(get)
}

func (s *SGetProtocolCacheValueImpl) GetInt32Value(meta *c_base.Meta) (int32, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return cvt.Int32E(get)
}

func (s *SGetProtocolCacheValueImpl) GetUintValue(meta *c_base.Meta) (uint, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return cvt.UintE(get)
}

func (s *SGetProtocolCacheValueImpl) GetUint32Value(meta *c_base.Meta) (uint32, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return cvt.Uint32E(get)
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Value(meta *c_base.Meta) (float32, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, errors.Errorf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}

	return cvt.Float32E(get)
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Values(metas ...*c_base.Meta) ([]float32, error) {
	list := make([]float32, len(metas))
	for i, poi := range metas {
		get, err := s.GetFloat32Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat64Value(meta *c_base.Meta) (float64, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	return gconv.Float64(get), nil
}

func (s *SGetProtocolCacheValueImpl) GetFloat64Values(metas ...*c_base.Meta) ([]float64, error) {
	list := make([]float64, len(metas))
	for i, poi := range metas {
		get, err := s.GetFloat64Value(poi)
		if err != nil {
			return nil, err
		}
		list[i] = get
	}
	return list, nil
}

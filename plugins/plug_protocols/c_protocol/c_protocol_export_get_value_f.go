package c_protocol

import (
	"common/c_base"
	"common/c_proto"
	"context"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
)

type SGetProtocolCacheValueImpl struct {
	id    string
	ctx   context.Context
	cache *gcache.Cache // 点位缓存
}

func NewGetProtocolCacheValue(ctx context.Context, id string, cache *gcache.Cache) c_proto.IGetProtocolCacheValue {
	return &SGetProtocolCacheValueImpl{
		id:    id,
		ctx:   ctx,
		cache: cache,
	}
}

func (s *SGetProtocolCacheValueImpl) GetValue(meta *c_base.Meta) (any, error) {
	get, err := s.cache.Get(s.ctx, meta)
	if err != nil {
		return nil, err
	}
	metaValue := &c_base.MetaValue{}
	err = get.Structs(metaValue)
	if err != nil {
		return nil, err
	}
	if metaValue.Value == nil {
		return nil, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}

	return gvar.New(metaValue.Value), err
}

func (s *SGetProtocolCacheValueImpl) GetBool(meta *c_base.Meta) (bool, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return false, err
	}
	if get == nil {
		return false, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return get.(*gvar.Var).Bool(), err
}

func (s *SGetProtocolCacheValueImpl) GetIntValue(meta *c_base.Meta) (int, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return get.(*gvar.Var).Int(), err
}

func (s *SGetProtocolCacheValueImpl) GetInt32Value(meta *c_base.Meta) (int32, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return get.(*gvar.Var).Int32(), err
}

func (s *SGetProtocolCacheValueImpl) GetUintValue(meta *c_base.Meta) (uint, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return get.(*gvar.Var).Uint(), err
}

func (s *SGetProtocolCacheValueImpl) GetUint32Value(meta *c_base.Meta) (uint32, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}
	return get.(*gvar.Var).Uint32(), err
}

func (s *SGetProtocolCacheValueImpl) GetFloat32Value(meta *c_base.Meta) (float32, error) {
	get, err := s.GetValue(meta)
	if err != nil {
		return 0, err
	}
	if get == nil {
		return 0, gerror.Newf("[%v-%s] 获取的值为空！", s.id, meta.Name)
	}

	return get.(*gvar.Var).Float32(), err
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

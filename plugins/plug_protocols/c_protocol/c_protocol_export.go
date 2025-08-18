package c_protocol

import (
	"c_protocol/internal/internal_meta"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/os/gcache"
	"reflect"
	"time"
)

func ReadTypeReadValue(d c_base.EReadType, bytes []byte, bitLength uint8, endianness c_base.ECharSequence) (any, error) {
	return internal_meta.ReadTypeReadValue(d, bytes, bitLength, endianness)
}

func ReadTypeTransform(d c_base.EReadType, value any, bitLength uint8, factor float32, offset int) any {
	return internal_meta.ReadTypeTransform(d, value, bitLength, factor, offset)
}

func ReadTypeRegisterSize(d c_base.EReadType) uint16 {
	return internal_meta.ReadTypeRegisterSize(d)
}

func ReadTypeEncoder(d c_base.EReadType, value int64, factor float32, offset int, endianness c_base.ECharSequence) []byte {
	return internal_meta.ReadTypeEncoder(d, value, factor, offset, endianness)
}

func ReadTypeGetReflectKind(d c_base.EReadType, bitLength uint8) reflect.Kind {
	return internal_meta.ReadTypeGetReflectKind(d, bitLength)
}

// MetaTransformAndCache 元数据转换并缓存
func MetaTransformAndCache(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, value any, cache *gcache.Cache, lifetime time.Duration) (any, error) {
	v := internal_meta.MetaProcess(meta, value)
	return internal_meta.CacheValue(ctx, deviceId, deviceType, protocol, meta, v, cache, lifetime)
}

// MetaTransformCanbus 解析can的数据
func MetaTransformCanbus(ctx context.Context, deviceId string, deviceType c_base.EDeviceType, protocol c_base.IProtocol, meta *c_base.Meta, canData []byte, cache *gcache.Cache, lifetime time.Duration) (any, error) {
	v, err := internal_meta.ParseCanbusData(canData, meta)
	if err != nil {
		return nil, err
	}

	return internal_meta.CacheValue(ctx, deviceId, deviceType, protocol, meta, v, cache, lifetime)
}

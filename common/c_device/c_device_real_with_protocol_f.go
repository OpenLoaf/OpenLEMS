package c_device

import (
	"common/c_base"

	"github.com/pkg/errors"

	"github.com/shockerli/cvt"
)

// 泛型类型转换函数，支持所有基本类型
func convertToType[T any](value any) (*T, error) {
	if value == nil {
		return nil, nil
	}

	// 检查是否是指针类型且为nil
	if ptr, ok := value.(*T); ok {
		return ptr, nil
	}

	// 使用cvt库进行类型转换，但返回指针
	var result T
	var err error

	switch any(result).(type) {
	case bool:
		if v, e := cvt.BoolE(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case int:
		if v, e := cvt.IntE(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case int8:
		if v, e := cvt.Int8E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case int16:
		if v, e := cvt.Int16E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case int32:
		if v, e := cvt.Int32E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case int64:
		if v, e := cvt.Int64E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case uint:
		if v, e := cvt.UintE(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case uint8:
		if v, e := cvt.Uint8E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case uint16:
		if v, e := cvt.Uint16E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case uint32:
		if v, e := cvt.Uint32E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case uint64:
		if v, e := cvt.Uint64E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case float32:
		if v, e := cvt.Float32E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	case float64:
		if v, e := cvt.Float64E(value); e != nil {
			err = e
		} else {
			result = any(v).(T)
		}
	default:
		// 对于其他类型，尝试直接转换
		if v, ok := value.(T); ok {
			result = v
		} else {
			err = errors.Errorf("unsupported type conversion from %T to %T", value, result)
		}
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocol(fc func(protocol P) (any, error)) (any, error) {
	if s.isProtocolNil() {
		return nil, errors.Errorf("protocol is nil")
	}
	//device.protocol.GetValue()
	v, err := fc(s.protocol)
	if err != nil {
		return nil, err
	}
	return v, nil
}

// 泛型方法，支持所有基本类型
func GetFromProtocolTyped[T any, P c_base.IProtocol](s *SRealDeviceImpl[P], fc func(protocol P) (any, error)) (*T, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	return convertToType[T](v)
}

// 为了保持向后兼容性，保留原有的具体类型方法
func (s *SRealDeviceImpl[P]) GetFromProtocolBool(fc func(protocol P) (any, error)) (*bool, error) {
	return GetFromProtocolTyped[bool, P](s, fc)
}

// 整数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolInt(fc func(protocol P) (any, error)) (*int, error) {
	return GetFromProtocolTyped[int, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt8(fc func(protocol P) (any, error)) (*int8, error) {
	return GetFromProtocolTyped[int8, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt16(fc func(protocol P) (any, error)) (*int16, error) {
	return GetFromProtocolTyped[int16, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt32(fc func(protocol P) (any, error)) (*int32, error) {
	return GetFromProtocolTyped[int32, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt64(fc func(protocol P) (any, error)) (*int64, error) {
	return GetFromProtocolTyped[int64, P](s, fc)
}

// 无符号整数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolUint(fc func(protocol P) (any, error)) (*uint, error) {
	return GetFromProtocolTyped[uint, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint8(fc func(protocol P) (any, error)) (*uint8, error) {
	return GetFromProtocolTyped[uint8, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint16(fc func(protocol P) (any, error)) (*uint16, error) {
	return GetFromProtocolTyped[uint16, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint32(fc func(protocol P) (any, error)) (*uint32, error) {
	return GetFromProtocolTyped[uint32, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint64(fc func(protocol P) (any, error)) (*uint64, error) {
	return GetFromProtocolTyped[uint64, P](s, fc)
}

// 浮点数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolFloat32(fc func(protocol P) (any, error)) (*float32, error) {
	return GetFromProtocolTyped[float32, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolFloat64(fc func(protocol P) (any, error)) (*float64, error) {
	return GetFromProtocolTyped[float64, P](s, fc)
}

func (s *SRealDeviceImpl[P]) GetFromPoint(point c_base.IPoint) (any, error) {
	return s.GetFromProtocol(func(protocol P) (any, error) {
		return protocol.GetValue(point)
	})
}

// 泛型方法，支持所有基本类型
func GetFromPointTyped[T any, P c_base.IProtocol](s *SRealDeviceImpl[P], point c_base.IPoint) (*T, error) {
	return GetFromProtocolTyped[T, P](s, func(protocol P) (any, error) {
		return protocol.GetValue(point)
	})
}

// 为了保持向后兼容性，保留原有的具体类型方法
func (s *SRealDeviceImpl[P]) GetFromPointBool(point c_base.IPoint) (*bool, error) {
	return GetFromPointTyped[bool, P](s, point)
}

// 整数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointInt(point c_base.IPoint) (*int, error) {
	return GetFromPointTyped[int, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt8(point c_base.IPoint) (*int8, error) {
	return GetFromPointTyped[int8, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt16(point c_base.IPoint) (*int16, error) {
	return GetFromPointTyped[int16, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt32(point c_base.IPoint) (*int32, error) {
	return GetFromPointTyped[int32, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt64(point c_base.IPoint) (*int64, error) {
	return GetFromPointTyped[int64, P](s, point)
}

// 无符号整数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointUint(point c_base.IPoint) (*uint, error) {
	return GetFromPointTyped[uint, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint8(point c_base.IPoint) (*uint8, error) {
	return GetFromPointTyped[uint8, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint16(point c_base.IPoint) (*uint16, error) {
	return GetFromPointTyped[uint16, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint32(point c_base.IPoint) (*uint32, error) {
	return GetFromPointTyped[uint32, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint64(point c_base.IPoint) (*uint64, error) {
	return GetFromPointTyped[uint64, P](s, point)
}

// 浮点数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointFloat32(point c_base.IPoint) (*float32, error) {
	return GetFromPointTyped[float32, P](s, point)
}

func (s *SRealDeviceImpl[P]) GetFromPointFloat64(point c_base.IPoint) (*float64, error) {
	return GetFromPointTyped[float64, P](s, point)
}

package c_device

import (
	"common/c_base"

	"github.com/pkg/errors"

	"github.com/shockerli/cvt"
)

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

func (s *SRealDeviceImpl[P]) GetFromProtocolBool(fc func(protocol P) (any, error)) (*bool, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.BoolP(v), nil
}

// 整数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolInt(fc func(protocol P) (any, error)) (*int, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.IntP(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt8(fc func(protocol P) (any, error)) (*int8, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int8P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt16(fc func(protocol P) (any, error)) (*int16, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int16P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt32(fc func(protocol P) (any, error)) (*int32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int32P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt64(fc func(protocol P) (any, error)) (*int64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int64P(v), nil
}

// 无符号整数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolUint(fc func(protocol P) (any, error)) (*uint, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.UintP(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint8(fc func(protocol P) (any, error)) (*uint8, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint8P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint16(fc func(protocol P) (any, error)) (*uint16, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint16P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint32(fc func(protocol P) (any, error)) (*uint32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint32P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint64(fc func(protocol P) (any, error)) (*uint64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint64P(v), nil
}

// 浮点数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolFloat32(fc func(protocol P) (any, error)) (*float32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Float32P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromProtocolFloat64(fc func(protocol P) (any, error)) (*float64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Float64P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPoint(point c_base.IPoint) (any, error) {
	return s.GetFromProtocol(func(protocol P) (any, error) {
		return protocol.GetValue(point)
	})
}

func (s *SRealDeviceImpl[P]) GetFromPointBool(point c_base.IPoint) (*bool, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.BoolP(v), nil
}

// 整数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointInt(point c_base.IPoint) (*int, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.IntP(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointInt8(point c_base.IPoint) (*int8, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int8P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointInt16(point c_base.IPoint) (*int16, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int16P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointInt32(point c_base.IPoint) (*int32, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int32P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointInt64(point c_base.IPoint) (*int64, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Int64P(v), nil
}

// 无符号整数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointUint(point c_base.IPoint) (*uint, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.UintP(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointUint8(point c_base.IPoint) (*uint8, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint8P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointUint16(point c_base.IPoint) (*uint16, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint16P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointUint32(point c_base.IPoint) (*uint32, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint32P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointUint64(point c_base.IPoint) (*uint64, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Uint64P(v), nil
}

// 浮点数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointFloat32(point c_base.IPoint) (*float32, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Float32P(v), nil
}

func (s *SRealDeviceImpl[P]) GetFromPointFloat64(point c_base.IPoint) (*float64, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return nil, err
	}
	if v == nil {
		return nil, nil
	}
	return cvt.Float64P(v), nil
}

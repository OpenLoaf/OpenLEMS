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

func (s *SRealDeviceImpl[P]) GetFromProtocolBool(fc func(protocol P) (any, error)) (bool, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return false, err
	}
	return cvt.BoolE(v)
}

// 整数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolInt(fc func(protocol P) (any, error)) (int, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.IntE(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt8(fc func(protocol P) (any, error)) (int8, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int8E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt16(fc func(protocol P) (any, error)) (int16, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int16E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt32(fc func(protocol P) (any, error)) (int32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int32E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolInt64(fc func(protocol P) (any, error)) (int64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int64E(v)
}

// 无符号整数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolUint(fc func(protocol P) (any, error)) (uint, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.UintE(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint8(fc func(protocol P) (any, error)) (uint8, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint8E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint16(fc func(protocol P) (any, error)) (uint16, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint16E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint32(fc func(protocol P) (any, error)) (uint32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint32E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolUint64(fc func(protocol P) (any, error)) (uint64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint64E(v)
}

// 浮点数类型函数
func (s *SRealDeviceImpl[P]) GetFromProtocolFloat32(fc func(protocol P) (any, error)) (float32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Float32E(v)
}

func (s *SRealDeviceImpl[P]) GetFromProtocolFloat64(fc func(protocol P) (any, error)) (float64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Float64E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPoint(point *c_base.Meta) (any, error) {
	return s.GetFromProtocol(func(protocol P) (any, error) {
		return protocol.GetValue(point)
	})
}

func (s *SRealDeviceImpl[P]) GetFromPointBool(point *c_base.Meta) (bool, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return false, err
	}
	return cvt.BoolE(v)
}

// 整数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointInt(point *c_base.Meta) (int, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.IntE(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt8(point *c_base.Meta) (int8, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Int8E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt16(point *c_base.Meta) (int16, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Int16E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt32(point *c_base.Meta) (int32, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Int32E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointInt64(point *c_base.Meta) (int64, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Int64E(v)
}

// 无符号整数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointUint(point *c_base.Meta) (uint, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.UintE(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint8(point *c_base.Meta) (uint8, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Uint8E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint16(point *c_base.Meta) (uint16, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Uint16E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint32(point *c_base.Meta) (uint32, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Uint32E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointUint64(point *c_base.Meta) (uint64, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Uint64E(v)
}

// 浮点数类型函数
func (s *SRealDeviceImpl[P]) GetFromPointFloat32(point *c_base.Meta) (float32, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Float32E(v)
}

func (s *SRealDeviceImpl[P]) GetFromPointFloat64(point *c_base.Meta) (float64, error) {
	v, err := s.GetFromPoint(point)
	if err != nil {
		return 0, err
	}
	return cvt.Float64E(v)
}

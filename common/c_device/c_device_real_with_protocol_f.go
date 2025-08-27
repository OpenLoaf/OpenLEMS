package c_device

import (
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

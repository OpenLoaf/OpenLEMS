package c_device

import (
	"gopkg.in/errgo.v2/fmt/errors"

	"github.com/shockerli/cvt"
)

func (s *SRealDevice[P]) GetFromProtocol(fc func(protocol P) (any, error)) (any, error) {
	if s.isProtocolNil() {
		return nil, errors.Newf("protocol is nil")
	}
	//device.protocol.GetValue()
	v, err := fc(s.protocol)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (s *SRealDevice[P]) GetFromProtocolBool(fc func(protocol P) (any, error)) (bool, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return false, err
	}
	return cvt.BoolE(v)
}

// 整数类型函数
func (s *SRealDevice[P]) GetFromProtocolInt(fc func(protocol P) (any, error)) (int, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.IntE(v)
}

func (s *SRealDevice[P]) GetFromProtocolInt8(fc func(protocol P) (any, error)) (int8, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int8E(v)
}

func (s *SRealDevice[P]) GetFromProtocolInt16(fc func(protocol P) (any, error)) (int16, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int16E(v)
}

func (s *SRealDevice[P]) GetFromProtocolInt32(fc func(protocol P) (any, error)) (int32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int32E(v)
}

func (s *SRealDevice[P]) GetFromProtocolInt64(fc func(protocol P) (any, error)) (int64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Int64E(v)
}

// 无符号整数类型函数
func (s *SRealDevice[P]) GetFromProtocolUint(fc func(protocol P) (any, error)) (uint, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.UintE(v)
}

func (s *SRealDevice[P]) GetFromProtocolUint8(fc func(protocol P) (any, error)) (uint8, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint8E(v)
}

func (s *SRealDevice[P]) GetFromProtocolUint16(fc func(protocol P) (any, error)) (uint16, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint16E(v)
}

func (s *SRealDevice[P]) GetFromProtocolUint32(fc func(protocol P) (any, error)) (uint32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint32E(v)
}

func (s *SRealDevice[P]) GetFromProtocolUint64(fc func(protocol P) (any, error)) (uint64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Uint64E(v)
}

// 浮点数类型函数
func (s *SRealDevice[P]) GetFromProtocolFloat32(fc func(protocol P) (any, error)) (float32, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Float32E(v)
}

func (s *SRealDevice[P]) GetFromProtocolFloat64(fc func(protocol P) (any, error)) (float64, error) {
	v, err := s.GetFromProtocol(fc)
	if err != nil {
		return 0, err
	}
	return cvt.Float64E(v)
}

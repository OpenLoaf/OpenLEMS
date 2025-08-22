package c_none

import (
	"canbus/p_canbus"
	"common/c_base"
	"common/c_proto"
	"context"
	"time"
)

type sNoneProtocol struct {
	sNoneAlarm
	sNoneDeviceRuntimeInfo
}

func (s *sNoneProtocol) RegisterHandler(handler func(ctx context.Context, status bool, isChange bool)) {

}

func (s *sNoneProtocol) GetStatus() *bool {
	return nil
}

func (s *sNoneProtocol) SetHigh() error {
	return NoneErr
}

func (s *sNoneProtocol) SetLow() error {
	return NoneErr
}

func (s *sNoneProtocol) RegisterReadTask(ctx context.Context, group *c_proto.SModbusTask, gs ...*c_proto.SModbusTask) {

}

func (s *sNoneProtocol) ReadSingleSync(meta *c_base.Meta, function c_proto.EModbusReadFunction, lifetime time.Duration, readCache bool) (any, error) {
	return nil, NoneErr
}

func (s *sNoneProtocol) ReadGroupSync(group *c_proto.SModbusTask, readCache bool, metas ...*c_base.Meta) ([]any, error) {
	return nil, NoneErr
}

func (s *sNoneProtocol) WriteSingleRegister(meta *c_base.Meta, value int32) error {
	return NoneErr
}

func (s *sNoneProtocol) WriteMultipleRegisters(group *c_proto.SModbusTask, values []int64) error {
	return NoneErr
}

func (s *sNoneProtocol) RegisterCanbusTask(group *p_canbus.SCanbusTask, gs ...*p_canbus.SCanbusTask) {

}

func (s *sNoneProtocol) GetValue(meta *c_base.Meta) (any, error) {
	return nil, NoneErr
}

func (s *sNoneProtocol) GetBool(meta *c_base.Meta) (bool, error) {
	return false, NoneErr
}

func (s *sNoneProtocol) GetIntValue(meta *c_base.Meta) (int, error) {
	return 0, NoneErr
}

func (s *sNoneProtocol) GetInt32Value(meta *c_base.Meta) (int32, error) {
	return 0, NoneErr
}

func (s *sNoneProtocol) GetUintValue(meta *c_base.Meta) (uint, error) {
	return 0, NoneErr
}

func (s *sNoneProtocol) GetUint32Value(meta *c_base.Meta) (uint32, error) {
	return 0, NoneErr
}

func (s *sNoneProtocol) GetFloat32Value(meta *c_base.Meta) (float32, error) {
	return 0, NoneErr
}

func (s *sNoneProtocol) GetFloat32Values(metas ...*c_base.Meta) ([]float32, error) {
	return nil, NoneErr
}

func (s *sNoneProtocol) GetFloat64Value(meta *c_base.Meta) (float64, error) {
	return 0, NoneErr
}

func (s *sNoneProtocol) GetFloat64Values(meta ...*c_base.Meta) ([]float64, error) {
	return nil, NoneErr
}

func (s *sNoneProtocol) SendMessage(task *p_canbus.SCanbusTask, values []int64) error {
	return NoneErr
}

func (s *sNoneProtocol) ProtocolListen() {

}

func (s *sNoneProtocol) IsActivate() bool {
	return true
}

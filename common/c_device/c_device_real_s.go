package c_device

import (
	"common/c_base"
	"common/c_enum"
	"context"
	"time"

	"github.com/pkg/errors"
)

type SRealDeviceImpl[P c_base.IProtocol] struct { // 真实设备
	DeviceCtx context.Context
	cancel    context.CancelFunc
	protocol  P
}

var _ c_base.IDevice = (*SRealDeviceImpl[c_base.IProtocol])(nil)

func NewRealDevice[P c_base.IProtocol](ctx context.Context, protocol P) (*SRealDeviceImpl[P], error) {
	deviceCtx, cancel := context.WithCancel(ctx)

	device := &SRealDeviceImpl[P]{
		DeviceCtx: deviceCtx,
		cancel:    cancel,
		protocol:  protocol,
	}

	return device, nil
}

func (s *SRealDeviceImpl[P]) IsVirtualDevice() bool {
	return false
}

func (s *SRealDeviceImpl[P]) GetAlarmLevel() c_enum.EAlarmLevel {
	if s.isProtocolNil() {
		return c_enum.EAlarmLevelError
	}
	return s.protocol.GetAlarmLevel()
}

func (s *SRealDeviceImpl[P]) GetAlarmList() []*c_base.SPointValue {
	if s.isProtocolNil() {
		return nil
	}
	return s.protocol.GetAlarmList()
}

func (s *SRealDeviceImpl[P]) UpdateAlarm(deviceId string, point c_base.IPoint, value any) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.UpdateAlarm(deviceId, point, value)
}

func (s *SRealDeviceImpl[P]) ResetAlarm() {
	if s.isProtocolNil() {
		return
	}
	s.protocol.ResetAlarm()
}

func (s *SRealDeviceImpl[P]) IgnoreClearAlarm(deviceId string, point string) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.IgnoreClearAlarm(deviceId, point)
}

func (s *SRealDeviceImpl[P]) RegisterAlarmHandlerFunc(alarmAction c_enum.EAlarmAction, handler func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isFirstHandler bool), sortValue ...int) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.RegisterAlarmHandlerFunc(alarmAction, handler)
}

func (s *SRealDeviceImpl[P]) GetProtocolStatus() c_enum.EProtocolStatus {
	if s.isProtocolNil() {
		return c_enum.EProtocolDisconnected
	}
	return s.protocol.GetProtocolStatus()
}

func (s *SRealDeviceImpl[P]) GetLastUpdateTime() *time.Time {
	if s.isProtocolNil() {
		return nil
	}
	return s.protocol.GetLastUpdateTime()
}

func (s *SRealDeviceImpl[P]) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.RegisterTask(task, tasks...)
}

func (s *SRealDeviceImpl[P]) GetServices() map[string]*c_base.SDriverService {
	return nil
}

func (s *SRealDeviceImpl[P]) GetPointValueList() []*c_base.SPointValue {
	return s.protocol.GetPointValueList()
}

func (s *SRealDeviceImpl[P]) GetConfig() *c_base.SDeviceConfig {
	return s.protocol.GetConfig()
}

func (s *SRealDeviceImpl[P]) isProtocolNil() bool {
	return any(s.protocol) == nil
}

func (s *SRealDeviceImpl[P]) ExecuteProtocolMethod(method func(protocol P) error) error {
	// 闭包执行协议方法，协议不存在将不会执行方法
	if s.isProtocolNil() {
		return errors.New("ExecuteProtocolMethod failed  protocol is nil")
	}

	return method(s.protocol)
}

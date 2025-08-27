package c_device

import (
	"common/c_base"
	"context"
	"github.com/pkg/errors"
	"time"
)

type SRealDeviceImpl[P c_base.IProtocol] struct { // 真实设备
	DeviceCtx    context.Context
	cancel       context.CancelFunc
	protocol     P
	deviceConfig *c_base.SDeviceConfig // 配置
}

var _ c_base.IDevice = (*SRealDeviceImpl[c_base.IProtocol])(nil)

func NewRealDevice[P c_base.IProtocol](ctx context.Context, deviceConfig *c_base.SDeviceConfig, protocol P) (*SRealDeviceImpl[P], error) {
	if deviceConfig == nil {
		// 必须有设备配置
		panic(errors.New("deviceConfig is nil"))
	}
	deviceCtx, cancel := context.WithCancel(ctx)

	device := &SRealDeviceImpl[P]{
		DeviceCtx:    deviceCtx,
		cancel:       cancel,
		protocol:     protocol,
		deviceConfig: deviceConfig,
	}

	return device, nil
}

func (s *SRealDeviceImpl[P]) GetAlarmLevel() c_base.EAlarmLevel {
	if s.isProtocolNil() {
		return c_base.EAlarmLevelError
	}
	return s.protocol.GetAlarmLevel()
}

func (s *SRealDeviceImpl[P]) GetAlarmList() []*c_base.MetaValueWrapper {
	if s.isProtocolNil() {
		return nil
	}
	return s.protocol.GetAlarmList()
}

func (s *SRealDeviceImpl[P]) UpdateAlarm(deviceId string, deviceType c_base.EDeviceType, meta *c_base.Meta, value any) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.UpdateAlarm(deviceId, deviceType, meta, value)
}

func (s *SRealDeviceImpl[P]) ResetAlarm() {
	if s.isProtocolNil() {
		return
	}
	s.protocol.ResetAlarm()
}

func (s *SRealDeviceImpl[P]) RegisterAlarmHandlerFunc(alarmAction c_base.EAlarmAction, handler func(alarm *c_base.MetaValueWrapper, currentMaxAlarmLevel c_base.EAlarmLevel, isFirstHandler bool), sortValue ...int) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.RegisterAlarmHandlerFunc(alarmAction, handler)
}

func (s *SRealDeviceImpl[P]) GetStatus() c_base.EProtocolStatus {
	if s.isProtocolNil() {
		return c_base.EProtocolDisconnected
	}
	return s.protocol.GetStatus()
}

func (s *SRealDeviceImpl[P]) GetLastUpdateTime() *time.Time {
	if s.isProtocolNil() {
		return nil
	}
	return s.protocol.GetLastUpdateTime()
}

func (s *SRealDeviceImpl[P]) RegisterTask(task c_base.ITask, tasks ...c_base.ITask) {
	if s.isProtocolNil() {
		return
	}
	s.protocol.RegisterTask(task, tasks...)
}

func (s *SRealDeviceImpl[P]) GetServices() map[string]*c_base.SDriverService {
	return nil
}

func (s *SRealDeviceImpl[P]) GetMetaValueList() []*c_base.MetaValueWrapper {
	return s.protocol.GetMetaValueList()
}

func (s *SRealDeviceImpl[P]) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
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

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

func (s *SRealDeviceImpl[P]) GetExportModbusPoints() []c_base.IPoint {
	return s.GetDevicePoints() // 直接使用设备point， 如果想自定义，就复写此方法
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

func (s *SRealDeviceImpl[P]) ClearAlarm(deviceId string, point string) {
	if s.isProtocolNil() {
		return
	}
	// IProtocol 继承 IAlarm，具体协议应实现 ClearAlarm；若未实现，将通过接口方法分发到协议实现
	s.protocol.ClearAlarm(deviceId, point)
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

func (s *SRealDeviceImpl[P]) GetProtocolPointValue(protocolPoint *c_base.SProtocolPoint) *c_base.SPointValue {
	if s.isProtocolNil() {
		return nil
	}
	return s.protocol.GetProtocolPointValue(protocolPoint)
}

func (s *SRealDeviceImpl[P]) GetConfig() *c_base.SDeviceConfig {
	return s.protocol.GetConfig()
}

// 实现新的IDevice接口方法 - 默认实现，子类可以覆盖
func (s *SRealDeviceImpl[P]) GetDevicePoints() []c_base.IPoint {
	return []c_base.IPoint{}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）- 默认实现，子类可以覆盖
func (s *SRealDeviceImpl[P]) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{}
}

func (s *SRealDeviceImpl[P]) isProtocolNil() bool {
	// 对于泛型类型，需要特殊处理 nil 检查
	// 由于泛型约束，我们需要通过接口来检查
	var zero P
	return any(s.protocol) == any(zero)
}

func (s *SRealDeviceImpl[P]) ExecuteProtocolMethod(method func(protocol P) error) error {
	// 闭包执行协议方法，协议不存在将不会执行方法
	if s.isProtocolNil() {
		return errors.New("ExecuteProtocolMethod failed  protocol is nil")
	}

	return method(s.protocol)
}

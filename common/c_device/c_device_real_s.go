package c_device

import (
	"common/c_base"
	"context"
	"errors"
	"reflect"
	"time"
)

type SRealDeviceImpl[P c_base.IProtocol] struct { // 真实设备
	c_base.IAlarm
	DeviceCtx    context.Context
	cancel       context.CancelFunc
	protocol     P
	deviceConfig *c_base.SDeviceConfig // 配置

	notifyChan chan<- *c_base.SAlarmDetail // 告警通知通道

	alarmHandlerList []func(maxAlarm c_base.EAlarmLevel, nowAlarm *c_base.SAlarmDetail)
}

func NewRealDevice[P c_base.IProtocol](ctx context.Context, deviceConfig *c_base.SDeviceConfig, protocol P) (*SRealDeviceImpl[P], error) {
	if deviceConfig == nil {
		// 必须有设备配置
		panic(errors.New("deviceConfig is nil"))
	}
	deviceCtx, cancel := context.WithCancel(ctx)
	var notifyChan = make(chan *c_base.SAlarmDetail)

	device := &SRealDeviceImpl[P]{
		DeviceCtx:        deviceCtx,
		cancel:           cancel,
		IAlarm:           c_base.NewAlarm(notifyChan),
		protocol:         protocol,
		deviceConfig:     deviceConfig,
		alarmHandlerList: make([]func(maxAlarm c_base.EAlarmLevel, nowAlarm *c_base.SAlarmDetail), 0),
	}

	go func() {
		for {
			select {
			case <-deviceCtx.Done():
				return
			case alarm := <-notifyChan:
				// 这里只解决了当前类的告警等级处理问题。没解决dirver那边如何处理告警
				for _, handler := range device.alarmHandlerList {
					go handler(c_base.ENone, alarm)
				}
			}
		}
	}()

	return device, nil
}

func (s *SRealDeviceImpl[P]) GetStatus() c_base.EProtocolStatus {
	if reflect.ValueOf(s.protocol).IsNil() {
		return c_base.EProtocolDisconnected
	}
	return s.protocol.GetStatus()
}

func (s *SRealDeviceImpl[P]) GetLastUpdateTime() *time.Time {
	if reflect.ValueOf(s.protocol).IsNil() {
		return nil
	}
	return s.protocol.GetLastUpdateTime()
}

func (s *SRealDeviceImpl[P]) RegisterTask(task c_base.ITask, tasks ...c_base.ITask) {
	if reflect.ValueOf(s.protocol).IsNil() {
		return
	}
	s.protocol.RegisterTask(task, tasks...)
}

func (s *SRealDeviceImpl[P]) GetServices() map[string]*c_base.SDriverService {
	return nil
}

// 注册告警处理器
func (s *SRealDeviceImpl[P]) RegisterAlarmHandler(handler func(maxAlarm c_base.EAlarmLevel, nowAlarm *c_base.SAlarmDetail)) {
	s.alarmHandlerList = append(s.alarmHandlerList, handler)
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

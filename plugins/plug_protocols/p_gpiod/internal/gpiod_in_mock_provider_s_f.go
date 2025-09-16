//go:build !linux

package internal

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

// sGpiodInMockProvider 因为Linux才可以操作gpio，所以模拟了一个输入provider。当打为非linux平台的时候，就会使用这个mock的
type sGpiodInMockProvider struct {
	c_base.IAlarm
	ctx          context.Context
	gpiodConfig  *c_proto.SGpiodProtocolConfig
	deviceConfig *c_base.SDeviceConfig

	// 状态管理
	currentStatus  *bool
	point          c_base.IPoint
	lastUpdateTime *time.Time

	handler func(status bool, isChange bool)
}

var _ c_proto.IGpiodProtocol = (*sGpiodInMockProvider)(nil)

// NewGpiodInProvider 创建新的GPIO输入provider
func NewGpiodInProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	// 解析协议配置
	gpiodConfig := &c_proto.SGpiodProtocolConfig{}
	if err := gconv.Scan(clientConfig.Params, gpiodConfig); err != nil {
		return nil, fmt.Errorf("failed to parse gpiod protocol config: %w", err)
	}

	// 验证方向必须是输入
	if gpiodConfig.Direction != c_enum.EGpioDirectionIn {
		return nil, fmt.Errorf("gpiod in provider only supports input direction, got: %v", gpiodConfig.Direction)
	}

	return &sGpiodInMockProvider{
		IAlarm:       c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		gpiodConfig:  gpiodConfig,
		deviceConfig: deviceConfig,
		ctx:          ctx,
	}, nil
}

func (s *sGpiodInMockProvider) InitGpioPoint(point c_base.IPoint) {
	s.point = point
}

func (s *sGpiodInMockProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	return c_enum.EProtocolConnected
}

func (s *sGpiodInMockProvider) GetLastUpdateTime() *time.Time {
	return s.lastUpdateTime
}

func (s *sGpiodInMockProvider) GetPointValueList() []*c_base.SPointValue {
	if s.point == nil {
		return nil
	}

	var point *c_base.SPointValue
	if s.currentStatus != nil {
		point = c_base.NewPointValue(s.deviceConfig.Id, s.point, *s.currentStatus)
	} else {
		point = c_base.NewPointValue(s.deviceConfig.Id, s.point, nil)
	}

	return []*c_base.SPointValue{point}
}

func (s *sGpiodInMockProvider) GetValue(point c_base.IPoint) (any, error) {
	if point == s.point {
		return s.currentStatus, nil
	}
	return nil, nil
}

func (s *sGpiodInMockProvider) ProtocolListen() {
	// Mock输入provider的监听逻辑
}

func (s *sGpiodInMockProvider) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *sGpiodInMockProvider) RegisterHandler(handler func(status bool, isChange bool)) {
	s.handler = handler
}

func (s *sGpiodInMockProvider) GetStatus() *bool {
	return s.currentStatus
}

func (s *sGpiodInMockProvider) SetHigh() error {
	return fmt.Errorf("cannot set value on input GPIO pin")
}

func (s *sGpiodInMockProvider) SetLow() error {
	return fmt.Errorf("cannot set value on input GPIO pin")
}

// updateStatus 更新GPIO状态和时间戳，并处理状态变化
func (s *sGpiodInMockProvider) updateStatus(status bool) {
	// 保存之前的状态
	last := s.currentStatus

	// 更新状态和时间戳
	s.currentStatus = &status
	now := time.Now()
	s.lastUpdateTime = &now

	// 更新告警状态
	if s.point != nil {
		s.UpdateAlarm(s.deviceConfig.Id, s.point, status)
	}

	// 判断是否有状态变化
	isChange := false
	if last == nil {
		isChange = true
	} else if *last != status {
		isChange = true
	}

	// 调用处理函数
	if s.handler != nil {
		s.handler(status, isChange)
	}
}

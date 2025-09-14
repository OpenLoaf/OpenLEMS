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

// sGpiodMockProvider 因为Linux才可以操作gpio，所以模拟了一个。当打为非linux平台的时候，就会使用这个mock的
type sGpiodMockProvider struct {
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

var _ c_proto.IGpiodProtocol = (*sGpiodMockProvider)(nil)

// NewGpiodProvider 创建新的GPIO provider
func NewGpiodProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	// 解析协议配置
	gpiodConfig := &c_proto.SGpiodProtocolConfig{}
	if err := gconv.Scan(clientConfig.Params, gpiodConfig); err != nil {
		return nil, fmt.Errorf("failed to parse gpiod protocol config: %w", err)
	}

	return &sGpiodMockProvider{
		IAlarm:       c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		gpiodConfig:  gpiodConfig,
		deviceConfig: deviceConfig,
	}, nil
}

func (s *sGpiodMockProvider) InitGpioPoint(point c_base.IPoint) {
	s.point = point
}

func (s *sGpiodMockProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	return c_enum.EProtocolMock
}

func (s *sGpiodMockProvider) GetLastUpdateTime() *time.Time {
	return s.lastUpdateTime
}

func (s *sGpiodMockProvider) GetPointValueList() []*c_base.SPointValue {
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

func (s *sGpiodMockProvider) GetValue(point c_base.IPoint) (any, error) {
	if point == s.point {
		return s.currentStatus, nil
	}
	return nil, nil
}

func (s *sGpiodMockProvider) ProtocolListen() {

}

func (s *sGpiodMockProvider) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *sGpiodMockProvider) RegisterHandler(handler func(status bool, isChange bool)) {
	s.handler = handler
}

func (s *sGpiodMockProvider) GetGpioStatus() *bool {
	return s.currentStatus
}

func (s *sGpiodMockProvider) SetHigh() error {
	s.updateStatus(high)
	return nil
}

func (s *sGpiodMockProvider) SetLow() error {
	s.updateStatus(low)
	return nil
}

// updateStatus 更新GPIO状态和时间戳，并处理状态变化
func (s *sGpiodMockProvider) updateStatus(status bool) {
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

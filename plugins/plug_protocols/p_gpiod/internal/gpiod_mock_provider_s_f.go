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
	gpiodConfig  *c_proto.SGpiodProtocolConfig
	deviceConfig *c_base.SDeviceConfig

	// 状态管理
	currentStatus  *bool
	lastUpdateTime *time.Time

	handler func(status bool)
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
		IAlarm:      c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		gpiodConfig: gpiodConfig,
	}, nil
}

func (s *sGpiodMockProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	return c_enum.EProtocolMock
}

func (s *sGpiodMockProvider) GetLastUpdateTime() *time.Time {
	return s.lastUpdateTime
}

func (s *sGpiodMockProvider) GetPointValueList() []*c_base.SPointValue {
	return []*c_base.SPointValue{
		c_base.NewPointValue(s.deviceConfig.Id, gpioPoint, s.currentStatus),
	}
}

func (s *sGpiodMockProvider) GetValue(point c_base.IPoint) (any, error) {
	return nil, nil
}

func (s *sGpiodMockProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {

}

func (s *sGpiodMockProvider) ProtocolListen() {

}

func (s *sGpiodMockProvider) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *sGpiodMockProvider) RegisterHandler(handler func(status bool)) {
	s.handler = handler
}

func (s *sGpiodMockProvider) GetGpioStatus() *bool {
	return s.currentStatus
}

func (s *sGpiodMockProvider) SetHigh() error {
	s.currentStatus = &high
	return nil
}

func (s *sGpiodMockProvider) SetLow() error {
	s.currentStatus = &low
	return nil
}

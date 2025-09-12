//go:build !linux

package internal

import (
	"common/c_base"
	"common/c_enum"
	"common/c_proto"
	"context"
	"time"
)

type sGpiodProvider struct {
}

func (s *sGpiodProvider) GetAlarmLevel() c_enum.EAlarmLevel {

	return c_enum.EAlarmLevelNone
}

func (s *sGpiodProvider) GetAlarmList() []*c_base.SPointValue {

	return nil
}

func (s *sGpiodProvider) UpdateAlarm(deviceId string, point c_base.IPoint, value any) {
}

func (s *sGpiodProvider) ResetAlarm() {
}

func (s *sGpiodProvider) IgnoreClearAlarm(deviceId string, point string) {
}

func (s *sGpiodProvider) RegisterAlarmHandlerFunc(alarmAction c_enum.EAlarmAction, handler func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isFirstHandler bool), sortValue ...int) {

}

func (s *sGpiodProvider) GetProtocolStatus() c_enum.EProtocolStatus {

	return c_enum.EProtocolDisconnected
}

func (s *sGpiodProvider) GetLastUpdateTime() *time.Time {

	return nil
}

func (s *sGpiodProvider) GetPointValueList() []*c_base.SPointValue {

	return nil
}

func (s *sGpiodProvider) GetValue(point c_base.IPoint) (any, error) {
	return nil, nil
}

func (s *sGpiodProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
}

func (s *sGpiodProvider) ProtocolListen() {

}

func (s *sGpiodProvider) GetConfig() *c_base.SDeviceConfig {

	return nil
}

func (s *sGpiodProvider) RegisterHandler(handler func(status bool)) {

}

func (s *sGpiodProvider) GetGpioStatus() *bool {

	return nil
}

func (s *sGpiodProvider) SetHigh() error {

	return nil
}

func (s *sGpiodProvider) SetLow() error {

	return nil
}

var _ c_proto.IGpiodProtocol = (*sGpiodProvider)(nil)

// NewGpiodProvider 创建新的GPIO provider
func NewGpiodProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {

	provider := &sGpiodProvider{}

	return provider, nil
}

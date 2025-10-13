package c_base

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"time"

	"github.com/shockerli/cvt"
)

type SPointValue struct {
	IPoint               //  点位
	deviceId   string    //  设备ID
	value      any       //  点位值
	happenTime time.Time //  发生时间

	isTrigger *bool
	level     c_enum.EAlarmLevel // 告警等级

}

// NewPointValue  创建点位Value
func NewPointValue(deviceId string, IPoint IPoint, value any) *SPointValue {
	return &SPointValue{
		deviceId:   deviceId,
		IPoint:     IPoint,
		value:      value,
		happenTime: time.Now(),
	}
}

func (s *SPointValue) GetDeviceId() string {
	return s.deviceId
}

func (s *SPointValue) GetValue() any {
	return s.value
}

func (s *SPointValue) GetLevel() c_enum.EAlarmLevel {
	if s.IPoint == nil {
		return c_enum.EAlarmLevelNone
	}
	if s.isTrigger == nil {
		s.IsAlarmTrigger()
	}
	return s.level
}

func (s *SPointValue) IsAlarmTrigger() bool {
	if s.isTrigger == nil {
		trigger, level, err := s.TriggerAlarm(s.value)
		if err != nil {
			ctx := context.WithValue(context.Background(), c_enum.ELogTypeDevice, s.deviceId)
			c_log.BizErrorf(ctx, "告警触发函数错误: %v", err)
			return false
		}
		s.isTrigger = &trigger
		s.level = level
	}
	return *s.isTrigger
}

// GetActualValueExplain 获取实际的值解释
func (s *SPointValue) GetActualValueExplain() (string, error) {
	if s.IPoint != nil {
		return s.IPoint.GetValueExplainByValue(s.value)
	}
	return cvt.StringE(s.value)
}

// GetActualValueType 获取实际的值类型, 如果是auto的将会转换为value的实际类型
func (s *SPointValue) GetActualValueType() c_enum.EValueType {
	if s.GetValueType() != c_enum.EAuto {
		return s.GetValueType()
	}
	return ResolvingValueType(s.value)
}

func (s *SPointValue) GetHappenTime() time.Time {
	return s.happenTime
}

func (s *SPointValue) SetHappenTime(happenTime time.Time) {
	s.happenTime = happenTime
}

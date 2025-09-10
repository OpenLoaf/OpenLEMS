package c_base

import (
	"common/c_enum"
	"time"

	"github.com/shockerli/cvt"
)

type SPointValue struct {
	IPoint                        //  点位
	level      c_enum.EAlarmLevel // 告警等级
	deviceId   string             //  设备ID
	value      any                //  点位值
	happenTime time.Time          //  发生时间
}

// NewPointValue  创建点位Value
func NewPointValue(deviceId string, IPoint IPoint, level c_enum.EAlarmLevel, value any) *SPointValue {
	return &SPointValue{
		deviceId:   deviceId,
		level:      level,
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
	return s.level
}

func (s *SPointValue) GetValueExplain() (string, error) {
	if s.IPoint != nil {
		if explainer, ok := s.IPoint.(interface {
			ValueExplain(value any) (string, error)
		}); ok {
			return explainer.ValueExplain(s.value)
		}
	}
	return cvt.StringE(s.value)
}

func (s *SPointValue) GetHappenTime() time.Time {
	return s.happenTime
}

func (s *SPointValue) SetHappenTime(happenTime time.Time) {
	s.happenTime = happenTime
}

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
	deviceType c_enum.EDeviceType //  设备类型
	value      any                //  点位值
	happenTime time.Time          //  发生时间
}

// NewPointValue  创建点位Value
func NewPointValue(deviceId string, deviceType c_enum.EDeviceType, IPoint IPoint, value any) *SPointValue {
	return &SPointValue{
		deviceId:   deviceId,
		deviceType: deviceType,
		IPoint:     IPoint,
		value:      value,
		happenTime: time.Now(),
	}
}

func (s *SPointValue) GetDeviceId() string {
	return s.deviceId
}

func (s *SPointValue) GetDeviceType() c_enum.EDeviceType {
	return s.deviceType
}

func (s *SPointValue) GetValue() any {
	return s.value
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

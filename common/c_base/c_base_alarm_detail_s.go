package c_base

import (
	"fmt"
	"time"
)

type SAlarmDetail struct {
	DeviceId   string      `json:"deviceId" dc:"设备ID"`
	DeviceType EDeviceType `json:"deviceType" dc:"设备类型"`
	Level      EAlarmLevel `json:"level" dc:"告警级别"`
	Meta       *Meta       `json:"meta" dc:"告警元数据"`
	HappenTime *time.Time  `json:"happenTime" dc:"发生时间"`
	IsTrigger  bool        `json:"isTrigger" dc:"是否触发"`
	Value      any         `json:"value" dc:"数值"`
}

func (s *SAlarmDetail) ToString() string {
	if s.IsTrigger {
		return fmt.Sprintf("++ 触发 ID：%s [%s], 点位: %s(%s),发生时间: %s,数值: %v", s.DeviceId, s.Level.String(), s.Meta.Name, s.Meta.Cn, s.HappenTime.Format("2006-01-02 15:04:05"), s.Value)
	} else {
		return fmt.Sprintf("-- 解除 ID：%s [%s], 点位: %s(%s),发生时间: %s,数值: %v", s.DeviceId, s.Level.String(), s.Meta.Name, s.Meta.Cn, s.HappenTime.Format("2006-01-02 15:04:05"), s.Value)
	}
}

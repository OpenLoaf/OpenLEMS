package gpio_in_basic_v1

import (
	"strings"

	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_type"

	"github.com/shockerli/cvt"
)

type sBasicGpioIn struct {
	*c_device.SRealDeviceImpl[c_proto.IGpiodProtocol]
	gpioPoint        *c_base.SPoint
	GpioDeviceConfig *c_proto.SGpioDeviceConfig
}

func (s *sBasicGpioIn) SetHigh() error {
	//  调用协议层的SetHigh方法
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetHigh()
	})
}

func (s *sBasicGpioIn) SetLow() error {
	// 调用协议层的SetLow方法
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetLow()
	})
}

var _ c_type.IGpioOut = (*sBasicGpioIn)(nil)

func (s *sBasicGpioIn) Shutdown() {

}

func (s *sBasicGpioIn) Init() error {

	err := s.GetConfig().ScanParams(s.GpioDeviceConfig)
	if err != nil {
		c_log.BizErrorf(s.DeviceCtx, "Init Device GpioDeviceConfig Error: %s", err.Error())
		return err
	}

	s.gpioPoint = &c_base.SPoint{
		Key:     "status",
		Name:    "状态",
		Group:   c_base.GroupTotal,
		Precise: 0,
		Hidden:  true,
	}

	// 设置状态解释
	if s.GpioDeviceConfig.HighTrigger {
		s.gpioPoint.ValueExplain = []*c_base.SFieldExplain{
			{Key: "true", Value: s.parseLabelName(s.GpioDeviceConfig.ClearName), Color: s.parseLabelColor(s.GpioDeviceConfig.ClearName)},
			{Key: "false", Value: s.parseLabelName(s.GpioDeviceConfig.TriggerName), Color: s.parseLabelColor(s.GpioDeviceConfig.TriggerName)},
		}
	} else {
		s.gpioPoint.ValueExplain = []*c_base.SFieldExplain{
			{Key: "false", Value: s.parseLabelName(s.GpioDeviceConfig.ClearName), Color: s.parseLabelColor(s.GpioDeviceConfig.ClearName)},
			{Key: "true", Value: s.parseLabelName(s.GpioDeviceConfig.TriggerName), Color: s.parseLabelColor(s.GpioDeviceConfig.TriggerName)},
		}
	}

	if s.GpioDeviceConfig.Level != c_enum.EAlarmLevelNone {
		// 触发告警
		s.gpioPoint.Trigger = func(value interface{}) (trigger bool, level c_enum.EAlarmLevel, err error) {
			trigger, err = cvt.BoolE(value)
			if !s.GpioDeviceConfig.HighTrigger {
				trigger = !trigger
			}
			level = s.GpioDeviceConfig.Level
			return
		}
	}

	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.InitGpioPoint(s.gpioPoint)
		return nil
	})
	return nil
}

func (s *sBasicGpioIn) RegisterHandler(handler func(status bool, isChange bool)) {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.RegisterHandler(handler)
		return nil
	})
}

func (s *sBasicGpioIn) GetStatus() *bool {
	v, err := s.GetFromProtocolBool(func(protocol c_proto.IGpiodProtocol) (any, error) {
		return protocol.GetStatus(), nil
	})
	if err != nil {
		return nil
	}
	return v
}

func (s *sBasicGpioIn) GetDevicePoints() []c_base.IPoint {
	// 返回GPIO协议点位
	return []c_base.IPoint{
		s.gpioPoint,
	}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (s *sBasicGpioIn) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		s.gpioPoint, // GPIO状态 - 最重要的状态信息
	}
}

// parseLabelName 解析标签名称，从 "名称|#颜色代码" 格式中提取名称
func (s *sBasicGpioIn) parseLabelName(label string) string {
	if strings.Contains(label, "|") {
		parts := strings.SplitN(label, "|", 2)
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	return label
}

// parseLabelColor 解析标签颜色，从 "名称|#颜色代码" 格式中提取颜色
func (s *sBasicGpioIn) parseLabelColor(label string) string {
	if strings.Contains(label, "|") {
		parts := strings.SplitN(label, "|", 2)
		if len(parts) > 1 {
			return strings.TrimSpace(parts[1])
		}
	}
	return ""
}

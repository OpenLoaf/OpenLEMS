package gpio_out_basic_v1

import (
	"errors"

	"common/c_base"
	"common/c_device"
	"common/c_log"
	"common/c_proto"
	"common/c_type"
)

type sBasicGpioOut struct {
	*c_device.SRealDeviceImpl[c_proto.IGpiodProtocol]

	GpioDeviceConfig *c_proto.SGpioDeviceConfig
}

var gpioPoint = &c_base.SPoint{
	Key:     "status",
	Name:    "状态",
	Group:   c_base.GroupTotal,
	Precise: 0,
	Hidden:  true,
}

var _ c_type.IGpioOut = (*sBasicGpioOut)(nil)

func (s *sBasicGpioOut) Shutdown() {

}

func (s *sBasicGpioOut) Init() error {

	err := s.GetConfig().ScanParams(s.GpioDeviceConfig)
	if err != nil {
		c_log.BizErrorf(s.DeviceCtx, "Init Device GpioDeviceConfig Error: %s", err.Error())
		return err
	}

	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.InitGpioPoint(gpioPoint)
		return nil
	})
	return nil
}

func (s *sBasicGpioOut) SetHigh() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetHigh()
	})
}

func (s *sBasicGpioOut) SetLow() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		return protocol.SetLow()
	})
}

func (s *sBasicGpioOut) RegisterHandler(handler func(status bool, isChange bool)) {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IGpiodProtocol) error {
		protocol.RegisterHandler(handler)
		return nil
	})
}

func (s *sBasicGpioOut) GetStatus() *bool {
	v, err := s.GetFromProtocolBool(func(protocol c_proto.IGpiodProtocol) (any, error) {
		return protocol.GetStatus(), nil
	})
	if err != nil {
		return nil
	}
	return v
}

func (s *sBasicGpioOut) StatusToggle() error {
	// 获取当前状态
	currentStatus := s.GetStatus()
	if currentStatus == nil {
		return errors.New("无法获取当前GPIO状态")
	}

	// 根据当前状态切换到相反状态
	if *currentStatus {
		// 当前是高电平，切换到低电平
		c_log.BizInfo(s.DeviceCtx, "从高电平切换到低电平")
		return s.SetLow()
	} else {
		// 当前是低电平，切换到高电平
		c_log.BizInfo(s.DeviceCtx, "从低电平切换到高电平")
		return s.SetHigh()
	}
}

// 实现新的IDevice接口方法
func (s *sBasicGpioOut) GetTelemetryPoints() []c_base.IPoint {
	// GPIO驱动没有遥测点位，返回空列表
	return []c_base.IPoint{}
}

func (s *sBasicGpioOut) GetProtocolPoints() []c_base.IPoint {
	// 返回GPIO协议点位
	return []c_base.IPoint{
		gpioPoint,
	}
}

func (s *sBasicGpioOut) GetConfigPoints() []*c_base.SConfigPoint {
	// 从配置结构体转换而来
	configPoints, err := c_base.BuildConfigPoints(s.GpioDeviceConfig)
	if err != nil {
		c_log.Errorf(s.DeviceCtx, "构建配置点位失败: %v", err)
		return []*c_base.SConfigPoint{}
	}
	return configPoints
}

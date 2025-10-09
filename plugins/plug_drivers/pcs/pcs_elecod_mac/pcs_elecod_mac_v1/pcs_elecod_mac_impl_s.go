package pcs_elecod_mac_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"common/c_type"
	"fmt"
	"pcs_elecod/pcs_elecod_mac_v1/elecod_mac_defined"
)

type sPcsElecodMac struct {
	*c_device.SRealDeviceImpl[c_proto.ICanbusProtocol]
	pcsConfig *sPcsElecodMacConfig
}

func (s *sPcsElecodMac) GetFrequency() (*float64, error) {
	return s.GetFromPointFloat64(elecod_mac_defined.AnalogGridFrequencyA)
}

var _ c_type.IPcs = (*sPcsElecodMac)(nil)

func (s *sPcsElecodMac) Init() error {
	// 解析配置
	s.pcsConfig = &sPcsElecodMacConfig{}
	err := s.GetConfig().ScanParams(s.pcsConfig)
	if err != nil {
		return fmt.Errorf("PcsElecodMac配置解析失败：内容:%v 原因: %s", s.GetConfig().Params, err.Error())
	}

	if s.pcsConfig.MacAddress == nil || s.pcsConfig.SelfAddress == nil {
		return fmt.Errorf("PcsElecodMac配置解析失败：缺少配置项！当前配置：%v", s.GetConfig().Params)
	}

	_ = s.ExecuteProtocolMethod(func(protocol c_proto.ICanbusProtocol) error {
		// 注册任务
		for _, task := range elecod_mac_defined.AnalogAllTasks {
			protocol.RegisterTask(task)
			c_log.Infof(s.DeviceCtx, "注册%v", task)
		}
		for _, task := range elecod_mac_defined.ConfigAllTasks {
			protocol.RegisterTask(task)
			c_log.Infof(s.DeviceCtx, "注册%v", task)
		}
		return nil
	})

	c_log.Info(s.DeviceCtx, "PcsElecodMac 初始化完成")
	return nil
}

func (s *sPcsElecodMac) Shutdown() {
	c_log.Info(s.DeviceCtx, "PcsElecodMac Shutdown")
}

func (s *sPcsElecodMac) SetReset() error {
	c_log.Warningf(s.DeviceCtx, "sPcsElecodMac SetReset() not support!")
	return nil
}

func (s *sPcsElecodMac) SetStatus(status c_enum.EEnergyStoreStatus) error {
	// TODO: 实现状态设置逻辑
	c_log.Warningf(s.DeviceCtx, "sPcsElecodMac SetStatus() not implemented!")
	return c_base.NotSupport
}

func (s *sPcsElecodMac) SetGridMode(mode c_enum.EGridMode) error {
	c_log.Warningf(s.DeviceCtx, "sPcsElecodMac SetGridMode() not support!")
	return c_base.NotSupport
}

func (s *sPcsElecodMac) GetStatus() (*c_enum.EEnergyStoreStatus, error) {
	// TODO: 实现状态获取逻辑
	status := c_enum.EPcsStatusUnknown
	return &status, c_base.NotSupport
}

func (s *sPcsElecodMac) GetGridMode() (*c_enum.EGridMode, error) {
	mode := c_enum.EGridOn
	return &mode, nil
}

func (s *sPcsElecodMac) SetPower(power int32) error {
	// TODO: 实现功率设置逻辑
	c_log.Warningf(s.DeviceCtx, "sPcsElecodMac SetPower() not implemented!")
	return c_base.NotSupport
}

func (s *sPcsElecodMac) SetReactivePower(power int32) error {
	// TODO: 实现无功功率设置逻辑
	c_log.Warningf(s.DeviceCtx, "sPcsElecodMac SetReactivePower() not implemented!")
	return c_base.NotSupport
}

func (s *sPcsElecodMac) SetPowerFactor(factor float32) error {
	c_log.Warningf(s.DeviceCtx, "sPcsElecodMac SetPowerFactor() not support!")
	return c_base.NotSupport
}

func (s *sPcsElecodMac) GetTargetPower() (*int32, error) {
	// TODO: 实现目标功率获取逻辑
	return nil, nil
}

func (s *sPcsElecodMac) GetTargetReactivePower() (*int32, error) {
	// TODO: 实现目标无功功率获取逻辑
	return nil, nil
}

func (s *sPcsElecodMac) GetTargetPowerFactor() (*float32, error) {
	factor := float32(-1)
	return &factor, nil
}

func (s *sPcsElecodMac) GetPower() (*float64, error) {
	// 使用新的协议方法获取功率值
	return s.GetFromPointFloat64(elecod_mac_defined.AnalogTotalActivePower)
}

func (s *sPcsElecodMac) GetApparentPower() (*float64, error) {
	return s.GetFromPointFloat64(elecod_mac_defined.AnalogTotalApparentPower)
}

func (s *sPcsElecodMac) GetReactivePower() (*float64, error) {
	return s.GetFromPointFloat64(elecod_mac_defined.AnalogTotalReactivePower)
}

func (s *sPcsElecodMac) GetRatedPower() (*uint32, error) {
	// TODO: 实现额定功率获取逻辑
	power := uint32(100) // 默认值
	return &power, nil
}

func (s *sPcsElecodMac) GetMaxInputPower() (*float32, error) {
	// TODO: 实现最大输入功率获取逻辑
	power := float32(100) // 默认值
	return &power, nil
}

func (s *sPcsElecodMac) GetMaxOutputPower() (*float32, error) {
	// TODO: 实现最大输出功率获取逻辑
	power := float32(100) // 默认值
	return &power, nil
}

func (s *sPcsElecodMac) GetTodayIncomingQuantity() (*float64, error) {
	// TODO: 实现今日充电量获取逻辑
	return nil, nil
}

func (s *sPcsElecodMac) GetHistoryIncomingQuantity() (*float64, error) {
	// TODO: 实现历史充电量获取逻辑
	return nil, nil
}

func (s *sPcsElecodMac) GetTodayOutgoingQuantity() (*float64, error) {
	// TODO: 实现今日放电量获取逻辑
	return nil, nil
}

func (s *sPcsElecodMac) GetHistoryOutgoingQuantity() (*float64, error) {
	// TODO: 实现历史放电量获取逻辑
	return nil, nil
}

func (s *sPcsElecodMac) GetIGBTTemperature() (*float32, error) {
	// TODO: 实现IGBT温度获取逻辑
	return nil, nil
}

// 实现新的IDevice接口方法
func (s *sPcsElecodMac) GetTelemetryPoints() []c_base.IPoint {
	// 返回CAN总线协议点位
	return []c_base.IPoint{
		elecod_mac_defined.AnalogTotalActivePower,
		elecod_mac_defined.AnalogTotalReactivePower,
		elecod_mac_defined.AnalogTotalApparentPower,
		elecod_mac_defined.AnalogTotalPowerFactor,
		// 可以继续添加其他协议点位
	}
}

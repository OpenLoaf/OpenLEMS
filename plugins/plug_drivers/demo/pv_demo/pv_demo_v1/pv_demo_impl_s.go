package pv_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
)

type sPvDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

func (s *sPvDemo) Init() error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(ReadTask)
		return nil
	})
}

func (s *sPvDemo) Shutdown() {

}

func (s *sPvDemo) GetStatus() (*c_enum.EPvStatus, error) {
	power, err := s.GetPower()
	if err != nil {
		status := c_enum.EPvUnknown
		return &status, err
	}
	if power == nil || *power == 0.0 {
		status := c_enum.EPvStandby
		return &status, nil
	}
	status := c_enum.EPvGenerate
	return &status, nil
}

func (s *sPvDemo) SetPower(power float64) error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		return protocol.WriteSingleRegister(PowerLimit, int32(power))
	})
}

func (s *sPvDemo) SetReactivePower(power float64) error {
	return c_base.NotSupport
}

func (s *sPvDemo) SetPowerFactor(factor float32) error {
	return c_base.NotSupport
}

func (s *sPvDemo) GetTargetPower() (*float64, error) {
	return s.GetFromPointFloat64(PowerLimit)
}

func (s *sPvDemo) GetTargetReactivePower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetTargetPowerFactor() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetPower() (*float64, error) {
	return s.GetFromPointFloat64(Power)
}

func (s *sPvDemo) GetApparentPower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetReactivePower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetDcPower() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetDcVoltage() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetDcCurrent() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetTodayIncomingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetHistoryIncomingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetTodayOutgoingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sPvDemo) GetHistoryOutgoingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(GeneratedEnergy)
}

func (s *sPvDemo) GetCapacity() (*uint32, error) {
	return s.GetFromPointUint32(InstalledCapacity)
}

func (s *sPvDemo) GetTemperature() (*uint32, error) {
	return s.GetFromPointUint32(Temperature)
}

func (s *sPvDemo) GetIrradiance() (*uint32, error) {
	return s.GetFromPointUint32(InstalledCapacity)
}

// 实现新的IDevice接口方法
func (s *sPvDemo) GetDevicePoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 协议点位
		Status,
		Power,
		GeneratedEnergy,
		OnOffState,
		PowerLimit,
		InstalledCapacity,
		Irradiance,
		Temperature,
		Efficiency,
	}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (s *sPvDemo) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		OnOffState,      // 开关状态
		Power,           // 发电功率 - 核心输出
		GeneratedEnergy, // 发电量 - 累计统计
		Temperature,     // 温度
		Irradiance,      // 当前辐射
	}
}

// GetExportModbusPoints 获取暴露出去的modbus点位
func (s *sPvDemo) GetExportModbusPoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 协议点位 - 控制参数
		Status,            // 设备状态
		Power,             // 功率
		GeneratedEnergy,   // 发电量
		OnOffState,        // 开关状态
		PowerLimit,        // 功率限制
		InstalledCapacity, // 装机容量
		Irradiance,        // 辐照度
		Temperature,       // 温度
		Efficiency,        // 效率
	}
}

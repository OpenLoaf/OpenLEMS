package ess_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
	"common/c_type"

	"github.com/shockerli/cvt"
)

type sEssDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

var _ c_type.IEnergyStore = (*sEssDemo)(nil)

func (s *sEssDemo) Init() error {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(ReadTask)
		return nil
	})
	return nil
}

func (s *sEssDemo) Shutdown() {
	_ = s.SetPower(0)
	_ = s.SetStatus(c_enum.EPcsStatusOff)
}

func (s *sEssDemo) GetCellMinTemp() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetCellMaxTemp() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetCellAvgTemp() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetCellMinVoltage() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetCellMaxVoltage() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetCellAvgVoltage() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetSoc() (*float32, error) {
	return s.GetFromPointFloat32(SOC)
}

func (s *sEssDemo) GetSoh() (*float32, error) {
	v := float32(100)
	return &v, nil
}

func (s *sEssDemo) GetCapacity() (*uint32, error) {
	return s.GetFromPointUint32(EnergyCapacity)
}

func (s *sEssDemo) GetCycleCount() (*uint, error) {
	return nil, nil
}

func (s *sEssDemo) GetDcPower() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) SetReset() error {
	return s.SetPower(0)
}

func (s *sEssDemo) SetStatus(status c_enum.EEnergyStoreStatus) error {
	if status == c_enum.EPcsStatusOff {
		_ = s.SetPower(0)

		return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
			return protocol.WriteSingleRegister(DeviceControl, 0)
		})
	}
	return nil
}

func (s *sEssDemo) GetFrequency() (*float64, error) {

	return nil, nil
}

func (s *sEssDemo) GetIGBTTemperature() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) SetGridMode(mode c_enum.EGridMode) error {
	return nil
}

func (s *sEssDemo) GetStatus() (*c_enum.EEnergyStoreStatus, error) {
	v, err := s.GetFromProtocol(func(protocol c_proto.IModbusProtocol) (any, error) {
		value, err := protocol.GetUintValue(Status)
		if err != nil {
			return nil, err
		}
		if value == nil {
			return nil, nil // 没有采集到数据
		}

		if v, err := cvt.Uint8E(*value); err == nil {
			switch v {
			case 0:
				return c_enum.EPcsStatusOff, nil
			case 1:
				return c_enum.EPcsStatusStandby, nil
			case 2:
				return c_enum.EPcsStatusCharge, nil
			case 3:
				return c_enum.EPcsStatusDischarge, nil
			case 4:
				return c_enum.EPcsStatusFault, nil
			}
		}
		return c_enum.EPcsStatusUnknown, err
	})
	if v == nil || err != nil {
		return nil, err
	}
	status := v.(c_enum.EEnergyStoreStatus)
	return &status, err
}

func (s *sEssDemo) GetGridMode() (*c_enum.EGridMode, error) {
	mode := c_enum.EGridOn
	return &mode, nil
}

func (s *sEssDemo) SetPower(power int32) error {
	return s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		return protocol.WriteSingleRegister(TargetPower, power)
	})
}

func (s *sEssDemo) SetReactivePower(power int32) error {
	return nil
}

func (s *sEssDemo) SetPowerFactor(factor float32) error {
	return nil
}

func (s *sEssDemo) GetTargetPower() (*int32, error) {
	return s.GetFromPointInt32(TargetPower)
}

func (s *sEssDemo) GetTargetReactivePower() (*int32, error) {
	return nil, nil
}

func (s *sEssDemo) GetTargetPowerFactor() (*float32, error) {
	return nil, nil
}

func (s *sEssDemo) GetPower() (*float64, error) {
	return s.GetFromPointFloat64(Power)
}

func (s *sEssDemo) GetApparentPower() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) GetReactivePower() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) GetRatedPower() (*uint32, error) {
	return s.GetFromPointUint32(PowerCapacity)
}

func (s *sEssDemo) GetMaxInputPower() (*float32, error) {
	return s.GetFromPointFloat32(MaxChargePower)
}

func (s *sEssDemo) GetMaxOutputPower() (*float32, error) {
	return s.GetFromPointFloat32(MaxDischargePower)
}

func (s *sEssDemo) GetTodayIncomingQuantity() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) GetHistoryIncomingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(ConsumedEnergy)
}

func (s *sEssDemo) GetTodayOutgoingQuantity() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) GetHistoryOutgoingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(ConsumedEnergy)
}

func (s *sEssDemo) GetFireEnvTemperature() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) GetCarbonMonoxideConcentration() (*float64, error) {
	return nil, nil
}

func (s *sEssDemo) HasSmoke() (*bool, error) {
	return nil, nil
}

// 实现新的IDevice接口方法
func (s *sEssDemo) GetDevicePoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 遥测点位
		telemetryPowerPoint,
		telemetrySocPoint,
		telemetryGeneratedEnergyPoint,
		telemetryConsumedEnergyPoint,
		// 协议点位
		Status,
		Power,
		SOC,
		GeneratedEnergy,
		ConsumedEnergy,
		MaxChargePower,
		MaxDischargePower,
		DeviceControl,
		TargetPower,
		PowerCapacity,
		EnergyCapacity,
		MinSOC,
		MaxSOC,
		ChargeEfficiency,
	}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (s *sEssDemo) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		Status,
		telemetryPowerPoint, // 功率 - 核心运行参数
		telemetrySocPoint,   // SOC - 电池电量百分比
	}
}

// GetExportModbusPoints 获取暴露出去的modbus点位
func (s *sEssDemo) GetExportModbusPoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 遥测点位 - 基本运行参数
		telemetryPowerPoint,           // 功率
		telemetrySocPoint,             // SOC - 电池电量百分比
		telemetryGeneratedEnergyPoint, // 累计放电量
		telemetryConsumedEnergyPoint,  // 累计用电量
		// 协议点位 - 控制参数
		Status,            // 设备状态
		Power,             // 功率
		SOC,               // SOC
		GeneratedEnergy,   // 发电量
		ConsumedEnergy,    // 用电量
		MaxChargePower,    // 最大充电功率
		MaxDischargePower, // 最大放电功率
		DeviceControl,     // 设备控制
		TargetPower,       // 目标功率
		PowerCapacity,     // 功率容量
		EnergyCapacity,    // 能量容量
		MinSOC,            // 最小SOC
		MaxSOC,            // 最大SOC
		ChargeEfficiency,  // 充电效率
	}
}

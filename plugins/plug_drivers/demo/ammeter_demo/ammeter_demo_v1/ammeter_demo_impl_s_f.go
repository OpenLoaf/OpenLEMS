package ammeter_demo_v1

import (
	"common/c_base"
	"common/c_device"
	"common/c_proto"
)

type sAmmeterDemo struct {
	*c_device.SRealDeviceImpl[c_proto.IModbusProtocol]
}

func (s *sAmmeterDemo) Init() error {
	_ = s.ExecuteProtocolMethod(func(protocol c_proto.IModbusProtocol) error {
		protocol.RegisterTask(ReadTask)
		return nil
	})
	return nil
}

func (s *sAmmeterDemo) Shutdown() {
}

func (s *sAmmeterDemo) GetUa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAVoltage)
}

func (s *sAmmeterDemo) GetUb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBVoltage)
}

func (s *sAmmeterDemo) GetUc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCVoltage)
}

func (s *sAmmeterDemo) GetIa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseACurrent)
}

func (s *sAmmeterDemo) GetIb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBCurrent)
}

func (s *sAmmeterDemo) GetIc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCCurrent)
}

func (s *sAmmeterDemo) GetPa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAActivePower)
}

func (s *sAmmeterDemo) GetPb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBActivePower)
}

func (s *sAmmeterDemo) GetPc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCActivePower)
}

func (s *sAmmeterDemo) GetPTotal() (*float32, error) {
	return s.GetFromPointFloat32(TotalActivePower)
}

func (s *sAmmeterDemo) GetQa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAReactivePower)
}

func (s *sAmmeterDemo) GetQb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBReactivePower)
}

func (s *sAmmeterDemo) GetQc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCReactivePower)
}

func (s *sAmmeterDemo) GetQTotal() (*float32, error) {
	return s.GetFromPointFloat32(TotalReactivePower)
}

func (s *sAmmeterDemo) GetSa() (*float32, error) {
	return s.GetFromPointFloat32(PhaseAApparentPower)
}

func (s *sAmmeterDemo) GetSb() (*float32, error) {
	return s.GetFromPointFloat32(PhaseBApparentPower)
}

func (s *sAmmeterDemo) GetSc() (*float32, error) {
	return s.GetFromPointFloat32(PhaseCApparentPower)
}

func (s *sAmmeterDemo) GetSTotal() (*float32, error) {
	return s.GetFromPointFloat32(TotalApparentPower)
}

func (s *sAmmeterDemo) GetPfa() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetPfb() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetPfc() (*float32, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetPfTotal() (*float32, error) {
	return s.GetFromPointFloat32(PowerFactor)
}

func (s *sAmmeterDemo) GetPtCt() (*float32, *float32, error) {
	return nil, nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetFrequency() (*float32, error) {
	return s.GetFromPointFloat32(Frequency)
}

func (s *sAmmeterDemo) GetTodayIncomingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetHistoryIncomingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(ForwardActiveEnergy)
}

func (s *sAmmeterDemo) GetTodayOutgoingQuantity() (*float64, error) {
	return nil, c_base.NotSupport
}

func (s *sAmmeterDemo) GetHistoryOutgoingQuantity() (*float64, error) {
	return s.GetFromPointFloat64(ReverseActiveEnergy)
}

// 实现新的IDevice接口方法
func (s *sAmmeterDemo) GetDevicePoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 协议点位
		Status,
		PhaseAVoltage,
		PhaseBVoltage,
		PhaseCVoltage,
		PhaseACurrent,
		PhaseBCurrent,
		PhaseCCurrent,
		PhaseAActivePower,
		PhaseBActivePower,
		PhaseCActivePower,
		PhaseAReactivePower,
		PhaseBReactivePower,
		PhaseCReactivePower,
		PhaseAApparentPower,
		PhaseBApparentPower,
		PhaseCApparentPower,
		TotalActivePower,
		TotalReactivePower,
		TotalApparentPower,
		ForwardActiveEnergy,
		ReverseActiveEnergy,
		Frequency,
		PowerFactor,
		RatedLineVoltage,
		RatedFrequency,
	}
}

// GetTelemetryPoints 获取主要遥测点位列表（只返回关键点位）
func (s *sAmmeterDemo) GetTelemetryPoints() []c_base.IPoint {
	return []c_base.IPoint{
		TotalActivePower,    // 总有功功率 - 功率测量
		TotalReactivePower,  // 总无功功率 - 功率测量
		TotalApparentPower,  // 总视在功率 - 功率测量
		Frequency,           // 频率 - 电网参数
		PowerFactor,         // 功率因数 - 电能质量
		ForwardActiveEnergy, // 正向有功电量 - 电量统计
		ReverseActiveEnergy, // 反向有功电量 - 电量统计
	}
}

// GetExportModbusPoints 获取暴露出去的modbus点位
func (s *sAmmeterDemo) GetExportModbusPoints() []c_base.IPoint {
	return []c_base.IPoint{
		// 协议点位 - 详细测量数据
		Status,              // 设备状态
		PhaseAVoltage,       // A相电压
		PhaseBVoltage,       // B相电压
		PhaseCVoltage,       // C相电压
		PhaseACurrent,       // A相电流
		PhaseBCurrent,       // B相电流
		PhaseCCurrent,       // C相电流
		TotalActivePower,    // 总有功功率
		TotalReactivePower,  // 总无功功率
		TotalApparentPower,  // 总视在功率
		ForwardActiveEnergy, // 正向有功电量
		ReverseActiveEnergy, // 反向有功电量
		Frequency,           // 频率
		PowerFactor,         // 功率因数
		RatedLineVoltage,    // 额定线电压
		RatedFrequency,      // 额定频率
	}
}

package c_device

import "common/c_base"

type IPvBase interface {
	GetGridFrequency() (float32, error) // 电网频率
	GetUa() (float32, error)            // A相电压
	GetUb() (float32, error)            // B相电压
	GetUc() (float32, error)            // C相电压
	GetIa() (float32, error)            // A相电流
	GetIb() (float32, error)            // B相电流
	GetIc() (float32, error)            // C相电流
	GetPa() (float32, error)            // A相有功功率
	GetPb() (float32, error)            // B相有功功率
	GetPc() (float32, error)            // C相有功功率

	SetPower(power float64) error         // 设置有功功率
	SetReactivePower(power float64) error // 设置无功功率
	SetPowerFactor(factor float32) error  // 设置功率因数
	GetTargetPower() float64              // 获取目标有功功率
	GetTargetReactivePower() float64      // 获取目标无功功率
	GetTargetPowerFactor() float32        // 获取目标功率因数

	GetPower() (float64, error)         // 有功功率
	GetApparentPower() (float64, error) // 视在功率
	GetReactivePower() (float64, error) // 无功功率

	GetDcPower() (float64, error)   // 直流功率
	GetDcVoltage() (float64, error) // 直流电压
	GetDcCurrent() (float64, error) // 直流电流

	GetTodayIncomingQuantity() (float64, error)   // 正向有功, 今日充电量
	GetHistoryIncomingQuantity() (float64, error) // 正向有功, 充电量
	GetTodayOutgoingQuantity() (float64, error)   // 反向有功, 今日放电量
	GetHistoryOutgoingQuantity() (float64, error) // 反向有功, 放电量
}

type IPv interface {
	c_base.IDevice
	IPvBase
}

// 拆开别写这
type SPvInfo struct {
	GridFrequency           float32 `json:"grid_frequency"`            // 电网频率
	Ua                      float32 `json:"ua"`                        // A相电压
	Ub                      float32 `json:"ub"`                        // B相电压
	Uc                      float32 `json:"uc"`                        // C相电压
	Ia                      float32 `json:"ia"`                        // A相电流
	Ib                      float32 `json:"ib"`                        // B相电流
	Ic                      float32 `json:"ic"`                        // C相电流
	Pa                      float32 `json:"pa"`                        // A相有功功率
	Pb                      float32 `json:"pb"`                        // B相有功功率
	Pc                      float32 `json:"pc"`                        // C相有功功率
	TargetPower             float64 `json:"target_power"`              // 目标有功功率
	TargetReactivePower     float64 `json:"target_reactive_power"`     // 目标无功功率
	TargetPowerFactor       float32 `json:"target_power_factor"`       // 目标功率因数
	Power                   float64 `json:"power"`                     // 有功功率
	ApparentPower           float64 `json:"apparent_power"`            // 视在功率
	ReactivePower           float64 `json:"reactive_power"`            // 无功功率
	DcPower                 float64 `json:"dc_power"`                  // 直流功率
	DcVoltage               float64 `json:"dc_voltage"`                // 直流电压
	DcCurrent               float64 `json:"dc_current"`                // 直流电流
	TodayIncomingQuantity   float64 `json:"today_incoming_quantity"`   // 今日充电量
	HistoryIncomingQuantity float64 `json:"history_incoming_quantity"` // 历史充电量
	TodayOutgoingQuantity   float64 `json:"today_outgoing_quantity"`   // 今日放电量
	HistoryOutgoingQuantity float64 `json:"history_outgoing_quantity"` // 历史放电量
}

func GetPvInfo(pv IPv) *SPvInfo {
	if pv == nil {
		return nil
	}
	info := &SPvInfo{}
	info.GridFrequency, _ = pv.GetGridFrequency()
	info.Ua, _ = pv.GetUa()
	info.Ub, _ = pv.GetUb()
	info.Uc, _ = pv.GetUc()
	info.Ia, _ = pv.GetIa()
	info.Ib, _ = pv.GetIb()
	info.Ic, _ = pv.GetIc()
	info.Pa, _ = pv.GetPa()
	info.Pb, _ = pv.GetPb()
	info.Pc, _ = pv.GetPc()
	info.TargetPower = pv.GetTargetPower()
	info.TargetReactivePower = pv.GetTargetReactivePower()
	info.TargetPowerFactor = pv.GetTargetPowerFactor()
	info.Power, _ = pv.GetPower()
	info.ApparentPower, _ = pv.GetApparentPower()
	info.ReactivePower, _ = pv.GetReactivePower()
	info.DcPower, _ = pv.GetDcPower()
	info.DcVoltage, _ = pv.GetDcVoltage()
	info.DcCurrent, _ = pv.GetDcCurrent()
	info.TodayIncomingQuantity, _ = pv.GetTodayIncomingQuantity()
	info.HistoryIncomingQuantity, _ = pv.GetHistoryIncomingQuantity()
	info.TodayOutgoingQuantity, _ = pv.GetTodayOutgoingQuantity()
	info.HistoryOutgoingQuantity, _ = pv.GetHistoryOutgoingQuantity()
	return info
}

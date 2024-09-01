package internal_group

import (
	"common_station/c_station"
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
)

type sGroupPv struct {
	*c_station.SGroupConfigImpl
	functionList []*c_base.SFunction

	ctx         context.Context
	pvs         []c_device.IPv
	rootPv      c_device.IPv
	ammeters    []c_device.IAmmeter
	rootAmmeter c_device.IAmmeter
}

func NewPv(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootPv c_device.IPv, pvs []c_device.IPv) c_station.IStationPv {
	if (rootAmmeter == nil || len(ammeters) == 0) || (rootPv == nil || len(pvs) == 0) {
		panic("创建StationLoad失败！缺少必要电表或者负载设备！")
	}
	instance := &sGroupPv{
		rootAmmeter: rootAmmeter,
		ammeters:    ammeters,
		rootPv:      rootPv,
		pvs:         pvs,
		ctx:         context.WithValue(ctx, "StationType", c_station.EGroupPv),
		functionList: []*c_base.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayOutgoingQuantity", Unit: "kWh", Remark: "当日发电量"},
			{FunctionName: "historyOutgoingQuantity", Unit: "kWh", Remark: "历史发电量"},
		},
	}
	return instance
}

func (s *sGroupPv) AllowControl() bool {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetFunctionList() []*c_base.SFunction {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetGridFrequency() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetGridVoltage() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetGridCurrent() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetGridPower() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) SetPower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) SetReactivePower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetTargetPower() float64 {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetTargetReactivePower() float64 {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupPv) GetChildren() []c_device.IPv {
	//TODO implement me
	panic("implement me")
}

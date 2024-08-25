package internal_group

import (
	"common_group/c_group"
	"context"
	"ems-plan/c_device"
	"ems-plan/c_telemetry"
)

type Pv struct {
	*c_group.SConfig
	functionList []*c_group.SFunction

	ctx         context.Context
	pvs         []c_device.IPv
	rootPv      c_device.IPv
	ammeters    []c_device.IAmmeter
	rootAmmeter c_device.IAmmeter
}

func NewPv(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter, rootPv c_device.IPv, pvs []c_device.IPv) c_group.IGroupPv {
	if (rootAmmeter == nil || len(ammeters) == 0) || (rootPv == nil || len(pvs) == 0) {
		panic("创建StationLoad失败！缺少必要电表或者负载设备！")
	}
	instance := &Pv{
		rootAmmeter: rootAmmeter,
		ammeters:    ammeters,
		rootPv:      rootPv,
		pvs:         pvs,
		ctx:         context.WithValue(ctx, "Group", c_group.EGroupPv),
		functionList: []*c_group.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayOutgoingQuantity", Unit: "kWh", Remark: "当日发电量"},
			{FunctionName: "historyOutgoingQuantity", Unit: "kWh", Remark: "历史发电量"},
		},
	}
	return instance
}

func (p *Pv) AllowControl() bool {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetFunctionList() []*c_group.SFunction {
	return p.functionList
}

func (p *Pv) GetGridFrequency() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetGridVoltage() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetGridCurrent() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetGridPower() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) SetPower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) SetReactivePower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetTargetPower() float64 {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetTargetReactivePower() float64 {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetDcStatisticsQuantity() c_telemetry.IStatisticsQuantity {
	//TODO implement me
	panic("implement me")
}

func (p *Pv) GetChildren() []c_device.IPv {
	//TODO implement me
	panic("implement me")
}

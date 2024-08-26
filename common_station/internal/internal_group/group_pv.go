package internal_group

import (
	"context"
	"ems-plan/c_device"
	"ems-plan/c_telemetry"
)

type sGroupPv struct {
	*c_station.SGroupConfigImpl
	functionList []*c_station.SFunction

	ctx         context.Context
	pvs         []c_base.IPv
	rootPv      c_base.IPv
	ammeters    []c_base.IAmmeter
	rootAmmeter c_base.IAmmeter
}

func NewPv(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter, rootPv c_base.IPv, pvs []c_base.IPv) c_station.IGroupPv {
	if (rootAmmeter == nil || len(ammeters) == 0) || (rootPv == nil || len(pvs) == 0) {
		panic("创建StationLoad失败！缺少必要电表或者负载设备！")
	}
	instance := &sGroupPv{
		rootAmmeter: rootAmmeter,
		ammeters:    ammeters,
		rootPv:      rootPv,
		pvs:         pvs,
		ctx:         context.WithValue(ctx, "Group", c_station.EGroupPv),
		functionList: []*c_station.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayOutgoingQuantity", Unit: "kWh", Remark: "当日发电量"},
			{FunctionName: "historyOutgoingQuantity", Unit: "kWh", Remark: "历史发电量"},
		},
	}
	return instance
}

func (p *sGroupPv) AllowControl() bool {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetFunctionList() []*c_station.SFunction {
	return p.functionList
}

func (p *sGroupPv) GetGridFrequency() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetGridVoltage() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetGridCurrent() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetGridPower() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) SetPower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) SetReactivePower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetTargetPower() float64 {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetTargetReactivePower() float64 {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetDcStatisticsQuantity() c_telemetry.IStatisticsQuantity {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) GetChildren() []c_base.IPv {
	//TODO implement me
	panic("implement me")
}

func (p *sGroupPv) HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

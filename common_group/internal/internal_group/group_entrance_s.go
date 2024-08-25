package internal_group

import (
	"common_group/c_group"
	"context"
	"ems-plan/c_device"
	"ems-plan/c_meta"
)

// sEntrance 场站总入口
type sEntrance struct {
	*c_group.SConfigImpl
	functionList []*c_group.SFunction

	ctx         context.Context
	ammeters    []c_device.IAmmeter // 分电表
	rootAmmeter c_device.IAmmeter   // 总电表
}

func NewEntrance(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter) c_group.IGroupEntrance {
	if rootAmmeter == nil || len(ammeters) == 0 {
		panic("创建StationEntrance失败！缺少必要电表！")
	}
	instance := &sEntrance{
		rootAmmeter: rootAmmeter,
		ammeters:    ammeters,
		ctx:         context.WithValue(ctx, "Group", c_group.EGroupEntrance),
		SConfigImpl: c_group.NewConfig(c_group.EGroupEntrance),
		functionList: []*c_group.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayIncomingQuantity", Unit: "kWh", Remark: "当日下网电量"},
			{FunctionName: "todayOutgoingQuantity", Unit: "kWh", Remark: "当日下网电量"},
			{FunctionName: "historyIncomingQuantity", Unit: "kWh", Remark: "历史下网电量"},
			{FunctionName: "historyOutgoingQuantity", Unit: "kWh", Remark: "历史下网电量"},
		},
	}
	return instance
}

func (e *sEntrance) AllowControl() bool {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetFunctionList() []*c_group.SFunction {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetGridFrequency() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetGridVoltage() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetGridCurrent() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetGridPower() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) SetPower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) SetReactivePower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetTargetPower() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetTargetReactivePower() float64 {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetPowerFactor() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) GetChildren() []c_device.IAmmeter {
	//TODO implement me
	panic("implement me")
}

func (e *sEntrance) HandleAlarm(self c_meta.SAlarmDetail, global c_meta.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

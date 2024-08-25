package internal_group

import (
	"common_group/c_group"
	"context"
	"ems-plan/c_device"
)

// sGenerator 发电机
type sGenerator struct {
	*c_group.SConfigImpl
	functionList []*c_group.SFunction

	ctx          context.Context
	allowControl bool

	rootAmmeter   c_device.IAmmeter // 总电表
	ammeters      []c_device.IAmmeter
	rootGenerator c_device.IGenerator
	generators    []c_device.IGenerator
}

func NewGenerator(ctx context.Context, rootAmmeter c_device.IAmmeter, ammeters []c_device.IAmmeter,
	rootGenerator c_device.IGenerator, generators []c_device.IGenerator) c_group.IGroupGenerator {

	if rootAmmeter == nil || len(ammeters) == 0 {
		panic("创建StationGenerator失败！缺少必要电表！")
	}

	instance := &sGenerator{
		rootAmmeter:  rootAmmeter,
		ammeters:     ammeters,
		ctx:          context.WithValue(ctx, "Group", c_group.EGroupGenerator),
		SConfigImpl:  c_group.NewConfig(c_group.EGroupGenerator),
		allowControl: rootGenerator != nil || len(generators) != 0,
		functionList: []*c_group.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayIncomingQuantity", Unit: "kWh", Remark: "当日用电量"},
			{FunctionName: "todayOutgoingQuantity", Unit: "kWh", Remark: "当日发电量"},
			{FunctionName: "historyIncomingQuantity", Unit: "kWh", Remark: "历史用电量"},
			{FunctionName: "historyOutgoingQuantity", Unit: "kWh", Remark: "历史发电量"},
		},
	}
	return instance
}

func (s *sGenerator) AllowControl() bool {
	return s.allowControl
}

func (s *sGenerator) GetFunctionList() []*c_group.SFunction {
	//TODO implement me
	panic("implement me")
}

func (s *sGenerator) GetChildren() []c_device.IGenerator {
	//TODO implement me
	panic("implement me")
}

func (s *sGenerator) HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

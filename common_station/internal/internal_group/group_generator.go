package internal_group

import (
	"context"
	"ems-plan/c_device"
)

// sGenerator 发电机
type sGenerator struct {
	*c_station.SGroupConfigImpl
	functionList []*c_station.SFunction

	ctx          context.Context
	allowControl bool

	rootAmmeter   c_base.IAmmeter // 总电表
	ammeters      []c_base.IAmmeter
	rootGenerator c_base.IGenerator
	generators    []c_base.IGenerator
}

func NewGenerator(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter,
	rootGenerator c_base.IGenerator, generators []c_base.IGenerator) c_station.IGroupGenerator {

	if rootAmmeter == nil || len(ammeters) == 0 {
		panic("创建StationGenerator失败！缺少必要电表！")
	}

	instance := &sGenerator{
		rootAmmeter:      rootAmmeter,
		ammeters:         ammeters,
		ctx:              context.WithValue(ctx, "Group", c_station.EGroupGenerator),
		SGroupConfigImpl: c_station.NewGroupConfig(c_station.EGroupGenerator),
		allowControl:     rootGenerator != nil || len(generators) != 0,
		functionList: []*c_station.SFunction{
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

func (s *sGenerator) GetFunctionList() []*c_station.SFunction {
	//TODO implement me
	panic("implement me")
}

func (s *sGenerator) GetChildren() []c_base.IGenerator {
	//TODO implement me
	panic("implement me")
}

func (s *sGenerator) HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

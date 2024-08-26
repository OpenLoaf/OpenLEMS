package internal_group

import (
	"context"
	"ems-plan/c_device"
)

type sGroupLoad struct {
	*c_station.SGroupConfigImpl
	functionList []*c_station.SFunction

	ctx         context.Context
	rootLoad    c_base.ILoad
	loads       []c_base.ILoad
	rootAmmeter c_base.IAmmeter   // 负载总电表
	ammeters    []c_base.IAmmeter // 分电表
}

func NewLoad(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter,
	rootLoad c_base.ILoad, loads []c_base.ILoad) c_station.IGroupLoad {
	instance := &sGroupLoad{
		rootAmmeter:      rootAmmeter,
		ammeters:         ammeters,
		rootLoad:         rootLoad,
		loads:            loads,
		ctx:              context.WithValue(ctx, "Group", c_station.EGroupLoad),
		SGroupConfigImpl: c_station.NewGroupConfig(c_station.EGroupLoad),
		functionList: []*c_station.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayIncomingQuantity", Unit: "kWh", Remark: "当日用电量"},
			{FunctionName: "historyIncomingQuantity", Unit: "kWh", Remark: "历史用电量"},
		},
	}

	return instance
}

func (s *sGroupLoad) AllowControl() bool {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetFunctionList() []*c_station.SFunction {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) SetPower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) SetReactivePower(power float64) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetTargetPower() float64 {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetTargetReactivePower() float64 {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetRatedPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) GetChildren() []c_base.ILoad {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupLoad) HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

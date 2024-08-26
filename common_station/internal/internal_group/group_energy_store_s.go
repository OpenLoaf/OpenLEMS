package internal_group

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/c_telemetry"
)

type sGroupEnergyStore struct {
	*c_station.SGroupConfigImpl
	functionList []*c_station.SFunction

	ctx          context.Context
	rootAmmeter  c_base.IAmmeter       // 储能总电表
	ammeters     []c_base.IAmmeter     // 电表列表（单个储能柜电表）
	energyStores []c_base.IEnergyStore // 储能列表

	targetPower         float64 // 目标功率
	targetReactivePower float64 // 目标无功功率
	targetPowerFactor   float32 // 目标功率因数
}

func NewGroupEnergyStore(ctx context.Context, rootAmmeter c_base.IAmmeter, ammeters []c_base.IAmmeter,
	energyStores []c_base.IEnergyStore) c_station.IGroupEnergyStore {
	if len(energyStores) == 0 {
		panic("创建StationEss失败！缺少必要储能柜设备！")
	}
	instance := &sGroupEnergyStore{
		rootAmmeter:      rootAmmeter,
		ammeters:         ammeters,
		energyStores:     energyStores,
		ctx:              context.WithValue(ctx, "Group", c_station.EGroupEnergyStore),
		SGroupConfigImpl: c_station.NewGroupConfig(c_station.EGroupEnergyStore),
		functionList: []*c_station.SFunction{
			{FunctionName: "power", Unit: "kW", Remark: "功率"},
			{FunctionName: "apparentPower", Unit: "kVA", Remark: "视在功率"},
			{FunctionName: "reactivePower", Unit: "kVar", Remark: "无功功率"},
			{FunctionName: "todayIncomingQuantity", FunctionNameI18nOverwrite: "essTodayCharge", Unit: "kWh", Remark: "当日充电量"},
			{FunctionName: "todayOutgoingQuantity", FunctionNameI18nOverwrite: "essTodayDischarge", Unit: "kWh", Remark: "当日放电量"},
			{FunctionName: "historyIncomingQuantity", FunctionNameI18nOverwrite: "essHistoryCharge", Unit: "kWh", Remark: "历史充电量"},
			{FunctionName: "historyOutgoingQuantity", FunctionNameI18nOverwrite: "essHistoryDischarge", Unit: "kWh", Remark: "历史放电量"},
		},
	}

	return instance
}

func (s *sGroupEnergyStore) AllowControl() bool {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetType() c_base.EDeviceType {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetIsMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetParams() map[string]string {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetFunctionList() []*c_station.SFunction {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetRatedPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetDcPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetDcVoltage() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetDcCurrent() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) SetBmsStatus(status c_base.EBmsStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetBmsStatus() (c_base.EBmsStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetSoc() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetSoh() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetCellTemp() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetCellVoltage() (float32, float32, float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetCapacity() (uint16, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetCycleCount() (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetFireEnvTemperature() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetCarbonMonoxideConcentration() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) HasSmoke() (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetDcStatisticsQuantity() c_telemetry.IStatisticsQuantity {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) GetChildren() []c_base.IEnergyStore {
	//TODO implement me
	panic("implement me")
}

func (s *sGroupEnergyStore) HandleAlarm(self c_base.SAlarmDetail, global c_base.SAlarmDetail) error {
	//TODO implement me
	panic("implement me")
}

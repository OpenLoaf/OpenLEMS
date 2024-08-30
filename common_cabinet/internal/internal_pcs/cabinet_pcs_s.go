package internal_pcs

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"github.com/gogf/gf/v2/os/gcache"
	"time"
)

// sCabinetPcs 实现了 c_device.IPcsBasic 接口
type sCabinetPcs struct {
	*c_base.SAlarmHandler
	ctx       context.Context
	cabinetId uint8
	rootPcs   c_device.IPcs
	pcsList   []c_device.IPcs
}

func NewPcs(ctx context.Context, cabinetId uint8, master c_device.IPcs, pcsList []c_device.IPcs) c_device.IPcsBasic {
	if len(pcsList) == 0 && master == nil {
		panic("master and pcsList cannot be empty at the same time")
	}

	/*info := newCabinetInfo(cabinetId, config.Pcs, _pcs,
	"sCabinetPcs", "{#canbintPcsName}",
	map[string]*plugin.PointInfo{
		"targetPower":           {I18nName: "targetPower", Unit: "kW"},
		"power":                 {I18nName: "power", Unit: "kW"},
		"apparentPower":         {I18nName: "apparentPower", Unit: "kVA"},
		"reactivePower":         {I18nName: "reactivePower", Unit: "kVar"},
		"todayIncomingQuantity": {I18nName: "todayIncomingQuantity", Unit: "kWh"},
		"todayOutgoingQuantity": {I18nName: "todayOutgoingQuantity", Unit: "kWh"},
	})*/

	instance := &sCabinetPcs{
		ctx:       context.WithValue(ctx, "DeviceName", "CabinetPcs_"+string(cabinetId)),
		cabinetId: cabinetId,
		//info:      info,
		rootPcs: master,
		pcsList: pcsList,
	}

	//device.RegisterInstance(instance)
	return instance
}

func (s *sCabinetPcs) GetId() string {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetType() c_base.EDeviceType {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetIsMaster() bool {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetName() string {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetDescription() c_base.SDescription {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetCache() *gcache.Cache {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetLastUpdateTime() *time.Time {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) Init(ctx context.Context, client c_base.IProtocol, cfg any) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) IsActivate() bool {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetRatedPower() (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetMaxInputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetMaxOutputPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sCabinetPcs) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

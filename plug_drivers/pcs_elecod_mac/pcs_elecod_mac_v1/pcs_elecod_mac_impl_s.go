package pcs_elecod_mac_v1

import (
	"canbus/p_canbus"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
)

type sPcsElecodBasic struct {
	p_canbus.ICanbusProtocol
	ctx          context.Context
	log          *glog.Logger
	deviceConfig *c_base.SDriverConfig
	*c_base.SDescription
}

func (s *sPcsElecodBasic) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.ICanbusProtocol = protocol.(p_canbus.ICanbusProtocol)

	for _, task := range analogAllTasks {
		s.RegisterRead(task)
	}
	for _, task := range configAllTasks {
		s.RegisterRead(task)
	}

}

func (s *sPcsElecodBasic) Destroy() {
	g.Log().Info(s.ctx, "Destroy")
}

func (s *sPcsElecodBasic) GetDriverType() c_base.EDeviceType {
	return c_base.EDevicePcs
}

func (s *sPcsElecodBasic) SetReset() error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetStatus(status c_base.EEnergyStoreStatus) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetGridMode(mode c_base.EGridMode) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetStatus() (c_base.EEnergyStoreStatus, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetGridMode() (c_base.EGridMode, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetPower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetReactivePower(power int32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) SetPowerFactor(factor float32) error {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTargetPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTargetReactivePower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTargetPowerFactor() float32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetApparentPower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetReactivePower() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetRatedPower() int32 {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetMaxInputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetMaxOutputPower() (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTodayIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetHistoryIncomingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetTodayOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sPcsElecodBasic) GetHistoryOutgoingQuantity() (float64, error) {
	//TODO implement me
	panic("implement me")
}

package c_default

//func NewSDeviceGetPcs[P c_base.IProtocol](ratePower uint32, device *c_device.SRealDeviceImpl[P]) *SDefaultDevicePcs[P] {
//	return &SDefaultDevicePcs[P]{
//		ratePower:       ratePower,
//		SRealDeviceImpl: device,
//	}
//}
//
//var _ c_type.IPcsBasic = (*SDefaultDevicePcs[c_base.IProtocol])(nil)
//
//type SDefaultDevicePcs[P c_base.IProtocol] struct {
//	ratePower                    uint32 // 额定功率
//	*c_device.SRealDeviceImpl[P]        // 真实设备
//}
//
//func (s *SDefaultDevicePcs[P]) GetIGBTTemperature() (*float32, error) {
//	return s.GetFromPointFloat32(VPointIGBTTemp)
//}
//
//func (s *SDefaultDevicePcs[P]) GetRatedPower() (*uint32, error) {
//	return &s.ratePower, nil
//}
//
//func (s *SDefaultDevicePcs[P]) GetMaxInputPower() (*float32, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *SDefaultDevicePcs[P]) GetMaxOutputPower() (*float32, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *SDefaultDevicePcs[P]) GetTargetPower() (*int32, error) {
//	return s.GetFromPointInt32(VPointTargetP)
//}
//
//func (s *SDefaultDevicePcs[P]) GetTargetReactivePower() (*int32, error) {
//	return s.GetFromPointInt32(VPointTargetQ)
//}
//
//func (s *SDefaultDevicePcs[P]) GetTargetPowerFactor() (*float32, error) {
//	return s.GetFromPointFloat32(VPointTargetPF)
//}
//
//func (s *SDefaultDevicePcs[P]) GetPower() (*float64, error) {
//	return s.GetFromPointFloat64(VPointP)
//}
//
//func (s *SDefaultDevicePcs[P]) GetApparentPower() (*float64, error) {
//	return s.GetFromPointFloat64(VPointQ)
//}
//
//func (s *SDefaultDevicePcs[P]) GetReactivePower() (*float64, error) {
//	return s.GetFromPointFloat64(VPointQ)
//}
//
//func (s *SDefaultDevicePcs[P]) GetTodayIncomingQuantity() (*float64, error) {
//	return s.GetFromPointFloat64(VPointTodayCharge)
//}
//
//func (s *SDefaultDevicePcs[P]) GetHistoryIncomingQuantity() (*float64, error) {
//	return s.GetFromPointFloat64(VPointTotalCharge)
//}
//
//func (s *SDefaultDevicePcs[P]) GetTodayOutgoingQuantity() (*float64, error) {
//	return s.GetFromPointFloat64(VPointTodayDischarge)
//}
//
//func (s *SDefaultDevicePcs[P]) GetHistoryOutgoingQuantity() (*float64, error) {
//	return s.GetFromPointFloat64(VPointTotalDischarge)
//}

package c_group

import "ems-plan/c_device"

type IConfig interface {
	c_device.IGetId
	GetCabinetId() uint8 // 0 means station
	GetGroupType() EType
}

func NewConfig(groupType EType) *SConfigImpl {
	if groupType == EGroupNan {
		panic("EGroupNan can't be a station")
	}
	return &SConfigImpl{
		Id:        string(groupType),
		CabinetId: 0,
		Group:     groupType,
	}
}

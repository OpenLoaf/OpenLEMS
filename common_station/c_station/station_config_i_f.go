package c_station

type IStationConfig interface {
	GetId() string
	GetCabinetId() uint8 // 0 means station
	GetGroupType() EType
}

func NewGroupConfig(groupType EType) *SGroupConfigImpl {
	if groupType == EGroupNan {
		panic("EGroupNan can't be a station")
	}
	return &SGroupConfigImpl{
		Id:        string(groupType),
		CabinetId: 0,
		Group:     groupType,
	}
}

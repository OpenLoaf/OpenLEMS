package c_device

type IConfig interface {
	IGetId
	GetCabinetId() uint8
	GetType() EType
	GetGroupType() EGroupType
	GetIsMaster() bool
	GetParams() map[string]string
}

func NewConfig(cabinetId uint8, deviceType EType, groupType EGroupType,
	master bool, params map[string]string) *SConfig {
	if groupType == EGroupNan {
		panic("EGroupNan can't be a station")
	}
	return &SConfig{
		Id:        string(groupType) + "_" + string(cabinetId),
		CabinetId: cabinetId,
		Type:      deviceType,
		Group:     groupType,
		Master:    master,
		Params:    params,
	}
}

package c_station

type SGroupConfigImpl struct {
	Id string `json:"id" dc:"ID"`

	CabinetId uint8 `json:"cabinetId" dc:"柜子ID"`
	Group     EType
}

func (s SGroupConfigImpl) GetId() string {
	return s.Id
}

func (s SGroupConfigImpl) GetCabinetId() uint8 {
	return s.CabinetId
}

func (s SGroupConfigImpl) GetGroupType() EType {
	return s.Group
}

package c_group

type SConfigImpl struct {
	Id        string `json:"id" dc:"ID"`
	CabinetId uint8  `json:"cabinetId" dc:"柜子ID"`
	Group     EType
}

func (s SConfigImpl) GetId() string {
	return s.Id
}

func (s SConfigImpl) GetCabinetId() uint8 {
	return s.CabinetId
}

func (s SConfigImpl) GetGroupType() EType {
	return s.Group
}

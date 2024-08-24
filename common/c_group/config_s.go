package c_group

type SConfig struct {
	Id        string `json:"id" dc:"ID"`
	CabinetId uint8  `json:"cabinetId" dc:"柜子ID"`
	Group     EType
}

func (s SConfig) GetId() string {
	return s.Id
}

func (s SConfig) GetCabinetId() uint8 {
	return s.CabinetId
}

func (s SConfig) GetGroupType() EType {
	return s.Group
}

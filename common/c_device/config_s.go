package c_device

type SConfigImpl struct {
	Id        string            `json:"id"`
	CabinetId uint8             `json:"cabinetId" dc:"柜子ID"`
	Type      EType             `json:"type" dc:"设备类型"`
	Group     EGroupType        `json:"group" dc:"组类型"`
	Master    bool              `json:"master" dc:"主设备"`
	Params    map[string]string `json:"params" dc:"其他参数"`
}

func (s *SConfigImpl) GetId() string {
	return s.Id
}

func (s *SConfigImpl) GetCabinetId() uint8 {
	return s.CabinetId
}

func (s *SConfigImpl) GetType() EType {
	return s.Type
}

func (s *SConfigImpl) GetGroupType() EGroupType {
	return s.Group
}

func (s *SConfigImpl) GetIsMaster() bool {
	return s.Master
}

func (s *SConfigImpl) GetParams() map[string]string {
	return s.Params
}

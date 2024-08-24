package c_device

type SConfig struct {
	Id        string            `json:"id"`
	CabinetId uint8             `json:"cabinetId" dc:"柜子ID"`
	Type      EType             `json:"type" dc:"设备类型"`
	Group     EGroupType        `json:"group" dc:"组类型"`
	Master    bool              `json:"master" dc:"主设备"`
	Params    map[string]string `json:"params" dc:"其他参数"`
}

func (s *SConfig) GetId() string {
	return s.Id
}

func (s *SConfig) GetCabinetId() uint8 {
	return s.CabinetId
}

func (s *SConfig) GetType() EType {
	return s.Type
}

func (s *SConfig) GetGroupType() EGroupType {
	return s.Group
}

func (s *SConfig) GetIsMaster() bool {
	return s.Master
}

func (s *SConfig) GetParams() map[string]string {
	return s.Params
}

package c_base

type IDriverConfig interface {
	GetName() string
	GetIsMaster() bool
	GetCabinetId() uint8
	IsEnable() bool
	GetGroup() EGroupType
}

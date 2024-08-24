package c_device

type IInstances interface {
	RegisterInstance(info IInfo)

	GetInstance(id string) IInfo

	DelInstance(id string)

	GetList() []IInfo
}

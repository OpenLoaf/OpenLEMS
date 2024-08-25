package c_device

type IInstances interface {
	RegisterInstance(info IInfo)

	FindById(id string) IInfo

	FindAll() []IInfo

	FindByType(t EType) []IInfo

	RemoveById(id string)
}

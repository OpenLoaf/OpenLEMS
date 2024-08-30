package c_base

type IDriverConfig interface {
	GetName() string
	GetIsMaster() bool
	GetCabinetId() uint8
	IsEnable() bool
	GetGroup() EGroupType
}

func ConvertConfig[T IDriverConfig](cfg any) T {
	var (
		config T
		ok     bool
	)
	if config, ok = cfg.(T); !ok {
		panic("配置文件转换失败！请检查配置文件！")
	}
	return config
}

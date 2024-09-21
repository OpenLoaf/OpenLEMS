package c_base

type IPush interface {
	// PushDevices 推送设备数据
	PushDevices(deviceId string, deviceType EDeviceType, fields map[string]any) error

	Close()
}

package c_base

type IStorage interface {
	// Save 保存设备数据
	Save(deviceId string, deviceType EDeviceType, fields map[string]any) error

	// SaveProtocolMetrics 保存协议指标数据
	SaveProtocolMetrics(protocolConfig *SProtocolConfig, metrics map[string]any) error

	Close()
	//FindByDeviceId(deviceId string) (map[string]any, error)
}

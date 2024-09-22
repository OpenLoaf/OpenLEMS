// //go:generate mockgen -source=storage_i.go -package=mock_c_base -destination=mock_c_base/storage_i.mock.go
package c_base

type IStorage interface {

	// SaveDevices 保存设备数据
	SaveDevices(deviceId string, deviceType EDeviceType, fields map[string]any) error

	// SaveProtocolMetrics 保存协议指标数据
	SaveProtocolMetrics(protocolConfig *SProtocolConfig, deviceConfig *SDriverConfig, metrics map[string]any) error

	// SaveSystemMetrics 保存系统指标数据
	SaveSystemMetrics(measurement string, tags map[string]string, metrics map[string]any) error

	Close()
	//FindByDeviceId(deviceId string) (map[string]any, error)
}

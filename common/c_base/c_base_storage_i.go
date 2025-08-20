// //go:generate mockgen -source=storage_i.go -package=mock_c_base -destination=mock_c_base/storage_i.mock.go
package c_base

import "common/c_chart"

type StorageType string

const (
	StorageTypeDevice   StorageType = "device"
	StorageTypeProtocol StorageType = "protocol"
	StorageTypeSystem   StorageType = "system"
)

type IStorage interface {

	// SaveDevices 保存设备数据
	SaveDevices(deviceId string, deviceType EDeviceType, fields map[string]any) error

	// SaveProtocolMetrics 保存协议指标数据
	SaveProtocolMetrics(protocolConfig *SProtocolConfig, deviceConfig *SDeviceConfig, metrics map[string]any) error

	// SaveSystemMetrics 保存系统指标数据
	SaveSystemMetrics(measurement string, tags map[string]string, metrics map[string]any) error

	// GetStorageData 获取存储数据
	GetStorageData(storageType StorageType, id string, pointKey []string, startTime, endTime *int, step int) (*c_chart.ChartData, error)

	// Close 关闭数据库
	Close()
}

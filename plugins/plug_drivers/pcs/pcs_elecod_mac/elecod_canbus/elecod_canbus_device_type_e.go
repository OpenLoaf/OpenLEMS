//go:generate stringer -type=DeviceType  -linecomment -trimprefix=DeviceType -output=elecod_canbus_device_type_e_string.go
package elecod_canbus

// 设备类型枚举
type DeviceType uint32

const (
	DeviceTypeBroadcast DeviceType = 0b000 // 广播地址
	DeviceTypeMAC       DeviceType = 0b001 // MAC
	DeviceTypeMDC       DeviceType = 0b010 // MDC
	DeviceTypeSTS       DeviceType = 0b011 // STS
	DeviceTypeScreen    DeviceType = 0b111 // 屏
)

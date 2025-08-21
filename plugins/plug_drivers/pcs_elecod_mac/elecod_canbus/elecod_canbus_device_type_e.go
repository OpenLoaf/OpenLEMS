package elecod_canbus

import "fmt"

// 设备类型枚举
type DeviceType uint32

const (
	DeviceTypeBroadcast DeviceType = 0b000 // 广播地址
	DeviceTypeMAC       DeviceType = 0b001 // MAC
	DeviceTypeMDC       DeviceType = 0b010 // MDC
	DeviceTypeSTS       DeviceType = 0b011 // STS
	DeviceTypeScreen    DeviceType = 0b111 // 屏
)

// 设备类型名称映射
var deviceTypeNames = map[DeviceType]string{
	DeviceTypeBroadcast: "广播地址",
	DeviceTypeMAC:       "MAC",
	DeviceTypeMDC:       "MDC",
	DeviceTypeSTS:       "STS",
	DeviceTypeScreen:    "屏",
}

func (d DeviceType) String() string {
	if name, exists := deviceTypeNames[d]; exists {
		return name
	}
	return fmt.Sprintf("未知类型(0x%X)", uint32(d))
}

// CANbus帧信息结构体
type CANFrameInfo struct {
	TargetDeviceType DeviceType
	TargetDeviceAddr uint32
	SourceDeviceType DeviceType
	SourceDeviceAddr uint32
	MessageType      MessageType
	ServiceCode      uint32
	Reserved         uint32
}

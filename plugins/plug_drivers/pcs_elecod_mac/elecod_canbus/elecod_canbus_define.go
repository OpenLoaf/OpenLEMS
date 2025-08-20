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

// 信息类型枚举
type MessageType uint32

const (
	MessageTypeConfig  MessageType = 0x1 // 配置信息
	MessageTypeControl MessageType = 0x2 // 控制信息
	MessageTypeAlarm   MessageType = 0x3 // 告警信息
	MessageTypeStatus  MessageType = 0x4 // 状态信息
	MessageTypeAnalog  MessageType = 0x5 // 模拟量信息
)

// 信息类型名称映射
var messageTypeNames = map[MessageType]string{
	MessageTypeConfig:  "配置信息",
	MessageTypeControl: "控制信息",
	MessageTypeAlarm:   "告警信息",
	MessageTypeStatus:  "状态信息",
	MessageTypeAnalog:  "模拟量信息",
}

func (m MessageType) String() string {
	if name, exists := messageTypeNames[m]; exists {
		return name
	}
	return fmt.Sprintf("未知信息类型(0x%X)", uint32(m))
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

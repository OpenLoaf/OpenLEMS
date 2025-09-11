//go:generate stringer -type=MessageType  -linecomment -trimprefix=MessageType -output=elecod_canbus_message_type_e_string.go
package elecod_canbus

// 信息类型枚举
type MessageType uint32

const (
	MessageTypeConfig  MessageType = 0x1 // 配置信息
	MessageTypeControl MessageType = 0x2 // 控制信息
	MessageTypeAlarm   MessageType = 0x3 // 告警信息
	MessageTypeStatus  MessageType = 0x4 // 状态信息
	MessageTypeAnalog  MessageType = 0x5 // 模拟量信息
)

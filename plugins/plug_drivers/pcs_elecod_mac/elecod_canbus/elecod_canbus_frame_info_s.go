package elecod_canbus

// CANbus帧信息结构体
type SCANFrameInfo struct {
	TargetDeviceType DeviceType
	TargetDeviceAddr uint32
	SourceDeviceType DeviceType
	SourceDeviceAddr uint32
	MessageType      MessageType
	ServiceCode      uint32
	Reserved         uint32
}

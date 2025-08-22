package c_base

type IProtocol interface {
	IAlarm             // 告警
	IDeviceRuntimeInfo // 设备运行信息

	ProtocolListen()  // 协议监听
	IsActivate() bool // 是否有效，无效一般是连接断了

}

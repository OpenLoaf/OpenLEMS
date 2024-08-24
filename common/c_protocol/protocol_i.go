package c_protocol

type IProtocol interface {
	GetType() EType // 获取协议类型

	Start() error     // 打开
	Close() error     // 关闭
	IsActivate() bool // 是否有效，无效一般是连接断了

	//Listen(notifyFunction func(notify *Notify)) // 启动监听
	//GetAlarmProvider() alarm.IProvider // 获取告警Provider

}

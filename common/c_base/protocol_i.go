package c_base

import (
	"time"
)

type IProtocol interface {
	IAlarm
	Init()            // 初始化
	Close()           //关闭
	IsActivate() bool // 是否有效，无效一般是连接断了

	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetLastUpdateTime() *time.Time         // 获取最后更新时间

	GetDeviceConfig() *SDriverConfig     // 获取设备配置
	GetProtocolConfig() *SProtocolConfig // 获取协议配置

}

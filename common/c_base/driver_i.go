package c_base

import (
	"time"
)

type IDriver interface {
	IAlarm
	Init(protocol IProtocol, deviceConfig *SDriverConfig)

	GetDriverType() EDeviceType            // 获取实现驱动的设备类型
	GetLastUpdateTime() *time.Time         // 获取最后更新时间
	GetMetaValueList() []*MetaValueWrapper // 获取元数据列表
	GetFunctionList() []*SFunction         // 获取功能列表
	GetDeviceConfig() *SDriverConfig       // 获取设备配置
	GetDescription() *SDescription         // 获取设备描述
}

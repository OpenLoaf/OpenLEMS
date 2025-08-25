package c_base

import "time"

type IDevice interface {
	IAlarm
	GetConfig() *SDeviceConfig

	GetMetaValueList() []*MetaValueWrapper // 获取所有缓存的数据列表
	GetLastUpdateTime() *time.Time         // 获取最后更新时间

}

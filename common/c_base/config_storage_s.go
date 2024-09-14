package c_base

type SStorageConfig struct {
	Enable      bool // 是否启用
	Type        EStorageType
	Url         string
	Database    string
	Username    string
	Password    string
	IntervalSec int32 // 多少秒保存一次，0不保存
	KeepDays    int32 // 保存多少天， 0或小于0 代表永久保存

	Params map[string]string
}

package c_base

type SStorageConfig struct {
	Enable                      bool         // 是否启用
	Type                        EStorageType // 类型
	Url                         string       // 地址
	Database                    string       // 数据库
	Username                    string       // 用户名
	Password                    string       // 密码
	ProtocolMetricsSurvivalDays int32        // 数据保存天数,0代表永久保存,-1代表不保存
	SystemMetricsSurvivalDays   int32        // 数据保存天数,0代表永久保存,-1代表不保存
	Params                      map[string]string
}

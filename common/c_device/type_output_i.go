package c_device

type EOutputType = string

// 定义基础的输出设备类型
const (
	RunningOutput EOutputType = "RunningOutput" // 运行输出
	WarningOutput EOutputType = "WarningOutput" // 告警输出
)

type IOutput interface {
	IInfo
	SetUp() (bool, error)
	SetDown() (bool, error)
}

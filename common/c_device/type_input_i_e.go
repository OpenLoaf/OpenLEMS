package c_device

type EInputType = string

// 定义基础的输入设备类型
const (
	EScramButton     EInputType = "ScramButton"     // 急停按钮
	EChargeButton    EInputType = "ChargeButton"    // 充电按钮
	EDischargeButton EInputType = "DischargeButton" // 放电按钮
)

type IInput interface {
	IInfo
	IsUp() (bool, error)
	IsDown() (bool, error)
}

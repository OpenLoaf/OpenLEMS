package c_telemetry

type IPowerLimit interface {
	GetRatedPower() (float64, error)     // 额定功率， -1代表未知
	GetMaxInputPower() (float64, error)  // 最大充电功率、最大输入功率限制
	GetMaxOutputPower() (float64, error) // 最大放电功率、最大输出功率限制
}

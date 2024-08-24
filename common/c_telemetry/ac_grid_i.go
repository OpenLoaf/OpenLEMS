package c_telemetry

type IAcGrid interface {
	GetGridFrequency() (float32, error)                 // 电网频率
	GetGridVoltage() (float32, float32, float32, error) // 电网电压 A、B、C
	GetGridCurrent() (float32, float32, float32, error) // 电网电流 A、B、C
	GetGridPower() (float32, float32, float32, error)   // 电网功率 A、B、C

}

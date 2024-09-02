package c_base

type SDescription struct {
	Brand     string        `json:"brand"`     // 品牌
	Model     string        `json:"model"`     // 型号
	Remark    string        `json:"remark"`    // 备注
	Telemetry []*STelemetry `json:"telemetry"` // 遥测
}

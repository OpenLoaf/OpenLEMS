package c_device

type SDescription struct {
	Brand  string `json:"brand"`  // 品牌
	Model  string `json:"model"`  // 型号
	Type   EType  `json:"type"`   // 设备类型
	Remark string `json:"remark"` // 备注
}

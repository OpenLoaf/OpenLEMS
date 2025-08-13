package c_base

// DriverInfo 驱动信息结构
type DriverInfo struct {
	Name        string        `json:"name"`        // 驱动名称
	Type        EDeviceType   `json:"type"`        // 驱动类型
	Description *SDescription `json:"description"` // 驱动描述
	Available   bool          `json:"available"`   // 是否可用
}

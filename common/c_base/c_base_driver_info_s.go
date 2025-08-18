package c_base

// SDriverInfo 驱动信息结构
type SDriverInfo struct {
	Name         string              `json:"name"`         // 驱动名称
	Type         EDeviceType         `json:"type"`         // 驱动类型
	Description  *SDriverDescription `json:"description"`  // 驱动描述
	Available    bool                `json:"available"`    // 是否可用
	Embedded     bool                `json:"embedded"`     // 是否是内嵌
	Path         string              `json:"path"`         // 路径
	HashCode     string              `json:"hashCode"`     // 哈希码
	FileSizeByte int64               `json:"fileSizeByte"` // 文件大小
}

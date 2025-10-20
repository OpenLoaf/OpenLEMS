package c_enum

// EPriceType 电价类型枚举
type EPriceType string

const (
	EPriceTypeValley     EPriceType = "valley"      // 谷电
	EPriceTypePeak       EPriceType = "peak"        // 峰电
	EPriceTypeFlat       EPriceType = "flat"        // 平电
	EPriceTypeSharp      EPriceType = "sharp"       // 尖峰
	EPriceTypeDeepValley EPriceType = "deep_valley" // 深谷
)

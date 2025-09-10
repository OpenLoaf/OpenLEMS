package c_enum

type EPolicyMode int

const (
	EPolicyModeAuto     = iota // // 全自动模式
	EPolicyModeSemiAuto        // 半自动模式
	EPolicyModeManual          // 纯手动模式
)

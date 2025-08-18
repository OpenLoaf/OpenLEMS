//go:generate stringer -type=EGridMode  -output=energy_store_grid_e_string.go
package c_base

type EGridMode int // 电网状态

const (
	EGridUnknown EGridMode = iota // 未知状态
	EGridOn                       // 并网
	EGridOff                      // 离网
	EGridSync                     // 同步中
)

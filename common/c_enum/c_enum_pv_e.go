//go:generate stringer -type=EPvStatus -trimprefix=EPv -output=c_enum_pv_e_string.go
package c_enum

type EPvStatus int // 连接状态
const (
	EPvUnknown    EPvStatus = iota //未知
	EPvOff                         // 关机
	EPvStandby                     // 待机
	EPvProcessing                  // 启动or关闭中，中间状态
	EPvGenerate                    // 发电中
	EPvFailed                      // 故障

)

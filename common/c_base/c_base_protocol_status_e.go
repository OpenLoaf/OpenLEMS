//go:generate stringer -type=EProtocolStatus -trimprefix=EProtocol -output=c_base_protocol_status_e_string.go
package c_base

type EProtocolStatus int // 连接状态
const (
	EProtocolDisconnected EProtocolStatus = iota // 连接断开
	EProtocolConnecting                          // 正在连接中
	EProtocolConnected                           // 连接成功
)

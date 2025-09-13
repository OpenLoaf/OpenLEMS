//go:generate stringer -type=EProtocolStatus -trimprefix=EProtocol -output=c_enum_protocol_e_string.go
package c_enum

type EProtocolStatus int // 连接状态
const (
	EProtocolDisconnected EProtocolStatus = iota // 连接断开
	EProtocolConnecting                          // 正在连接中
	EProtocolConnected                           // 连接成功
	EProtocolMock                                // 模拟
)

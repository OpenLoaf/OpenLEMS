package c_enum

type ESystemGroupType string

const (
	ELogTypeEms        ESystemGroupType = "ems"
	ELogTypeDevice     ESystemGroupType = "device"
	ELogTypeProtocol   ESystemGroupType = "protocol"
	ELogTypePolicy     ESystemGroupType = "policy"
	ELogTypeAutomation ESystemGroupType = "automation" // 自动化
	ELogTypeRemote     ESystemGroupType = "remote"     // 远程协议

)

func (e ESystemGroupType) String() string {
	return string(e)
}

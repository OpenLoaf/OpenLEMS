package c_enum

type ELogType string

const (
	ELogTypeEms        ELogType = "Ems"
	ELogTypeDevice     ELogType = "Device"
	ELogTypeProtocol   ELogType = "Protocol"
	ELogTypePolicy     ELogType = "Policy"
	ELogTypeAutomation ELogType = "Automation" // 自动化

)

func (e ELogType) String() string {
	return string(e)
}

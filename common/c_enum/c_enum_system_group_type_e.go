package c_enum

type ESystemGroupType string

const (
	ELogTypeEms        ESystemGroupType = "Ems"
	ELogTypeDevice     ESystemGroupType = "Device"
	ELogTypeProtocol   ESystemGroupType = "Protocol"
	ELogTypePolicy     ESystemGroupType = "Policy"
	ELogTypeAutomation ESystemGroupType = "Automation" // 自动化

)

func (e ESystemGroupType) String() string {
	return string(e)
}

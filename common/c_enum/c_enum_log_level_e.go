package c_enum

type ELogLevel string

const (
	Debug ELogLevel = "DEBUG"
	Info  ELogLevel = "INFO"
	Warn  ELogLevel = "WARN"
	Error ELogLevel = "ERROR"
)

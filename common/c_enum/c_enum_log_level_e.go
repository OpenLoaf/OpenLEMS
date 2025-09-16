package c_enum

type ELogLevel string

const (
	Debug ELogLevel = "debug"
	Info  ELogLevel = "info"
	Warn  ELogLevel = "warn"
	Error ELogLevel = "error"
)

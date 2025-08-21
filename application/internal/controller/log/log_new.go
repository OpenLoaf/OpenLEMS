package log

import (
	apilog "application/api/log"
)

type ControllerV1 struct{}

func NewV1() apilog.ILogV1 { return &ControllerV1{} }

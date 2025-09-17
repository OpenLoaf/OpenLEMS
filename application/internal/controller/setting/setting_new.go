package setting

import (
	"application/api/setting"
)

type ControllerV1 struct{}

func NewV1() setting.ISettingV1 {
	return &ControllerV1{}
}

package c_none

import (
	"common/c_base"
	"time"
)

type sNoneDeviceRuntimeInfo struct{}

func (s *sNoneDeviceRuntimeInfo) IsPhysical() bool {
	return false
}

func (s *sNoneDeviceRuntimeInfo) GetDeviceConfig() *c_base.SDeviceConfig {
	return nil
}

func (s *sNoneDeviceRuntimeInfo) GetMetaValueList() []*c_base.MetaValueWrapper {
	return nil
}

func (s *sNoneDeviceRuntimeInfo) GetLastUpdateTime() *time.Time {
	return nil
}

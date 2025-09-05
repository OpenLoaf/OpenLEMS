package p_policy_mircogrid

import (
	"common/c_base"
	"p_policy_mircogrid/internal"
)

func GetMicrogridConfigFields() []*c_base.SConfigStructFields {
	f, err := c_base.BuildConfigStructFields(&internal.SPolicyMircogridConfig{})
	if err != nil {
		return nil
	}
	return f
}

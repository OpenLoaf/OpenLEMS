package c_proto

import (
	"common/c_base"
	"fmt"
	"testing"
)

func TestConfigFields(t *testing.T) {
	config := &SModbusDeviceConfig{}
	fs, err := c_base.BuildConfigStructFields(config)
	if err != nil {
		t.Errorf("Build config struct fail")
	}
	for _, field := range fs {
		fmt.Println(field)
	}
}

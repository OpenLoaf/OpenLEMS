package cmd

import (
	"context"
	"ems-plan/c_base"
	"fmt"
	"testing"
)

func TestStationConfig(t *testing.T) {
	ctx := context.TODO()
	linkConfigList, configPath, err := c_base.GetConfigList[c_base.SDriverConfig](ctx, "", c_base.StationsKey)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(linkConfigList)
	fmt.Println(configPath)
}

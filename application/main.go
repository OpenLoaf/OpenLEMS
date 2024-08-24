package main

import (
	"ems-plan/c_group"
	"fmt"
)

type TestConfig struct {
	c_group.SConfig
}

func main() {
	config := c_group.NewConfig(c_group.EGroupPv)

	fmt.Printf("config: %+v", config)

}

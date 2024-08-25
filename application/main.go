package main

import (
	"common_group/c_group"
	"fmt"
	"runtime"
)

type TestConfig struct {
	c_group.SConfig
}

func main() {
	config := c_group.NewConfig(c_group.EGroupPv)

	fmt.Printf("config: %+v", config)

	fmt.Println(runtime.GOMAXPROCS(0))

	abc := func(say string) {
		fmt.Println(say)
	}

	abc("sdf")

	cache := map[string]func(string2 string){}
	cache["a"] = func(string2 string) {
		fmt.Println(string2)
	}

}

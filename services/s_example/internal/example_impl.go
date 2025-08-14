package internal

import (
	"fmt"
	"s_example/example_interface"
)

type Hello struct {
}

func NewHay() example_interface.IExample {
	return &Hello{}
}

func (h *Hello) Hay() {
	fmt.Println("Hello")
}

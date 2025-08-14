package impl

import "fmt"

type Hello struct {
}

func NewHay() Hello {
	return Hello{}
}

func (h *Hello) Hay() {
	fmt.Println("Hello")
}

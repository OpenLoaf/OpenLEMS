package s_example

import (
	"s_example/example_interface"
	"s_example/internal"
)

func NewExample() example_interface.IExample {
	return internal.NewHay()
}

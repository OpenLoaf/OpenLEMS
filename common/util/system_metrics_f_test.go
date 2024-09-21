package util

import (
	"fmt"
	"testing"
)

func TestGetSystemMetrics(t *testing.T) {
	metrics := GetSystemMetrics()
	for key, value := range metrics {
		fmt.Printf("%s: %v\n", key, value)
	}
}

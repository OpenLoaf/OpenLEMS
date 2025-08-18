package internal

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

func TestGetProcessInfo(t *testing.T) {
	processInfo := GetProcessInfo()
	for key, value := range processInfo {
		fmt.Printf("%s: %v\n", key, value)
	}
}

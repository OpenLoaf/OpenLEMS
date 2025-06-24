package service

import (
	"context"
	"testing"
)

func TestNewConfigManage(t *testing.T) {

	NewConfigManage(context.Background(), 1).GetDeviceConfig(context.Background())
}

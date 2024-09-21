package internal

import (
	"common/c_base"
	"context"
	"testing"
)

func TestNewInfluxdb2(t *testing.T) {

	NewInfluxdb2(context.Background(), &c_base.SStorageConfig{
		Enable:                    true,
		Type:                      "influxdb2",
		Url:                       "http://localhost:8086",
		SystemMetricsSurvivalDays: 30,
		Params: map[string]string{
			"Token": "vdfoWUugqfivUqnxD2rECypyHUQCHL0taOZfNkH4iS25wxUkwqHhmb8AyszIMhcZJ2oOl2NpLdPJr7GpH2tYLw==",
			"Org":   "hexsolar",
		},
	})

}

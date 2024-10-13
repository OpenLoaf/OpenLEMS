package internal_config_db

import (
	"common/c_base"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"testing"
)

func setDefaultDbConfig(dbName string) {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Link:   fmt.Sprintf("sqlite::@file(../../../tmp/%s.sqlite3)", dbName),
				Prefix: "hex_",
				Debug:  true,
			},
		},
	})
}

func TestGetDriverConfig(t *testing.T) {
	setDefaultDbConfig("service_sqlite_test")

	config := GetDriverConfig()
	deepPrintDeviceConfig(config)
}

func deepPrintDeviceConfig(config *c_base.SDriverConfig) {
	if config == nil {
		fmt.Println("获取到的config为空！")
		return
	}
	fmt.Printf("Id: %s, Name: %s \n", config.Id, config.Name)
	if len(config.DeviceChildren) == 0 {
		return
	}
	for _, child := range config.DeviceChildren {
		deepPrintDeviceConfig(child)
	}
}

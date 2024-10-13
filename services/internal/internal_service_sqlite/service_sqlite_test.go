package internal_service_sqlite

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func setDefaultDbConfig(dbName string) {
	gdb.SetConfig(gdb.Config{
		"default": gdb.ConfigGroup{
			gdb.ConfigNode{
				Link:  fmt.Sprintf("sqlite::@file(../../../tmp/%s.sqlite3)", dbName),
				Debug: true,
			},
		},
	})
}

func TestCreateDB(t *testing.T) {
	ctx := context.Background()

	setDefaultDbConfig("service_sqlite_test")

	// 创建用户表
	_, err := g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL,
			email TEXT NOT NULL
		)
	`)
	if err != nil {
		t.Errorf("create table failed: %v", err)
	}
}

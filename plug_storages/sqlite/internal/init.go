package internal

import (
	"os"
	"path/filepath"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func Init() {
	ctx := gctx.New()

	// 确保数据库目录存在
	dbPath := "out/db.sqlite3"
	dbDir := filepath.Dir(dbPath)
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			g.Log().Fatalf(ctx, "创建数据库目录失败: %v", err)
		}
		g.Log().Infof(ctx, "创建数据库目录: %s", dbDir)
	}

	// 测试数据库连接
	if _, err := g.DB().Exec(ctx, "SELECT 1"); err != nil {
		g.Log().Fatalf(ctx, "数据库连接失败: %v", err)
	}

	// 创建协议表
	_, err := g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS protocol (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			address TEXT NOT NULL,
			timeout INTEGER NOT NULL,
			log_level TEXT NOT NULL,
			params TEXT
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建分组表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS "group" (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			version TEXT NOT NULL
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建设备表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS device (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			pid INTEGER NOT NULL,
			gid INTEGER NOT NULL,
			protocol_id INTEGER,
			name TEXT NOT NULL,
			driver TEXT,
			log_level TEXT,
			enable BOOLEAN,
			params TEXT
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Info(ctx, "Tables created successfully")
}

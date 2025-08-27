package internal

import (
	"os"
	"path/filepath"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var DefaultSqliteDbPath = gcmd.GetOpt("db-path", "./out/db.sqlite3").String()

func Init() {
	initConfigDatabase()
}

// 初始化配置数据库
func initConfigDatabase() {
	ctx := gctx.New()

	// 确保数据库目录存在
	dbPath := DefaultSqliteDbPath
	dbDir := filepath.Dir(dbPath)
	if _, err := os.Stat(dbDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dbDir, 0755); err != nil {
			g.Log().Fatalf(ctx, "创建数据库目录失败: %+v", err)
		}
		g.Log().Infof(ctx, "创建数据库目录: %s", dbDir)
	}

	// 测试数据库连接（如果文件不存在，SQLite会自动创建）
	if _, err := g.DB().Exec(ctx, "SELECT 1"); err != nil {
		g.Log().Fatalf(ctx, "数据库连接失败: %+v", err)
	}

	// 创建协议表
	_, err := g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS protocol (
			id VARCHAR(255) PRIMARY KEY ,
			name VARCHAR(255) NOT NULL,
			type VARCHAR(255) NOT NULL,
			address VARCHAR(255) NOT NULL,
			timeout INTEGER NOT NULL,
			log_level VARCHAR(255) NOT NULL,
			sort INTEGER DEFAULT 0,
			params TEXT,
		    enable BOOLEAN DEFAULT 1,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建设备表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS device (
			id VARCHAR(255) PRIMARY KEY ,
			pid VARCHAR(255) NOT NULL DEFAULT '0',
			protocol_id VARCHAR(255),
			name VARCHAR(255) NOT NULL,
			driver VARCHAR(255),
			policy VARCHAR(255),
			policy_params VARCHAR(255),
			log_level VARCHAR(255),
			enabled BOOLEAN DEFAULT 1,
			params TEXT,
		    strategy VARCHAR(255),
		    storage_enable BOOLEAN DEFAULT 1,
			storage_interval_sec INTEGER NOT NULL DEFAULT 60,
			sort INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建设置表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS setting (
			id VARCHAR(255) PRIMARY KEY,
			value VARCHAR(255),
		    is_public BOOLEAN DEFAULT 0,
			enabled BOOLEAN default 1,
			remark VARCHAR(255),
			sort INTEGER DEFAULT 0,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Info(ctx, "Config tables created successfully")
}

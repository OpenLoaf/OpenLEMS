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
			timeout INTEGER NOT NULL DEFAULT 30,
			log_level VARCHAR(255) NOT NULL DEFAULT 'INFO',
			sort INTEGER DEFAULT 0,
			params TEXT DEFAULT '{}',
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
			log_level VARCHAR(255) DEFAULT 'INFO',
			enabled BOOLEAN DEFAULT 1,
			params TEXT DEFAULT '{}',
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
			value TEXT,
		    is_public BOOLEAN DEFAULT 0,
			enabled BOOLEAN default 1,
			remark VARCHAR(255),
			sort INTEGER DEFAULT 0,
			group_name VARCHAR(255) DEFAULT 'system',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建告警历史表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS alarm_history (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			device_id VARCHAR(255) NOT NULL,
			source_device_id VARCHAR(255) NOT NULL,
			point VARCHAR(255) NOT NULL,
			point_name VARCHAR(255) NOT NULL,
			level VARCHAR(255) NOT NULL ,
			detail TEXT,
			trigger_at DATETIME NOT NULL,
			clear_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建告警忽略表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS alarm_ignore (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			device_id VARCHAR(255) NOT NULL,
			source_device_id VARCHAR(255) NOT NULL,
			point VARCHAR(255) NOT NULL,
			point_name VARCHAR(255) NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建日志表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS log (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			type VARCHAR(255) NOT NULL,
			device_id VARCHAR(255),
			level VARCHAR(255) NOT NULL,
			content TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Info(ctx, "Config tables created successfully")
}

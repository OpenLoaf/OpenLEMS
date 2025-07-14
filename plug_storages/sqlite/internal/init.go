package internal

import (
	"os"
	"path/filepath"

	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
)

var DefaultSqliteDbPath = gcmd.GetOpt("sqlite-path", "./out/db.sqlite3").String()

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
			g.Log().Fatalf(ctx, "创建数据库目录失败: %v", err)
		}
		g.Log().Infof(ctx, "创建数据库目录: %s", dbDir)
	}

	// 检查数据库文件是否存在
	dbExists := true
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		dbExists = false
		g.Log().Infof(ctx, "数据库文件不存在，将创建新的数据库文件: %s", dbPath)
	}

	// 测试数据库连接（如果文件不存在，SQLite会自动创建）
	if _, err := g.DB().Exec(ctx, "SELECT 1"); err != nil {
		g.Log().Fatalf(ctx, "数据库连接失败: %v", err)
	}

	// 如果是新创建的数据库，记录日志
	if !dbExists {
		g.Log().Infof(ctx, "成功创建新的数据库文件: %s", dbPath)
	}

	// 创建协议表
	_, err := g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS protocol (
			id VARCHAR(255) PRIMARY KEY,
			name VARCHAR(255) NOT NULL,
			type INTEGER NOT NULL,
			address VARCHAR(255) NOT NULL,
			timeout INTEGER NOT NULL,
			log_level VARCHAR(255) NOT NULL,
			params TEXT
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建协议类型表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS protocol_type (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			remark VARCHAR(255)
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 插入协议类型初始数据
	protocolTypes := []struct {
		name   string
		remark string
	}{
		{"modbus_tcp", "Modbus TCP协议"},
		{"modbus_rtu", "Modbus RTU协议"},
		{"canbus", "CAN总线协议"},
		{"canbus_tcp", "CAN总线TCP协议"},
		{"gpio", "GPIO协议"},
		{"virtual", "虚拟协议"},
	}

	for _, pt := range protocolTypes {
		// 检查是否已存在该协议类型
		count, err := g.DB().GetValue(ctx, "SELECT COUNT(*) FROM protocol_type WHERE name = ?", pt.name)
		if err != nil {
			g.Log().Fatal(ctx, err)
		}
		if count.Int() == 0 {
			// 插入协议类型数据
			_, err = g.DB().Exec(ctx, "INSERT INTO protocol_type (name, remark) VALUES (?, ?)", pt.name, pt.remark)
			if err != nil {
				g.Log().Fatal(ctx, err)
			}
			g.Log().Infof(ctx, "插入协议类型: %s - %s", pt.name, pt.remark)
		}
	}

	// 创建设备表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS device (
			id VARCHAR(255) PRIMARY KEY,
			pid VARCHAR(255) NOT NULL,
			protocol_id VARCHAR(255),
			name VARCHAR(255) NOT NULL,
			driver VARCHAR(255),
			log_level VARCHAR(255),
			enable BOOLEAN,
			params TEXT,
			retention_days INTEGER NOT NULL
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	// 创建设置表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS setting (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(255) NOT NULL,
			value VARCHAR(255),
			enable BOOLEAN,
			remark VARCHAR(255)
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Info(ctx, "Config tables created successfully")
}

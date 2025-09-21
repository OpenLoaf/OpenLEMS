package internal

import (
	"context"
	"os"
	"path/filepath"

	"s_db/s_db_basic"
	"s_db/s_db_model"

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
		    enabled BOOLEAN DEFAULT 1,
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
			enable_debug BOOLEAN DEFAULT 0,
		    manual_mode BOOLEAN DEFAULT FALSE,
			params TEXT DEFAULT '{}',
		    storage_enable BOOLEAN DEFAULT 1,
			storage_interval_sec INTEGER NOT NULL DEFAULT 60,
			external_param TEXT DEFAULT '{}',
			sort INTEGER DEFAULT 0,
		    enabled BOOLEAN DEFAULT 1,
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
			group_name VARCHAR(255) DEFAULT 'system',
			value TEXT,
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

	// 创建自动化表
	_, err = g.DB().Exec(ctx, `
		CREATE TABLE IF NOT EXISTS automation (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			start_time DATETIME,
			end_time DATETIME,
			time_range_type VARCHAR(255) ,
			time_range_value VARCHAR(255),
			trigger_rule TEXT NOT NULL,
			execute_rule TEXT NOT NULL,
			execution_interval INTEGER DEFAULT 0,
			enabled BOOLEAN DEFAULT 1,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		g.Log().Fatal(ctx, err)
	}

	g.Log().Info(ctx, "Database tables created successfully")

	// 初始化系统设置数据
	initSystemSettings(context.Background())
}

// initSystemSettings 初始化系统设置数据
func initSystemSettings(ctx context.Context) {
	// 获取所有系统设置定义
	systemSettings := []*s_db_basic.SSystemSettingDefine{
		s_db_basic.SystemSettingActiveDeviceRootId,
		s_db_basic.SystemSettingActivePolicyId,
		s_db_basic.SystemSettingAutomationInternalMilliseconds,
		s_db_basic.SystemSettingDeviceRetentionDays,
		s_db_basic.SystemSettingSystemRetentionDays,
		s_db_basic.SystemSettingLogRetentionDays,
		s_db_basic.SystemSettingSystemEnableDebugLog,
		s_db_basic.SystemSettingLicenseKey,
		s_db_basic.SystemSettingUserPassword,
		s_db_basic.SystemSettingAdminPassword,
		s_db_basic.SystemSettingPasswordLength,
		s_db_basic.SystemSettingMqttConfigList,
		s_db_basic.SystemSettingModbusConfig,
	}

	// 遍历所有系统设置定义，创建默认设置记录
	for index, settingDefine := range systemSettings {
		// 检查设置是否已存在
		existingSetting := &s_db_model.SSettingModel{}
		err := existingSetting.GetById(ctx, settingDefine.Id)
		if err != nil {
			// 设置不存在，创建新的设置记录
			newSetting := &s_db_model.SSettingModel{
				SDatabaseBasic: s_db_model.SDatabaseBasic{
					Id: settingDefine.Id,
				},
				Value:    settingDefine.DefaultValue,
				IsPublic: settingDefine.IsPublic,
				Enabled:  true,
				Remark:   settingDefine.Remark,
				Sort:     index, // 使用数组中的顺序作为排序
				Group:    settingDefine.Group,
			}

			err = newSetting.Create(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "创建系统设置失败 - 设置ID: %s, 错误: %+v", settingDefine.Id, err)
			} else {
				g.Log().Infof(ctx, "创建系统设置成功 - 设置ID: %s, 默认值: %s", settingDefine.Id, settingDefine.DefaultValue)
			}
		} else {
			// 设置已存在，检查是否需要更新分组、备注、排序和公开状态信息
			needUpdate := false
			if existingSetting.Group != settingDefine.Group {
				existingSetting.Group = settingDefine.Group
				needUpdate = true
			}
			if existingSetting.Remark != settingDefine.Remark {
				existingSetting.Remark = settingDefine.Remark
				needUpdate = true
			}
			if existingSetting.IsPublic != settingDefine.IsPublic {
				existingSetting.IsPublic = settingDefine.IsPublic
				needUpdate = true
			}
			if existingSetting.Sort != index {
				existingSetting.Sort = index
				needUpdate = true
			}

			if needUpdate {
				err = existingSetting.Update(ctx)
				if err != nil {
					g.Log().Errorf(ctx, "更新系统设置失败 - 设置ID: %s, 错误: %+v", settingDefine.Id, err)
				} else {
					g.Log().Infof(ctx, "更新系统设置成功 - 设置ID: %s", settingDefine.Id)
				}
			}
		}
	}

	g.Log().Info(ctx, "System settings initialized successfully")
}

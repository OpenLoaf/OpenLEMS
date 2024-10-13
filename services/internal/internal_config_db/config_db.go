package internal_config_db

import (
	"common/c_base"
	_ "github.com/gogf/gf/contrib/drivers/sqlite/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func init() {
	// 判断数据库是否存在，如果不存在，就新建

}

func GetDriverConfig() *c_base.SDriverConfig {
	var driverConfig *c_base.SDriverConfig

	err := g.Model(c_base.SDriverConfig{}).WithAll().Where("parent_id", nil).Scan(&driverConfig)
	if err != nil {
		return nil
	}
	return driverConfig
}

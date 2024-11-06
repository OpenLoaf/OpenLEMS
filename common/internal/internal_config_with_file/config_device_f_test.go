package internal_config_with_file

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
	"testing"
)

// SDriverConfig 基础设备配置
type SDriverConfig struct {
	gmeta.Meta         `orm:"table:device"`
	Id                 string            `json:"id,omitempty" orm:"id"`                            // 设备ID
	Name               string            `json:"name,omitempty" orm:"name"`                        // 设备名称
	ProtocolId         string            `json:"protocolId,omitempty" orm:"protocolId"`            // 协议配置ID,如果配置了肯定是实体设备
	Driver             string            `json:"driver,omitempty" orm:"driver"`                    // 驱动名称
	IsEnable           bool              `json:"enable" orm:"isEnable"`                            // 是否启用
	StorageIntervalSec int32             `json:"storageIntervalSec" orm:"storageIntervalSec"`      // 存储间隔(秒),0代表默认1分钟，负数代表不存储
	LogLevel           string            `json:"logLevel,omitempty" orm:"logLevel"`                // 日志等级
	Params             map[string]string `json:"params,omitempty" orm:"params"`                    // 额外参数
	DeviceChildren     []*SDriverConfig  `json:"deviceChildren,omitempty" orm:"with:parent_id=id"` // 子设备
}

func TestSConfig_GetDriverConfig(t *testing.T) {
	ctx := context.Background()
	c := &SConfig{
		deviceCfgName: "device",
	}

	//config := c.GetDriverConfig(ctx)
	//g.Log().Infof(ctx, "config: %v", config)

	config, configPath, err := getConfig[SDriverConfig](ctx, c.deviceCfgName, "device")

	if err != nil {
		t.Errorf("GetDriverConfig() error = %v", err)
		return
	}

	g.Log().Infof(ctx, "加载设备%s \n%+v,", configPath, config)
}

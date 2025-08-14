package internal

import (
	"common/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
)

type SCanbusConfig struct {
	BaudRate  uint32   // 波特率
	FilterIds []uint32 // 过滤ID列表
}

// GetCanbusConfig 从协议配置中获取canbus配置
func GetCanbusConfig(protocolConfig *c_base.SProtocolConfig) *SCanbusConfig {
	if protocolConfig == nil {
		panic(gerror.Newf("protocolConfig is nil"))
	}
	canbusConfig := &SCanbusConfig{}
	err := gconv.Scan(protocolConfig.Params, canbusConfig)
	if err != nil {
		panic(gerror.Newf("canbusConfig params error: %s", err.Error()))

	}
	return canbusConfig
}

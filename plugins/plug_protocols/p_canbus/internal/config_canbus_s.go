package internal

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"context"

	"github.com/gogf/gf/v2/util/gconv"
)

// GetCanbusConfig 从协议配置中获取canbus配置
func GetCanbusConfig(protocolConfig *c_base.SProtocolConfig) *c_proto.SCanbusConfig {
	if protocolConfig == nil {
		return nil
	}
	canbusConfig := &c_proto.SCanbusConfig{}
	err := gconv.Scan(protocolConfig.Params, canbusConfig)
	if err != nil {
		c_log.BizErrorf(context.Background(), "canbusConfig params error: %s", err.Error())
	}
	return canbusConfig
}

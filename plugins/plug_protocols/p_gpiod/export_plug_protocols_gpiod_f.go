package p_gpiod

import (
	"common/c_base"
	"common/c_log"
	"common/c_proto"
	"context"
	"p_gpiod/internal"
)

// NewGpiodInProvider 创建GPIO输入provider，一个设备一个provider，同时gpio不能重复被使用
func NewGpiodInProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	return internal.NewGpiodInProvider(ctx, clientConfig, deviceConfig)
}

// NewGpiodOutProvider 创建GPIO输出provider，一个设备一个provider，同时gpio不能重复被使用
func NewGpiodOutProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	return internal.NewGpiodOutProvider(ctx, clientConfig, deviceConfig)
}

// GetGpiodDeviceConfigFields 获取gpiod的设备配置
func GetGpiodDeviceConfigFields() []*c_base.SFieldDefinition {
	gpiodDeviceConfig := &c_proto.SGpiodProtocolConfig{}
	fields, err := c_base.BuildConfigStructFields(gpiodDeviceConfig)
	if err != nil {
		c_log.BizErrorf(context.Background(), "解析GPIO配置信息结构失败！")
		return nil
	}
	return fields
}

package internal_config_with_sqlite

import (
	"common/c_base"
	"context"
)

func (c *SConfig) GetDriverConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig {
	return c.instance.GetDeviceConfig(ctx, activeDeviceRootId)
}

func (c *SConfig) GetProtocolsConfig(ctx context.Context) []*c_base.SProtocolConfig {
	return c.instance.GetProtocolConfig(ctx)
}

func (c *SConfig) GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig {
	list := c.GetProtocolsConfig(ctx)
	for _, protocol := range list {
		if protocol.Id == id {
			return protocol
		}
	}
	return nil
}

package impl

import (
	"common"
	"common/c_base"
	"context"
	"s_db/s_db_interface"
	"sync"
)

type driverConfigServiceImpl struct {
	manage s_db_interface.IConfigService
}

var (
	driverConfigServiceInstance common.IDriverConfigService
	driverConfigServiceOnce     sync.Once
)

func GetDriverConfigService() common.IDriverConfigService {
	driverConfigServiceOnce.Do(func() {
		driverConfigServiceInstance = &driverConfigServiceImpl{
			manage: GetConfigService(),
		}
	})
	return driverConfigServiceInstance
}

func (d *driverConfigServiceImpl) GetDriverConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig {
	return d.manage.GetDeviceConfig(ctx, activeDeviceRootId)
}

//func (d *driverConfigService) GetStorageConfig(ctx context.Context) *c_base.SStorageConfig {
//	return d.manage.Get
//}

func (d *driverConfigServiceImpl) GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig {
	return d.manage.GetProtocolsConfigList(ctx)
}

func (d *driverConfigServiceImpl) GetProtocolsConfigMap(ctx context.Context) map[string]*c_base.SProtocolConfig {
	list := d.manage.GetProtocolsConfigList(ctx)
	m := make(map[string]*c_base.SProtocolConfig)
	for _, protocol := range list {
		m[protocol.Id] = protocol
	}
	return m
}

func (d *driverConfigServiceImpl) GetProtocolById(ctx context.Context, id string) *c_base.SProtocolConfig {
	//TODO implement me
	panic("implement me")
}

package impl

import (
	"common"
	"common/c_base"
	"context"
	"s_db/s_db_basic"
	"sync"
)

// driverConfigServiceImpl 实现 common.IDriverConfigServ 接口。通过依赖注入控制反转
type driverConfigServiceImpl struct {
	manage s_db_basic.IConfigService
}

var (
	driverConfigServiceInstance common.IDriverConfigServ
	driverConfigServiceOnce     sync.Once
)

func GetDriverConfigService() common.IDriverConfigServ {
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

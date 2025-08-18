package impl

import (
	"common/c_base"
	"context"
	"s_db/basic"
	"s_db/model"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

type sDeviceServiceImpl struct {
	deviceModel *model.SDeviceModel
	ctx         context.Context
}

func (s *sDeviceServiceImpl) GetRootDeviceId() string {
	return GetSettingService().GetSettingValueByKey(s.ctx, basic.ActiveDeviceRootId, basic.DefaultActiveDeviceRootId)
}

var (
	deviceManageInstance basic.IDeviceService
	deviceManageOnce     sync.Once
)

func GetDeviceService() basic.IDeviceService {
	deviceManageOnce.Do(func() {
		deviceManageInstance = &sDeviceServiceImpl{
			ctx: context.Background(),
		}
	})
	return deviceManageInstance
}

func (s *sDeviceServiceImpl) GetDeviceConfigs(ctx context.Context, parentId string) ([]*c_base.SDeviceConfig, error) {

	// devices, err := s.GetAllDevicesOrderBySortAndEnable(ctx, true)

	list, err := g.DB().GetAll(ctx, `
		WITH RECURSIVE DeviceDescendants AS (
			SELECT *
			FROM device
			WHERE id = ?  OR pid = ? AND enabled =true
			UNION ALL
			SELECT d.*
			FROM device AS d
					 JOIN DeviceDescendants AS dd ON d.pid = dd.id
			where d.enabled = true
		)
		SELECT * FROM DeviceDescendants
`, g.Slice{parentId, parentId})

	if err != nil {
		return nil, err
	}

	var devices []*c_base.SDeviceConfig
	if convErr := list.Structs(&devices); convErr != nil {
		return nil, convErr
	}

	return devices, nil
}

// DeleteById 根据ID删除设备记录
func (s *sDeviceServiceImpl) DeleteDeviceById(ctx context.Context, id string) error {
	_, err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldId, id).Delete()
	return err
}

// GetAll 获取所有设备记录
func (s *sDeviceServiceImpl) GetAllDevices(ctx context.Context) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Scan(&devices)
	return devices, err
}

// GetByCondition 根据条件获取设备记录
func (s *sDeviceServiceImpl) GetDevicesByCondition(ctx context.Context, condition g.Map) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(condition).Scan(&devices)
	return devices, err
}

func (s *sDeviceServiceImpl) GetRecursiveDevicesByPid(ctx context.Context, pid string) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldPid, pid).Scan(&devices)
	if err != nil {
		return nil, err
	}

	for _, device := range devices {
		subDevices, err := s.GetRecursiveDevicesByPid(ctx, device.Id)
		if err != nil {
			return nil, err
		}
		devices = append(devices, subDevices...)
	}

	return devices, nil
}

// GetDevicesById 根据ID获取设备记录
func (s *sDeviceServiceImpl) GetDeviceById(ctx context.Context, id string) (*model.SDeviceModel, error) {
	var device *model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldId, id).Scan(&device)
	return device, err
}

// GetDevicesByPid 根据父设备ID获取子设备列表
func (s *sDeviceServiceImpl) GetDevicesByPid(ctx context.Context, pid string) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldPid, pid).Scan(&devices)
	return devices, err
}

// GetByProtocolId 根据协议ID获取设备列表
func (s *sDeviceServiceImpl) GetDevicesByProtocolId(ctx context.Context, protocolId string) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldProtocolId, protocolId).Scan(&devices)
	return devices, err
}

// GetEnabledDevices 获取所有启用的设备
func (s *sDeviceServiceImpl) GetEnabledDevices(ctx context.Context) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldEnable, true).Scan(&devices)
	return devices, err
}

// Count 获取设备总数
func (s *sDeviceServiceImpl) CountDevices(ctx context.Context) (int, error) {
	count, err := g.Model(s.deviceModel).Ctx(ctx).Count()
	return count, err
}

// CountByCondition 根据条件获取设备数量
func (s *sDeviceServiceImpl) CountDevicesByCondition(ctx context.Context, condition g.Map) (int, error) {
	count, err := g.Model(s.deviceModel).Ctx(ctx).Where(condition).Count()
	return count, err
}

// Paginate 分页获取设备列表
func (s *sDeviceServiceImpl) PaginateDevices(ctx context.Context, page, pageSize int) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Page(page, pageSize).Scan(&devices)
	return devices, err
}

// GetAllDevicesOrderBySort 获取所有设备记录，按sort字段排序
func (s *sDeviceServiceImpl) GetAllDevicesOrderBySort(ctx context.Context) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Order(model.FieldSort).Scan(&devices)
	return devices, err
}

// GetAllDevicesOrderBySortAndEnable 获取所有设备记录，按sort字段排序，enable参数控制是否只获取启用的设备
func (s *sDeviceServiceImpl) GetAllDevicesOrderBySortAndEnable(ctx context.Context, enable bool) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldEnable, enable).Order(model.FieldSort).Scan(&devices)
	return devices, err
}

// GetDevicesByPidOrderBySort 根据父设备ID获取子设备列表，按sort字段排序
func (s *sDeviceServiceImpl) GetDevicesByPidOrderBySort(ctx context.Context, pid string) ([]*model.SDeviceModel, error) {
	var devices []*model.SDeviceModel
	err := g.Model(s.deviceModel).Ctx(ctx).Where(model.FieldPid, pid).Order(model.FieldSort).Scan(&devices)
	return devices, err
}

package impl

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sSettingServiceImpl struct {
}

var (
	configManageInstance s_db_basic.ISettingService
	configManageOnce     sync.Once
)

func GetSettingService() s_db_basic.ISettingService {
	configManageOnce.Do(func() {
		configManageInstance = &sSettingServiceImpl{}
	})
	return configManageInstance
}

func (s *sSettingServiceImpl) GetAllSettings(ctx context.Context) ([]*s_db_model.SSettingModel, error) {
	// 调用模型层的 GetAllSettings 方法获取所有设置
	settings, err := s_db_model.GetAllSettings(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取所有设置失败 - 错误: %+v", err)
		return nil, err
	}

	c_log.Debugf(ctx, "成功获取所有设置，共 %d 条记录", len(settings))
	return settings, nil
}

// GetAllSettingsByGroup 根据分组获取所有设置
func (s *sSettingServiceImpl) GetAllSettingsByGroup(ctx context.Context, group string) ([]*s_db_model.SSettingModel, error) {
	// 调用模型层的 GetSettingsByGroup 方法获取指定分组的设置
	settings, err := s_db_model.GetSettingsByGroup(ctx, group)
	if err != nil {
		g.Log().Errorf(ctx, "获取分组设置失败 - 分组: %s, 错误: %+v", group, err)
		return nil, err
	}

	c_log.Debugf(ctx, "成功获取分组设置 - 分组: %s, 共 %d 条记录", group, len(settings))
	return settings, nil
}

// GetSettingById 根据ID获取设置详情
func (s *sSettingServiceImpl) GetSettingById(ctx context.Context, id string) (*s_db_model.SSettingModel, error) {
	setting := &s_db_model.SSettingModel{}
	err := setting.GetById(ctx, id)
	if err != nil {
		c_log.Debugf(ctx, "获取设置详情失败 - 设置ID: %s, 错误: %s", id, err.Error())
		return nil, nil
	}

	c_log.Debugf(ctx, "成功获取设置详情 - 设置ID: %s", id)
	return setting, nil
}

// 获取设置配置通过名称
func (s *sSettingServiceImpl) GetSettingValueById(ctx context.Context, id string) string {
	setting := &s_db_model.SSettingModel{}
	// 通过 id 获取设置，如果设置不存在，则返回空字符串
	err := setting.GetById(ctx, id)

	if err != nil {
		g.Log().Warningf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", id, err)
		return ""
	}

	// 检查设置是否启用
	if !setting.Enabled {
		g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
		return ""
	}

	return setting.GetValue()
}

// 获取设置配置通过名称，支持默认值和分组
func (s *sSettingServiceImpl) GetSettingValueByIdWithDefaultValue(ctx context.Context, id, group, defaultValue string, remark ...string) string {
	setting := &s_db_model.SSettingModel{}
	// 通过 id 获取设置，如果设置不存在，则创建默认设置
	err := setting.GetById(ctx, id)

	if err != nil {
		g.Log().Warningf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", id, err)
		setting.SDatabaseBasic = s_db_model.SDatabaseBasic{
			Id:        id,
			CreatedAt: gtime.Now(),
			UpdatedAt: gtime.Now(),
		}
		setting.Value = defaultValue
		setting.Enabled = true
		setting.Sort = 999
		setting.Group = group

		if len(remark) > 0 {
			setting.Remark = remark[0]
		}

		err = setting.Create(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "保存设置失败！设置名称：%s，值：%v 错误：%v", id, defaultValue, err)
		}
		c_log.Infof(ctx, "保存默认设置成功！设置名称：%s，值：%s，分组：%s", id, defaultValue, group)
		return defaultValue
	}

	// 检查设置是否启用
	if !setting.Enabled {
		g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
		return defaultValue
	}

	return setting.GetValue()
}

// 设置设置值通过名称
func (s *sSettingServiceImpl) SetSettingValueById(ctx context.Context, id string, value string) error {
	setting := &s_db_model.SSettingModel{}
	err := setting.GetById(ctx, id)
	if err != nil {
		g.Log().Errorf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", id, err)
		return err
	}
	setting.SetValue(value)
	_ = setting.Update(ctx)
	return nil
}

// GetRootDeviceId 获取根设备ID
func (s *sSettingServiceImpl) GetRootDeviceId(ctx context.Context) string {
	return s.GetSettingValueByIdWithDefaultValue(ctx, s_db_basic.SettingActiveDeviceRootIdKey, c_enum.ESettingGroupSystem, s_db_basic.DefaultActiveDeviceRootId, "默认激活的根设备ID")
}

// GetRootPolicyId 获取激活的策略ID
func (s *sSettingServiceImpl) GetRootPolicyId(ctx context.Context) string {
	return s.GetSettingValueById(ctx, s_db_basic.SettingActivePolicyIdKey)
}

// GetPublicEnabledSettings 获取公开且启用的设置
func (s *sSettingServiceImpl) GetPublicEnabledSettings(ctx context.Context) ([]*s_db_model.SSettingModel, error) {
	// 调用模型层的 GetPublicEnabledSettings 方法
	settings, err := s_db_model.GetPublicEnabledSettings(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取公开且启用的设置失败 - 错误: %+v", err)
		return nil, err
	}

	c_log.Infof(ctx, "成功获取公开且启用的设置，共 %d 条记录", len(settings))
	return settings, nil
}

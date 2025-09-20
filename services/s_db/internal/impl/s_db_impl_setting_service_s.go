package impl

import (
	"common/c_enum"
	"common/c_log"
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gtime"
)

type sSettingServiceImpl struct {
	settingCache  *gcache.Cache // 设置缓存
	cacheDuration time.Duration // 缓存过期时间
}

const (
	// 缓存相关常量
	DefaultCacheDuration = 5 * time.Minute // 默认缓存过期时间
	CacheKeyPrefix       = "setting:"      // 缓存键前缀
)

var (
	configManageInstance s_db_basic.ISettingService
	configManageOnce     sync.Once
)

func GetSettingService() s_db_basic.ISettingService {
	configManageOnce.Do(func() {
		configManageInstance = &sSettingServiceImpl{
			settingCache:  gcache.New(),         // 创建设置缓存
			cacheDuration: DefaultCacheDuration, // 设置缓存过期时间
		}
	})
	return configManageInstance
}

// getCacheKey 生成缓存键
func (s *sSettingServiceImpl) getCacheKey(id string) string {
	return CacheKeyPrefix + id
}

// getFromCache 从缓存获取设置数据
func (s *sSettingServiceImpl) getFromCache(ctx context.Context, id string) (*s_db_model.SSettingModel, bool) {
	key := s.getCacheKey(id)
	value, err := s.settingCache.Get(ctx, key)
	if err != nil || value == nil {
		return nil, false
	}

	// 从 gvar.Var 中获取实际值
	actualValue := value.Val()
	if setting, ok := actualValue.(*s_db_model.SSettingModel); ok {
		c_log.Debugf(ctx, "从缓存获取设置成功 - 设置ID: %s", id)
		return setting, true
	}

	return nil, false
}

// setToCache 设置缓存
func (s *sSettingServiceImpl) setToCache(ctx context.Context, id string, setting *s_db_model.SSettingModel) {
	key := s.getCacheKey(id)
	err := s.settingCache.Set(ctx, key, setting, s.cacheDuration)
	if err != nil {
		c_log.Warningf(ctx, "设置缓存失败 - 设置ID: %s, 错误: %v", id, err)
	} else {
		c_log.Debugf(ctx, "设置缓存成功 - 设置ID: %s", id)
	}
}

// clearAllCache 清除所有缓存
func (s *sSettingServiceImpl) clearAllCache(ctx context.Context) {
	s.settingCache.Clear(ctx)
	c_log.Debugf(ctx, "已清除所有设置缓存")
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
	// 先尝试从缓存获取
	if setting, found := s.getFromCache(ctx, id); found {
		return setting, nil
	}

	// 缓存未命中，从数据库获取
	setting := &s_db_model.SSettingModel{}
	err := setting.GetById(ctx, id)
	if err != nil {
		c_log.Debugf(ctx, "获取设置详情失败 - 设置ID: %s, 错误: %s", id, err.Error())
		return nil, nil
	}

	// 将结果存入缓存
	s.setToCache(ctx, id, setting)

	c_log.Debugf(ctx, "成功获取设置详情 - 设置ID: %s", id)
	return setting, nil
}

// 获取设置配置通过名称
func (s *sSettingServiceImpl) GetSettingValueById(ctx context.Context, id string) string {
	// 先尝试从缓存获取
	if setting, found := s.getFromCache(ctx, id); found {
		// 检查设置是否启用
		if !setting.Enabled {
			g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
			return ""
		}
		return setting.GetValue()
	}

	// 缓存未命中，从数据库获取
	setting := &s_db_model.SSettingModel{}
	err := setting.GetById(ctx, id)
	if err != nil {
		g.Log().Warningf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", id, err)
		return ""
	}

	// 将结果存入缓存
	s.setToCache(ctx, id, setting)

	// 检查设置是否启用
	if !setting.Enabled {
		g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
		return ""
	}

	return setting.GetValue()
}

// 获取设置配置通过名称，支持默认值和分组
func (s *sSettingServiceImpl) GetSettingValueByIdWithDefaultValue(ctx context.Context, id, group, defaultValue string, remark ...string) string {
	// 先尝试从缓存获取
	if setting, found := s.getFromCache(ctx, id); found {
		// 检查设置是否启用
		if !setting.Enabled {
			g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
			return defaultValue
		}
		return setting.GetValue()
	}

	// 缓存未命中，从数据库获取
	setting := &s_db_model.SSettingModel{}
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
		} else {
			// 创建成功后，将新设置存入缓存
			s.setToCache(ctx, id, setting)
		}
		c_log.Infof(ctx, "保存默认设置成功！设置名称：%s，值：%s，分组：%s", id, defaultValue, group)
		return defaultValue
	}

	// 将结果存入缓存
	s.setToCache(ctx, id, setting)

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
	err = setting.Update(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "更新设置失败 - 设置名称: %s, 错误: %v", id, err)
		return err
	}

	// 更新成功后，清除所有缓存
	s.clearAllCache(ctx)

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

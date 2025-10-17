package impl

import (
	"common/c_log"
	"context"
	"s_db/s_db_basic"
	"s_db/s_db_model"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
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
func (s *sSettingServiceImpl) GetSettingValueById(ctx context.Context, id string) *string {
	// 先尝试从缓存获取
	if setting, found := s.getFromCache(ctx, id); found {
		// 检查设置是否启用
		if !setting.Enabled {
			g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
			return nil
		}
		value := setting.GetValue()
		return &value
	}

	// 缓存未命中，从数据库获取
	setting := &s_db_model.SSettingModel{}
	err := setting.GetById(ctx, id)
	if err != nil {
		g.Log().Warningf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", id, err)
		return nil
	}

	// 将结果存入缓存
	s.setToCache(ctx, id, setting)

	// 检查设置是否启用
	if !setting.Enabled {
		g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", id)
		return nil
	}

	value := setting.GetValue()
	return &value
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
func (s *sSettingServiceImpl) GetRootDeviceId(ctx context.Context) *string {
	return s.GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingActiveDeviceRootId)
}

// GetRootPolicyId 获取激活的策略ID
func (s *sSettingServiceImpl) GetRootPolicyId(ctx context.Context) *string {
	return s.GetSettingValueBySystemSettingDefine(ctx, s_db_basic.SystemSettingActivePolicyId)
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

// GetSettingValueBySystemSettingDefine 通过系统设置定义获取设置值
func (s *sSettingServiceImpl) GetSettingValueBySystemSettingDefine(ctx context.Context, settingDefine *s_db_basic.SSystemSettingDefine) *string {
	if settingDefine == nil {
		c_log.Warning(ctx, "系统设置定义为空")
		return nil
	}

	// 先尝试从缓存获取
	if setting, found := s.getFromCache(ctx, settingDefine.Id); found {
		// 检查设置是否启用
		if !setting.Enabled {
			g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", settingDefine.Id)
			return &settingDefine.DefaultValue
		}
		c_log.Debugf(ctx, "通过系统设置定义获取设置值成功 - 设置ID: %s, 值: %s", settingDefine.Id, setting.GetValue())
		value := setting.GetValue()
		return &value
	}

	// 缓存未命中，从数据库获取
	setting := &s_db_model.SSettingModel{}
	err := setting.GetById(ctx, settingDefine.Id)
	if err != nil {
		g.Log().Errorf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", settingDefine.Id, err)
		// 系统启动时已初始化所有设置，如果获取失败说明数据库有问题
		return &settingDefine.DefaultValue
	}

	// 将结果存入缓存
	s.setToCache(ctx, settingDefine.Id, setting)

	// 检查设置是否启用
	if !setting.Enabled {
		g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", settingDefine.Id)
		return &settingDefine.DefaultValue
	}

	c_log.Debugf(ctx, "通过系统设置定义获取设置值成功 - 设置ID: %s, 值: %s", settingDefine.Id, setting.GetValue())
	value := setting.GetValue()
	return &value
}

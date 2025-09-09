package impl

import (
	"common/c_base"
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

	g.Log().Infof(ctx, "成功获取所有设置，共 %d 条记录", len(settings))
	return settings, nil
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
func (s *sSettingServiceImpl) GetSettingValueByIdWithDefaultValue(ctx context.Context, id, group, defaultValue string) string {
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

		err = setting.Create(ctx)
		if err != nil {
			g.Log().Errorf(ctx, "保存设置失败！设置名称：%s，值：%v 错误：%v", id, defaultValue, err)
		}
		g.Log().Infof(ctx, "保存默认设置成功！设置名称：%s，值：%s，分组：%s", id, defaultValue, group)
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
	return s.GetSettingValueByIdWithDefaultValue(ctx, s_db_basic.SettingActiveDeviceRootIdKey, c_base.ESettingGroupSystem, s_db_basic.DefaultActiveDeviceRootId)
}

// GetRootPolicyId 获取激活的策略ID
func (s *sSettingServiceImpl) GetRootPolicyId(ctx context.Context) string {
	return s.GetSettingValueById(ctx, s_db_basic.SettingActivePolicyIdKey)
}

/*func (s *sSettingServiceImpl) GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig {
	protocols, err := GetProtocolService().GetAllProtocols(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取协议配置失败 - 错误: %+v", err)
		return nil
	}

	if len(protocols) == 0 {
		g.Log().Warningf(ctx, "未找到任何协议配置")
		return nil
	}

	protocolConfigs := make([]*c_base.SProtocolConfig, 0, len(protocols))
	for _, protocol := range protocols {

		params, err := protocol.GetParamsMap()
		if err != nil {
			g.Log().Errorf(context.Background(), "获取协议参数失败 - 协议ID: %d, 协议名称: %s, 错误: %v",
				protocol.Id, protocol.Name, err)
			continue
		}

		protocolConfig := &c_base.SProtocolConfig{
			Id:       protocol.Id,
			Type:     c_base.EProtocolType(protocol.Type),
			SourceAddress:  protocol.SourceAddress,
			Timeout:  protocol.Timeout,
			LogLevel: protocol.LogLevel,
			Params:   params,
		}
		protocolConfigs = append(protocolConfigs, protocolConfig)
	}

	// 添加调试打印 - 使用JSON格式清楚显示protocolConfigs结构
	if len(protocolConfigs) > 0 {
		protocolsJSON, err := json.MarshalIndent(protocolConfigs, "", "  ")
		if err != nil {
			g.Log().Errorf(ctx, "序列化protocolConfigs为JSON失败: %+v", err)
		} else {
			g.Log().Infof(ctx, "协议配置结构:\n%s", string(protocolsJSON))
		}

		// 额外打印一些关键信息
		g.Log().Infof(ctx, "协议配置数量: %d", len(protocolConfigs))
		for i, protocol := range protocolConfigs {
			g.Log().Infof(ctx, "协议[%d]: ID=%s, SProtocolModel=%s",
				i, protocol.Id, protocol.Type)
		}
	} else {
		g.Log().Infof(ctx, "没有找到任何协议配置")
	}
	return protocolConfigs
}*/

/*
// BuildDeviceTree 递归构建设备树结构
func BuildDeviceTree(devices []*model.SDeviceModel) []*c_base.SDeviceConfig {
	// 建立索引
	idToDevice := make(map[string]*model.SDeviceModel, len(devices))
	pidToChildren := make(map[string][]*model.SDeviceModel, len(devices))
	for _, d := range devices {
		idToDevice[d.Id] = d
		pidToChildren[d.Pid] = append(pidToChildren[d.Pid], d)
	}

	// 根节点：pid 不存在于 idToDevice 或者 pid 为空/"0"
	var roots []*model.SDeviceModel
	for _, d := range devices {
		if d.Pid == "" || d.Pid == basic.DefaultActiveDeviceRootId || idToDevice[d.Pid] == nil {
			roots = append(roots, d)
		}
	}

	var buildNode func(device *model.SDeviceModel) *c_base.SDeviceConfig
	buildNode = func(device *model.SDeviceModel) *c_base.SDeviceConfig {
		// 获取设备参数
		params, err := device.GetParamsMap()
		if err != nil {
			g.Log().Errorf(context.Background(), "获取设备参数失败 - 设备ID: %s, 设备名称: %s, 参数原始值: %s, 错误: %v", device.Id, device.Name, device.Params, err)
			return nil
		}
		node := &c_base.SDeviceConfig{
			Id:         device.Id,
			ProtocolId: device.ProtocolId,
			Name:       device.Name,
			Driver:     device.Driver,
			LogLevel:   device.LogLevel,
			Enabled:    device.Enabled,
			Params:     params,
		}
		// 子节点
		for _, cm := range pidToChildren[device.Id] {
			if child := buildNode(cm); child != nil {
				node.DeviceChildren = append(node.DeviceChildren, child)
			}
		}
		if len(node.DeviceChildren) == 0 {
			node.DeviceChildren = nil
		}
		return node
	}

	var tree []*c_base.SDeviceConfig
	for _, r := range roots {
		if n := buildNode(r); n != nil {
			tree = append(tree, n)
		}
	}
	if len(tree) == 0 {
		return nil
	}
	return tree
}
*/

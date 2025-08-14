package impl

import (
	"context"
	"encoding/json"
	"s_db/s_db_interface"
	"s_db/s_db_model"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"common/c_base"
)

type sConfigServiceImpl struct {
}

var (
	configManageInstance s_db_interface.IConfigService
	configManageOnce     sync.Once
)

func GetConfigService() s_db_interface.IConfigService {
	configManageOnce.Do(func() {
		configManageInstance = &sConfigServiceImpl{}
	})
	return configManageInstance
}

func (s *sConfigServiceImpl) GetDeviceConfig(ctx context.Context, activeDeviceRootId string) *c_base.SDriverConfig {
	//devices, err := model.GetDevicesByCondition(ctx, g.Map{})
	var (
		deviceRootId string
		err          error
	)

	if activeDeviceRootId == "" {
		deviceRootId, err = s_db_model.GetSettingValueByName("active_device_root_id")
		g.Log().Noticef(ctx, "GetDeviceConfig From DB! active_device_root_id: %v", activeDeviceRootId)
	} else {
		deviceRootId = activeDeviceRootId
	}

	g.Log().Infof(ctx, "Activce SDeviceModel Root Id: %s", deviceRootId)

	if err != nil {
		g.Log().Errorf(ctx, "获取激活的设备父ID配置失败 - 错误: %v", err)
		return nil
	}
	rootDevice, err := GetDeviceService().GetDeviceById(ctx, deviceRootId)
	if err != nil {
		g.Log().Errorf(ctx, "获取激活的设备父ID配置失败 - 错误: %v", err)
		return nil
	}

	g.Log().Infof(ctx, "激活的设备父ID配置: %s", deviceRootId)
	devices, err := GetDeviceService().GetRecursiveDevicesByPid(ctx, deviceRootId)
	devices = append(devices, rootDevice)

	if err != nil {
		g.Log().Errorf(ctx, "获取设备配置失败 - 错误: %v", err)
		return nil
	}

	if len(devices) == 0 {
		g.Log().Warningf(ctx, "未找到任何设备配置")
		return nil
	}

	tree := BuildDeviceTree(ctx, devices, "0")

	//添加调试打印 - 使用JSON格式清楚显示tree结构
	if tree != nil {
		treeJSON, err := json.MarshalIndent(tree, "", "  ")
		if err != nil {
			g.Log().Errorf(ctx, "序列化tree为JSON失败: %v", err)
		} else {
			g.Log().Infof(ctx, "设备树结构:\n%s", string(treeJSON))
		}

		// 额外打印一些关键信息
		g.Log().Infof(ctx, "根设备: ID=%s, Name=%s, Driver=%s, IsEnable=%t",
			tree.Id, tree.Name, tree.Driver, tree.IsEnable)

		if len(tree.DeviceChildren) > 0 {
			g.Log().Infof(ctx, "子设备数量: %d", len(tree.DeviceChildren))
			for i, child := range tree.DeviceChildren {
				g.Log().Infof(ctx, "子设备[%d]: ID=%s, Name=%s, Driver=%s",
					i, child.Id, child.Name, child.Driver)
			}
		}

		// 使用树形结构打印
		g.Log().Infof(ctx, "设备树层级结构:")
		PrintDeviceTree(ctx, tree, 0)
	}

	return tree
}

func (s *sConfigServiceImpl) GetProtocolsConfigList(ctx context.Context) []*c_base.SProtocolConfig {
	protocols, err := GetProtocolService().GetAllProtocols(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "获取协议配置失败 - 错误: %v", err)
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
			Protocol: c_base.EProtocolType(protocol.Type),
			Address:  protocol.Address,
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
			g.Log().Errorf(ctx, "序列化protocolConfigs为JSON失败: %v", err)
		} else {
			g.Log().Infof(ctx, "协议配置结构:\n%s", string(protocolsJSON))
		}

		// 额外打印一些关键信息
		g.Log().Infof(ctx, "协议配置数量: %d", len(protocolConfigs))
		for i, protocol := range protocolConfigs {
			g.Log().Infof(ctx, "协议[%d]: ID=%s, SProtocolModel=%s",
				i, protocol.Id, protocol.Protocol)
		}
	} else {
		g.Log().Infof(ctx, "没有找到任何协议配置")
	}
	return protocolConfigs
}

// 获取设置配置通过名称
func (s *sConfigServiceImpl) GetSettingValueByName(ctx context.Context, name string) string {
	setting := &s_db_model.SSettingModel{}
	// 通过 name 获取设置，如果设置不存在，则返回空字符串
	err := setting.GetByName(ctx, name)
	if err != nil {
		g.Log().Errorf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", name, err)
		return ""
	}

	// 检查设置是否启用
	if !setting.Enable {
		g.Log().Warningf(ctx, "设置已禁用 - 设置名称: %s", name)
		return ""
	}

	return setting.GetValue()
}

// 设置设置值通过名称
func (s *sConfigServiceImpl) SetSettingValueByName(ctx context.Context, name string, value string) error {
	setting := &s_db_model.SSettingModel{}
	err := setting.GetByName(ctx, name)
	if err != nil {
		g.Log().Errorf(ctx, "获取设置失败 - 设置名称: %s, 错误: %v", name, err)
		return err
	}
	setting.SetValue(value)
	setting.Update(ctx)
	return nil
}

// BuildDeviceTree 递归构建设备树结构
func BuildDeviceTree(ctx context.Context, devices []*s_db_model.SDeviceModel, parentId string) *c_base.SDriverConfig {
	var tree []*c_base.SDriverConfig

	for _, device := range devices {
		if device.Pid == parentId {
			// 获取设备参数
			params, err := device.GetParamsMap()
			if err != nil {
				g.Log().Errorf(context.Background(), "获取设备参数失败 - 设备ID: %d, 设备名称: %s, 参数原始值: %s, 错误: %v",
					device.Id, device.Name, device.Params, err)
				continue
			}

			var deviceId string
			if device.Pid == "0" { // Pid 为零的是根元素
				deviceId = "root"
			} else {
				deviceId = device.Id
			}
			// 创建驱动配置
			driverConfig := &c_base.SDriverConfig{
				Id:         deviceId,
				ProtocolId: device.ProtocolId,
				Name:       device.Name,
				Driver:     device.Driver,
				LogLevel:   device.LogLevel,
				IsEnable:   device.Enable,
				Params:     params,
			}

			// 递归获取子设备
			children := BuildDeviceTree(ctx, devices, device.Id)
			if children != nil {
				driverConfig.DeviceChildren = []*c_base.SDriverConfig{children}
			}

			tree = append(tree, driverConfig)
		}
	}

	// 修复潜在的数组越界问题
	if len(tree) == 0 {
		return nil
	}
	return tree[0]
}

// PrintDeviceTree 打印设备树的层级结构
func PrintDeviceTree(ctx context.Context, config *c_base.SDriverConfig, level int) {
	if config == nil {
		return
	}

	// 生成缩进
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}

	// 打印当前设备信息
	g.Log().Infof(ctx, "%s├─ [%s] %s", indent, config.Id, config.Name)
	g.Log().Infof(ctx, "%s   Driver: %s, SProtocolModel: %s, Enable: %t",
		indent, config.Driver, config.ProtocolId, config.IsEnable)

	// 打印参数
	if len(config.Params) > 0 {
		g.Log().Infof(ctx, "%s   Params:", indent)
		for key, value := range config.Params {
			g.Log().Infof(ctx, "%s     %s: %s", indent, key, value)
		}
	}

	// 递归打印子设备
	for _, child := range config.DeviceChildren {
		PrintDeviceTree(ctx, child, level+1)
	}
}

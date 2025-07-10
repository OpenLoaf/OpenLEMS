package service

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gogf/gf/v2/frame/g"

	"common/c_base"
	"sqlite/model"
)

type IConfigManage interface {
	GetDeviceConfig(ctx context.Context) *c_base.SDriverConfig
	GetProtocolConfig(ctx context.Context) []*c_base.SProtocolConfig
	GetSettingValueByName(ctx context.Context, name string) string
	SetSettingValueByName(ctx context.Context, name string, value string) error
}

type sConfigManage struct {
}

var (
	instance IConfigManage
	mu       sync.RWMutex
)

func NewConfigManage(ctx context.Context) IConfigManage {
	// 先用读锁检查实例是否已存在
	mu.RLock()
	if instance != nil {
		mu.RUnlock()
		return instance
	}
	mu.RUnlock()

	// 用写锁创建新实例
	mu.Lock()
	defer mu.Unlock()

	// 双重检查，防止并发创建
	if instance != nil {
		return instance
	}

	// 创建新的单例实例
	instance = &sConfigManage{}
	return instance
}

// ClearConfigManageInstance 清理ConfigManage实例
func ClearConfigManageInstance() {
	mu.Lock()
	defer mu.Unlock()
	instance = nil
}

func (s *sConfigManage) GetDeviceConfig(ctx context.Context) *c_base.SDriverConfig {
	devices, err := model.GetDevicesByCondition(ctx, g.Map{})
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

func (s *sConfigManage) GetProtocolConfig(ctx context.Context) []*c_base.SProtocolConfig {
	protocols, err := model.GetAllProtocols(ctx)
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
			Protocol: c_base.EProtocolType(protocol.Name),
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
			g.Log().Infof(ctx, "协议[%d]: ID=%s, Protocol=%s",
				i, protocol.Id, protocol.Protocol)
		}
	} else {
		g.Log().Infof(ctx, "没有找到任何协议配置")
	}
	return protocolConfigs
}

// 获取设置配置通过名称
func (s *sConfigManage) GetSettingValueByName(ctx context.Context, name string) string {
	setting := &model.Setting{}
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
func (s *sConfigManage) SetSettingValueByName(ctx context.Context, name string, value string) error {
	setting := &model.Setting{}
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
func BuildDeviceTree(ctx context.Context, devices []*model.Device, parentId string) *c_base.SDriverConfig {
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
	g.Log().Infof(ctx, "%s   Driver: %s, Protocol: %s, Enable: %t",
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

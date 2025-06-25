package service

import (
	"common/c_base"
	"context"
	"encoding/json"
	"sqlite/model"

	"github.com/gogf/gf/v2/frame/g"
)

type IConfigManage interface {
	GetDeviceConfig(ctx context.Context) *c_base.SDriverConfig
	GetProtocolConfig(ctx context.Context) []*c_base.SProtocolConfig
}

type sConfigManage struct {
	gId uint
}

func NewConfigManage(ctx context.Context, gId uint) IConfigManage {
	return &sConfigManage{
		gId: gId,
	}
}

func (s *sConfigManage) GetDeviceConfig(ctx context.Context) *c_base.SDriverConfig {
	devices, err := model.GetDevicesByCondition(ctx, g.Map{
		"gid": s.gId,
	})
	if err != nil {
		panic(err)
	}

	if len(devices) == 0 {
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
		panic(err)
	}

	if len(protocols) == 0 {
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
	if config.Params != nil && len(config.Params) > 0 {
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

package device

import (
	v1 "application/api/device/v1"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
	"fmt"
	"sort"
)

// GetDevicePointsDefinition 获取设备全部点位定义
func (c *ControllerV1) GetDevicePointsDefinition(ctx context.Context, req *v1.GetDevicePointsDefinitionReq) (res *v1.GetDevicePointsDefinitionRes, err error) {
	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		// 系统初始化中，返回空数据
		return &v1.GetDevicePointsDefinitionRes{Groups: []*c_base.SPointGroup{}, Fields: []*v1.SDevicePointField{}}, nil
	}

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		return &v1.GetDevicePointsDefinitionRes{Groups: []*c_base.SPointGroup{}, Fields: []*v1.SDevicePointField{}}, nil
	}

	// 检查是否为虚拟设备
	if device.IsVirtualDevice() {
		return c.getVirtualDevicePointsDefinition(device)
	}

	// 分别获取设备点位和遥测点位
	devicePoints := device.GetDevicePoints()
	telemetryPoints := device.GetTelemetryPoints()

	if len(devicePoints) == 0 && len(telemetryPoints) == 0 {
		return &v1.GetDevicePointsDefinitionRes{Groups: []*c_base.SPointGroup{}, Fields: []*v1.SDevicePointField{}}, nil
	}

	fields := make([]*c_base.SFieldDefinition, 0, len(devicePoints)+len(telemetryPoints))

	// 处理设备点位（保持原有group设置）
	for _, p := range devicePoints {
		if fd := c_base.ConvertIPointToFieldDefinition(p); fd != nil {
			fields = append(fields, fd)
		}
	}

	// 处理遥测点位（强制设置group为"i18n:common.summary"）
	for _, p := range telemetryPoints {
		if fd := c_base.ConvertIPointToFieldDefinition(p); fd != nil {
			// 强制设置group为"i18n:common.summary"
			fd.Group = c_base.GroupTotal
			fields = append(fields, fd)
		}
	}

	// 基于 group+key 去重
	uniqueFields := deduplicateFieldsByGroupAndKey(fields)

	// 构建分组列表并排序
	groupMap := make(map[string]*c_base.SPointGroup)
	for _, f := range uniqueFields {
		if f != nil && f.Group != nil {
			if _, ok := groupMap[f.Group.GroupKey]; !ok {
				// 拷贝一份，避免外部修改
				g := *f.Group
				groupMap[g.GroupKey] = &g
			}
		}
	}
	groups := make([]*c_base.SPointGroup, 0, len(groupMap))
	for _, g := range groupMap {
		groups = append(groups, g)
	}
	sort.Slice(groups, func(i, j int) bool { return groups[i].GroupSort < groups[j].GroupSort })

	// 转换为响应字段类型（使用 groupKey）
	resFields := make([]*v1.SDevicePointField, 0, len(uniqueFields))
	for _, f := range uniqueFields {
		if f == nil {
			continue
		}
		var groupKey string
		if f.Group != nil {
			groupKey = f.Group.GroupKey
		}
		resFields = append(resFields, &v1.SDevicePointField{
			Key:                f.Key,
			Name:               f.Name,
			GroupKey:           groupKey,
			ValueType:          f.ValueType,
			ComponentType:      f.ComponentType,
			Step:               f.Step,
			Required:           f.Required,
			Unit:               f.Unit,
			Min:                f.Min,
			Max:                f.Max,
			Default:            f.Default,
			ParamExplain:       f.ParamExplain,
			Regex:              f.Regex,
			RegexFailedMessage: f.RegexFailedMessage,
			Description:        f.Description,
		})
	}

	return &v1.GetDevicePointsDefinitionRes{Groups: groups, Fields: resFields}, nil
}

// getVirtualDevicePointsDefinition 获取虚拟设备点位定义
func (c *ControllerV1) getVirtualDevicePointsDefinition(device c_base.IDevice) (*v1.GetDevicePointsDefinitionRes, error) {
	deviceConfig := device.GetConfig()
	if deviceConfig == nil {
		return &v1.GetDevicePointsDefinitionRes{Groups: []*c_base.SPointGroup{}, Fields: []*v1.SDevicePointField{}}, nil
	}

	var allFields []*c_base.SFieldDefinition

	// 遍历所有子设备
	for _, childConfig := range deviceConfig.ChildDeviceConfig {
		childDevice := common.GetDeviceManager().GetDeviceById(childConfig.Id)
		if childDevice == nil {
			continue
		}

		// 获取子设备的遥测点位
		telemetryPoints := childDevice.GetTelemetryPoints()
		deviceName := childConfig.Name

		for _, point := range telemetryPoints {
			if point == nil || point.IsHidden() {
				continue
			}

			// 转换为FieldDefinition
			if fd := c_base.ConvertIPointToFieldDefinition(point); fd != nil {
				// 设置分组信息：设备名:汇总
				fd.Group = &c_base.SPointGroup{GroupKey: fmt.Sprintf("%s:汇总", deviceName), GroupName: fmt.Sprintf("%s:汇总", deviceName), GroupSort: 0}
				allFields = append(allFields, fd)
			}
		}
	}

	// 基于 group+key 去重
	uniqueFields := deduplicateFieldsByGroupAndKey(allFields)

	// 构建分组列表并排序
	groupMap := make(map[string]*c_base.SPointGroup)
	for _, f := range uniqueFields {
		if f != nil && f.Group != nil {
			if _, ok := groupMap[f.Group.GroupKey]; !ok {
				g := *f.Group
				groupMap[g.GroupKey] = &g
			}
		}
	}
	groups := make([]*c_base.SPointGroup, 0, len(groupMap))
	for _, g := range groupMap {
		groups = append(groups, g)
	}
	sort.Slice(groups, func(i, j int) bool { return groups[i].GroupSort < groups[j].GroupSort })

	// 转换字段
	resFields := make([]*v1.SDevicePointField, 0, len(uniqueFields))
	for _, f := range uniqueFields {
		if f == nil {
			continue
		}
		var groupKey string
		if f.Group != nil {
			groupKey = f.Group.GroupKey
		}
		resFields = append(resFields, &v1.SDevicePointField{
			Key:                f.Key,
			Name:               f.Name,
			GroupKey:           groupKey,
			ValueType:          f.ValueType,
			ComponentType:      f.ComponentType,
			Step:               f.Step,
			Required:           f.Required,
			Unit:               f.Unit,
			Min:                f.Min,
			Max:                f.Max,
			Default:            f.Default,
			ParamExplain:       f.ParamExplain,
			Regex:              f.Regex,
			RegexFailedMessage: f.RegexFailedMessage,
			Description:        f.Description,
		})
	}

	return &v1.GetDevicePointsDefinitionRes{Groups: groups, Fields: resFields, IsVirtualDevice: true}, nil
}

// deduplicateFieldsByGroupAndKey 根据 group 与 key 进行去重，保留首次出现的字段
func deduplicateFieldsByGroupAndKey(fields []*c_base.SFieldDefinition) []*c_base.SFieldDefinition {
	if len(fields) <= 1 {
		return fields
	}

	seen := make(map[string]struct{}, len(fields))
	unique := make([]*c_base.SFieldDefinition, 0, len(fields))
	for _, f := range fields {
		if f == nil {
			continue
		}
		var groupKey string
		if f.Group != nil {
			groupKey = f.Group.GroupKey
		}
		key := groupKey + "\x00" + f.Key
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		unique = append(unique, f)
	}
	return unique
}

package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
	"fmt"
	"sort"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {
	if common.GetDeviceManager().Status() == c_enum.EStateInit {
		// 系统还在初始化中
		return &v1.GetRealDeviceCacheRes{}, nil
	}

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		return &v1.GetRealDeviceCacheRes{}, nil
	}

	groupCacheMap := make(map[string]*entity.SSingleDeviceGroup)

	list := device.GetPointValueList()
	//var list = make([]*c_base.SPointValue, 0)
	list = append(list, c_base.GetAllTelemetryPoint(device)...)

	for _, v := range list {

		groupName := ""

		// 如果点位的设备ID和当前请求的设备ID不一样，说明是子设备。名称修改为设备名+组名
		if req.DeviceId != v.GetDeviceId() {
			groupName = common.GetDeviceManager().GetDeviceNameById(v.GetDeviceId())
		}
		if v.IPoint == nil {
			continue
		}

		if v.IPoint.GetGroup() != nil {
			if groupName != "" {
				groupName = fmt.Sprintf("%s:%s", groupName, v.IPoint.GetGroup().GroupName)
			} else {
				groupName = v.IPoint.GetGroup().GroupName
			}
		} else {
			if groupName != "" {
				groupName = fmt.Sprintf("%s:其他", groupName)
			} else {
				groupName = "其他"
			}
		}

		// 缓存group
		group, exist := groupCacheMap[groupName]
		if !exist {
			groupSort := 99999
			if v.IPoint.GetGroup() != nil {
				groupSort = v.IPoint.GetGroup().GroupSort
			}

			group = &entity.SSingleDeviceGroup{
				GroupName: groupName,
				GroupSort: groupSort,
				Values:    make([]*entity.SSingleDeviceValue, 0),
			}
			groupCacheMap[groupName] = group
		}
		d := &entity.SSingleDeviceValue{}
		_ = gconv.Scan(v, d)

		// 获取状态解释
		if explain, err := v.IPoint.GetValueExplain(v.GetValue()); err == nil && explain != "" {
			d.StatueExplain = explain
		}
		group.Values = append(group.Values, d)
	}

	// 排序
	groups := make([]*entity.SSingleDeviceGroup, 0, len(groupCacheMap))
	for _, group := range groupCacheMap {
		groups = append(groups, group)
	}

	sort.Slice(groups, func(i, j int) bool {
		if groups[i].GroupSort == groups[j].GroupSort {
			return groups[i].GroupName < groups[j].GroupName
		}
		return groups[i].GroupSort < groups[j].GroupSort
	})

	return &v1.GetRealDeviceCacheRes{
		DeviceServerState: device.GetProtocolStatus().String(),
		AlarmLevel:        device.GetAlarmLevel().String(),
		LastUpdateTime:    device.GetLastUpdateTime(),
		Groups:            groups,
	}, nil
}

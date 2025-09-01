package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"context"
	"fmt"
	"sort"

	"github.com/gogf/gf/v2/util/gconv"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {
	if common.GetDeviceManager().Status() == c_base.EStateInit {
		// 系统还在初始化中
		return &v1.GetRealDeviceCacheRes{}, nil
	}

	device := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if device == nil {
		return &v1.GetRealDeviceCacheRes{}, nil
	}

	groupCacheMap := make(map[string]*entity.SSingleDeviceGroup)
	list := device.GetMetaValueList()
	for _, v := range list {

		groupName := ""

		// 如果点位的设备ID和当前请求的设备ID不一样，说明是子设备。名称修改为设备名+组名
		if req.DeviceId != v.DeviceId {
			groupName = common.GetDeviceManager().GetDeviceNameById(v.DeviceId)
		}

		if v.Meta.Group != nil {
			if groupName != "" {
				groupName = fmt.Sprintf("%s:%s", groupName, v.Meta.Group.GroupName)
			} else {
				groupName = v.Meta.Group.GroupName
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
			if v.Meta.Group != nil {
				groupSort = v.Meta.Group.GroupSort
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
		if v.Meta.SystemType == c_base.SUseReadType {
			d.Meta.SystemType = d.Meta.ReadType
		}
		if v.Meta.StatusExplain != nil {
			d.StatueExplain = v.Meta.StatusExplain(v.Value)
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

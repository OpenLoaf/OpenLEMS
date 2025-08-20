package device

import (
	v1 "application/api/device/v1"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	"sort"
)

func (c *ControllerV1) GetRealDeviceCache(ctx context.Context, req *v1.GetRealDeviceCacheReq) (res *v1.GetRealDeviceCacheRes, err error) {

	deviceWrapper := common.GetDeviceManager().GetDeviceById(req.DeviceId)
	if deviceWrapper == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}
	deviceInstance := deviceWrapper.GetDeviceInstance()
	if deviceInstance == nil {
		return nil, gerror.NewCode(gcode.CodeNotFound)
	}

	var alarmLevel = c_base.ENone
	if deviceWrapper.GetDeviceState() != c_base.EStateError &&
		deviceWrapper.GetDeviceState() != c_base.EStateInit {
		// todo 修改
		alarmLevel = deviceInstance.GetAlarmLevel()
	}

	res = &v1.GetRealDeviceCacheRes{
		DeviceServerState: deviceWrapper.GetDeviceState().String(),
		AlarmLevel:        alarmLevel.String(),
	}
	res.LastUpdateTime = deviceInstance.GetLastUpdateTime()

	//driverDescription := deviceInstance.GetDriverDescription()

	//for _, t := range driverDescription.Telemetry {
	//
	//}
	//
	//driverDescription.GetAllTelemetry(deviceInstance)

	groupCacheMap := make(map[string]*entity.SSingleDeviceGroup)

	for _, v := range deviceInstance.GetMetaValueList() {
		groupName := ""
		if v.Meta.Group != nil {
			groupName = v.Meta.Group.GroupName
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
		return groups[i].GroupSort < groups[j].GroupSort
	})

	res.Groups = groups
	return res, nil
}

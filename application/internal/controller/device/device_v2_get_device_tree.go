package device

import (
	v2 "application/api/device/v2"
	"application/internal/model/entity"
	"common"
	"common/c_base"
	"context"
	"s_db"
)

func (c *ControllerV2) GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error) {

	parentId := "0"
	if req.ActiveRootOnly {
		parentId = s_db.GetDeviceService().GetRootDeviceId()
	}

	// 从数据库中获取设备列表
	deviceList, err := s_db.GetDeviceService().GetDeviceConfigs(ctx, parentId)
	if err != nil {
		return nil, err
	}

	deviceTree := BuildDeviceTree(deviceList)
	return &v2.GetDeviceTreeRes{
		DeviceTree: deviceTree,
	}, nil
}

// BuildDeviceTree 将设备列表构造成树结构
func BuildDeviceTree(devices []*c_base.SDeviceConfig) []*entity.SDeviceTree {
	// 建立索引
	deviceMap := make(map[string]*entity.SDeviceTree, len(devices))
	var allKeys = make([]string, 0, len(devices))
	for _, d := range devices {

		device := common.GetDeviceManager().GetDeviceById(d.Id)
		var deviceTree = &entity.SDeviceTree{
			Children: make([]*entity.SDeviceTree, 0),
		}
		if device != nil {
			deviceTree.SDeviceDetail = device.GetDeviceDetail()
		} else {
			deviceTree.SDeviceDetail = c_base.NewDeviceDetailNoInstance(d, nil, nil)
		}

		deviceMap[d.Id] = deviceTree

		allKeys = append(allKeys, d.Id)
	}

	// 建立父子关系
	for _, key := range allKeys {
		node := deviceMap[key]
		// 为每个节点寻找父节点
		if parentNode, exist := deviceMap[node.Pid]; node.Pid != key && exist {
			parentNode.Children = append(parentNode.Children, node)
		}
	}

	// 找出所有根节点（没有父节点的节点）
	var roots []*entity.SDeviceTree
	for _, key := range allKeys {
		node := deviceMap[key]
		if _, exist := deviceMap[node.Pid]; node.Pid != key && !exist {
			roots = append(roots, node)
		}
	}

	return roots
}

package device

import (
	"common"
	"context"
	"s_db"
	"s_db/s_db_model"

	"github.com/gogf/gf/v2/frame/g"

	v2 "application/api/device/v2"
	"application/internal/model/entity"
)

func (c *ControllerV2) GetDeviceTree(ctx context.Context, req *v2.GetDeviceTreeReq) (res *v2.GetDeviceTreeRes, err error) {

	parentId := "0"
	if req.ActiveRootOnly {
		parentId = s_db.GetDeviceService().GetRootDeviceId()
	}

	// 从数据库中获取设备列表
	deviceList, err := s_db.GetDeviceService().GetDeviceList(ctx, parentId)
	if err != nil {
		return nil, err
	}

	deviceTree := BuildDeviceTree(ctx, req.RunningOnly, deviceList)
	return &v2.GetDeviceTreeRes{
		DeviceTree: deviceTree,
	}, nil
}

// BuildDeviceTree 将设备列表构造成树结构
func BuildDeviceTree(ctx context.Context, runningOnly bool, devices []*s_db_model.SDeviceModel) []*entity.SDeviceTree {
	// 建立索引
	idToDevice := make(map[string]*s_db_model.SDeviceModel, len(devices))
	pidToChildren := make(map[string][]*s_db_model.SDeviceModel, len(devices))
	for _, d := range devices {
		idToDevice[d.Id] = d
		pidToChildren[d.Pid] = append(pidToChildren[d.Pid], d)
	}

	// 根节点：pid 不存在于 idToDevice 或者 pid 为空/"0"
	var roots []*s_db_model.SDeviceModel
	for _, d := range devices {
		if d.Pid == "" || d.Pid == "0" || idToDevice[d.Pid] == nil {
			roots = append(roots, d)
		}
	}

	var buildNode func(dev *s_db_model.SDeviceModel) *entity.SDeviceTree
	buildNode = func(dev *s_db_model.SDeviceModel) *entity.SDeviceTree {
		// 获取设备参数
		if _, err := dev.GetParamsMap(); err != nil {
			g.Log().Errorf(context.Background(), "获取设备参数失败 - 设备ID: %s, 设备名称: %s, 参数原始值: %s, 错误: %v", dev.Id, dev.Name, dev.Params, err)
			return nil
		}

		isRunning := false
		lastUpdateTime := ""
		if d := common.GetRunningDeviceById(dev.Id); d != nil {
			isRunning = true
			if t := d.GetLastUpdateTime(); t != nil {
				lastUpdateTime = t.Format("2006-01-02 15:04:05")
			}
		}

		if runningOnly && !isRunning {
			return nil
		}

		node := &entity.SDeviceTree{
			DeviceId:       dev.Id,
			ProtocolId:     dev.ProtocolId,
			DevicePid:      dev.Pid,
			DeviceName:     dev.Name,
			DeviceDriver:   dev.Driver,
			LogLevel:       dev.LogLevel,
			Enable:         dev.Enable,
			Sort:           dev.Sort,
			IsRunning:      isRunning,
			LastUpdateTime: lastUpdateTime,
		}

		// 子节点
		for _, cm := range pidToChildren[dev.Id] {
			if child := buildNode(cm); child != nil {
				node.DeviceChildren = append(node.DeviceChildren, child)
			}
		}
		if len(node.DeviceChildren) == 0 {
			node.DeviceChildren = nil
		}
		return node
	}

	var tree []*entity.SDeviceTree
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

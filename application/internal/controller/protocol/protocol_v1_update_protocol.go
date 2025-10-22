package protocol

import (
	"common"
	"common/c_log"
	"context"
	"s_db"

	v1 "application/api/protocol/v1"
)

func (c *ControllerV1) UpdateProtocol(ctx context.Context, req *v1.UpdateProtocolReq) (res *v1.UpdateProtocolRes, err error) {
	// 记录业务操作开始
	c_log.BizInfo(ctx, "协议更新操作开始", "protocolId", req.ProtocolId)

	// 将请求参数构造为 map
	updateData := map[string]interface{}{
		"protocolName":     req.ProtocolName,
		"protocolType":     req.ProtocolType,
		"protocolAddress":  req.ProtocolAddress,
		"protocolPort":     req.ProtocolPort,
		"protocolTimeout":  req.ProtocolTimeout,
		"protocolLogLevel": req.ProtocolLogLevel,
		"protocolParams":   req.ProtocolParams,
	}

	err = s_db.GetProtocolService().UpdateProtocol(ctx, req.ProtocolId, updateData)
	if err != nil {
		c_log.BizError(ctx, "协议更新失败", "protocolId", req.ProtocolId, "error", err)
		return nil, err
	}

	// 记录协议配置变更需要重启设备服务
	c_log.BizInfo(ctx, "协议配置变更需要重启设备服务", "protocolId", req.ProtocolId)

	// 重启设备管理器以应用协议配置更改
	common.GetDeviceManager().Restart()

	c_log.BizInfo(ctx, "设备服务重启完成", "protocolId", req.ProtocolId)

	// 记录业务操作成功
	c_log.BizInfo(ctx, "协议更新操作成功", "protocolId", req.ProtocolId)

	res = &v1.UpdateProtocolRes{
		ProtocolId: req.ProtocolId,
	}
	return res, nil
}

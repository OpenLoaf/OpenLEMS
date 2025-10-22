package protocol

import (
	"common"
	"common/c_log"
	"context"
	"s_db"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "application/api/protocol/v1"

	"github.com/gogf/gf/v2/errors/gcode"
)

func (c *ControllerV1) DeleteProtocol(ctx context.Context, req *v1.DeleteProtocolReq) (res *v1.DeleteProtocolRes, err error) {
	// 记录业务操作开始
	c_log.BizInfo(ctx, "协议删除操作开始", "protocolId", req.ProtocolId)

	// 删除协议
	err = s_db.GetProtocolService().DeleteProtocol(ctx, req.ProtocolId)
	if err != nil {
		c_log.BizError(ctx, "协议删除失败", "protocolId", req.ProtocolId, "error", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "删除协议失败")
	}

	// 记录协议删除需要重启设备服务
	c_log.BizInfo(ctx, "协议删除需要重启设备服务", "protocolId", req.ProtocolId)

	// 重启设备管理器以应用协议删除更改
	common.GetDeviceManager().Restart()

	c_log.BizInfo(ctx, "设备服务重启完成", "protocolId", req.ProtocolId)

	// 记录业务操作成功
	c_log.BizInfo(ctx, "协议删除操作成功", "protocolId", req.ProtocolId)

	res = &v1.DeleteProtocolRes{
		ProtocolId: req.ProtocolId,
	}
	return res, nil
}

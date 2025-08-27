package protocol

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"s_db"

	v1 "application/api/protocol/v1"
	"github.com/gogf/gf/v2/errors/gcode"
)

func (c *ControllerV1) DeleteProtocol(ctx context.Context, req *v1.DeleteProtocolReq) (res *v1.DeleteProtocolRes, err error) {

	// 删除协议
	err = s_db.GetProtocolService().DeleteProtocol(ctx, req.ProtocolId)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "删除协议失败")
	}

	res = &v1.DeleteProtocolRes{
		ProtocolId: req.ProtocolId,
	}
	return res, nil
}

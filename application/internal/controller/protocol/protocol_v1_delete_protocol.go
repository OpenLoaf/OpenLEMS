package protocol

import (
	"context"
	"s_db"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "application/api/protocol/v1"
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

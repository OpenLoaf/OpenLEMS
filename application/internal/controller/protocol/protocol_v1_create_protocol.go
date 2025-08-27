package protocol

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"s_db"

	v1 "application/api/protocol/v1"
	"github.com/gogf/gf/v2/errors/gcode"
)

func (c *ControllerV1) CreateProtocol(ctx context.Context, req *v1.CreateProtocolReq) (res *v1.CreateProtocolRes, err error) {
	// 将请求数据转换为map
	data := map[string]interface{}{
		"protocolName":     req.ProtocolName,
		"protocolType":     req.ProtocolType,
		"protocolAddress":  req.ProtocolAddress,
		"protocolPort":     req.ProtocolPort,
		"protocolTimeout":  req.ProtocolTimeout,
		"protocolLogLevel": req.ProtocolLogLevel,
		"protocolParams":   req.ProtocolParams,
	}

	// 创建协议
	protocolId, err := s_db.GetProtocolService().CreateProtocol(ctx, data)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "创建协议失败")
	}

	return &v1.CreateProtocolRes{
		ProtocolId: protocolId,
	}, nil
}

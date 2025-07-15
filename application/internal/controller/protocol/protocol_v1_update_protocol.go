package protocol

import (
	"context"

	v1 "application/api/protocol/v1"
	"sqlite"
)

func (c *ControllerV1) UpdateProtocol(ctx context.Context, req *v1.UpdateProtocolReq) (res *v1.UpdateProtocolRes, err error) {
	protocolManage := sqlite.NewProtocolManage(ctx)

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

	err = protocolManage.UpdateProtocol(ctx, req.ProtocolId, updateData)
	if err != nil {
		return nil, err
	}

	res = &v1.UpdateProtocolRes{
		ProtocolId: req.ProtocolId,
	}
	return res, nil
}

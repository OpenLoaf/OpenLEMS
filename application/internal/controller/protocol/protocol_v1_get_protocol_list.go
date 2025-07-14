package protocol

import (
	"context"

	v1 "application/api/protocol/v1"
	"application/internal/model/entity"
	"sqlite"
)

func (c *ControllerV1) GetProtocolList(ctx context.Context, req *v1.GetProtocolListReq) (res *v1.GetProtocolListRes, err error) {
	protocolManage := sqlite.NewProtocolManage(ctx)
	protocols, err := protocolManage.GetProtocolList(ctx, req.Type)
	if err != nil {
		return nil, err
	}

	// 转换为 entity.SProtocol 类型
	var entityProtocols []*entity.SProtocol
	for _, protocol := range protocols {
		entityProtocols = append(entityProtocols, &entity.SProtocol{
			ProtocolId:       protocol.Id,
			ProtocolName:     protocol.Name,
			ProtocolType:     protocol.Type,
			ProtocolAddress:  protocol.Address,
			ProtocolTimeout:  int(protocol.Timeout),
			ProtocolLogLevel: protocol.LogLevel,
			ProtocolParams:   protocol.Params,
		})
	}

	res = &v1.GetProtocolListRes{
		ProtocolList: entityProtocols,
		Total:        len(entityProtocols),
	}
	return res, nil
}

package protocol

import (
	"application/internal/model/entity"
	"context"
	"net"
	"s_db"
	"strconv"
	"strings"

	v1 "application/api/protocol/v1"
)

func (c *ControllerV1) GetProtocolList(ctx context.Context, req *v1.GetProtocolListReq) (res *v1.GetProtocolListRes, err error) {
	protocols, err := s_db.GetProtocolService().GetProtocolList(ctx, req.Type)
	if err != nil {
		return nil, err
	}

	// 转换为 entity.SProtocol 类型
	var entityProtocols []*entity.SProtocol
	for _, protocol := range protocols {
		// 安全地解析地址和端口
		address, port := parseAddressAndPort(protocol.Address)

		//protocolActive := common.GetDeviceManager().IsProtocolActive(protocol.Id)

		entityProtocols = append(entityProtocols, &entity.SProtocol{
			ProtocolId:       protocol.Id,
			ProtocolName:     protocol.Name,
			ProtocolType:     protocol.Type,
			ProtocolAddress:  address,
			ProtocolPort:     port,
			ProtocolTimeout:  int(protocol.Timeout),
			ProtocolLogLevel: protocol.LogLevel,
			ProtocolParams:   protocol.Params,
			ProtocolActive:   true, // todo 完善
		})
	}

	res = &v1.GetProtocolListRes{
		ProtocolList: entityProtocols,
		Total:        len(entityProtocols),
	}
	return res, nil
}

// parseAddressAndPort 安全地解析地址和端口
func parseAddressAndPort(addressStr string) (string, int) {
	// 处理空字符串
	if addressStr == "" {
		return "", 0
	}

	// 尝试使用 net.SplitHostPort 解析（支持IPv6）
	host, portStr, err := net.SplitHostPort(addressStr)
	if err != nil {
		// 如果解析失败，可能是没有端口号的情况
		// 检查是否包含冒号
		if strings.Contains(addressStr, ":") {
			// 包含冒号但解析失败，可能是格式错误
			// 尝试简单的分割作为后备方案
			parts := strings.Split(addressStr, ":")
			if len(parts) >= 2 {
				// 取第一部分作为地址，最后一部分作为端口
				host = strings.Join(parts[:len(parts)-1], ":")
				portStr = parts[len(parts)-1]
			} else {
				// 只有一个部分，当作地址处理
				return addressStr, 0
			}
		} else {
			// 没有冒号，整个字符串当作地址
			return addressStr, 0
		}
	}

	// 解析端口号
	port, err := strconv.Atoi(portStr)
	if err != nil {
		// 端口号解析失败，返回0
		return host, 0
	}

	return host, port
}

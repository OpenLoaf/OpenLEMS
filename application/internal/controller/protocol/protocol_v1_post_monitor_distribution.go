package protocol

import (
	v1 "application/api/protocol/v1"
	"common"
	"context"
	"s_db"
)

// PostProtocolMonitorDistribution 获取协议状态分布数据
func (c *ControllerV1) PostProtocolMonitorDistribution(ctx context.Context, req *v1.PostProtocolMonitorDistributionReq) (res *v1.PostProtocolMonitorDistributionRes, err error) {
	// 获取协议列表
	protocols, err := s_db.GetProtocolService().GetProtocolList(ctx, "")
	if err != nil {
		return nil, err
	}

	// 应用筛选条件
	filteredProtocols := protocols
	if len(req.ProtocolTypes) > 0 {
		filteredProtocols = filterProtocolsByTypes(filteredProtocols, req.ProtocolTypes)
	}

	// 统计协议类型分布
	typeDistribution := make(map[string]int)
	statusDistribution := v1.ProtocolStatusDistribution{}

	for _, protocol := range filteredProtocols {
		// 统计协议类型分布
		typeDistribution[protocol.Type]++

		// 统计协议状态分布 - 使用 IsProtocolActive 判断协议是否激活
		if common.GetDeviceManager().IsProtocolActive(protocol.Id) {
			statusDistribution.Active++
		} else {
			statusDistribution.Inactive++
		}
	}

	return &v1.PostProtocolMonitorDistributionRes{
		TypeDistribution:   typeDistribution,
		StatusDistribution: statusDistribution,
	}, nil
}

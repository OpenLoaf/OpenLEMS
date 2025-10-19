package protocol

import (
	v1 "application/api/protocol/v1"
	"common"
	"common/c_base"
	"context"
	"s_db"
	"s_db/s_db_model"
	"s_storage"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
)

// PostProtocolMonitorOverview 获取协议监控概览数据
func (c *ControllerV1) PostProtocolMonitorOverview(ctx context.Context, req *v1.PostProtocolMonitorOverviewReq) (res *v1.PostProtocolMonitorOverviewRes, err error) {
	// 获取协议列表
	protocols, err := s_db.GetProtocolService().GetProtocolList(ctx, "")
	if err != nil {
		return nil, err
	}

	// 应用筛选条件
	filteredProtocols := protocols
	if len(req.ProtocolIds) > 0 {
		filteredProtocols = filterProtocolsByIds(filteredProtocols, req.ProtocolIds)
	}
	if len(req.ProtocolTypes) > 0 {
		filteredProtocols = filterProtocolsByTypes(filteredProtocols, req.ProtocolTypes)
	}

	// 计算统计数据
	totalProtocols := len(filteredProtocols)
	activeProtocols := 0
	abnormalProtocols := 0
	var totalSuccessRate float64
	var totalResponseMs float64
	validProtocols := 0

	// 获取最新性能指标
	now := time.Now().UnixMilli()
	startTime := now - int64(5*60*1000) // 最近5分钟
	endTime := now

	for _, protocol := range filteredProtocols {
		// 使用 IsProtocolActive 判断协议是否激活
		if common.GetDeviceManager().IsProtocolActive(protocol.Id) {
			activeProtocols++
		}

		// 查询协议性能指标
		chart, err := common.GetStorageInstance().GetStorageData(
			c_base.StorageTypeProtocol,
			protocol.Id,
			[]string{s_storage.ProtocolMetricSuccessRate, s_storage.ProtocolMetricAvgResponseMs},
			&startTime,
			&endTime,
			60000, // 1分钟步长
		)
		if err != nil || chart == nil || len(chart.Series) == 0 {
			continue
		}

		// 计算该协议的平均指标
		protocolSuccessRate := 0.0
		protocolResponseMs := 0.0
		hasData := false

		for _, series := range chart.Series {
			if len(series.Data) > 0 {
				hasData = true
				if series.Name == s_storage.ProtocolMetricSuccessRate {
					// 取最后一个值
					if rate := gconv.Float64(series.Data[len(series.Data)-1]); rate > 0 {
						protocolSuccessRate = rate
					}
				} else if series.Name == s_storage.ProtocolMetricAvgResponseMs {
					// 取最后一个值
					if ms := gconv.Float64(series.Data[len(series.Data)-1]); ms > 0 {
						protocolResponseMs = ms
					}
				}
			}
		}

		if hasData {
			totalSuccessRate += protocolSuccessRate
			totalResponseMs += protocolResponseMs
			validProtocols++

			// 判断是否为异常协议（成功率低于90%或响应时间超过1000ms）
			if protocolSuccessRate < 90.0 || protocolResponseMs > 1000.0 {
				abnormalProtocols++
			}
		}
	}

	// 计算平均值
	avgSuccessRate := 0.0
	avgResponseMs := 0.0
	if validProtocols > 0 {
		avgSuccessRate = totalSuccessRate / float64(validProtocols)
		avgResponseMs = totalResponseMs / float64(validProtocols)
	}

	return &v1.PostProtocolMonitorOverviewRes{
		TotalProtocols:    totalProtocols,
		ActiveProtocols:   activeProtocols,
		AvgSuccessRate:    avgSuccessRate,
		AvgResponseMs:     avgResponseMs,
		AbnormalProtocols: abnormalProtocols,
	}, nil
}

// filterProtocolsByIds 按协议ID筛选
func filterProtocolsByIds(protocols []*s_db_model.SProtocolModel, ids []string) []*s_db_model.SProtocolModel {
	var filtered []*s_db_model.SProtocolModel
	idMap := make(map[string]bool)
	for _, id := range ids {
		idMap[id] = true
	}

	for _, protocol := range protocols {
		if idMap[protocol.Id] {
			filtered = append(filtered, protocol)
		}
	}
	return filtered
}

// filterProtocolsByTypes 按协议类型筛选
func filterProtocolsByTypes(protocols []*s_db_model.SProtocolModel, types []string) []*s_db_model.SProtocolModel {
	var filtered []*s_db_model.SProtocolModel
	typeMap := make(map[string]bool)
	for _, t := range types {
		typeMap[t] = true
	}

	for _, protocol := range protocols {
		if typeMap[protocol.Type] {
			filtered = append(filtered, protocol)
		}
	}
	return filtered
}

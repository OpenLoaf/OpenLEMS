package protocol

import (
	"context"

	v1 "application/api/protocol/v1"
)

// GetGpioSysfsScan 扫描所有 GPIO 协议（type = gpio）的 SourceAddress，并调用 ScanGpioSysfs 聚合返回
func (c *ControllerV1) GetGpioSysfsScan(ctx context.Context, req *v1.GetGpioSysfsScanReq) (res *v1.GetGpioSysfsScanRes, err error) {
	// 获取所有 GPIO 协议配置
	//protocols, err := s_db.GetProtocolService().GetProtocolList(ctx, string(c_enum.EGpioSysfs))
	//if err != nil {
	//	return nil, err
	//}

	result := &v1.GetGpioSysfsScanRes{Items: make([]*v1.GetGpioSysfsScanResItem, 0)}

	//for _, p := range protocols {
	//	// 协议 SourceAddress 作为 sysfs 根目录
	//	if p.Addr == "" {
	//		continue
	//	}
	//	scan, scanErr := gpio_sysfs.ScanGpioSysfs(ctx, p.Addr)
	//	if scanErr != nil {
	//		// 不中断整个流程，忽略该地址错误
	//		continue
	//	}
	//	result.Items = append(result.Items, &v1.GetGpioSysfsScanResItem{
	//		Root:  p.Addr,
	//		Chips: scan.Chips,
	//		Gpios: scan.Gpios,
	//	})
	//}

	return result, nil
}

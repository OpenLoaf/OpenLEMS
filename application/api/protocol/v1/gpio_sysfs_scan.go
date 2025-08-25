package v1

import (
	//"gpio_sysfs/p_gpio_sysfs"

	"github.com/gogf/gf/v2/frame/g"
)

// GetGpioSysfsScanReq 请求：扫描所有 GPIO 协议配置下的 sysfs 信息
type GetGpioSysfsScanReq struct {
	g.Meta `path:"/protocol/gpio/sysfs/scan" method:"get" tags:"协议相关" summary:"扫描GPIO Sysfs 信息（按协议地址）"`
}

// GetGpioSysfsScanResItem 单个根目录的扫描结果
type GetGpioSysfsScanResItem struct {
	Root string `json:"root" dc:"扫描根目录(协议Address)"`
	//Chips []*p_gpio_sysfs.SGpioChipInfo `json:"chips"`
	//Gpios []*p_gpio_sysfs.SGpioInfo     `json:"gpios"`
}

// GetGpioSysfsScanRes 响应：多根目录扫描结果
type GetGpioSysfsScanRes struct {
	Items []*GetGpioSysfsScanResItem `json:"items" dc:"按协议地址聚合的扫描结果"`
}

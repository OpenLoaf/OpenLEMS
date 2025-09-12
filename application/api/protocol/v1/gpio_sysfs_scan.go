package v1

import (
	//"gpio_sysfs/p_gpio_sysfs"

	"github.com/gogf/gf/v2/frame/g"
)

// GetGpioSysfsScanReq 请求：扫描所有 GPIO 协议配置下的 sysfs 信息
type GetGpioSysfsScanReq struct {
	g.Meta `path:"/protocol/gpio/sysfs/scan" method:"get" tags:"协议相关" summary:"扫描GPIO Sysfs 信息（按协议地址）"`
}

// SGpioInfo GPIO信息结构体
type SGpioInfo struct {
	PinNumber  int    `json:"pinNumber" dc:"GPIO引脚编号"`
	Direction  string `json:"direction" dc:"GPIO方向：in(输入)/out(输出)"`
	Value      int    `json:"value" dc:"GPIO当前值：0(低电平)/1(高电平)"`
	Active     bool   `json:"active" dc:"GPIO是否可用"`
	ProtocolId string `json:"protocolId" dc:"关联的协议ID"`
	DeviceId   string `json:"deviceId" dc:"关联的设备ID"`
	DeviceName string `json:"deviceName" dc:"关联的设备名称"`
}

// GetGpioSysfsScanRes 响应：多根目录扫描结果
type GetGpioSysfsScanRes struct {
	GpioList []*SGpioInfo `json:"gpioList" dc:"GPIO列表"`
	Total    int          `json:"total" dc:"GPIO总数"`
}

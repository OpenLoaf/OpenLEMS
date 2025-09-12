package internal

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
	"context"
	"p_base"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"go.einride.tech/can"
)

type CanbusProtocolProvider struct {
	c_base.IAlarm
	c_base.IProtocolCacheValue

	ctx  context.Context // 上下文
	once sync.Once

	deviceId string

	receiverChan    <-chan can.Frame       // 接收通道
	transmitterChan chan<- can.Frame       // 发送通道
	canTaskList     []*c_proto.SCanbusTask // 任务列表

	lastUpdateTime *time.Time // 最后更新时间
	deviceConfig   *c_base.SDeviceConfig
	protocolConfig *c_base.SProtocolConfig
	protocolStatus c_enum.EProtocolStatus // 协议状态，canbus只要初始化成功，就是连接成功

	//metricProtocol *sMetricProtocol // 统计协议
}

var _ c_proto.ICanbusProtocol = (*CanbusProtocolProvider)(nil)

func NewCanbusProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, receiverChan <-chan can.Frame, transmitterChan chan<- can.Frame) (c_proto.ICanbusProtocol, error) {
	if protocolConfig == nil {
		return nil, gerror.Newf("Canbus协议：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name)
	}
	if deviceConfig == nil {
		return nil, gerror.Newf("Canbus协议：%s 的设备配置不能为空！", protocolConfig.Id)
	}

	provider := &CanbusProtocolProvider{
		IProtocolCacheValue: p_base.NewGetProtocolCacheValue(ctx, deviceConfig.Id),
		IAlarm:              c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		deviceId:            deviceConfig.Id,

		once:            sync.Once{},
		ctx:             ctx,
		receiverChan:    receiverChan,
		transmitterChan: transmitterChan,
		canTaskList:     make([]*c_proto.SCanbusTask, 0),
		deviceConfig:    deviceConfig,
		protocolConfig:  protocolConfig,
		protocolStatus:  c_enum.EProtocolConnecting,
	}

	return provider, nil
}

func (c *CanbusProtocolProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	return c.protocolStatus
}

func (c *CanbusProtocolProvider) GetConfig() *c_base.SDeviceConfig {
	return c.deviceConfig
}

func (c *CanbusProtocolProvider) GetLastUpdateTime() *time.Time {
	return c.lastUpdateTime
}

package internal

import (
	"canbus/p_canbus"
	"common"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"go.einride.tech/can"
	"net"
	"sync"
	"time"
)

type CanbusProtocolProvider struct {
	*c_base.SAlarmHandler                 // 告警
	ctx                   context.Context // 上下文
	once                  sync.Once       // 只执行一次Init方法

	connect         net.Conn         // 链接
	receiverChan    <-chan can.Frame // 接收通道
	transmitterChan chan<- can.Frame // 发送通道

	cache              *gcache.Cache // 点位缓存
	log                *glog.Logger  // 日志
	canbusRwMutex      sync.RWMutex  // 读写锁
	lastUpdateTime     *time.Time    // 最后更新时间
	deviceType         c_base.EDeviceType
	deviceConfig       *c_base.SDriverConfig
	canbusDeviceConfig *p_canbus.SCanbusDeviceConfig
	protocolConfig     *c_base.SProtocolConfig

	//metricProtocol *sMetricProtocol // 统计协议
}

func (c *CanbusProtocolProvider) Init() {
	c.once.Do(func() {
		device := common.GetDeviceById(c.deviceConfig.Id)
		if device != nil {
			c.deviceType = device.GetDriverType()
		}

		go func() {
			for {
				select {
				case <-c.ctx.Done():
					return

				case frame := <-c.receiverChan:
					g.Log().Infof(c.ctx, "内部接收到数据：%v", frame)
				}
			}
		}()

	})

}

func (c *CanbusProtocolProvider) Close() {
	c.cache.Clear(c.Ctx)
}

func (c *CanbusProtocolProvider) IsActivate() bool {
	return true
}

func (c *CanbusProtocolProvider) GetMetaValueList() []*c_base.MetaValueWrapper {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetLastUpdateTime() *time.Time {
	return c.lastUpdateTime
}

func (c *CanbusProtocolProvider) GetDeviceConfig() *c_base.SDriverConfig {
	return c.deviceConfig
}

func (c *CanbusProtocolProvider) GetProtocolConfig() *c_base.SProtocolConfig {
	return c.protocolConfig
}

func (c *CanbusProtocolProvider) RegisterRead(ctx context.Context, group *p_canbus.SCanbusTask, gs ...*p_canbus.SCanbusTask) error {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetBool(meta *c_base.Meta) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetIntValue(meta *c_base.Meta) (int, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetInt32Value(meta *c_base.Meta) (int32, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetUintValue(meta *c_base.Meta) (uint, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetUint32Value(meta *c_base.Meta) (uint32, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetFloat32Value(meta *c_base.Meta) (float32, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetFloat32Values(metas ...*c_base.Meta) ([]float32, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetFloat64Value(meta *c_base.Meta) (float64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetFloat64Values(meta ...*c_base.Meta) ([]float64, error) {
	//TODO implement me
	panic("implement me")
}

func (c *CanbusProtocolProvider) GetCanbusDeviceConfig() *p_canbus.SCanbusDeviceConfig {
	return c.canbusDeviceConfig
}

func NewCanbusProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig, receiverChan <-chan can.Frame, transmitterChan chan<- can.Frame) (p_canbus.ICanbusProtocol, error) {
	if protocolConfig == nil {
		panic(gerror.Newf("Canbus协议：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name))
	}
	if deviceConfig == nil {
		panic(gerror.Newf("Canbus协议：%s 的设备配置不能为空！", protocolConfig.Id))
	}
	canbusDeviceConfig := &p_canbus.SCanbusDeviceConfig{}
	err := gconv.Scan(deviceConfig.Params, canbusDeviceConfig)
	if err != nil {
		panic(gerror.Newf("设备[%s]的Param参数配置错误：%v 无法转换为SCanbusDeviceConfig", deviceConfig.Id, err))
	}

	provider := &CanbusProtocolProvider{
		once:               sync.Once{},
		ctx:                ctx,
		deviceConfig:       deviceConfig,
		protocolConfig:     protocolConfig,
		canbusDeviceConfig: canbusDeviceConfig,
		receiverChan:       receiverChan,
		transmitterChan:    transmitterChan,

		//modbusReadChan:     make(chan *p_canbus.SCanbusTask),
		//preQuery: make(map[string]bool),
		cache: gcache.New(),
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),
		//metricProtocol: newMetricProtocol(ctx, protocolConfig, deviceConfig),
	}

	return provider, nil
}

package internal

import (
	"canbus/p_canbus"
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
	once                  sync.Once

	connect         net.Conn         // 链接
	receiverChan    <-chan can.Frame // 接收通道
	transmitterChan chan<- can.Frame // 发送通道
	//canTaskMap      map[uint32]*p_canbus.SCanbusTask // 任务map
	canTaskList []*p_canbus.SCanbusTask // 任务列表

	cache              *gcache.Cache // 点位缓存
	log                *glog.Logger  // 日志
	canbusRwMutex      sync.RWMutex  // 读写锁
	lastUpdateTime     *time.Time    // 最后更新时间
	deviceType         c_base.EDeviceType
	deviceConfig       *c_base.SDeviceConfig
	canbusDeviceConfig *p_canbus.SCanbusDeviceConfig
	protocolConfig     *c_base.SProtocolConfig

	protocolState c_base.EServerState
	//metricProtocol *sMetricProtocol // 统计协议
}

var _ c_base.IProtocol = (*CanbusProtocolProvider)(nil)

func (c *CanbusProtocolProvider) IsPhysical() bool {
	return true
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

func (c *CanbusProtocolProvider) GetDeviceConfig() *c_base.SDeviceConfig {
	return c.deviceConfig
}

func (c *CanbusProtocolProvider) GetProtocolConfig() *c_base.SProtocolConfig {
	return c.protocolConfig
}

func (c *CanbusProtocolProvider) GetCanbusDeviceConfig() *p_canbus.SCanbusDeviceConfig {
	return c.canbusDeviceConfig
}

func NewCanbusProvider(ctx context.Context, deviceType c_base.EDeviceType, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, receiverChan <-chan can.Frame, transmitterChan chan<- can.Frame) (p_canbus.ICanbusProtocol, error) {
	if protocolConfig == nil {
		return nil, gerror.Newf("Canbus协议：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name)
	}
	if deviceConfig == nil {
		return nil, gerror.Newf("Canbus协议：%s 的设备配置不能为空！", protocolConfig.Id)
	}
	canbusDeviceConfig := &p_canbus.SCanbusDeviceConfig{}
	err := gconv.Scan(deviceConfig.Params, canbusDeviceConfig)
	if err != nil {
		return nil, gerror.Wrapf(err, "设备[%s]的Param参数配置错误：%v 无法转换为SCanbusDeviceConfig", deviceConfig.Id)
	}

	provider := &CanbusProtocolProvider{
		once:               sync.Once{},
		ctx:                ctx,
		deviceType:         deviceType,
		deviceConfig:       deviceConfig,
		protocolConfig:     protocolConfig,
		canbusDeviceConfig: canbusDeviceConfig,
		receiverChan:       receiverChan,
		transmitterChan:    transmitterChan,
		//canTaskMap:         make(map[uint32]*p_canbus.SCanbusTask),
		canTaskList: make([]*p_canbus.SCanbusTask, 0),
		cache:       gcache.New(),
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),
		//metricProtocol: newMetricProtocol(ctx, protocolConfig, deviceConfig),
	}

	return provider, nil
}

package internal

import (
	"canbus/p_canbus"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/container/garray"
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
	// TODO 此处和其他的合并到一个处理方法中
	// 排序
	_sortValues := garray.NewSortedArray(func(v1, v2 interface{}) int {
		v1Meta := v1.(*c_base.MetaValueWrapper).Meta
		v2Meta := v2.(*c_base.MetaValueWrapper).Meta

		if v1Meta.Addr > v2Meta.Addr {
			return 1
		} else {
			if v1Meta.Addr == v2Meta.Addr {
				// 比对别的
				if v1Meta.ReadType > v2Meta.ReadType {
					return 1
				}
			}

			return -1
		}
	})

	metas, err := c.cache.Keys(c.Ctx)
	if err != nil {
		return nil
	}

	for _, meta := range metas {
		_varValue, err := c.cache.Get(c.Ctx, meta) // MetaValue类型
		if err != nil {
			continue
		}

		metaValue := &c_base.MetaValue{}
		err = _varValue.Structs(metaValue)
		if err != nil {
			g.Log().Errorf(c.ctx, "解析缓存值失败：%v", err)
			continue
		}

		_sortValues.Add(&c_base.MetaValueWrapper{
			DeviceId:   c.deviceConfig.Id,
			DeviceType: c.deviceType,
			Meta:       meta.(*c_base.Meta),
			Value:      metaValue.Value,
			HappenTime: metaValue.HappenTime,
		})
	}
	//_sortValues = _sortValues.Sort()

	result := make([]*c_base.MetaValueWrapper, _sortValues.Len())
	for i, v := range _sortValues.Slice() {
		result[i] = v.(*c_base.MetaValueWrapper)
	}

	return result
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

package internal

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
	"context"
	"net"
	"p_base"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"go.einride.tech/can"
)

type CanbusProtocolProvider struct {
	c_base.IAlarm
	c_base.IProtocolCacheValue

	ctx  context.Context // 上下文
	once sync.Once

	deviceId string

	connect         net.Conn               // 链接
	receiverChan    <-chan can.Frame       // 接收通道
	transmitterChan chan<- can.Frame       // 发送通道
	canTaskList     []*c_proto.SCanbusTask // 任务列表

	lastUpdateTime *time.Time // 最后更新时间
	deviceConfig   *c_base.SDeviceConfig
	protocolConfig *c_base.SProtocolConfig

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
	}

	return provider, nil
}

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

		// 先比较 Sort
		if v1Meta.Sort > v2Meta.Sort {
			return 1
		} else if v1Meta.Sort < v2Meta.Sort {
			return -1
		}

		// Sort 相等时，再比较 Addr
		if v1Meta.Addr > v2Meta.Addr {
			return 1
		} else if v1Meta.Addr < v2Meta.Addr {
			return -1
		}

		// Addr 相等时，比较 ReadType
		if v1Meta.ReadType > v2Meta.ReadType {
			return 1
		} else if v1Meta.ReadType < v2Meta.ReadType {
			return -1
		}

		return 0
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

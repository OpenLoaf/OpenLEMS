package internal

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"context"
	"p_base"
	"sync"
	"time"

	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/pkg/errors"
	"github.com/torykit/go-modbus"
)

type ModbusProtocolProvider struct {
	c_base.IAlarm
	c_base.IProtocolCacheValue

	ctx  context.Context // 上下文
	once sync.Once       // 只执行一次Init方法

	deviceId   string
	deviceType c_enum.EDeviceType

	client             modbus.Client   // modbus的通讯
	preQuery           map[string]bool // 预读
	cache              *gcache.Cache   // 点位缓存
	modbusRwMutex      sync.RWMutex    // 读写锁
	lastUpdateTime     *time.Time      // 最后更新时间
	modbusDeviceConfig *c_proto.SModbusDeviceConfig
	protocolConfig     *c_base.SProtocolConfig

	metricProtocol *sMetricProtocol // 统计协议
}

var _ c_proto.IModbusProtocol = (*ModbusProtocolProvider)(nil)

func NewModbusProvider(ctx context.Context, deviceType c_enum.EDeviceType, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig, client any) (c_proto.IModbusProtocol, error) {
	if protocolConfig == nil {
		panic(errors.Errorf("Modbus设备：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name))
	}
	if deviceConfig == nil {
		panic(errors.Errorf("modbus协议：%s 的设备配置不能为空！", protocolConfig.Id))
	}
	modbusDeviceConfig := &c_proto.SModbusDeviceConfig{}
	err := gconv.Scan(deviceConfig.Params, modbusDeviceConfig)
	if err != nil {
		panic(errors.Errorf("设备[%s]的Param参数配置错误：%v 无法转换为SModbusDeviceConfig", deviceConfig.Id, err))
	}
	if modbusDeviceConfig.UnitId == 0 {
		// unitId默认禁止为0
		panic(errors.Errorf("Modbus设备[%s]的UnitId不能为0", deviceConfig.Id))
	}
	cache := gcache.New()
	provider := &ModbusProtocolProvider{
		IProtocolCacheValue: p_base.NewGetProtocolCacheValue(ctx, deviceConfig.Id, deviceType, cache),
		IAlarm:              c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		deviceId:            deviceConfig.Id,
		deviceType:          deviceType,

		once:               sync.Once{},
		ctx:                ctx,
		protocolConfig:     protocolConfig,
		modbusDeviceConfig: modbusDeviceConfig,
		preQuery:           make(map[string]bool),
		cache:              cache,
		metricProtocol:     newMetricProtocol(ctx, protocolConfig, deviceConfig),
	}
	if client != nil {
		provider.client = client.(modbus.Client)
	} else {
		panic(errors.Errorf("modbus client is nil, please init the modbus client"))
	}

	//if deviceConfig.LogLevel != "" {
	//	err := provider.log.SetLevelStr(deviceConfig.LogLevel)
	//	if err != nil {
	//		provider.log.Level(glog.LEVEL_INFO)
	//	}
	//} else {
	//	// 设置默认的日志等级为上一级接口配置的日志等级
	//	err := provider.log.SetLevelStr(protocolConfig.GetLogLevel())
	//	if err != nil {
	//		provider.log.Level(glog.LEVEL_INFO)
	//	}
	//}

	return provider, nil
}

func (p *ModbusProtocolProvider) GetDeviceConfigFields() []*c_base.SConfigStructFields {
	modbusDeviceConfig := &c_proto.SModbusDeviceConfig{}
	fields, err := c_base.BuildConfigStructFields(modbusDeviceConfig)
	if err != nil {
		c_log.BizErrorf(p.ctx, "解析Modbus配置信息结构失败！")
		return nil
	}
	return fields
}

func (p *ModbusProtocolProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	if p.client == nil {
		return c_enum.EProtocolDisconnected
	}
	if p.client.IsConnected() {
		return c_enum.EProtocolConnected
	}
	return c_enum.EProtocolConnecting
}
func (p *ModbusProtocolProvider) GetLastUpdateTime() *time.Time {
	return p.lastUpdateTime
}

func (p *ModbusProtocolProvider) GetPointValueList() []*c_base.SPointValue {
	// 排序
	_sortValues := garray.NewSortedArray(func(v1, v2 interface{}) int {
		v1Meta := v1.(*c_base.SPointValue).IPoint.(*c_proto.SModbusPoint)
		v2Meta := v2.(*c_base.SPointValue).IPoint.(*c_proto.SModbusPoint)

		// 先比较 Sort
		if v1Meta.GetSort() > v2Meta.GetSort() {
			return 1
		} else if v1Meta.GetSort() < v2Meta.GetSort() {
			return -1
		}

		// Sort 相等时，再比较 Addr
		if v1Meta.Addr > v2Meta.Addr {
			return 1
		} else {
			if v1Meta.Addr == v2Meta.Addr {
				// 比对别的
				if v1Meta.DataAccess.ValueType > v2Meta.DataAccess.ValueType {
					return 1
				}
			}

			return -1
		}
	})

	metas, err := p.cache.Keys(p.ctx)
	if err != nil {
		return nil
	}

	for _, meta := range metas {
		_varValue, err := p.cache.Get(p.ctx, meta) // MetaValue类型
		if err != nil {
			continue
		}

		pointValue := &c_base.SPointValue{}
		err = _varValue.Structs(pointValue)
		if err != nil {
			g.Log().Errorf(p.ctx, "解析缓存值失败：%+v", err)
			continue
		}

		_sortValues.Add(pointValue)
	}
	//_sortValues = _sortValues.Sort()

	result := make([]*c_base.SPointValue, _sortValues.Len())
	for i, v := range _sortValues.Slice() {
		result[i] = v.(*c_base.SPointValue)
	}

	return result
}

package internal

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/torykit/go-modbus"
	"plug_protocol_modbus/p_modbus"
	"sync"
	"time"
)

type ModbusProtocolProvider struct {
	*c_base.SAlarmHandler                 // 告警
	ctx                   context.Context // 上下文
	once                  sync.Once       // 只执行一次Init方法
	//deviceId              string                     // 设备名称
	//unitId                uint8                      // 设备的unitId
	modbusReadChan chan *p_modbus.ModbusGroup // 查询用的通道
	client         modbus.Client              // modbus的通讯
	preQuery       map[string]bool            // 预读
	cache          *gcache.Cache              // 点位缓存
	log            *glog.Logger               // 日志
	modbusRwMutex  sync.RWMutex               // 读写锁
	lastUpdateTime *time.Time                 // 最后更新时间

	deviceConfig       *c_base.SDriverConfig
	modbusDeviceConfig *p_modbus.SModbusDeviceConfig
	protocolConfig     *c_base.SProtocolConfig
}

func NewModbusProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDriverConfig, client any) (p_modbus.IModbusProtocol, error) {
	if protocolConfig == nil {
		panic(gerror.Newf("GPIO设备：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name))
	}
	if deviceConfig == nil {
		panic(gerror.Newf("modbus协议：%s 的设备配置不能为空！", protocolConfig.Id))
	}
	modbusDeviceConfig := &p_modbus.SModbusDeviceConfig{}
	err := gconv.Scan(deviceConfig.Params, modbusDeviceConfig)
	if err != nil {
		panic(gerror.Newf("设备[%s]的Param参数配置错误：%v 无法转换为SModbusDeviceConfig", deviceConfig.Id, err))
	}
	if modbusDeviceConfig.UnitId == 0 {
		// unitId默认禁止为0
		panic(gerror.Newf("Modbus设备[%s]的UnitId不能为0", deviceConfig.Id))
	}

	provider := &ModbusProtocolProvider{
		once:               sync.Once{},
		ctx:                ctx,
		deviceConfig:       deviceConfig,
		protocolConfig:     protocolConfig,
		modbusDeviceConfig: modbusDeviceConfig,
		modbusReadChan:     make(chan *p_modbus.ModbusGroup),
		preQuery:           make(map[string]bool),
		cache:              gcache.New(),
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),
	}
	if client != nil {
		provider.client = client.(modbus.Client)
	} else {
		panic(gerror.New("modbus client is nil, please init the modbus client"))
	}

	if deviceConfig.LogLevel != "" {
		err := provider.log.SetLevelStr(deviceConfig.LogLevel)
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	} else {
		// 设置默认的日志等级为上一级接口配置的日志等级
		err := provider.log.SetLevelStr(protocolConfig.GetLogLevel())
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	}

	return provider, nil
}

func (p *ModbusProtocolProvider) GetDeviceConfig() *c_base.SDriverConfig {
	return p.deviceConfig
}

func (p *ModbusProtocolProvider) GetProtocolConfig() *c_base.SProtocolConfig {
	return p.protocolConfig
}

func (p *ModbusProtocolProvider) GetModbusDeviceConfig() *p_modbus.SModbusDeviceConfig {
	return p.modbusDeviceConfig
}

func (p *ModbusProtocolProvider) GetLastUpdateTime() *time.Time {
	return p.lastUpdateTime
}

func (p *ModbusProtocolProvider) IsActivate() bool {
	return p.client.IsConnected()
}

func (p *ModbusProtocolProvider) Close() error {
	return p.client.Close()
}

func (p *ModbusProtocolProvider) GetMetaValueList() []*c_base.MetaValueWrapper {
	// 排序
	_sortValues := garray.NewSortedArray(func(v1, v2 interface{}) int {

		if v1.(*c_base.MetaValueWrapper).Meta.Addr > v2.(*c_base.MetaValueWrapper).Meta.Addr {
			return 1
		} else {
			return -1
		}
	})

	metas, err := p.cache.Keys(p.Ctx)
	if err != nil {
		return nil
	}

	for _, meta := range metas {
		_varValue, err := p.cache.Get(p.Ctx, meta) // MetaValue类型
		if err != nil {
			continue
		}

		metaValue := &c_base.MetaValue{}
		err = _varValue.Structs(metaValue)
		if err != nil {
			g.Log().Errorf(p.ctx, "解析缓存值失败：%v", err)
			continue
		}

		_sortValues.Add(&c_base.MetaValueWrapper{
			DeviceId:   p.deviceConfig.Id,
			DeviceType: p.deviceConfig.Type,
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

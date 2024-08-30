package internal

import (
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/torykit/go-modbus"
	"plug_protocol_modbus/p_modbus"
	"sync"
	"time"
)

type ModbusProvider struct {
	ctx                   context.Context            // 上下文
	once                  sync.Once                  // 只执行一次Init方法
	deviceId              string                     // 设备名称
	deviceType            c_base.EDeviceType         // 设备类型
	unitId                uint8                      // 设备的unitId
	modbusReadChan        chan *p_modbus.ModbusGroup // 查询用的通道
	client                modbus.Client              // modbus的通讯
	preQuery              map[string]bool            // 预读
	cache                 *gcache.Cache              // 点位缓存
	log                   *glog.Logger               // 日志
	printCacheValue       bool                       // 打印缓存值
	modbusRwMutex         sync.RWMutex               // 读写锁
	lastUpdateTime        *time.Time                 // 最后更新时间
	*c_base.SAlarmHandler                            // 告警
}

func NewModbusProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *p_modbus.SModbusDeviceConfig, client any) (p_modbus.IModbusProtocol, error) {
	provider := &ModbusProvider{
		once:            sync.Once{},
		ctx:             ctx,
		deviceId:        deviceConfig.Id,
		unitId:          deviceConfig.UnitId,
		printCacheValue: deviceConfig.PrintCacheValue,
		modbusReadChan:  make(chan *p_modbus.ModbusGroup),
		preQuery:        make(map[string]bool),
		cache:           gcache.New(),
		SAlarmHandler: &c_base.SAlarmHandler{
			Ctx: ctx,
		},
		log: g.Log(deviceConfig.Id),
	}
	if client != nil {
		provider.client = client.(modbus.Client)
	} else {
		panic("modbus client is nil, please init the modbus client")
	}

	if deviceConfig.LogLevel != "" {
		err := provider.log.SetLevelStr(deviceConfig.LogLevel)
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	} else {
		// 设置默认的日志等级为上一级接口配置的日志等级
		err := provider.log.SetLevelStr(clientConfig.GetLogLevel())
		if err != nil {
			provider.log.Level(glog.LEVEL_INFO)
		}
	}

	return provider, nil
}

func (p *ModbusProvider) GetId() string {
	return p.deviceId
}

func (p *ModbusProvider) GetType() c_base.EDeviceType {
	return p.deviceType
}

func (p *ModbusProvider) GetCache() *gcache.Cache {
	return p.cache
}

func (p *ModbusProvider) GetLastUpdateTime() *time.Time {
	return p.lastUpdateTime
}

func (p *ModbusProvider) IsActivate() bool {
	return p.client.IsConnected()
}

func (p *ModbusProvider) Close() error {
	return p.client.Close()
}

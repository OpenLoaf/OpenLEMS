package driver

import (
	"context"
	common "ems-plan"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/torykit/go-modbus"
	"plug_protocol_gpio_sysfs"
	"plug_protocol_modbus"
	"reflect"
)

type DeviceCmd struct {
	ctx                  context.Context
	cancelFunc           context.CancelFunc
	modbusClientCache    map[string]modbus.Client
	driverCache          map[string]c_base.IDriver
	pluginNewMethodCache map[string]reflect.Method
}

func NewDeviceCmd(ctx context.Context) *DeviceCmd {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Driver")
	ctx, cancelFunc := context.WithCancel(ctx)
	return &DeviceCmd{
		ctx:         ctx,
		cancelFunc:  cancelFunc,
		driverCache: make(map[string]c_base.IDriver),
	}
}

func (d *DeviceCmd) Start() {
	// 捕捉panic
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(d.ctx, "启动驱动失败！原因：%v", err)
			d.Stop()
		}
	}()
	driver := d.InitDriver(d.ctx, common.GetDriverConfig(d.ctx))
	if driver == nil {
		g.Log().Warningf(d.ctx, "没有可用的驱动被加载！")
	}
}

func (d *DeviceCmd) Stop() {
	// 关闭所有client
	for _, client := range d.modbusClientCache {
	_:
		client.Close()
	}
	d.modbusClientCache = make(map[string]modbus.Client)
	d.cancelFunc()
}

func (d *DeviceCmd) InitDriver(ctx context.Context, config *c_base.SDriverConfig) c_base.IDriver {
	if err := config.Check(); err != nil {
		panic(err)
	}

	if config.Enable == false {
		g.Log().Noticef(ctx, "设备[%s]未启用！", config.Name)
		return nil
	}

	// 先递归初始化子设备
	if config.DeviceChildren != nil {
		for _, _device := range config.DeviceChildren {
			d.InitDriver(ctx, _device)
		}
	}

	var protocolProvider c_base.IProtocol
	// 设备初始化
	if config.ProtocolId != "" {
		protocolProvider = d.getProtocolProvider(ctx, config)
	} else {
		ctx = context.WithValue(ctx, c_base.ConstCtxKeyProtocolId, "Virtual")
	}

	driver := d.getDriver(ctx, config)
	driver.Init(protocolProvider, config)

	if protocolProvider != nil {
		protocolProvider.Init()
	}

	common.RegisterDevice(driver)
	return driver
}

func (d *DeviceCmd) getProtocolProvider(ctx context.Context, deviceConfig *c_base.SDriverConfig) c_base.IProtocol {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyProtocolId, deviceConfig.ProtocolId)

	protocolId := deviceConfig.ProtocolId

	// 从配置中获取协议
	protocolConfig := common.GetProtocolById(ctx, protocolId)
	if protocolConfig == nil {
		panic(gerror.Newf("协议Id: %s 配置信息不存在！", protocolId))
	}
	if err := protocolConfig.Check(); err != nil {
		panic(err)
	}

	// 初始化协议
	switch protocolConfig.GetProtocol() {
	case c_base.EModbusRtu, c_base.EModbusTcp:
		// 从缓存中获取client，如果没有就新建后放入缓存
		var client modbus.Client
		if _client, exist := d.modbusClientCache[protocolConfig.Id]; exist {
			client = _client
		} else {
			// 创建client
			client = plug_protocol_modbus.NewModbusClient(ctx, protocolConfig)
		}

		modbusProvider, err := plug_protocol_modbus.NewModbusProvider(ctx, protocolConfig, deviceConfig, client)
		if err != nil {
			panic(err)
		}

		return modbusProvider
	case c_base.ECanbusTcp:
	case c_base.ECanbus:
	case c_base.EGpioSysfs:
		gpioSysfsProtocol, err := plug_protocol_gpio_sysfs.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
		if err != nil {
			panic(err)
		}
		return gpioSysfsProtocol
	}

	panic(gerror.Newf("未知的协议类型：%s", protocolConfig.GetProtocol()))
}

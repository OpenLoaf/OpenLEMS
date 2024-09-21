package driver

import (
	"common"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/torykit/go-modbus"
	"gpio_sysfs"
	"influxdb_2"
	protocolModbus "modbus"
	"os"
)

type SDeviceCmd struct {
	ctx               context.Context
	cancelFunc        context.CancelFunc
	modbusClientCache map[string]modbus.Client
	//driverCache          map[string]c_base.IDriver
	//pluginNewMethodCache map[string]reflect.Method
}

func NewDeviceCmd(ctx context.Context) *SDeviceCmd {
	ctx = context.WithValue(ctx, c_base.ConstCtxKeyGroupName, "Driver")
	ctx, cancelFunc := context.WithCancel(ctx)
	return &SDeviceCmd{
		ctx:        ctx,
		cancelFunc: cancelFunc,
		//driverCache: make(map[string]c_base.IDriver),
	}
}

func (d *SDeviceCmd) Start() {
	// 捕捉panic
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(d.ctx, "启动驱动失败！原因：%v", err)
			d.Stop()
		}
	}()

	// 初始化存储
	//common.InitStorage(d.ctx, influxdb_1.NewStorageInstance(d.ctx, common.GetStorageConfig(d.ctx)))

	common.RegisterStorageInstance(func(ctx context.Context) c_base.IStorage {
		return influxdb_2.NewStorageInstance(ctx, common.GetStorageConfig(d.ctx))
	})

	d.InitDriver(common.GetDriverConfig(d.ctx), common.GetProtocolsConfigList(d.ctx))
}

func (d *SDeviceCmd) Stop() {
	// 关闭存储
	common.CloseStorage()
	// 关闭所有client
	for _, driver := range common.GetDeviceAll() {
		driver.Destroy()
	}
	d.modbusClientCache = make(map[string]modbus.Client)
	d.cancelFunc()
}

func (d *SDeviceCmd) InitDriver(config *c_base.SDriverConfig, protocolConfigList []*c_base.SProtocolConfig) c_base.IDriver {
	if err := config.Check(); err != nil {
		panic(err)
	}

	if config.Enable == false {
		g.Log().Noticef(d.ctx, "设备[%s]未启用！", config.Name)
		return nil
	}

	// 先递归初始化子设备
	if config.DeviceChildren != nil {
		for _, _device := range config.DeviceChildren {
			d.InitDriver(_device, protocolConfigList)
		}
	}

	if config.Id == "root" {
		return nil
	}

	var protocolProvider c_base.IProtocol
	// 设备初始化
	if config.ProtocolId != "" {
		d.ctx = context.WithValue(d.ctx, c_base.ConstCtxKeyProtocolId, config.ProtocolId)
		var protocolConfig *c_base.SProtocolConfig
		for _, _protocolConfig := range protocolConfigList {
			if _protocolConfig.Id == config.ProtocolId {
				protocolConfig = _protocolConfig
				break
			}
		}
		protocolProvider = d.getProtocolProvider(d.ctx, config, protocolConfig)
	} else {
		d.ctx = context.WithValue(d.ctx, c_base.ConstCtxKeyProtocolId, "Virtual")
	}

	driver := d.getDriver(d.ctx, config)
	if driver == nil {
		g.Log().Errorf(d.ctx, "设备[%s]驱动加载失败！", config.Name)
		return nil
	}

	driver.Init(protocolProvider, config)

	g.Log().Noticef(d.ctx, "设备[%s]驱动加载初始化完毕！\n  设备信息: %s", config.Name, driver.GetDescription())

	if protocolProvider != nil {
		protocolProvider.Init()
	}

	common.RegisterDevice(driver)

	common.StorageTimerSaveDeviceMetrics(d.ctx, config.StorageIntervalSec, driver)

	return driver
}

// Block 阻塞进程
func (d *SDeviceCmd) Block() {
	gproc.AddSigHandlerShutdown(func(sig os.Signal) {
		g.Log().Noticef(d.ctx, "接收到信号：%s", sig.String())
		d.Stop()
	})

	gproc.Listen()
}

func (d *SDeviceCmd) getProtocolProvider(ctx context.Context, deviceConfig *c_base.SDriverConfig, protocolConfig *c_base.SProtocolConfig) c_base.IProtocol {
	// 从配置中获取协议
	//protocolConfig := common.GetProtocolById(ctx, protocolId)
	if protocolConfig == nil {
		panic(gerror.Newf("协议Id: %s 配置信息不存在！", deviceConfig.ProtocolId))
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
			client = protocolModbus.NewModbusClient(ctx, protocolConfig)
		}

		modbusProvider, err := protocolModbus.NewModbusProvider(ctx, protocolConfig, deviceConfig, client)
		if err != nil {
			panic(err)
		}

		return modbusProvider
	case c_base.ECanbusTcp:
	case c_base.ECanbus:
	case c_base.EGpioSysfs:
		gpioSysfsProtocol, err := gpio_sysfs.NewGpioSysfsProvider(ctx, protocolConfig, deviceConfig)
		if err != nil {
			panic(err)
		}
		return gpioSysfsProtocol
	}

	panic(gerror.Newf("未知的协议类型：%s", protocolConfig.GetProtocol()))
}

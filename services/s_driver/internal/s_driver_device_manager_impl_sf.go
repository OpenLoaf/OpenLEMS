package internal

import (
	"common"
	"common/c_base"
	"common/c_device"
	"common/c_log"
	"common/c_proto"
	"context"
	"s_db"
	"s_storage"
	"sort"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
)

// SDeviceManager 通用驱动管理器实现
type SDeviceManager struct {
	parentCtx  context.Context
	ctx        context.Context
	cancelFunc context.CancelFunc
	state      c_base.EServerState // 状态

	protocolClientCache map[string]any // 协议客户端缓存

	deviceConfig      []*c_base.SDeviceConfig
	deviceInstanceMap map[string]c_base.IDevice // 设备实例
}

var (
	// 编译期断言内部实现满足对外接口
	_ common.IDeviceManager = (*SDeviceManager)(nil)

	driverManagerInstance *SDeviceManager
	driverManagerInitOnce sync.Once
)

// NewSingleDriverManager 创建驱动管理器
func NewSingleDriverManager(parentCtx context.Context) *SDeviceManager {
	driverManagerInitOnce.Do(func() {
		driverManagerInstance = &SDeviceManager{
			parentCtx: parentCtx,
		}
	})
	return driverManagerInstance
}

func (d *SDeviceManager) GetDeviceById(deviceId string) c_base.IDevice {
	if dw, exist := d.deviceInstanceMap[deviceId]; exist {
		return dw
	}
	return nil
}

func (d *SDeviceManager) GetDeviceConfigById(deviceId string) *c_base.SDeviceConfig {
	return d.FindDevice(deviceId)
}

func (d *SDeviceManager) GetAllDriversInfo() []*c_base.SDriverInfo {
	allDriversMap := GetAllDriversInfo()

	// 将map转换为切片
	var drivers []*c_base.SDriverInfo
	for _, driver := range allDriversMap {
		drivers = append(drivers, driver)
	}

	// 按名称排序
	sort.Slice(drivers, func(i, j int) bool {
		return drivers[i].Name < drivers[j].Name
	})

	return drivers
}

func (d *SDeviceManager) IteratorAssAllDevicesWrapper(deviceWrapper func(config *c_base.SDeviceConfig, device c_base.IDevice) bool) {
	flatList := d.GetFlatList()
	for _, config := range flatList {
		if !deviceWrapper(config, d.deviceInstanceMap[config.Id]) {
			break
		}
	}
}

func (d *SDeviceManager) Start() {
	if d.state == c_base.EStateRunning {
		c_log.BizErrorf(d.ctx, "服务启动失败，服务已经在运行状态中了")
		return
	}

	d.ctx, d.cancelFunc = context.WithCancel(d.parentCtx)
	d.state = c_base.EStateInit
	d.protocolClientCache = make(map[string]any)
	d.deviceInstanceMap = make(map[string]c_base.IDevice)

	rootDeviceID := s_db.GetDeviceService().GetRootDeviceId()
	deviceConfigs, err := s_db.GetDeviceService().GetEnableDeviceConfigsWithRecursion(d.ctx, rootDeviceID)
	if err != nil {
		d.state = c_base.EStateError
		g.Log().Errorf(d.ctx, "初始化失败！获取设备配置失败！%v", err)
		return
	}
	protocolConfigs, err := s_db.GetProtocolService().GetAllProtocolConfigs(d.ctx)
	if err != nil {
		d.state = c_base.EStateError
		g.Log().Errorf(d.ctx, "初始化失败！获取协议配置失败！%v", err)
		return
	}

	var protocolConfigMap = make(map[string]*c_base.SProtocolConfig)
	for _, protocolConfig := range protocolConfigs {
		protocolConfigMap[protocolConfig.Id] = protocolConfig
	}

	for _, deviceConfig := range deviceConfigs {
		if deviceConfig.Id == deviceConfig.Pid {
			g.Log().Errorf(d.ctx, "错误的设备[%s]配置，id和pid一致！", deviceConfig.Name)
			continue
		}
		// 设置协议配置
		deviceConfig.ProtocolConfig = protocolConfigMap[deviceConfig.ProtocolId]
		// 添加驱动信息
		deviceConfig.DriverInfo, _ = GetDriverInfo(deviceConfig.Driver)
	}

	// 构建树形结构
	d.deviceConfig = d.BuildTree(deviceConfigs)

	// 从底部开始初始化设备
	d.ExecuteFromBottom(func(deviceConfig *c_base.SDeviceConfig) {
		c_log.BizInfof(d.ctx, "加载设备：[%s] 准备初始化！", deviceConfig.Name)

		if deviceConfig.Enabled == false {
			c_log.BizInfof(d.ctx, "设备[%s]未启用！", deviceConfig.Name)
			return
		}
		if deviceConfig.DriverInfo == nil {
			c_log.BizInfof(d.ctx, "设备[%s]驱动未找到！", deviceConfig.Name)
			// todo 虚拟节点启动失败的话，子节点都需要关机 （可以做成全局参数）
			return
		}

		deviceCtx := context.WithValue(d.ctx, c_base.ConstCtxKeyDeviceId, deviceConfig.Id)
		deviceCtx = context.WithValue(deviceCtx, c_base.ConstCtxKeyDeviceName, deviceConfig.Name)

		if deviceConfig.ProtocolConfig == nil {
			// 虚拟设备，创建不会失败
			device := c_device.NewVirtualDevice(deviceCtx, deviceConfig)
			// 物理设备
			driver, err := getDriver(deviceConfig.Driver, device)
			if driver == nil || err != nil {
				c_log.BizErrorf(d.ctx, "虚拟设备[%s]驱动加载失败！原因：%s", deviceConfig.Name, err.Error())
				return
			}
			err = driver.Init()
			if err != nil {
				c_log.BizErrorf(d.ctx, "虚拟设备[%s]初始化失败！原因：%s", deviceConfig.Name, err.Error())
				return
			}
			d.deviceInstanceMap[deviceConfig.Id] = driver
			c_log.BizInfof(d.ctx, "虚拟设备[%s]初始化成功！", deviceConfig.Name)
		} else {

			protocolProvider, err := d.getProtocolProvider(deviceCtx, deviceConfig)
			if protocolProvider == nil || err != nil {
				// todo 添加日志，创建连接失败了
				c_log.BizErrorf(deviceCtx, "创建协议实例失败! 协议ID: %s 原因：%s", deviceConfig.ProtocolId, err.Error())
				d.state = c_base.EStateError
				return
			}
			device, err := c_device.NewRealDevice(deviceCtx, deviceConfig, protocolProvider.(c_proto.IModbusProtocol))
			if err != nil {
				c_log.BizErrorf(d.ctx, "设备[%s] 初始化失败！原因：%s", deviceConfig.Name, err.Error())
				return
			}

			// 物理设备
			driver, err := getDriver(deviceConfig.Driver, device)
			if driver == nil || err != nil {
				c_log.BizErrorf(d.ctx, "设备[%s]驱动加载失败！原因：%s", deviceConfig.Name, err.Error())
				return
			}
			err = driver.Init()
			if err != nil {
				c_log.BizErrorf(d.ctx, "设备[%s]初始化失败！原因：%s", deviceConfig.Name, err.Error())
				return
			}
			protocolProvider.ProtocolListen() // 启动监听

			d.deviceInstanceMap[deviceConfig.Id] = driver
			c_log.BizInfof(d.ctx, "设备[%s]初始化成功！", deviceConfig.Name)
		}
	})

	// 从底部开始注册存储驱动
	d.ExecuteFromBottom(func(deviceConfig *c_base.SDeviceConfig) {
		if deviceConfig.StorageEnable {
			s_storage.RegisterStorageDriver(deviceConfig)
		}
	})

	c_log.BizInfof(d.ctx, "全部设备启动成功！")
	d.state = c_base.EStateRunning

	d.state = c_base.EStateRunning
}

// GetDriverInfo 获取指定驱动的详细信息
func (d *SDeviceManager) GetDriverInfo(driverName string) (*c_base.SDriverInfo, error) {
	return GetDriverInfo(driverName)
}

// GetDriversByType 根据设备类型获取驱动信息
func (d *SDeviceManager) GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []*c_base.SDriverInfo {
	return GetDriversByType(ctx, deviceType)
}

// GetSupportedDeviceTypes 获取支持的设备类型列表
func (d *SDeviceManager) GetSupportedDeviceTypes(ctx context.Context) []c_base.EDeviceType {
	driversInfo := d.GetAllDriversInfo()
	typeMap := make(map[c_base.EDeviceType]bool)

	for _, driver := range driversInfo {
		if driver.Enabled && driver.Type != "" {
			typeMap[driver.Type] = true
		}
	}

	var deviceTypes []c_base.EDeviceType
	for deviceType := range typeMap {
		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes
}

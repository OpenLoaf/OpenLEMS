package internal

import (
	"common"
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"context"
	"s_db"
	"s_storage"
	"sort"
	"sync"

	"github.com/gogf/gf/v2/container/glist"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
)

// SDeviceManager 通用驱动管理器实现
type SDeviceManager struct {
	parentCtx  context.Context
	ctx        context.Context
	cancelFunc context.CancelFunc
	state      c_enum.EServerState // 状态

	protocolClientCache *gmap.StrAnyMap // 协议客户端缓存（线程安全）

	deviceConfigTree  *glist.List     // 设备配置树形节点（线程安全）
	deviceInstanceMap *gmap.StrAnyMap // 设备实例（线程安全）
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
			parentCtx:           parentCtx,
			protocolClientCache: gmap.NewStrAnyMap(true), // 线程安全
			deviceConfigTree:    glist.New(true),         // 线程安全
			deviceInstanceMap:   gmap.NewStrAnyMap(true), // 线程安全
		}
	})
	return driverManagerInstance
}

func (m *SDeviceManager) GetDeviceNameById(deviceId string) string {
	config := m.GetDeviceConfigById(deviceId)
	if config == nil {
		return ""
	}
	return config.Name
}

func (m *SDeviceManager) GetDeviceConfigTree() []*c_base.SDeviceConfig {
	// 将线程安全的列表转换为切片
	var result []*c_base.SDeviceConfig
	m.deviceConfigTree.Iterator(func(e *glist.Element) bool {
		if config, ok := e.Value.(*c_base.SDeviceConfig); ok {
			result = append(result, config)
		}
		return true
	})
	return result
}

func (m *SDeviceManager) GetDeviceById(deviceId string) c_base.IDevice {
	if device := m.deviceInstanceMap.Get(deviceId); device != nil {
		if dev, ok := device.(c_base.IDevice); ok {
			return dev
		}
	}
	return nil
}

func (m *SDeviceManager) GetDeviceConfigById(deviceId string) *c_base.SDeviceConfig {
	return m.FindDevice(deviceId)
}

func (m *SDeviceManager) GetAllDriversInfo() []*c_base.SDriverInfo {
	allDriversMap := GetAllDriversInfo()

	// 将map转换为切片
	var drivers []*c_base.SDriverInfo
	for key, driver := range allDriversMap {
		if driver.Name == "" {
			driver.Name = key
		}
		drivers = append(drivers, driver)
	}

	// 按名称排序
	sort.Slice(drivers, func(i, j int) bool {
		return drivers[i].Name < drivers[j].Name
	})

	return drivers
}

func (m *SDeviceManager) IteratorAllDevices(deviceWrapper func(config *c_base.SDeviceConfig, device c_base.IDevice) bool) {
	flatList := m.GetFlatList()
	for _, config := range flatList {
		var device c_base.IDevice
		if dev := m.deviceInstanceMap.Get(config.Id); dev != nil {
			if d, ok := dev.(c_base.IDevice); ok {
				device = d
			}
		}
		if !deviceWrapper(config, device) {
			break
		}
	}
}

// IteratorChildDevicesById 按设备ID遍历该设备及所有子设备
func (m *SDeviceManager) IteratorChildDevicesById(deviceId string, iterator func(config *c_base.SDeviceConfig, device c_base.IDevice) bool) {
	if deviceId == "" || iterator == nil {
		m.IteratorAllDevices(iterator)
		return
	}
	start := m.GetDeviceConfigById(deviceId)
	if start == nil {
		return
	}
	// 递归遍历子树（包含起始节点）
	var walk func(node *c_base.SDeviceConfig) bool
	walk = func(node *c_base.SDeviceConfig) bool {
		if node == nil {
			return true
		}
		var device c_base.IDevice
		if dev := m.deviceInstanceMap.Get(node.Id); dev != nil {
			if d, ok := dev.(c_base.IDevice); ok {
				device = d
			}
		}
		if cont := iterator(node, device); !cont {
			return false
		}
		for _, child := range node.ChildDeviceConfig {
			if !walk(child) {
				return false
			}
		}
		return true
	}
	_ = walk(start)
}

func (m *SDeviceManager) IteratorParentDevicesById(deviceId string, iterator func(config *c_base.SDeviceConfig, device c_base.IDevice) bool) {
	if deviceId == "" || iterator == nil {
		return
	}
	current := m.GetDeviceConfigById(deviceId)
	if current == nil {
		return
	}
	visited := make(map[string]bool)
	for current != nil {
		// 防止异常配置导致的循环
		if visited[current.Id] {
			break
		}
		visited[current.Id] = true
		var device c_base.IDevice
		if dev := m.deviceInstanceMap.Get(current.Id); dev != nil {
			if d, ok := dev.(c_base.IDevice); ok {
				device = d
			}
		}
		if cont := iterator(current, device); !cont {
			break
		}
		if current.Pid == "" || current.Pid == current.Id {
			break
		}
		current = m.GetDeviceConfigById(current.Pid)
	}
}

func (m *SDeviceManager) Start() {
	if m.state == c_enum.EStateRunning {
		c_log.BizErrorf(m.ctx, "服务启动失败，服务已经在运行状态中了")
		return
	}

	m.ctx, m.cancelFunc = context.WithCancel(m.parentCtx)
	m.state = c_enum.EStateInit
	// 清空线程安全的容器
	m.protocolClientCache.Clear()
	m.deviceInstanceMap.Clear()
	m.deviceConfigTree.Clear()

	rootDeviceID := s_db.GetSettingService().GetRootDeviceId(m.ctx)
	deviceConfigs, err := s_db.GetDeviceService().GetDeviceConfigsWithRecursion(m.ctx, rootDeviceID)
	if err != nil {
		m.state = c_enum.EStateError
		g.Log().Errorf(m.ctx, "初始化失败！获取设备配置失败！%+v", err)
		return
	}
	protocolConfigs, err := s_db.GetProtocolService().GetAllProtocolConfigs(m.ctx)
	if err != nil {
		m.state = c_enum.EStateError
		g.Log().Errorf(m.ctx, "初始化失败！获取协议配置失败！%+v", err)
		return
	}

	var protocolConfigMap = make(map[string]*c_base.SProtocolConfig)
	for _, protocolConfig := range protocolConfigs {
		protocolConfigMap[protocolConfig.Id] = protocolConfig
	}

	for _, deviceConfig := range deviceConfigs {
		if deviceConfig.Id == deviceConfig.Pid {
			g.Log().Errorf(m.ctx, "错误的设备[%s]配置，id和pid一致！", deviceConfig.Name)
			continue
		}
		// 设置协议配置
		deviceConfig.ProtocolConfig = protocolConfigMap[deviceConfig.ProtocolId]
		// 添加驱动信息
		deviceConfig.DriverInfo, _ = GetDriverInfo(deviceConfig.Driver)
	}

	// 构建树形结构
	treeConfigs := m.BuildTree(deviceConfigs)
	// 将树形结构添加到线程安全的列表中
	for _, config := range treeConfigs {
		m.deviceConfigTree.PushBack(config)
	}

	// 递归初始化设备（从最底部开始）
	m.InitializeDevicesRecursively(treeConfigs)

	// 递归注册存储驱动
	m.RegisterStorageDriversRecursively(treeConfigs)

	m.state = c_enum.EStateRunning
}

// BuildVirtualDevice 创建虚拟设备
func (m *SDeviceManager) BuildVirtualDevice(deviceCtx context.Context, deviceConfig *c_base.SDeviceConfig) {
	// 先判断子设备是否都注册成功，如果成功了就创建。否则不创建虚拟设备
	for _, child := range deviceConfig.ChildDeviceConfig {
		if child.Enabled == false {
			continue
		}
		if !m.deviceInstanceMap.Contains(child.Id) {
			c_log.BizErrorf(deviceCtx, "设备启动失败！原因：子设备[%s(%s) ]启动失败!", child.Name, child.Id)
			return
		}
	}

	// 虚拟设备，创建不会失败
	device := c_device.NewVirtualDevice(deviceCtx, deviceConfig)
	// 物理设备
	driver, err := getDriver(deviceConfig.Driver, device)
	if err != nil {
		c_log.BizErrorf(deviceCtx, "虚拟设备[%s]驱动加载失败！原因：%s", deviceConfig.Name, err.Error())
		return
	}
	if driver == nil {
		c_log.BizErrorf(deviceCtx, "虚拟设备[%s]驱动加载失败！原因driver为空", deviceConfig.Name)
		return
	}
	err = driver.Init()
	if err != nil {
		c_log.BizErrorf(deviceCtx, "虚拟设备[%s]初始化失败！原因：%s", deviceConfig.Name, err.Error())
		return
	}
	m.deviceInstanceMap.Set(deviceConfig.Id, driver)
	c_log.BizInfof(deviceCtx, "虚拟设备[%s]初始化成功！", deviceConfig.Name)
}

// BuildRealDevice 创建物理设备连接
func (m *SDeviceManager) BuildRealDevice(deviceCtx context.Context, deviceConfig *c_base.SDeviceConfig) {
	protocolProvider, err := m.getProtocolProvider(deviceCtx, deviceConfig)
	if protocolProvider == nil || err != nil {
		// todo 添加日志，创建连接失败了
		c_log.BizErrorf(deviceCtx, "创建协议实例失败! 协议ID: %s 原因：%s", deviceConfig.ProtocolId, err)
		return
	}

	var device c_base.IDevice
	switch deviceConfig.ProtocolConfig.GetProtocol() {
	case c_enum.EModbusRtu, c_enum.EModbusTcp:
		device, err = c_device.NewRealDevice(deviceCtx, protocolProvider.(c_proto.IModbusProtocol))
	case c_enum.ECanbus, c_enum.ECanbusUdp:
		device, err = c_device.NewRealDevice(deviceCtx, protocolProvider.(c_proto.ICanbusProtocol))
	case c_enum.EGpioIn, c_enum.EGpioOut:
		//device, err = c_device.NewRealGpio(deviceCtx, protocolProvider.(c_proto.IGpiodProtocol))
		device, err = c_device.NewRealDevice(deviceCtx, protocolProvider.(c_proto.IGpiodProtocol))
	}

	if err != nil {
		c_log.BizErrorf(deviceCtx, "设备[%s] 初始化失败！原因：%s", deviceConfig.Name, err.Error())
		return
	}
	if device == nil {
		c_log.BizErrorf(deviceCtx, "设备[%s] 初始化失败！原因：device为空", deviceConfig.Name)
		return
	}

	// 物理设备
	driver, err := getDriver(deviceConfig.Driver, device)
	if driver == nil || err != nil {
		c_log.BizErrorf(deviceCtx, "设备[%s]驱动加载失败！原因：%s", deviceConfig.Name, err.Error())
		return
	}
	err = driver.Init()
	if err != nil {
		c_log.BizErrorf(deviceCtx, "设备[%s]初始化失败！原因：%s", deviceConfig.Name, err.Error())
		return
	}
	protocolProvider.ProtocolListen() // 启动监听

	m.deviceInstanceMap.Set(deviceConfig.Id, driver)
	c_log.BizInfof(deviceCtx, "设备[%s]初始化成功！", deviceConfig.Name)
}

// GetDriverInfo 获取指定驱动的详细信息
func (m *SDeviceManager) GetDriverInfo(driverName string) (*c_base.SDriverInfo, error) {
	return GetDriverInfo(driverName)
}

// GetDriversByType 根据设备类型获取驱动信息
func (m *SDeviceManager) GetDriversByType(ctx context.Context, deviceType c_enum.EDeviceType) []*c_base.SDriverInfo {
	return GetDriversByType(ctx, deviceType)
}

// GetSupportedDeviceTypes 获取支持的设备类型列表
func (m *SDeviceManager) GetSupportedDeviceTypes(ctx context.Context) []c_enum.EDeviceType {
	driversInfo := m.GetAllDriversInfo()
	typeMap := make(map[c_enum.EDeviceType]bool)

	for _, driver := range driversInfo {
		if driver.Enabled && driver.Type != "" {
			typeMap[driver.Type] = true
		}
	}

	var deviceTypes []c_enum.EDeviceType
	for deviceType := range typeMap {
		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes
}

// InitializeDevicesRecursively 递归初始化设备，从最底部开始
func (m *SDeviceManager) InitializeDevicesRecursively(devices []*c_base.SDeviceConfig) {
	for _, deviceConfig := range devices {
		// 递归初始化子设备
		if len(deviceConfig.ChildDeviceConfig) > 0 {
			m.InitializeDevicesRecursively(deviceConfig.ChildDeviceConfig)
		}

		// 初始化当前设备
		m.initializeSingleDevice(deviceConfig)
	}
}

// initializeSingleDevice 初始化单个设备
func (m *SDeviceManager) initializeSingleDevice(deviceConfig *c_base.SDeviceConfig) {
	c_log.BizInfof(m.ctx, "加载设备：[%s] 准备初始化！", deviceConfig.Name)

	// 检查当前设备是否启用
	if !deviceConfig.Enabled {
		c_log.BizWarningf(m.ctx, "设备[%s]未启用，跳过初始化！", deviceConfig.Name)
		return
	}

	// 检查父设备是否启用（如果父设备未启用，子设备也不应该初始化）
	if deviceConfig.Pid != "" && deviceConfig.Pid != deviceConfig.Id {
		parentConfig := m.GetDeviceConfigById(deviceConfig.Pid)
		if parentConfig != nil && !parentConfig.Enabled {
			c_log.BizWarningf(m.ctx, "设备[%s]的父设备[%s]未启用，跳过初始化！", deviceConfig.Name, parentConfig.Name)
			return
		}
	}

	// 检查驱动信息
	if deviceConfig.DriverInfo == nil {
		c_log.BizErrorf(m.ctx, "设备[%s]驱动未找到！", deviceConfig.Name)
		return
	}

	deviceCtx := context.WithValue(m.ctx, c_enum.ELogTypeDevice, deviceConfig.Id)
	deviceCtx = context.WithValue(deviceCtx, c_enum.ELogTypeEms, deviceConfig.Name)

	// 根据协议配置选择初始化方式
	if deviceConfig.ProtocolConfig == nil {
		m.BuildVirtualDevice(deviceCtx, deviceConfig)
	} else {
		m.BuildRealDevice(deviceCtx, deviceConfig)
	}
}

// RegisterStorageDriversRecursively 递归注册存储驱动
func (m *SDeviceManager) RegisterStorageDriversRecursively(devices []*c_base.SDeviceConfig) {
	for _, deviceConfig := range devices {
		// 递归注册子设备的存储驱动
		if len(deviceConfig.ChildDeviceConfig) > 0 {
			m.RegisterStorageDriversRecursively(deviceConfig.ChildDeviceConfig)
		}

		// 注册当前设备的存储驱动
		if deviceConfig.StorageEnable {
			s_storage.RegisterStorageDriver(deviceConfig)
		}
	}
}

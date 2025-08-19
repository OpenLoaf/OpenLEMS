package internal

import (
	"common"
	"common/c_base"
	"context"
	"github.com/gogf/gf/v2/container/gtree"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gutil"
	"s_db"
	"s_storage"
	"sync"
)

// SDeviceManager 通用驱动管理器实现
type SDeviceManager struct {
	ctx        context.Context
	cancelFunc context.CancelFunc
	state      c_base.EServerState // 状态

	protocolClientCache map[string]any // 协议客户端缓存
	deviceWrapperTree   *gtree.AVLTree // c_base.IDeviceWrapper 设备树
}

func (d *SDeviceManager) IteratorAssAllDevicesWrapper(process func(deviceWrapper c_base.IDeviceWrapper)) {
	if d.deviceWrapperTree == nil {
	}
	d.deviceWrapperTree.IteratorAsc(func(key, value any) bool {
		if v, ok := value.(c_base.IDeviceWrapper); ok {
			process(v)
		}
		return true
	})
}

func (d *SDeviceManager) GetDeviceById(deviceId string) c_base.IDevice {
	dw := d.deviceWrapperTree.Get(deviceId)
	if dw == nil {
		return nil
	}
	return dw.(c_base.IDeviceWrapper).GetDeviceInstance()
}

func (d *SDeviceManager) Start(parentCtx context.Context) {
	d.ctx, d.cancelFunc = context.WithCancel(parentCtx)
	d.state = c_base.EStateInit
	d.protocolClientCache = make(map[string]any)
	d.deviceWrapperTree = gtree.NewAVLTree(gutil.ComparatorString)

	rootDeviceID := s_db.GetDeviceService().GetRootDeviceId()
	deviceConfigs, err := s_db.GetDeviceService().GetDeviceConfigs(d.ctx, rootDeviceID)
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

		// 更新设备信息
		driverInfo, _ := GetDriverInfo(d.ctx, deviceConfig.Driver)

		d.deviceWrapperTree.Set(deviceConfig.Id, &SDeviceWrapper{
			deviceConfig:   deviceConfig,
			driverInfo:     driverInfo,
			protocolConfig: protocolConfigMap[deviceConfig.ProtocolId],
			instance:       nil,
			deviceState:    c_base.EStateInit,
		})
	}

	// 反向递归，初始化设备
	d.deviceWrapperTree.IteratorDesc(func(k, v any) bool {
		deviceWrapper := v.(*SDeviceWrapper)
		deviceConfig := deviceWrapper.deviceConfig
		protocolConfig := deviceWrapper.protocolConfig
		ctx := context.WithValue(d.ctx, c_base.ConstCtxKeyDeviceId, deviceConfig.Name)

		if deviceConfig.Enabled == false {
			g.Log().Noticef(d.ctx, "设备[%s]未启用！", deviceConfig.Name)
			deviceWrapper.UpdateState(c_base.EStateStopped)
			return true
		}

		driver := getDriver(ctx, deviceConfig)
		if driver == nil {
			g.Log().Errorf(d.ctx, "设备[%s]驱动加载失败！", deviceConfig.Name)
			deviceWrapper.UpdateState(c_base.EStateError)
			return true
		}
		deviceWrapper.instance = driver

		var protocolProvider c_base.IProtocol
		if deviceConfig.ProtocolId != "" {
			// 创建协议provider
			protocolProvider, err = d.getProtocolProvider(ctx, driver.GetDriverType(), deviceConfig, protocolConfig)
			if protocolProvider == nil || err != nil {
				// todo 添加日志，创建连接失败了
				deviceWrapper.UpdateState(c_base.EStateError)
				return true
			}
		}
		// 初始化设备
		driver.InitDevice(deviceConfig, protocolProvider, d.GetChildDeviceInstance(deviceConfig.Id))
		// 协议监听
		driver.ProtocolListen()

		if deviceConfig.StorageEnable {
			//common.GetStorageInstance().
			s_storage.RegisterStorageDriver(deviceConfig.StorageIntervalSec, driver)
		}

		deviceWrapper.UpdateState(c_base.EStateRunning)
		g.Log().Noticef(ctx, "设备[%s]驱动加载初始化完毕！\n  设备信息: %s", deviceConfig.Name, driver.GetDriverDescription())
		return true
	})

	d.state = c_base.EStateRunning
}

var (
	// 编译期断言内部实现满足对外接口
	_ common.IDeviceManager = (*SDeviceManager)(nil)

	driverManagerInstance *SDeviceManager
	driverManagerInitOnce sync.Once
)

// GetDriverManager 创建驱动管理器
func GetDriverManager() *SDeviceManager {
	driverManagerInitOnce.Do(func() {
		driverManagerInstance = &SDeviceManager{
			//protocolClientCache: make(map[string]any),
		}
	})
	return driverManagerInstance
}

// GetAllDriverNames 获取所有驱动名称
func (d *SDeviceManager) GetAllDriverNames() []string {
	return GetAllDriverNames()
}

// GetAllDriversInfo 获取所有驱动的详细信息
func (d *SDeviceManager) GetAllDriversInfo(ctx context.Context) []c_base.SDriverInfo {
	return GetAllDriversInfo(ctx)
}

// GetDriverInfo 获取指定驱动的详细信息
func (d *SDeviceManager) GetDriverInfo(ctx context.Context, driverName string) (*c_base.SDriverInfo, error) {
	return GetDriverInfo(ctx, driverName)
}

// GetDriversByType 根据设备类型获取驱动信息
func (d *SDeviceManager) GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.SDriverInfo {
	return GetDriversByType(ctx, deviceType)
}

//// CreateDriver 创建驱动实例
//func (d *SDeviceManager) CreateDriver(ctx context.Context, deviceConfig *c_base.SDeviceConfig) (c_base.IDevice, error) {
//	defer func() {
//		if r := recover(); r != nil {
//			// 将panic转换为error
//			if err, ok := r.(error); ok {
//				panic(err)
//			} else {
//				panic(r)
//			}
//		}
//	}()
//
//	driver := getDriver(ctx, deviceConfig)
//	return driver, nil
//}

// IsDriverAvailable 检查驱动是否可用
func (d *SDeviceManager) IsDriverAvailable(ctx context.Context, driverName string) bool {
	driverInfo, err := d.GetDriverInfo(ctx, driverName)
	if err != nil {
		return false
	}
	return driverInfo.Available
}

// GetDriverDescription 获取驱动描述信息
func (d *SDeviceManager) GetDriverDescription(ctx context.Context, driverName string) (*c_base.SDriverDescription, error) {
	driverInfo, err := d.GetDriverInfo(ctx, driverName)
	if err != nil {
		return nil, err
	}
	return driverInfo.Description, nil
}

// GetSupportedDeviceTypes 获取支持的设备类型列表
func (d *SDeviceManager) GetSupportedDeviceTypes(ctx context.Context) []c_base.EDeviceType {
	driversInfo := d.GetAllDriversInfo(ctx)
	typeMap := make(map[c_base.EDeviceType]bool)

	for _, driver := range driversInfo {
		if driver.Available && driver.Type != "" {
			typeMap[driver.Type] = true
		}
	}

	var deviceTypes []c_base.EDeviceType
	for deviceType := range typeMap {
		deviceTypes = append(deviceTypes, deviceType)
	}

	return deviceTypes
}

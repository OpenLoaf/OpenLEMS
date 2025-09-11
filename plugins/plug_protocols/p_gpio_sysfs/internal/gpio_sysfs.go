package internal

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_proto"
	"context"
	"fmt"
	"p_base"
	"p_gpio_sysfs/p_gpio_sysfs"
	"sync"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
)

// SGpioPoint GPIO点位结构
type SGpioPoint struct {
	*c_base.SPoint
}

func (s *SGpioPoint) GetDataAccess() *c_base.SDataAccess {
	// GPIO 点位不需要数据访问配置，返回 nil
	return nil
}

type SGpioSysfsProvider struct {
	c_base.IAlarm
	c_base.IProtocolCacheValue

	ctx  context.Context // 上下文
	once sync.Once       // 只执行一次Init方法

	deviceId string

	// GPIO 相关配置
	deviceConfig       *c_base.SDeviceConfig                                 // 设备基础配置
	deviceGpioConfig   *c_proto.SDeviceGpioConfig                            // 设备扩展配置
	protocolConfig     *c_base.SProtocolConfig                               // 协议配置
	protocolGpioConfig *p_gpio_sysfs.SProtocolGpioConfig                     // 协议扩展配置
	gpioDirPath        string                                                // gpio路径
	gpioValuePath      string                                                // gpio值路径
	handler            func(ctx context.Context, status bool, isChange bool) // 处理函数

	// 状态管理
	status         bool                   // 结果
	lastUpdateTime *time.Time             // 最后更新时间
	protocolStatus c_enum.EProtocolStatus // 协议状态

	// 元数据
	meta *SGpioPoint // 元数据
}

var _ c_proto.IGpioSysfsProtocol = (*SGpioSysfsProvider)(nil)

func (s *SGpioSysfsProvider) IsPhysical() bool {
	return true
}

func (s *SGpioSysfsProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	return s.protocolStatus
}

func NewGpioSysfsProvider(ctx context.Context, protocolConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpioSysfsProtocol, error) {
	if protocolConfig == nil {
		return nil, gerror.Newf("GPIO设备：[%s]%s 的协议配置不能为空！", deviceConfig.Id, deviceConfig.Name)
	}
	if protocolConfig.Address == "" {
		return nil, gerror.Newf("GPIO设备：[%s]%s 的协议配置地址不能为空！", deviceConfig.Id, deviceConfig.Name)
	}
	if deviceConfig == nil {
		return nil, gerror.Newf("GPIO协议：%s 的设备配置不能为空！", protocolConfig.Id)
	}

	var deviceGpioConfig = &c_proto.SDeviceGpioConfig{}
	err := gconv.Scan(deviceConfig.Params, deviceGpioConfig)
	if err != nil {
		return nil, gerror.Newf("设备[%s]的Param参数配置错误：%v 无法转换为SGpioSysfsDeviceConfig", deviceConfig.Id, err)
	}
	if deviceGpioConfig.Io == 0 {
		// io默认禁止为0
		return nil, gerror.Newf("GPIO设备[%s]的IO不能为空或者0", deviceConfig.Id)
	}

	provider := &SGpioSysfsProvider{
		IProtocolCacheValue: p_base.NewGetProtocolCacheValue(ctx, deviceConfig.Id),
		IAlarm:              c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		deviceId:            deviceConfig.Id,

		once:             sync.Once{},
		ctx:              ctx,
		deviceConfig:     deviceConfig,
		deviceGpioConfig: deviceGpioConfig,
		protocolConfig:   protocolConfig,
		gpioDirPath:      gfile.Join(protocolConfig.Address, fmt.Sprintf("gpio%d", deviceGpioConfig.Io)),
		protocolStatus:   c_enum.EProtocolConnecting,

		meta: &SGpioPoint{
			SPoint: &c_base.SPoint{
				Key:     "status",
				Name:    deviceConfig.Name,
				Unit:    "",
				Desc:    fmt.Sprintf("Gpio Id: %s %s", deviceConfig.Id, deviceConfig.Name),
				Sort:    0,
				Min:     0,
				Max:     1,
				Precise: 0,
				Trigger: nil,
			},
		},
	}
	provider.gpioValuePath = gfile.Join(provider.gpioDirPath, GpioPathValue)

	provider.protocolGpioConfig = &p_gpio_sysfs.SProtocolGpioConfig{}
	err = gconv.Scan(protocolConfig.Params, provider.protocolGpioConfig)
	if err != nil {
		return nil, gerror.Newf("协议[%s]的Param参数配置错误：%v 无法转换为SGpioProtocolConfig", protocolConfig.Id, err)
	}

	return provider, nil
}

func (s *SGpioSysfsProvider) GetDeviceConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *SGpioSysfsProvider) GetPointValueList() []*c_base.SPointValue {
	if s.lastUpdateTime == nil {
		return []*c_base.SPointValue{}
	}
	return []*c_base.SPointValue{
		c_base.NewPointValue(s.deviceConfig.Id, s.meta, s.status),
	}
}

func (s *SGpioSysfsProvider) ProtocolListen() {
	if s.deviceGpioConfig.Direction == c_proto.EGpioDirectionNone {
		panic(gerror.Newf("direction方向未设置！"))
	}
	s.once.Do(func() {

		// 先判断Path是否存在
		if !gfile.Exists(s.gpioDirPath) {
			// 判断一下export是否存在，如果存在，就export导出一下端口
			exportFilePath := gfile.Join(s.protocolConfig.Address, "export")
			if gfile.Exists(exportFilePath) {
				// 执行导出方法
				gproc.MustShellExec(s.ctx, fmt.Sprintf("echo %d > %s", s.deviceGpioConfig.Io, exportFilePath))
			} else {
				panic(gerror.Newf("导出的方法不存在！%s", exportFilePath))
			}
			// 再次判断是否成功
			if !gfile.Exists(s.gpioDirPath) {
				panic(gerror.Newf("export导出后%s还是不存在！", s.gpioDirPath))
			}
		}
		if !gfile.Exists(s.gpioValuePath) {
			panic(gerror.Newf("gpio值路径不存在！%s", s.gpioValuePath))
		}

		// 设置方向
		err := gfile.PutContents(gfile.Join(s.gpioDirPath, GpioPathDirection), string(s.deviceGpioConfig.Direction))
		if err != nil {
			panic(gerror.Newf("设置方向失败！%v", err))
		}

		// 设置协议状态为已连接
		s.protocolStatus = c_enum.EProtocolConnected

		// 如果是In类型的，说明需要时刻监听
		if s.deviceGpioConfig.Direction == c_proto.EGpioDirectionIn {
			gtimer.SetInterval(s.ctx, 200*time.Millisecond, func(ctx context.Context) {
				// 读取值
				s.isHighForce()
			})
		} else {
			gtimer.SetInterval(s.ctx, 3000*time.Millisecond, func(ctx context.Context) {
				// out类型的 3秒读取一次值
				s.isHighForce()
			})
		}
	})
}

func (s *SGpioSysfsProvider) RegisterHandler(handler func(ctx context.Context, status bool, isChange bool)) {
	if s.deviceGpioConfig.Direction != c_proto.EGpioDirectionIn {
		panic(gerror.Newf("只有输入类型的GPIO才能注册Handler"))
	}
	s.handler = handler
}

func (s *SGpioSysfsProvider) GetId() string {
	return s.deviceConfig.Id
}

func (s *SGpioSysfsProvider) IsActivate() bool {
	return true
}

func (s *SGpioSysfsProvider) GetValue(point c_base.IPoint) (any, error) {
	if point == nil {
		return nil, gerror.New("point is nil")
	}
	// 对于GPIO，我们只有一个点位，直接返回状态
	return s.status, nil
}

func (s *SGpioSysfsProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	// GPIO 协议不需要注册任务，因为它是单一状态
}

func (s *SGpioSysfsProvider) GetLastUpdateTime() *time.Time {
	return s.lastUpdateTime
}

func (s *SGpioSysfsProvider) GetGpioStatus() *bool {
	// 如果缓存中的数据获取时间大于1秒或者，或者缓存中无值就force获取
	if s.lastUpdateTime == nil || time.Now().Sub(*s.lastUpdateTime) > time.Second {
		return s.isHighForce()
	}
	return &s.status
}

func (s *SGpioSysfsProvider) isHighForce() *bool {
	value := gfile.GetContents(s.gpioValuePath)

	status := false
	if gstr.Trim(value) == "1" {
		s.process(true)
		status = true
	} else {
		s.process(false)
		status = false
	}
	return &status
}

func (s *SGpioSysfsProvider) process(status bool) {
	now := time.Now()
	s.lastUpdateTime = &now
	// 如果是反向的，就取反
	if s.deviceGpioConfig.Reverse {
		status = !status
	}

	if s.handler != nil {
		s.handler(s.ctx, status, s.status != status)
		//g.Log().Debugf(s.ctx, "GPIO %s value: %v, func:%v", s.deviceConfig.Id, status, s.handler)
	}

	if s.status != status {
		s.status = status

		// 触发告警
		if s.deviceGpioConfig.Level != c_enum.EAlarmLevelNone && s.deviceGpioConfig.Direction == c_proto.EGpioDirectionIn {
			// 使用 IAlarm 接口的 UpdateAlarm 方法
			s.IAlarm.UpdateAlarm(s.GetId(), s.meta, status)
		}

	}
}

func (s *SGpioSysfsProvider) SetHigh() error {
	// 设置为高电平
	return s.setValue("1")
}

func (s *SGpioSysfsProvider) SetLow() error {
	// 设置为低电平
	return s.setValue("0")
}

func (s *SGpioSysfsProvider) setValue(value string) error {
	// 设置为低电平
	if s.deviceGpioConfig.Direction == c_proto.EGpioDirectionIn {
		g.Log().Debugf(s.ctx, "GPIO Direction %s is [in], can't output", s.GetId())
		return nil
	}
	return gfile.PutContents(s.gpioValuePath, value)
}

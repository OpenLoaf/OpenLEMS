package internal

import (
	"common/c_base"
	"common/c_enum"
	"common/c_proto"
	"context"
	"fmt"
	"os"
	"sync"
	"syscall"
	"time"
	"unsafe"
)

// GPIO相关常量
const (
	GPIOHANDLE_REQUEST_INPUT       = 0x1
	GPIOHANDLE_REQUEST_OUTPUT      = 0x2
	GPIOHANDLE_REQUEST_ACTIVE_LOW  = 0x4
	GPIOHANDLE_REQUEST_OPEN_DRAIN  = 0x8
	GPIOHANDLE_REQUEST_OPEN_SOURCE = 0x10

	GPIO_GET_LINEHANDLE_IOCTL        = 0xC16B03
	GPIOHANDLE_GET_LINE_VALUES_IOCTL = 0xC040B408
	GPIOHANDLE_SET_LINE_VALUES_IOCTL = 0xC040B409
)

// GPIO句柄请求结构体
type gpioHandleRequest struct {
	LineOffsets   [64]uint32
	Flags         uint32
	DefaultValues [64]uint8
	ConsumerLabel [32]byte
	Lines         uint32
	Fd            int32
}

// GPIO值结构体
type gpioHandleData struct {
	Values [64]uint8
}

type sGpiodProvider struct {
	c_base.IAlarm
	gpiodConfig *c_proto.SGpiodProtocolConfig

	// GPIO相关字段
	chipFd *os.File
	lineFd *os.File

	// 状态管理
	mu             sync.RWMutex
	isInitialized  bool
	lastStatus     *bool
	lastUpdateTime time.Time

	// 事件处理
	handler      func(status bool)
	ctx          context.Context
	cancel       context.CancelFunc
	eventChannel chan bool
}

// NewGpiodProvider 创建新的GPIO提供者
func NewGpiodProvider(alarm c_base.IAlarm, config *c_proto.SGpiodProtocolConfig) *sGpiodProvider {
	return &sGpiodProvider{
		IAlarm:      alarm,
		gpiodConfig: config,
	}
}

// Close 关闭GPIO提供者，清理资源
func (s *sGpiodProvider) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isInitialized {
		s.cleanup()
	}

	return nil
}

func (s *sGpiodProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.isInitialized {
		return c_enum.EProtocolDisconnected
	}

	return c_enum.EProtocolConnected
}

func (s *sGpiodProvider) GetLastUpdateTime() *time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if s.lastUpdateTime.IsZero() {
		return nil
	}

	return &s.lastUpdateTime
}

func (s *sGpiodProvider) GetPointValueList() []*c_base.SPointValue {
	//TODO implement me
	panic("implement me")
}

func (s *sGpiodProvider) GetValue(point c_base.IPoint) (any, error) {
	//TODO implement me
	panic("implement me")
}

func (s *sGpiodProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	//TODO implement me
	panic("implement me")
}

func (s *sGpiodProvider) ProtocolListen() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.isInitialized {
		return
	}

	// 创建上下文用于控制生命周期
	s.ctx, s.cancel = context.WithCancel(context.Background())
	s.eventChannel = make(chan bool, 10)

	// 初始化GPIO芯片
	chipName := fmt.Sprintf("/dev/gpiochip%d", s.gpiodConfig.ChipIndex)
	chipFd, err := os.OpenFile(chipName, os.O_RDWR, 0)
	if err != nil {
		// 记录错误但不触发告警，因为这是初始化错误
		fmt.Printf("无法打开GPIO芯片 %s: %v\n", chipName, err)
		return
	}
	s.chipFd = chipFd

	// 根据配置初始化GPIO引脚
	err = s.initializeGpioLine()
	if err != nil {
		// 记录错误但不触发告警，因为这是初始化错误
		fmt.Printf("初始化GPIO引脚失败: %v\n", err)
		s.cleanup()
		return
	}

	s.isInitialized = true

	// 启动事件监听协程
	go s.eventListener()
}

func (s *sGpiodProvider) RegisterHandler(handler func(status bool)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handler = handler
}

func (s *sGpiodProvider) GetGpioStatus() *bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if !s.isInitialized || s.lineFd == nil {
		return nil
	}

	// 对于输入模式，返回当前读取的状态
	if s.gpiodConfig.Direction == c_enum.EGpioDirectionIn {
		return s.lastStatus
	}

	// 对于输出模式，读取当前输出值
	value, err := s.readGpioStatus()
	if err != nil {
		return nil
	}

	return &value
}

func (s *sGpiodProvider) SetHigh() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isInitialized || s.lineFd == nil {
		return fmt.Errorf("GPIO引脚未初始化")
	}

	if s.gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		return fmt.Errorf("只有输出模式的GPIO引脚才能设置电平")
	}

	// 根据低电平有效配置设置值
	value := uint8(1)
	if s.gpiodConfig.LowActive {
		value = 0 // 低电平有效时，高电平对应0
	}

	// 使用ioctl设置GPIO值
	data := gpioHandleData{}
	data.Values[0] = value
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, s.lineFd.Fd(), GPIOHANDLE_SET_LINE_VALUES_IOCTL, uintptr(unsafe.Pointer(&data)))
	if errno != 0 {
		return fmt.Errorf("设置GPIO高电平失败: %v", errno)
	}

	// 更新状态
	s.lastStatus = &[]bool{true}[0]
	s.lastUpdateTime = time.Now()

	return nil
}

func (s *sGpiodProvider) SetLow() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.isInitialized || s.lineFd == nil {
		return fmt.Errorf("GPIO引脚未初始化")
	}

	if s.gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		return fmt.Errorf("只有输出模式的GPIO引脚才能设置电平")
	}

	// 根据低电平有效配置设置值
	value := uint8(0)
	if s.gpiodConfig.LowActive {
		value = 1 // 低电平有效时，低电平对应1
	}

	// 使用ioctl设置GPIO值
	data := gpioHandleData{}
	data.Values[0] = value
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, s.lineFd.Fd(), GPIOHANDLE_SET_LINE_VALUES_IOCTL, uintptr(unsafe.Pointer(&data)))
	if errno != 0 {
		return fmt.Errorf("设置GPIO低电平失败: %v", errno)
	}

	// 更新状态
	s.lastStatus = &[]bool{false}[0]
	s.lastUpdateTime = time.Now()

	return nil
}

// initializeGpioLine 根据配置初始化GPIO引脚
func (s *sGpiodProvider) initializeGpioLine() error {
	pinOffset := uint32(s.gpiodConfig.Pin)

	// 根据方向配置初始化引脚
	switch s.gpiodConfig.Direction {
	case c_enum.EGpioDirectionIn:
		// 输入模式：配置为输入
		// 使用GPIO v1 UAPI请求输入引脚
		req := gpioHandleRequest{
			LineOffsets:   [64]uint32{pinOffset},
			Flags:         GPIOHANDLE_REQUEST_INPUT,
			DefaultValues: [64]uint8{0},
			ConsumerLabel: [32]byte{'e', 'm', 's', '-', 'p', 'l', 'a', 'n'},
			Lines:         1,
		}

		err := s.gpioIoctl(GPIO_GET_LINEHANDLE_IOCTL, unsafe.Pointer(&req))
		if err != nil {
			return fmt.Errorf("请求输入引脚失败: %v", err)
		}

		// 创建文件描述符
		s.lineFd = os.NewFile(uintptr(req.Fd), fmt.Sprintf("gpio-line-%d", pinOffset))

	case c_enum.EGpioDirectionOut:
		// 输出模式：配置为输出，初始值为低电平
		req := gpioHandleRequest{
			LineOffsets:   [64]uint32{pinOffset},
			Flags:         GPIOHANDLE_REQUEST_OUTPUT,
			DefaultValues: [64]uint8{0}, // 初始值为低电平
			ConsumerLabel: [32]byte{'e', 'm', 's', '-', 'p', 'l', 'a', 'n'},
			Lines:         1,
		}

		err := s.gpioIoctl(GPIO_GET_LINEHANDLE_IOCTL, unsafe.Pointer(&req))
		if err != nil {
			return fmt.Errorf("请求输出引脚失败: %v", err)
		}

		// 创建文件描述符
		s.lineFd = os.NewFile(uintptr(req.Fd), fmt.Sprintf("gpio-line-%d", pinOffset))

	default:
		return fmt.Errorf("不支持的GPIO方向: %s", s.gpiodConfig.Direction)
	}

	return nil
}

// gpioIoctl 执行GPIO ioctl操作
func (s *sGpiodProvider) gpioIoctl(cmd uintptr, arg unsafe.Pointer) error {
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, s.chipFd.Fd(), cmd, uintptr(arg))
	if errno != 0 {
		return fmt.Errorf("ioctl failed: %v", errno)
	}
	return nil
}

// cleanup 清理资源
func (s *sGpiodProvider) cleanup() {
	if s.cancel != nil {
		s.cancel()
	}

	if s.lineFd != nil {
		s.lineFd.Close()
		s.lineFd = nil
	}

	if s.chipFd != nil {
		s.chipFd.Close()
		s.chipFd = nil
	}

	if s.eventChannel != nil {
		close(s.eventChannel)
		s.eventChannel = nil
	}

	s.isInitialized = false
}

// eventListener 事件监听协程
func (s *sGpiodProvider) eventListener() {
	if s.gpiodConfig.Direction != c_enum.EGpioDirectionIn {
		return // 只有输入模式才需要监听事件
	}

	ticker := time.NewTicker(100 * time.Millisecond) // 100ms轮询间隔
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			// 读取当前GPIO状态
			status, err := s.readGpioStatus()
			if err != nil {
				continue
			}

			// 检查状态是否发生变化
			s.mu.Lock()
			statusChanged := s.lastStatus == nil || *s.lastStatus != status
			if statusChanged {
				s.lastStatus = &status
				s.lastUpdateTime = time.Now()

				// 通知状态变化
				if s.handler != nil {
					go s.handler(status)
				}
			}
			s.mu.Unlock()
		}
	}
}

// readGpioStatus 读取GPIO状态
func (s *sGpiodProvider) readGpioStatus() (bool, error) {
	if s.lineFd == nil {
		return false, fmt.Errorf("GPIO引脚未初始化")
	}

	// 使用ioctl读取GPIO值
	data := gpioHandleData{}
	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, s.lineFd.Fd(), GPIOHANDLE_GET_LINE_VALUES_IOCTL, uintptr(unsafe.Pointer(&data)))
	if errno != 0 {
		return false, fmt.Errorf("读取GPIO值失败: %v", errno)
	}

	value := data.Values[0]

	// 根据低电平有效配置调整返回值
	if s.gpiodConfig.LowActive {
		return value == 0, nil // 低电平有效时，0表示高电平
	}
	return value == 1, nil // 高电平有效时，1表示高电平
}

var _ c_proto.IGpiodProtocol = (*sGpiodProvider)(nil)

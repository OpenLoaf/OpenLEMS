package internal

import (
	"common/c_base"
	"common/c_enum"
	"common/c_proto"
	"fmt"
	"sync"
	"time"

	"github.com/warthog618/go-gpiocdev"
)

type sGpiodProvider struct {
	c_base.IAlarm
	gpiodConfig *c_proto.SGpiodProtocolConfig

	// GPIO相关资源
	line     *gpiocdev.Line
	chipName string

	// 状态管理
	currentStatus *bool
	lastUpdate    *time.Time
	handler       func(status bool)

	// 并发控制
	mu sync.RWMutex
}

// NewGpiodProvider 创建新的GPIO provider
func NewGpiodProvider(config *c_proto.SGpiodProtocolConfig, alarm c_base.IAlarm) *sGpiodProvider {
	return &sGpiodProvider{
		IAlarm:      alarm,
		gpiodConfig: config,
		chipName:    fmt.Sprintf("gpiochip%d", config.ChipIndex),
	}
}

// initializeGPIO 初始化GPIO引脚
func (s *sGpiodProvider) initializeGPIO() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 如果已经初始化，先清理
	if s.line != nil {
		s.cleanupGPIO()
	}

	// 根据配置直接请求GPIO引脚
	var line *gpiocdev.Line
	var err error
	pinOffset := int(s.gpiodConfig.Pin)

	switch s.gpiodConfig.Direction {
	case c_enum.EGpioDirectionIn:
		// 输入引脚配置
		options := []gpiocdev.LineReqOption{
			gpiocdev.AsInput,
		}

		// 如果配置了低电平有效，设置active low
		if s.gpiodConfig.LowActive {
			options = append(options, gpiocdev.AsActiveLow)
		}

		// 如果设置了状态变化处理函数，添加边沿检测
		if s.handler != nil {
			options = append(options, gpiocdev.WithBothEdges, gpiocdev.WithEventHandler(s.handleGPIOEvent))
		}

		line, err = gpiocdev.RequestLine(s.chipName, pinOffset, options...)

	case c_enum.EGpioDirectionOut:
		// 输出引脚配置
		options := []gpiocdev.LineReqOption{
			gpiocdev.AsOutput(0), // 初始值设为低电平
		}

		// 如果配置了低电平有效，设置active low
		if s.gpiodConfig.LowActive {
			options = append(options, gpiocdev.AsActiveLow)
		}

		line, err = gpiocdev.RequestLine(s.chipName, pinOffset, options...)

	default:
		return fmt.Errorf("unsupported GPIO direction: %v", s.gpiodConfig.Direction)
	}

	if err != nil {
		return fmt.Errorf("failed to request GPIO line %d from %s: %w", pinOffset, s.chipName, err)
	}

	s.line = line
	now := time.Now()
	s.lastUpdate = &now

	return nil
}

// cleanupGPIO 清理GPIO资源
func (s *sGpiodProvider) cleanupGPIO() {
	if s.line != nil {
		s.line.Close()
		s.line = nil
	}
}

// handleGPIOEvent 处理GPIO状态变化事件
func (s *sGpiodProvider) handleGPIOEvent(evt gpiocdev.LineEvent) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 更新状态
	status := evt.Type == gpiocdev.LineEventRisingEdge
	s.currentStatus = &status
	now := time.Now()
	s.lastUpdate = &now

	// 调用处理函数
	if s.handler != nil {
		go s.handler(status)
	}
}

func (s *sGpiodProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 如果GPIO未初始化，返回断开状态
	if s.line == nil {
		return c_enum.EProtocolDisconnected
	}

	// 如果最后更新时间超过5分钟，认为连接可能有问题
	if s.lastUpdate != nil && time.Since(*s.lastUpdate) > 5*time.Minute {
		return c_enum.EProtocolConnecting
	}

	return c_enum.EProtocolConnected
}

func (s *sGpiodProvider) GetLastUpdateTime() *time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lastUpdate
}

func (s *sGpiodProvider) GetPointValueList() []*c_base.SPointValue {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 如果GPIO未初始化，返回空列表
	if s.line == nil {
		return []*c_base.SPointValue{}
	}

	// 返回当前GPIO状态
	status := s.GetGpioStatus()
	if status == nil {
		return []*c_base.SPointValue{}
	}

	value := 0
	if *status {
		value = 1
	}

	// 使用NewPointValue构造函数创建SPointValue
	pointValue := c_base.NewPointValue("gpiod-device", nil, value)
	if s.lastUpdate != nil {
		pointValue.SetHappenTime(*s.lastUpdate)
	}

	return []*c_base.SPointValue{pointValue}
}

func (s *sGpiodProvider) GetValue(point c_base.IPoint) (any, error) {
	status := s.GetGpioStatus()
	if status == nil {
		return nil, fmt.Errorf("GPIO not initialized or unavailable")
	}

	// 返回布尔值或数值
	if *status {
		return 1, nil
	}
	return 0, nil
}

func (s *sGpiodProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	// GPIO协议通常不需要任务注册，因为它是事件驱动的
	// 如果需要定期读取状态，可以在这里实现
}

func (s *sGpiodProvider) ProtocolListen() {
	// 初始化GPIO
	if err := s.initializeGPIO(); err != nil {
		// 记录错误，但不阻塞程序运行
		fmt.Printf("Failed to initialize GPIO: %v\n", err)
		return
	}

	// 如果是输入引脚，读取初始状态
	if s.gpiodConfig.Direction == c_enum.EGpioDirectionIn && s.line != nil {
		if value, err := s.line.Value(); err == nil {
			s.mu.Lock()
			status := value == 1
			s.currentStatus = &status
			now := time.Now()
			s.lastUpdate = &now
			s.mu.Unlock()
		}
	}
}

func (s *sGpiodProvider) RegisterHandler(handler func(status bool)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.handler = handler

	// 如果GPIO已经初始化且是输入引脚，需要重新初始化以添加事件处理
	if s.line != nil && s.gpiodConfig.Direction == c_enum.EGpioDirectionIn {
		// 重新初始化以添加事件处理
		go func() {
			if err := s.initializeGPIO(); err != nil {
				fmt.Printf("Failed to reinitialize GPIO with handler: %v\n", err)
			}
		}()
	}
}

func (s *sGpiodProvider) GetGpioStatus() *bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 如果GPIO未初始化，返回nil
	if s.line == nil {
		return nil
	}

	// 如果是输入引脚，实时读取状态
	if s.gpiodConfig.Direction == c_enum.EGpioDirectionIn {
		if value, err := s.line.Value(); err == nil {
			status := value == 1
			return &status
		}
	}

	// 返回缓存的状态
	return s.currentStatus
}

func (s *sGpiodProvider) SetHigh() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查是否是输出引脚
	if s.gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		return fmt.Errorf("cannot set value on input GPIO pin")
	}

	// 检查GPIO是否已初始化
	if s.line == nil {
		return fmt.Errorf("GPIO not initialized")
	}

	// 设置高电平
	if err := s.line.SetValue(1); err != nil {
		return fmt.Errorf("failed to set GPIO high: %w", err)
	}

	// 更新状态
	status := true
	s.currentStatus = &status
	now := time.Now()
	s.lastUpdate = &now

	return nil
}

func (s *sGpiodProvider) SetLow() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查是否是输出引脚
	if s.gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		return fmt.Errorf("cannot set value on input GPIO pin")
	}

	// 检查GPIO是否已初始化
	if s.line == nil {
		return fmt.Errorf("GPIO not initialized")
	}

	// 设置低电平
	if err := s.line.SetValue(0); err != nil {
		return fmt.Errorf("failed to set GPIO low: %w", err)
	}

	// 更新状态
	status := false
	s.currentStatus = &status
	now := time.Now()
	s.lastUpdate = &now

	return nil
}

// Close 关闭GPIO资源
func (s *sGpiodProvider) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.cleanupGPIO()
	return nil
}

var _ c_proto.IGpiodProtocol = (*sGpiodProvider)(nil)

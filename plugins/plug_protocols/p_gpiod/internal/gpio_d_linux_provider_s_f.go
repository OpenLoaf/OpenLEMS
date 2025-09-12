//go:build linux

package internal

import (
	"common/c_base"
	"common/c_device"
	"common/c_enum"
	"common/c_log"
	"common/c_proto"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/warthog618/go-gpiocdev"
)

type sGpiodProvider struct {
	c_base.IAlarm

	ctx context.Context
	mu  sync.RWMutex

	gpiodConfig  *c_proto.SGpiodProtocolConfig
	deviceConfig *c_base.SDeviceConfig

	// GPIO相关资源
	line     *gpiocdev.Line
	chipName string

	// 状态管理
	currentStatus  *bool
	lastUpdateTime *time.Time
	handler        func(status bool)
	protocolStatus c_enum.EProtocolStatus
}

var _ c_proto.IGpiodProtocol = (*sGpiodProvider)(nil)

// NewGpiodProvider 创建新的GPIO provider
func NewGpiodProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	// 解析协议配置
	gpiodConfig := &c_proto.SGpiodProtocolConfig{}
	if err := gconv.Scan(clientConfig.Params, gpiodConfig); err != nil {
		return nil, fmt.Errorf("failed to parse gpiod protocol config: %w", err)
	}

	provider := &sGpiodProvider{
		IAlarm:         c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		gpiodConfig:    gpiodConfig,
		chipName:       fmt.Sprintf("gpiochip%d", gpiodConfig.ChipIndex),
		ctx:            ctx,
		deviceConfig:   deviceConfig,
		protocolStatus: c_enum.EProtocolDisconnected,
	}

	return provider, nil
}

func (s *sGpiodProvider) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

// initializeGPIO 初始化GPIO引脚
func (s *sGpiodProvider) initializeGPIO() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 如果已经初始化，先清理
	if s.line != nil {
		s.cleanupGPIO()
	}

	pinOffset := int(s.gpiodConfig.Pin)
	var options []gpiocdev.LineReqOption

	switch s.gpiodConfig.Direction {
	case c_enum.EGpioDirectionIn:
		options = []gpiocdev.LineReqOption{gpiocdev.AsInput}
		if s.gpiodConfig.LowActive {
			options = append(options, gpiocdev.AsActiveLow)
		}
		if s.handler != nil {
			options = append(options, gpiocdev.WithBothEdges, gpiocdev.WithEventHandler(s.handleGPIOEvent))
		}

	case c_enum.EGpioDirectionOut:
		options = []gpiocdev.LineReqOption{gpiocdev.AsOutput(0)}
		if s.gpiodConfig.LowActive {
			options = append(options, gpiocdev.AsActiveLow)
		}

	default:
		return fmt.Errorf("unsupported GPIO direction: %v", s.gpiodConfig.Direction)
	}

	line, err := gpiocdev.RequestLine(s.chipName, pinOffset, options...)
	if err != nil {
		return fmt.Errorf("failed to request GPIO line %d from %s: %w", pinOffset, s.chipName, err)
	}

	s.line = line
	s.protocolStatus = c_enum.EProtocolConnected
	return nil
}

// cleanupGPIO 清理GPIO资源
func (s *sGpiodProvider) cleanupGPIO() {
	if s.line != nil {
		s.line.Close()
		s.line = nil
	}
	s.protocolStatus = c_enum.EProtocolDisconnected
}

// handleGPIOEvent 处理GPIO状态变化事件
func (s *sGpiodProvider) handleGPIOEvent(evt gpiocdev.LineEvent) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 更新状态
	status := evt.Type == gpiocdev.LineEventRisingEdge
	s.updateStatus(status)

	// 调用处理函数
	if s.handler != nil {
		go s.handler(status)
	}
}

func (s *sGpiodProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.protocolStatus
}

func (s *sGpiodProvider) GetLastUpdateTime() *time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lastUpdateTime
}

func (s *sGpiodProvider) GetPointValueList() []*c_base.SPointValue {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 如果协议未连接，返回空列表
	if s.protocolStatus != c_enum.EProtocolConnected {
		return []*c_base.SPointValue{}
	}

	// 返回当前GPIO状态
	status := s.getGpioStatusUnsafe()
	if status == nil {
		return []*c_base.SPointValue{}
	}

	pointValue := c_base.NewPointValue(s.deviceConfig.Id, nil, *status)
	if s.lastUpdateTime != nil {
		pointValue.SetHappenTime(*s.lastUpdateTime)
	}

	return []*c_base.SPointValue{pointValue}
}

func (s *sGpiodProvider) GetValue(point c_base.IPoint) (any, error) {
	if s.protocolStatus != c_enum.EProtocolConnected {
		return nil, fmt.Errorf("GPIO protocol not connected")
	}

	status := s.GetGpioStatus()
	if status == nil {
		return nil, fmt.Errorf("GPIO status unavailable")
	}

	return *status, nil
}

func (s *sGpiodProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	// GPIO协议通常不需要任务注册，因为它是事件驱动的
	// 如果需要定期读取状态，可以在这里实现
}

func (s *sGpiodProvider) ProtocolListen() {
	c_log.Infof(s.ctx, "Starting GPIO protocol listen on chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)

	// 初始化GPIO
	if err := s.initializeGPIO(); err != nil {
		// 记录错误，但不阻塞程序运行
		c_log.Error(s.ctx, "Failed to initialize GPIO", err)
		return
	}

	c_log.Infof(s.ctx, "GPIO initialized successfully on chip %s, pin %d, direction: %v",
		s.chipName, s.gpiodConfig.Pin, s.gpiodConfig.Direction)

	// 如果是输入引脚，读取初始状态
	if s.gpiodConfig.Direction == c_enum.EGpioDirectionIn && s.line != nil {
		if value, err := s.line.Value(); err == nil {
			s.mu.Lock()
			status := value == 1
			s.updateStatus(status)
			s.mu.Unlock()
			c_log.Debugf(s.ctx, "Initial GPIO state: %v", status)
		} else {
			c_log.Warning(s.ctx, "Failed to read initial GPIO value", err)
		}
	}
}

func (s *sGpiodProvider) RegisterHandler(handler func(status bool)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c_log.Infof(s.ctx, "Registering GPIO event handler for chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)
	s.handler = handler

	// 如果GPIO已经连接且是输入引脚，需要重新初始化以添加事件处理
	if s.protocolStatus == c_enum.EProtocolConnected && s.gpiodConfig.Direction == c_enum.EGpioDirectionIn {
		c_log.Debug(s.ctx, "Reinitializing GPIO to add event handler")
		go func() {
			if err := s.initializeGPIO(); err != nil {
				c_log.Error(s.ctx, "Failed to reinitialize GPIO with handler", err)
			} else {
				c_log.Info(s.ctx, "GPIO reinitialized with event handler successfully")
			}
		}()
	}
}

func (s *sGpiodProvider) GetGpioStatus() *bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.getGpioStatusUnsafe()
}

// getGpioStatusUnsafe 获取GPIO状态（不加锁，调用者需确保已加锁）
func (s *sGpiodProvider) getGpioStatusUnsafe() *bool {
	// 如果协议未连接，返回nil
	if s.protocolStatus != c_enum.EProtocolConnected {
		return nil
	}

	// 实时读取GPIO状态（输入和输出引脚都读取）
	if value, err := s.line.Value(); err == nil {
		status := value == 1
		return &status
	}

	// 如果读取失败，返回缓存的状态
	return s.currentStatus
}

// updateStatus 更新GPIO状态和时间戳（不加锁，调用者需确保已加锁）
func (s *sGpiodProvider) updateStatus(status bool) {
	s.currentStatus = &status
	now := time.Now()
	s.lastUpdateTime = &now
}

func (s *sGpiodProvider) SetHigh() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查是否是输出引脚
	if s.gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		c_log.Warningf(s.ctx, "Attempted to set high on input GPIO pin %d", s.gpiodConfig.Pin)
		return fmt.Errorf("cannot set value on input GPIO pin")
	}

	// 检查协议是否已连接
	if s.protocolStatus != c_enum.EProtocolConnected {
		c_log.Error(s.ctx, "GPIO protocol not connected when trying to set high")
		return fmt.Errorf("GPIO protocol not connected")
	}

	// 设置高电平
	if err := s.line.SetValue(1); err != nil {
		c_log.Errorf(s.ctx, "Failed to set GPIO high on pin %d: %v", s.gpiodConfig.Pin, err)
		return fmt.Errorf("failed to set GPIO high: %w", err)
	}

	// 读取实际值确认设置成功
	if actualValue, err := s.line.Value(); err == nil {
		actualStatus := actualValue == 1
		s.updateStatus(actualStatus)
		c_log.Debugf(s.ctx, "GPIO pin %d set to high, actual value: %v", s.gpiodConfig.Pin, actualStatus)
	} else {
		// 如果读取失败，使用期望值更新状态
		s.updateStatus(true)
		c_log.Warningf(s.ctx, "Failed to read GPIO value after setting high on pin %d: %v", s.gpiodConfig.Pin, err)
	}

	c_log.Debugf(s.ctx, "GPIO pin %d set to high", s.gpiodConfig.Pin)
	return nil
}

func (s *sGpiodProvider) SetLow() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 检查是否是输出引脚
	if s.gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		c_log.Warningf(s.ctx, "Attempted to set low on input GPIO pin %d", s.gpiodConfig.Pin)
		return fmt.Errorf("cannot set value on input GPIO pin")
	}

	// 检查协议是否已连接
	if s.protocolStatus != c_enum.EProtocolConnected {
		c_log.Error(s.ctx, "GPIO protocol not connected when trying to set low")
		return fmt.Errorf("GPIO protocol not connected")
	}

	// 设置低电平
	if err := s.line.SetValue(0); err != nil {
		c_log.Errorf(s.ctx, "Failed to set GPIO low on pin %d: %v", s.gpiodConfig.Pin, err)
		return fmt.Errorf("failed to set GPIO low: %w", err)
	}

	// 读取实际值确认设置成功
	if actualValue, err := s.line.Value(); err == nil {
		actualStatus := actualValue == 1
		s.updateStatus(actualStatus)
		c_log.Debugf(s.ctx, "GPIO pin %d set to low, actual value: %v", s.gpiodConfig.Pin, actualStatus)
	} else {
		// 如果读取失败，使用期望值更新状态
		s.updateStatus(false)
		c_log.Warningf(s.ctx, "Failed to read GPIO value after setting low on pin %d: %v", s.gpiodConfig.Pin, err)
	}

	c_log.Debugf(s.ctx, "GPIO pin %d set to low", s.gpiodConfig.Pin)
	return nil
}

// Close 关闭GPIO资源
func (s *sGpiodProvider) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	c_log.Infof(s.ctx, "Closing GPIO resources for chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)
	s.cleanupGPIO()
	return nil
}

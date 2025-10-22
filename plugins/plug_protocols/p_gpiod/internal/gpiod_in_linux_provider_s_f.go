//go:build linux && gpio_enable

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

type sGpiodInLinuxProvider struct {
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
	point          c_base.IPoint
	lastUpdateTime *time.Time
	handler        func(status bool, isChange bool)
	protocolStatus c_enum.EProtocolStatus
}

var _ c_proto.IGpiodProtocol = (*sGpiodInLinuxProvider)(nil)

// NewGpiodInProvider 创建新的GPIO输入provider
func NewGpiodInProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	// 解析协议配置
	gpiodConfig := &c_proto.SGpiodProtocolConfig{}
	if err := gconv.Scan(clientConfig.Params, gpiodConfig); err != nil {
		return nil, fmt.Errorf("failed to parse gpiod protocol config: %w", err)
	}

	// 验证方向必须是输入
	if gpiodConfig.Direction != c_enum.EGpioDirectionIn {
		return nil, fmt.Errorf("gpiod in provider only supports input direction, got: %v", gpiodConfig.Direction)
	}

	provider := &sGpiodInLinuxProvider{
		IAlarm:         c_device.NewAlarmImpl(ctx, deviceConfig.Id, deviceConfig.Pid),
		gpiodConfig:    gpiodConfig,
		chipName:       fmt.Sprintf("gpiochip%d", gpiodConfig.ChipIndex),
		ctx:            ctx,
		deviceConfig:   deviceConfig,
		protocolStatus: c_enum.EProtocolDisconnected,
	}

	// 初始化GPIO
	if err := provider.initializeGPIO(); err != nil {
		return nil, err
	}

	return provider, nil
}

func (s *sGpiodInLinuxProvider) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *sGpiodInLinuxProvider) InitGpioPoint(point c_base.IPoint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.point = point
}

// initializeGPIO 初始化GPIO输入引脚
func (s *sGpiodInLinuxProvider) initializeGPIO() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 如果已经初始化，先清理
	if s.line != nil {
		s.cleanupGPIO()
	}

	pinOffset := int(s.gpiodConfig.Pin)
	var options []gpiocdev.LineReqOption

	// 输入引脚配置
	options = []gpiocdev.LineReqOption{gpiocdev.AsInput}
	if s.gpiodConfig.LowActive {
		options = append(options, gpiocdev.AsActiveLow)
	}
	options = append(options, gpiocdev.WithBothEdges, gpiocdev.WithEventHandler(s.handleGPIOEvent))

	line, err := gpiocdev.RequestLine(s.chipName, pinOffset, options...)
	if err != nil {
		return fmt.Errorf("failed to request GPIO input line %d from %s: %w", pinOffset, s.chipName, err)
	}

	s.line = line
	s.protocolStatus = c_enum.EProtocolConnected
	return nil
}

// cleanupGPIO 清理GPIO资源
func (s *sGpiodInLinuxProvider) cleanupGPIO() {
	if s.line != nil {
		s.line.Close()
		s.line = nil
	}
	s.protocolStatus = c_enum.EProtocolDisconnected
}

// handleGPIOEvent 处理GPIO状态变化事件
func (s *sGpiodInLinuxProvider) handleGPIOEvent(evt gpiocdev.LineEvent) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 更新状态并处理状态变化
	status := evt.Type == gpiocdev.LineEventRisingEdge
	c_log.Infof(s.ctx, "gpiochip event type: %v", evt.Type)
	s.updateStatus(status)
}

func (s *sGpiodInLinuxProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.protocolStatus
}

func (s *sGpiodInLinuxProvider) GetLastUpdateTime() *time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lastUpdateTime
}

func (s *sGpiodInLinuxProvider) GetProtocolPointValue(point c_base.IPoint) *c_base.SPointValue {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 如果协议未连接或没有点位，返回nil
	if s.protocolStatus != c_enum.EProtocolConnected || s.point == nil {
		return nil
	}

	// GPIO是单点位设备，检查point是否匹配当前点位
	if point == nil || point.GetKey() != s.point.GetKey() {
		return nil
	}

	// 返回当前GPIO状态
	status := s.GetStatusUnsafe()
	if status == nil {
		return nil
	}

	pointValue := c_base.NewPointValue(s.deviceConfig.Id, s.point, *status)
	if s.lastUpdateTime != nil {
		pointValue.SetHappenTime(*s.lastUpdateTime)
	}

	return pointValue
}

func (s *sGpiodInLinuxProvider) GetValue(point c_base.IPoint) (any, error) {
	if s.protocolStatus != c_enum.EProtocolConnected {
		return nil, fmt.Errorf("GPIO protocol not connected")
	}

	status := s.GetStatus()
	if status == nil {
		return nil, fmt.Errorf("GPIO status unavailable")
	}

	return *status, nil
}

func (s *sGpiodInLinuxProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	// GPIO输入协议通常不需要任务注册，因为它是事件驱动的
	// 如果需要定期读取状态，可以在这里实现
}

func (s *sGpiodInLinuxProvider) ProtocolListen() {
	c_log.Infof(s.ctx, "Starting GPIO Input [%s-%s] protocol listen on chip %s, pin %d", s.deviceConfig.Id, s.deviceConfig.ProtocolId, s.chipName, s.gpiodConfig.Pin)

	c_log.Infof(s.ctx, "GPIO Input initialized successfully on chip %s, pin %d",
		s.chipName, s.gpiodConfig.Pin)

	// 读取初始状态
	if s.line != nil {
		if value, err := s.line.Value(); err == nil {
			s.mu.Lock()
			status := value == 1
			s.updateStatus(status)
			s.mu.Unlock()
			c_log.Debugf(s.ctx, "Initial GPIO input state: %v", status)
		} else {
			c_log.Warning(s.ctx, "Failed to read initial GPIO input value", err)
		}
	}
}

func (s *sGpiodInLinuxProvider) RegisterHandler(handler func(status bool, isChange bool)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c_log.Infof(s.ctx, "Registering GPIO input event handler for chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)
	s.handler = handler

	// 如果GPIO已经连接，需要重新初始化以添加事件处理
	if s.protocolStatus == c_enum.EProtocolConnected {
		c_log.Debug(s.ctx, "Reinitializing GPIO input to add event handler")
		go func() {
			if err := s.initializeGPIO(); err != nil {
				c_log.Error(s.ctx, "Failed to reinitialize GPIO input with handler", err)
			} else {
				c_log.Info(s.ctx, "GPIO input reinitialized with event handler successfully")
			}
		}()
	}
}

func (s *sGpiodInLinuxProvider) GetStatus() *bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.GetStatusUnsafe()
}

// GetStatusUnsafe 获取GPIO状态（不加锁，调用者需确保已加锁）
func (s *sGpiodInLinuxProvider) GetStatusUnsafe() *bool {
	// 如果协议未连接，返回nil
	if s.protocolStatus != c_enum.EProtocolConnected {
		return nil
	}

	// 手动模式下直接从缓存读取状态
	if s.deviceConfig.ManualMode {
		return s.currentStatus
	}

	// 自动模式下实时读取GPIO状态
	if value, err := s.line.Value(); err == nil {
		status := value == 1
		// 更新缓存状态和时间戳
		s.currentStatus = &status
		now := time.Now()
		s.lastUpdateTime = &now
		return &status
	}

	// 如果读取失败，返回缓存的状态
	return s.currentStatus
}

// updateStatus 更新GPIO状态和时间戳，并处理状态变化（不加锁，调用者需确保已加锁）
func (s *sGpiodInLinuxProvider) updateStatus(status bool) {
	// 保存之前的状态
	last := s.currentStatus

	// 更新状态和时间戳
	s.currentStatus = &status
	now := time.Now()
	s.lastUpdateTime = &now

	// 更新告警状态
	if s.point != nil {
		s.UpdateAlarm(s.deviceConfig.Id, s.point, status)
	}

	// 判断是否有状态变化
	isChange := false
	if last == nil {
		isChange = true
	} else if *last != status {
		isChange = true
	}

	// 调用处理函数
	if s.handler != nil {
		go s.handler(status, isChange)
	}
}

func (s *sGpiodInLinuxProvider) SetHigh() error {
	// 检查是否为手动模式
	if !s.deviceConfig.ManualMode {
		c_log.Warningf(s.ctx, "Attempted to set high on input GPIO pin %d - operation not supported in auto mode", s.gpiodConfig.Pin)
		return fmt.Errorf("cannot set value on input GPIO pin in auto mode")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// 手动模式下直接更新缓存状态
	status := true
	s.updateStatus(status)
	c_log.Infof(s.ctx, "Manually set GPIO input pin %d to HIGH in manual mode", s.gpiodConfig.Pin)
	return nil
}

func (s *sGpiodInLinuxProvider) SetLow() error {
	// 检查是否为手动模式
	if !s.deviceConfig.ManualMode {
		c_log.Warningf(s.ctx, "Attempted to set low on input GPIO pin %d - operation not supported in auto mode", s.gpiodConfig.Pin)
		return fmt.Errorf("cannot set value on input GPIO pin in auto mode")
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	// 手动模式下直接更新缓存状态
	status := false
	s.updateStatus(status)
	c_log.Infof(s.ctx, "Manually set GPIO input pin %d to LOW in manual mode", s.gpiodConfig.Pin)
	return nil
}

// Close 关闭GPIO资源
func (s *sGpiodInLinuxProvider) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	c_log.Infof(s.ctx, "Closing GPIO input resources for chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)
	s.cleanupGPIO()
	return nil
}

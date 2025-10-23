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

type sGpiodOutLinuxProvider struct {
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

var _ c_proto.IGpiodProtocol = (*sGpiodOutLinuxProvider)(nil)

// NewGpiodOutProvider 创建新的GPIO输出provider
func NewGpiodOutProvider(ctx context.Context, clientConfig *c_base.SProtocolConfig, deviceConfig *c_base.SDeviceConfig) (c_proto.IGpiodProtocol, error) {
	// 解析协议配置
	gpiodConfig := &c_proto.SGpiodProtocolConfig{}
	if err := gconv.Scan(clientConfig.Params, gpiodConfig); err != nil {
		return nil, fmt.Errorf("failed to parse gpiod protocol config: %w", err)
	}

	// 验证方向必须是输出
	if gpiodConfig.Direction != c_enum.EGpioDirectionOut {
		return nil, fmt.Errorf("gpiod out provider only supports output direction, got: %v", gpiodConfig.Direction)
	}

	provider := &sGpiodOutLinuxProvider{
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

func (s *sGpiodOutLinuxProvider) GetConfig() *c_base.SDeviceConfig {
	return s.deviceConfig
}

func (s *sGpiodOutLinuxProvider) InitGpioPoint(point c_base.IPoint) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.point = point
}

// initializeGPIO 初始化GPIO输出引脚
func (s *sGpiodOutLinuxProvider) initializeGPIO() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 如果已经初始化，先清理
	if s.line != nil {
		s.cleanupGPIO()
	}

	pinOffset := int(s.gpiodConfig.Pin)
	var options []gpiocdev.LineReqOption

	// 输出引脚配置，初始值为低电平
	options = []gpiocdev.LineReqOption{gpiocdev.AsOutput(0)}
	if s.gpiodConfig.LowActive {
		options = append(options, gpiocdev.AsActiveLow)
	}

	line, err := gpiocdev.RequestLine(s.chipName, pinOffset, options...)
	if err != nil {
		return fmt.Errorf("failed to request GPIO output line %d from %s: %w", pinOffset, s.chipName, err)
	}

	s.line = line
	s.protocolStatus = c_enum.EProtocolConnected

	// 初始化状态为低电平
	s.currentStatus = &[]bool{false}[0]
	now := time.Now()
	s.lastUpdateTime = &now

	return nil
}

// cleanupGPIO 清理GPIO资源
func (s *sGpiodOutLinuxProvider) cleanupGPIO() {
	if s.line != nil {
		s.line.Close()
		s.line = nil
	}
	s.protocolStatus = c_enum.EProtocolDisconnected
}

func (s *sGpiodOutLinuxProvider) GetProtocolStatus() c_enum.EProtocolStatus {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.protocolStatus
}

func (s *sGpiodOutLinuxProvider) GetLastUpdateTime() *time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.lastUpdateTime
}

func (s *sGpiodOutLinuxProvider) GetProtocolPointValue(point c_base.IPoint) *c_base.SPointValue {
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

func (s *sGpiodOutLinuxProvider) GetValue(point c_base.IPoint) (any, error) {
	if s.protocolStatus != c_enum.EProtocolConnected {
		return nil, fmt.Errorf("GPIO protocol not connected")
	}

	status := s.GetStatus()
	if status == nil {
		return nil, fmt.Errorf("GPIO status unavailable")
	}

	return *status, nil
}

func (s *sGpiodOutLinuxProvider) RegisterTask(task c_base.IPointTask, tasks ...c_base.IPointTask) {
	// GPIO输出协议通常不需要任务注册
}

func (s *sGpiodOutLinuxProvider) ProtocolListen() {
	c_log.Infof(s.ctx, "Starting GPIO Output [%s-%s] protocol listen on chip %s, pin %d", s.deviceConfig.Id, s.deviceConfig.ProtocolId, s.chipName, s.gpiodConfig.Pin)

	c_log.Infof(s.ctx, "GPIO Output initialized successfully on chip %s, pin %d",
		s.chipName, s.gpiodConfig.Pin)
}

func (s *sGpiodOutLinuxProvider) RegisterHandler(handler func(status bool, isChange bool)) {
	s.mu.Lock()
	defer s.mu.Unlock()

	c_log.Infof(s.ctx, "Registering GPIO output handler for chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)
	s.handler = handler
}

func (s *sGpiodOutLinuxProvider) GetStatus() *bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.GetStatusUnsafe()
}

// GetStatusUnsafe 获取GPIO状态（不加锁，调用者需确保已加锁）
func (s *sGpiodOutLinuxProvider) GetStatusUnsafe() *bool {
	// 如果协议未连接，返回nil
	if s.protocolStatus != c_enum.EProtocolConnected {
		return nil
	}

	// 实时读取GPIO状态
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
func (s *sGpiodOutLinuxProvider) updateStatus(status bool) {
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

func (s *sGpiodOutLinuxProvider) SetHigh() error {
	s.mu.Lock()
	defer s.mu.Unlock()

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

func (s *sGpiodOutLinuxProvider) SetLow() error {
	s.mu.Lock()
	defer s.mu.Unlock()

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
func (s *sGpiodOutLinuxProvider) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	c_log.Infof(s.ctx, "Closing GPIO output resources for chip %s, pin %d", s.chipName, s.gpiodConfig.Pin)
	s.cleanupGPIO()
	return nil
}

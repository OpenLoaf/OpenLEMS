package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"math"
	"strings"
	"sync"
	"time"

	"common"
	"common/c_base"
	"common/c_enum"
	"common/c_log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/pkg/errors"
)

// SMqttClientStatus MQTT客户端状态结构体
type SMqttClientStatus struct {
	Config        *SMqttConfig `json:"config"`        // MQTT配置信息
	IsConnected   bool         `json:"isConnected"`   // 是否已连接
	IsRunning     bool         `json:"isRunning"`     // 是否正在运行
	ClientID      string       `json:"clientId"`      // 客户端ID
	SystemNumber  string       `json:"systemNumber"`  // 系统序列号
	Topic         string       `json:"topic"`         // 当前使用的Topic
	DeviceCount   int          `json:"deviceCount"`   // 设备数量
	UploadPeriod  int          `json:"uploadPeriod"`  // 上传周期（秒）
	LastPublishAt *time.Time   `json:"lastPublishAt"` // 最后发布时间
	PublishCount  int64        `json:"publishCount"`  // 发布消息总数
	ErrorCount    int64        `json:"errorCount"`    // 错误次数
	StartTime     *time.Time   `json:"startTime"`     // 启动时间
}

// SMqttClient MQTT客户端封装
type SMqttClient struct {
	config        *SMqttConfig       // MQTT配置
	client        mqtt.Client        // MQTT客户端
	ticker        *time.Ticker       // 定时器
	ctx           context.Context    // 上下文
	cancel        context.CancelFunc // 取消函数
	formatter     IDataFormatter     // 数据格式化器
	systemNumber  string             // 系统序列号
	isRunning     bool               // 运行状态
	startTime     *time.Time         // 启动时间
	lastPublishAt *time.Time         // 最后发布时间
	publishCount  int64              // 发布消息总数
	errorCount    int64              // 错误次数
	countMutex    sync.RWMutex       // 计数器读写锁，保证并发安全
}

// NewMqttClient 创建新的MQTT客户端
func NewMqttClient(config *SMqttConfig, formatter IDataFormatter, systemNumber string) *SMqttClient {
	return &SMqttClient{
		config:       config,
		formatter:    formatter,
		systemNumber: systemNumber,
	}
}

// safeIncrementPublishCount 安全递增发布计数器，防止溢出
func (c *SMqttClient) safeIncrementPublishCount() {
	c.countMutex.Lock()
	defer c.countMutex.Unlock()

	// 检查是否接近int64最大值，提前重置避免溢出
	if c.publishCount >= math.MaxInt64-1000 {
		c_log.Warningf(c.ctx, "MQTT发布计数器即将溢出，重置为0: 当前值=%d", c.publishCount)
		c.publishCount = 0
	} else {
		c.publishCount++
	}
}

// safeIncrementErrorCount 安全递增错误计数器，防止溢出
func (c *SMqttClient) safeIncrementErrorCount() {
	c.countMutex.Lock()
	defer c.countMutex.Unlock()

	// 检查是否接近int64最大值，提前重置避免溢出
	if c.errorCount >= math.MaxInt64-1000 {
		c_log.Warningf(c.ctx, "MQTT错误计数器即将溢出，重置为0: 当前值=%d", c.errorCount)
		c.errorCount = 0
	} else {
		c.errorCount++
	}
}

// getPublishCount 安全获取发布计数器值
func (c *SMqttClient) getPublishCount() int64 {
	c.countMutex.RLock()
	defer c.countMutex.RUnlock()
	return c.publishCount
}

// getErrorCount 安全获取错误计数器值
func (c *SMqttClient) getErrorCount() int64 {
	c.countMutex.RLock()
	defer c.countMutex.RUnlock()
	return c.errorCount
}

// Start 启动MQTT客户端
func (c *SMqttClient) Start(ctx context.Context) error {
	c.ctx, c.cancel = context.WithCancel(ctx)
	// 设置远程协议类型上下文
	c.ctx = context.WithValue(c.ctx, c_enum.ELogTypeRemote, c_base.ConstRemoteMqtt)

	// 创建MQTT客户端
	opts := mqtt.NewClientOptions()

	// 根据SSL配置选择协议
	var brokerURL string
	if c.config.UseSSL {
		brokerURL = fmt.Sprintf("ssl://%s:%d", c.config.ServerAddress, c.config.ServerPort)
	} else {
		brokerURL = fmt.Sprintf("tcp://%s:%d", c.config.ServerAddress, c.config.ServerPort)
	}
	opts.AddBroker(brokerURL)
	opts.SetClientID(fmt.Sprintf("lems_%s", c.systemNumber))
	opts.SetAutoReconnect(true)

	// 使用配置的超时和重连参数
	connectTimeout := time.Duration(c.config.ConnectTimeout) * time.Second
	reconnectInterval := time.Duration(c.config.ReconnectInterval) * time.Second
	keepAliveTimeout := time.Duration(c.config.KeepAliveTimeout) * time.Second

	opts.SetConnectTimeout(connectTimeout)
	opts.SetConnectRetryInterval(reconnectInterval)
	opts.SetMaxReconnectInterval(reconnectInterval)
	opts.SetKeepAlive(keepAliveTimeout)
	opts.SetPingTimeout(10 * time.Second)

	// 配置SSL/TLS（如果启用）
	if c.config.UseSSL {
		tlsConfig := &tls.Config{
			InsecureSkipVerify: c.config.InsecureSkipVerify,
		}
		opts.SetTLSConfig(tlsConfig)
	}

	// 设置用户名和密码认证（如果提供）
	if c.config.Username != "" {
		opts.SetUsername(c.config.Username)
		if c.config.Password != "" {
			opts.SetPassword(c.config.Password)
		}
	}

	// 设置连接丢失处理
	opts.SetConnectionLostHandler(func(client mqtt.Client, err error) {
		c_log.Warningf(c.ctx, "MQTT连接丢失: %v", err)
	})

	// 设置连接成功处理
	opts.SetOnConnectHandler(func(client mqtt.Client) {
		clientId := fmt.Sprintf("lems_%s", c.systemNumber)
		protocol := "TCP"
		if c.config.UseSSL {
			protocol = "SSL/TLS"
		}

		if c.config.Username != "" {
			c_log.BizInfof(c.ctx, "MQTT连接成功: %s:%d (ClientID: %s, 协议: %s, 用户: %s)", c.config.ServerAddress, c.config.ServerPort, clientId, protocol, c.config.Username)
		} else {
			c_log.BizInfof(c.ctx, "MQTT连接成功: %s:%d (ClientID: %s, 协议: %s, 匿名连接)", c.config.ServerAddress, c.config.ServerPort, clientId, protocol)
		}
	})

	c.client = mqtt.NewClient(opts)

	// 连接到MQTT服务器
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		clientId := fmt.Sprintf("lems_%s", c.systemNumber)
		authInfo := ""
		if c.config.Username != "" {
			authInfo = fmt.Sprintf(" (用户: %s)", c.config.Username)
		}
		return errors.Wrapf(token.Error(), "连接MQTT服务器失败: %s:%d (ClientID: %s)%s", c.config.ServerAddress, c.config.ServerPort, clientId, authInfo)
	}

	// 启动定时推送
	c.ticker = time.NewTicker(time.Duration(c.config.UploadPeriod) * time.Second)
	go c.publishLoop()

	// 设置运行状态和启动时间
	c.isRunning = true
	now := time.Now()
	c.startTime = &now

	clientId := fmt.Sprintf("lems_%s", c.systemNumber)
	c_log.BizInfof(c.ctx, "MQTT客户端启动成功: %s:%d (ClientID: %s), 推送周期: %d秒, 连接超时: %d秒, 重连间隔: %d秒, 保活超时: %d秒",
		c.config.ServerAddress, c.config.ServerPort, clientId, c.config.UploadPeriod,
		c.config.ConnectTimeout, c.config.ReconnectInterval, c.config.KeepAliveTimeout)
	return nil
}

// Stop 停止MQTT客户端
func (c *SMqttClient) Stop() error {
	if c.cancel != nil {
		c.cancel()
	}

	if c.ticker != nil {
		c.ticker.Stop()
	}

	if c.client != nil && c.client.IsConnected() {
		c.client.Disconnect(250)
	}

	// 设置运行状态为false
	c.isRunning = false

	clientId := fmt.Sprintf("lems_%s", c.systemNumber)
	c_log.BizInfof(c.ctx, "MQTT客户端已停止: %s:%d (ClientID: %s)", c.config.ServerAddress, c.config.ServerPort, clientId)
	return nil
}

// publishLoop 定时推送循环
func (c *SMqttClient) publishLoop() {
	for {
		select {
		case <-c.ctx.Done():
			c_log.Debugf(c.ctx, "MQTT推送循环已停止")
			return
		case <-c.ticker.C:
			if err := c.publishDeviceData(); err != nil {
				c_log.Errorf(c.ctx, "推送设备数据失败: %v", err)
			}
		}
	}
}

// publishDeviceData 推送设备数据
func (c *SMqttClient) publishDeviceData() error {
	if !c.client.IsConnected() {
		return errors.New("MQTT客户端未连接")
	}

	// 遍历配置的设备ID列表
	for _, deviceId := range c.config.DeviceIds {
		// 获取设备实例
		device := common.GetDeviceManager().GetDeviceById(deviceId)
		if device == nil {
			c_log.Warningf(c.ctx, "设备不存在: %s", deviceId)
			continue
		}

		// 检查设备连接状态
		if device.GetProtocolStatus() != common.GetDeviceManager().GetDeviceById(deviceId).GetProtocolStatus() {
			c_log.Debugf(c.ctx, "设备未连接，跳过推送: %s", deviceId)
			continue
		}

		// 获取设备所有点位数据
		deviceData := c_base.GetAllTelemetry(device)
		if len(deviceData) == 0 {
			c_log.Debugf(c.ctx, "设备无数据，跳过推送: %s", deviceId)
			continue
		}

		// 格式化数据
		jsonData, err := c.formatter.Format(c.ctx, deviceId, deviceData, c.systemNumber)
		if err != nil {
			c_log.Errorf(c.ctx, "格式化设备数据失败: 设备ID=%s, 错误=%v", deviceId, err)
			continue
		}

		// 构建topic
		topic := c.buildTopic()
		if topic == "" {
			c_log.Errorf(c.ctx, "构建topic失败: 设备ID=%s", deviceId)
			continue
		}

		// 发布MQTT消息
		if token := c.client.Publish(topic, 0, false, jsonData); token.Wait() && token.Error() != nil {
			c.safeIncrementErrorCount()
			c_log.Errorf(c.ctx, "发布MQTT消息失败: 设备ID=%s, Topic=%s, 错误=%v", deviceId, topic, token.Error())
			continue
		}

		// 更新发布统计
		c.safeIncrementPublishCount()
		now := time.Now()
		c.lastPublishAt = &now

		c_log.BizDebugf(c.ctx, "成功推送设备数据: 设备ID=%s, Topic=%s, 数据长度=%d", deviceId, topic, len(jsonData))
	}

	return nil
}

// buildTopic 构建MQTT topic
func (c *SMqttClient) buildTopic() string {
	// 如果配置了自定义pushChannel且rewriteChannel=true，使用自定义
	if c.config.RewriteChannel && c.config.PushChannel != "" {
		return c.config.PushChannel
	}

	// 使用formatter的模板，替换{system_number}为实际值
	topicTemplate := c.formatter.GetTopicTemplate()
	return strings.Replace(topicTemplate, "{system_number}", c.systemNumber, -1)
}

// GetStatus 获取MQTT客户端状态
func (c *SMqttClient) GetStatus() *SMqttClientStatus {
	clientId := fmt.Sprintf("lems_%s", c.systemNumber)
	topic := c.buildTopic()

	// 检查连接状态
	isConnected := false
	if c.client != nil {
		isConnected = c.client.IsConnected()
	}

	return &SMqttClientStatus{
		Config:        c.config,
		IsConnected:   isConnected,
		IsRunning:     c.isRunning,
		ClientID:      clientId,
		SystemNumber:  c.systemNumber,
		Topic:         topic,
		DeviceCount:   len(c.config.DeviceIds),
		UploadPeriod:  c.config.UploadPeriod,
		LastPublishAt: c.lastPublishAt,
		PublishCount:  c.getPublishCount(),
		ErrorCount:    c.getErrorCount(),
		StartTime:     c.startTime,
	}
}

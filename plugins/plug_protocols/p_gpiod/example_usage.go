package p_gpiod

import (
	"common/c_base"
	"common/c_enum"
	"common/c_proto"
	"fmt"
	"log"
	"time"

	"p_gpiod/internal"
)

// 示例告警实现
type exampleAlarm struct{}

func (e *exampleAlarm) GetAlarmLevel() c_enum.EAlarmLevel {
	return c_enum.EAlarmLevelNone
}

func (e *exampleAlarm) GetAlarmList() []*c_base.SPointValue {
	return nil
}

func (e *exampleAlarm) UpdateAlarm(deviceId string, point c_base.IPoint, value any) {
	// 示例实现
}

func (e *exampleAlarm) ResetAlarm() {
	// 示例实现
}

func (e *exampleAlarm) IgnoreClearAlarm(deviceId string, point string) {
	// 示例实现
}

func (e *exampleAlarm) RegisterAlarmHandlerFunc(alarmAction c_enum.EAlarmAction, handler func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isFirstHandler bool), sortValue ...int) {
	// 示例实现
}

func ExampleUsage() {
	// 创建GPIO配置
	config := &c_proto.SGpiodProtocolConfig{
		Direction: c_enum.EGpioDirectionOut, // 输出模式
		ChipIndex: 0,                        // gpiochip0
		Pin:       18,                       // GPIO18引脚
		LowActive: false,                    // 高电平有效
	}

	// 创建告警实例
	alarm := &exampleAlarm{}

	// 创建GPIO提供者
	provider := internal.NewGpiodProvider(alarm, config)

	// 初始化GPIO
	provider.ProtocolListen()

	// 等待初始化完成
	time.Sleep(100 * time.Millisecond)

	// 检查连接状态
	status := provider.GetProtocolStatus()
	fmt.Printf("GPIO连接状态: %v\n", status)

	if status == c_enum.EProtocolConnected {
		// 设置高电平
		err := provider.SetHigh()
		if err != nil {
			log.Printf("设置高电平失败: %v", err)
		} else {
			fmt.Println("GPIO设置为高电平")
		}

		// 读取状态
		gpioStatus := provider.GetGpioStatus()
		if gpioStatus != nil {
			fmt.Printf("当前GPIO状态: %v\n", *gpioStatus)
		}

		// 等待2秒
		time.Sleep(2 * time.Second)

		// 设置低电平
		err = provider.SetLow()
		if err != nil {
			log.Printf("设置低电平失败: %v", err)
		} else {
			fmt.Println("GPIO设置为低电平")
		}

		// 再次读取状态
		gpioStatus = provider.GetGpioStatus()
		if gpioStatus != nil {
			fmt.Printf("当前GPIO状态: %v\n", *gpioStatus)
		}
	}

	// 清理资源
	err := provider.Close()
	if err != nil {
		log.Printf("关闭GPIO提供者失败: %v", err)
	}

	fmt.Println("GPIO示例完成")
}

// InputExample 输入模式示例
func InputExample() {
	// 创建输入模式配置
	config := &c_proto.SGpiodProtocolConfig{
		Direction: c_enum.EGpioDirectionIn, // 输入模式
		ChipIndex: 0,                       // gpiochip0
		Pin:       24,                      // GPIO24引脚
		LowActive: false,                   // 高电平有效
	}

	// 创建告警实例
	alarm := &exampleAlarm{}

	// 创建GPIO提供者
	provider := internal.NewGpiodProvider(alarm, config)

	// 注册状态变化处理函数
	provider.RegisterHandler(func(status bool) {
		fmt.Printf("GPIO状态变化: %v\n", status)
	})

	// 初始化GPIO
	provider.ProtocolListen()

	// 等待初始化完成
	time.Sleep(100 * time.Millisecond)

	// 检查连接状态
	status := provider.GetProtocolStatus()
	fmt.Printf("GPIO连接状态: %v\n", status)

	if status == c_enum.EProtocolConnected {
		// 读取当前状态
		gpioStatus := provider.GetGpioStatus()
		if gpioStatus != nil {
			fmt.Printf("当前GPIO状态: %v\n", *gpioStatus)
		}

		// 监听状态变化10秒
		fmt.Println("监听GPIO状态变化10秒...")
		time.Sleep(10 * time.Second)
	}

	// 清理资源
	err := provider.Close()
	if err != nil {
		log.Printf("关闭GPIO提供者失败: %v", err)
	}

	fmt.Println("GPIO输入示例完成")
}

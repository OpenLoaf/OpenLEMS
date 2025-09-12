package protocol

import (
	"common"
	"common/c_base"
	"common/c_enum"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"s_db"
	"s_db/s_db_model"
	"strconv"
	"strings"

	v1 "application/api/protocol/v1"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"periph.io/x/conn/v3/gpio"
	"periph.io/x/conn/v3/gpio/gpioreg"
	"periph.io/x/host/v3"
)

// GetGpioSysfsScan 扫描所有 GPIO 协议（type = gpio）的 SourceAddress，并调用 ScanGpioSysfs 聚合返回
func (c *ControllerV1) GetGpioSysfsScan(ctx context.Context, req *v1.GetGpioSysfsScanReq) (res *v1.GetGpioSysfsScanRes, err error) {
	g.Log().Infof(ctx, "开始扫描GPIO Sysfs信息")

	// 1. 获取所有GPIO协议配置
	protocols, err := s_db.GetProtocolService().GetProtocolList(ctx, string(c_enum.EGpioSysfs))
	if err != nil {
		g.Log().Errorf(ctx, "获取GPIO协议列表失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "获取GPIO协议列表失败")
	}
	g.Log().Infof(ctx, "找到 %d 个GPIO协议配置", len(protocols))

	// 2. 初始化periph.io，包括sysfs驱动
	_, err = host.Init()
	if err != nil {
		g.Log().Errorf(ctx, "初始化periph.io失败: %+v", err)
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "初始化periph.io失败")
	}
	g.Log().Infof(ctx, "periph.io初始化成功")

	// 2.1 尝试初始化sysfs驱动（如果可用）
	// 注意：periph.io的sysfs驱动需要手动导入才能使用
	// 这里我们通过直接扫描sysfs文件系统来补充periph.io的不足

	// 3. 获取设备管理器
	deviceManager := common.GetDeviceManager()
	if deviceManager == nil {
		g.Log().Errorf(ctx, "设备管理器未初始化")
		return nil, gerror.NewCode(gcode.CodeInternalError, "设备管理器未初始化")
	}

	// 4. 扫描所有GPIO引脚
	var gpioList []*v1.SGpioInfo
	scannedPins := make(map[int]bool) // 避免重复扫描

	// 4.1 扫描已配置的GPIO协议
	for _, protocol := range protocols {
		gpioInfo, err := c.scanConfiguredGpio(ctx, protocol, deviceManager)
		if err != nil {
			g.Log().Warningf(ctx, "扫描已配置GPIO失败 [%s]: %+v", protocol.Id, err)
			continue
		}
		if gpioInfo != nil {
			gpioList = append(gpioList, gpioInfo)
			scannedPins[gpioInfo.PinNumber] = true
		}
	}

	// 4.2 扫描系统中其他可用的GPIO引脚
	// 在Linux系统上，优先使用sysfs扫描，因为它更可靠
	sysfsGpios, err := c.scanSysfsGpios(ctx, scannedPins)
	if err != nil {
		g.Log().Warningf(ctx, "sysfs扫描GPIO失败: %+v", err)
	} else {
		gpioList = append(gpioList, sysfsGpios...)
		g.Log().Infof(ctx, "通过sysfs发现 %d 个GPIO引脚", len(sysfsGpios))
	}

	// 尝试扫描GPIO芯片信息，发现更多可用的GPIO引脚
	chipGpios, err := c.scanGpioChips(ctx, scannedPins)
	if err != nil {
		g.Log().Warningf(ctx, "扫描GPIO芯片失败: %+v", err)
	} else if len(chipGpios) > 0 {
		gpioList = append(gpioList, chipGpios...)
		g.Log().Infof(ctx, "通过GPIO芯片发现 %d 个GPIO引脚", len(chipGpios))
	}

	// 最后尝试使用periph.io扫描（作为补充）
	availableGpios, err := c.scanAvailableGpios(ctx, scannedPins)
	if err != nil {
		g.Log().Warningf(ctx, "periph.io扫描可用GPIO失败: %+v", err)
	} else if len(availableGpios) > 0 {
		gpioList = append(gpioList, availableGpios...)
		g.Log().Infof(ctx, "通过periph.io发现 %d 个GPIO引脚", len(availableGpios))
	}

	// 5. 构建响应
	res = &v1.GetGpioSysfsScanRes{
		GpioList: gpioList,
		Total:    len(gpioList),
	}

	g.Log().Infof(ctx, "GPIO扫描完成，共发现 %d 个GPIO引脚", len(gpioList))
	return res, nil
}

// scanConfiguredGpio 扫描已配置的GPIO协议
func (c *ControllerV1) scanConfiguredGpio(ctx context.Context, protocol *s_db_model.SProtocolModel, deviceManager common.IDeviceManager) (*v1.SGpioInfo, error) {
	// 解析协议参数获取GPIO编号
	var gpioConfig struct {
		Io        int    `json:"io"`
		Direction string `json:"direction"`
	}

	if protocol.Params != "" {
		err := gconv.Scan(protocol.Params, &gpioConfig)
		if err != nil {
			return nil, fmt.Errorf("解析协议参数失败: %v", err)
		}
	}

	if gpioConfig.Io == 0 {
		return nil, fmt.Errorf("GPIO编号未配置")
	}

	// 获取GPIO引脚
	pinName := fmt.Sprintf("GPIO%d", gpioConfig.Io)
	pin := gpioreg.ByName(pinName)
	if pin == nil {
		return nil, fmt.Errorf("GPIO引脚 %s 不可用", pinName)
	}

	// 读取GPIO状态
	var value int
	var direction string
	var active bool

	// 尝试读取GPIO值
	level := pin.Read()
	value = 0
	if level == gpio.High {
		value = 1
	}

	// 设置方向
	if gpioConfig.Direction != "" {
		direction = gpioConfig.Direction
	} else {
		direction = "unknown"
	}

	active = true

	// 查找关联的设备
	var deviceId, deviceName string
	deviceManager.IteratorAllDevices(func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if config.ProtocolId == protocol.Id {
			deviceId = config.Id
			deviceName = config.Name
			return false // 停止迭代
		}
		return true
	})

	return &v1.SGpioInfo{
		PinNumber:  gpioConfig.Io,
		Direction:  direction,
		Value:      value,
		Active:     active,
		ProtocolId: protocol.Id,
		DeviceId:   deviceId,
		DeviceName: deviceName,
	}, nil
}

// scanAvailableGpios 扫描系统中其他可用的GPIO引脚
func (c *ControllerV1) scanAvailableGpios(ctx context.Context, scannedPins map[int]bool) ([]*v1.SGpioInfo, error) {
	var gpioList []*v1.SGpioInfo

	// 扫描常见的GPIO引脚范围
	for i := 0; i < 100; i++ {
		if scannedPins[i] {
			continue // 跳过已扫描的引脚
		}

		pinName := fmt.Sprintf("GPIO%d", i)
		pin := gpioreg.ByName(pinName)
		if pin != nil {
			// 尝试读取GPIO值
			var value int
			var active bool

			// 尝试读取GPIO状态
			level := pin.Read()
			value = 0
			if level == gpio.High {
				value = 1
			}
			active = true

			gpioInfo := &v1.SGpioInfo{
				PinNumber:  i,
				Direction:  "unknown",
				Value:      value,
				Active:     active,
				ProtocolId: "",
				DeviceId:   "",
				DeviceName: "",
			}
			gpioList = append(gpioList, gpioInfo)
		}
	}

	return gpioList, nil
}

// scanSysfsGpios 直接扫描Linux sysfs文件系统获取GPIO信息
func (c *ControllerV1) scanSysfsGpios(ctx context.Context, scannedPins map[int]bool) ([]*v1.SGpioInfo, error) {
	var gpioList []*v1.SGpioInfo

	// 扫描 /sys/class/gpio/ 目录
	gpioDir := "/sys/class/gpio"
	if _, err := os.Stat(gpioDir); os.IsNotExist(err) {
		g.Log().Debugf(ctx, "GPIO sysfs目录不存在: %s", gpioDir)
		return gpioList, nil
	}

	g.Log().Debugf(ctx, "开始扫描sysfs GPIO目录: %s", gpioDir)

	// 读取gpio目录下的所有文件
	files, err := os.ReadDir(gpioDir)
	if err != nil {
		return nil, fmt.Errorf("读取GPIO目录失败: %v", err)
	}

	// 正则表达式匹配gpio数字
	gpioRegex := regexp.MustCompile(`^gpio(\d+)$`)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// 检查是否是gpio数字目录
		matches := gpioRegex.FindStringSubmatch(file.Name())
		if len(matches) != 2 {
			continue
		}

		// 解析GPIO编号
		gpioNum, err := strconv.Atoi(matches[1])
		if err != nil {
			continue
		}

		// 跳过已扫描的引脚
		if scannedPins[gpioNum] {
			continue
		}

		// 读取GPIO信息
		gpioInfo, err := c.readGpioInfoFromSysfs(ctx, gpioNum, file.Name())
		if err != nil {
			g.Log().Debugf(ctx, "读取GPIO %d 信息失败: %v", gpioNum, err)
			continue
		}

		gpioList = append(gpioList, gpioInfo)
	}

	return gpioList, nil
}

// readGpioInfoFromSysfs 从sysfs文件系统读取GPIO信息
func (c *ControllerV1) readGpioInfoFromSysfs(ctx context.Context, gpioNum int, gpioName string) (*v1.SGpioInfo, error) {
	gpioPath := fmt.Sprintf("/sys/class/gpio/%s", gpioName)

	// 读取方向
	direction := "unknown"
	directionFile := filepath.Join(gpioPath, "direction")
	if data, err := os.ReadFile(directionFile); err == nil {
		direction = strings.TrimSpace(string(data))
	}

	// 读取当前值
	value := 0
	valueFile := filepath.Join(gpioPath, "value")
	if data, err := os.ReadFile(valueFile); err == nil {
		if val, err := strconv.Atoi(strings.TrimSpace(string(data))); err == nil {
			value = val
		}
	}

	// 检查GPIO是否可用（通过检查是否可以读取value文件）
	active := true
	if _, err := os.Stat(valueFile); os.IsNotExist(err) {
		active = false
	}

	return &v1.SGpioInfo{
		PinNumber:  gpioNum,
		Direction:  direction,
		Value:      value,
		Active:     active,
		ProtocolId: "",
		DeviceId:   "",
		DeviceName: "",
	}, nil
}

// scanGpioChips 扫描GPIO芯片信息，发现可用的GPIO引脚
func (c *ControllerV1) scanGpioChips(ctx context.Context, scannedPins map[int]bool) ([]*v1.SGpioInfo, error) {
	var gpioList []*v1.SGpioInfo

	// 扫描 /sys/class/gpio/ 目录下的gpiochip文件
	gpioDir := "/sys/class/gpio"
	if _, err := os.Stat(gpioDir); os.IsNotExist(err) {
		g.Log().Debugf(ctx, "GPIO目录不存在，无法扫描芯片: %s", gpioDir)
		return gpioList, nil
	}

	g.Log().Debugf(ctx, "开始扫描GPIO芯片信息")

	files, err := os.ReadDir(gpioDir)
	if err != nil {
		return nil, fmt.Errorf("读取GPIO目录失败: %v", err)
	}

	// 正则表达式匹配gpiochip
	chipRegex := regexp.MustCompile(`^gpiochip(\d+)$`)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		// 检查是否是gpiochip目录
		matches := chipRegex.FindStringSubmatch(file.Name())
		if len(matches) != 2 {
			continue
		}

		// 读取芯片信息
		chipInfo, err := c.readGpioChipInfo(ctx, file.Name())
		if err != nil {
			g.Log().Debugf(ctx, "读取GPIO芯片 %s 信息失败: %v", file.Name(), err)
			continue
		}

		// 为芯片中的每个GPIO引脚创建信息
		for i := chipInfo.Base; i < chipInfo.Base+chipInfo.Ngpio; i++ {
			if scannedPins[i] {
				continue // 跳过已扫描的引脚
			}

			// 尝试读取GPIO引脚信息
			gpioInfo, err := c.readGpioFromChip(ctx, i, chipInfo)
			if err != nil {
				continue // 如果无法读取，跳过这个引脚
			}

			gpioList = append(gpioList, gpioInfo)
		}
	}

	return gpioList, nil
}

// GpioChipInfo GPIO芯片信息
type GpioChipInfo struct {
	Name  string
	Base  int
	Ngpio int
	Label string
}

// readGpioChipInfo 读取GPIO芯片信息
func (c *ControllerV1) readGpioChipInfo(ctx context.Context, chipName string) (*GpioChipInfo, error) {
	chipPath := fmt.Sprintf("/sys/class/gpio/%s", chipName)

	// 读取base
	base := 0
	baseFile := filepath.Join(chipPath, "base")
	if data, err := os.ReadFile(baseFile); err == nil {
		if val, err := strconv.Atoi(strings.TrimSpace(string(data))); err == nil {
			base = val
		}
	}

	// 读取ngpio
	ngpio := 0
	ngpioFile := filepath.Join(chipPath, "ngpio")
	if data, err := os.ReadFile(ngpioFile); err == nil {
		if val, err := strconv.Atoi(strings.TrimSpace(string(data))); err == nil {
			ngpio = val
		}
	}

	// 读取label
	label := chipName
	labelFile := filepath.Join(chipPath, "label")
	if data, err := os.ReadFile(labelFile); err == nil {
		label = strings.TrimSpace(string(data))
	}

	return &GpioChipInfo{
		Name:  chipName,
		Base:  base,
		Ngpio: ngpio,
		Label: label,
	}, nil
}

// readGpioFromChip 从芯片信息中读取GPIO引脚信息
func (c *ControllerV1) readGpioFromChip(ctx context.Context, gpioNum int, chipInfo *GpioChipInfo) (*v1.SGpioInfo, error) {
	// 尝试通过periph.io读取GPIO信息
	pinName := fmt.Sprintf("GPIO%d", gpioNum)
	pin := gpioreg.ByName(pinName)

	var value int
	var active bool
	var direction string = "unknown"

	if pin != nil {
		// 使用periph.io读取GPIO状态
		level := pin.Read()
		value = 0
		if level == gpio.High {
			value = 1
		}
		active = true
	} else {
		// 如果periph.io无法访问，尝试通过sysfs读取
		gpioName := fmt.Sprintf("gpio%d", gpioNum)
		gpioPath := fmt.Sprintf("/sys/class/gpio/%s", gpioName)

		// 检查GPIO是否已导出
		if _, err := os.Stat(gpioPath); os.IsNotExist(err) {
			// GPIO未导出，尝试导出
			exportFile := "/sys/class/gpio/export"
			if _, err := os.Stat(exportFile); err == nil {
				// 尝试导出GPIO（需要root权限）
				err := os.WriteFile(exportFile, []byte(strconv.Itoa(gpioNum)), 0644)
				if err != nil {
					// 导出失败，可能是权限问题或GPIO不可用
					return nil, fmt.Errorf("无法导出GPIO %d: %v", gpioNum, err)
				}
			}
		}

		// 读取GPIO信息
		valueFile := filepath.Join(gpioPath, "value")
		if data, err := os.ReadFile(valueFile); err == nil {
			if val, err := strconv.Atoi(strings.TrimSpace(string(data))); err == nil {
				value = val
				active = true
			}
		}

		// 读取方向
		directionFile := filepath.Join(gpioPath, "direction")
		if data, err := os.ReadFile(directionFile); err == nil {
			direction = strings.TrimSpace(string(data))
		}
	}

	return &v1.SGpioInfo{
		PinNumber:  gpioNum,
		Direction:  direction,
		Value:      value,
		Active:     active,
		ProtocolId: "",
		DeviceId:   "",
		DeviceName: fmt.Sprintf("%s-GPIO%d", chipInfo.Label, gpioNum),
	}, nil
}

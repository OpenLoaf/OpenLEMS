//go:build !dev && !windows
// +build !dev,!windows

package internal

import (
	"common/c_base"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"plugin"
	"strings"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"

	"github.com/gogf/gf/v2/errors/gerror"
)

type SConfig struct {
	driverPath string // 驱动存放路径
}

var (
	configInitOnce = sync.Once{}
	ConfigInstance *SConfig
)

func init() {
	gcmd.GetOptWithEnv("device")
}

func NewConfigInstance(deviceCfgName, driverPath string) *SConfig {
	if deviceCfgName == "" {
		deviceCfgName = "device"
	}
	if driverPath == "" {
		driverPath = "out/drivers"
	}
	configInitOnce.Do(func() {
		ConfigInstance = &SConfig{
			driverPath: driverPath,
		}
	})
	return ConfigInstance
}

func (c *SConfig) OpenPlugin(ctx context.Context, path string, symName ...string) (plugin.Symbol, string, error) {
	functionName := c_base.ConstNewPluginFunctionName
	if len(symName) != 0 {
		functionName = symName[0]
	}
	//return openPlugin(ctx, gfile.Join(c.driverPath, path), functionName)

	fullPath := gfile.Join(c.driverPath, path)
	if !gfile.Exists(fullPath) {
		panic(gerror.Newf("插件%s不存在", fullPath))
	}
	g.Log().Infof(ctx, "加载插件：%s", fullPath)
	// 打开插件
	p, err := plugin.Open(fullPath)
	if err != nil {
		panic(gerror.Newf("打开插件%s失败,symName%s。失败原因：%v", path, symName, err))
	}

	// 查找并使用结构体和函数
	symbol, err := p.Lookup(functionName)
	return symbol, fullPath, nil
}

// GetAllDriverNames 获取所有驱动名称
func GetAllDriverNames() []string {
	var driverNames []string

	// 扫描当前目录下的所有.driver文件
	files, err := os.ReadDir(".")
	if err != nil {
		return driverNames
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".driver") {
			driverName := strings.TrimSuffix(file.Name(), ".driver")
			driverNames = append(driverNames, driverName)
		}
	}

	return driverNames
}

// GetAllDriversInfo 获取所有驱动的详细信息
func GetAllDriversInfo(ctx context.Context) []c_base.SDriverInfo {
	var driversInfo []c_base.SDriverInfo
	driverNames := GetAllDriverNames()

	for _, driverName := range driverNames {
		driverInfo := c_base.SDriverInfo{
			Name:      driverName,
			Available: true,
			Embedded:  false,
		}

		// 尝试加载驱动获取详细信息
		driverPath := fmt.Sprintf("%s.driver", driverName)

		func() {
			defer func() {
				if r := recover(); r != nil {
					// 如果加载驱动失败，标记为不可用
					driverInfo.Available = false
				}
			}()

			config := NewConfigInstance("", "")
			pluginSymbol, fullPath, err := config.OpenPlugin(ctx, driverPath)
			if err != nil {
				driverInfo.Available = false
				return
			}

			driverInfo.Path = fullPath

			// 设置文件大小
			if stat, statErr := os.Stat(fullPath); statErr == nil {
				driverInfo.FileSizeByte = stat.Size()
			}

			// 计算文件哈希（SHA-256）
			if f, openErr := os.Open(fullPath); openErr == nil {
				defer f.Close()
				h := sha256.New()
				if _, copyErr := io.Copy(h, f); copyErr == nil {
					driverInfo.HashCode = hex.EncodeToString(h.Sum(nil))
				}
			}

			if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDevice); ok {
				driver := driverNewFunction(ctx)
				if driver != nil {
					driverInfo.Type = driver.GetDriverType()
					driverInfo.Description = driver.GetDriverDescription()
				}
			} else {
				driverInfo.Available = false
			}
		}()

		driversInfo = append(driversInfo, driverInfo)
	}

	return driversInfo
}

// GetDriverInfo 获取指定驱动的详细信息
func GetDriverInfo(ctx context.Context, driverName string) (*c_base.SDriverInfo, error) {
	driverPath := fmt.Sprintf("%s.driver", driverName)

	// 检查驱动文件是否存在
	if _, err := os.ReadFile(driverPath); err != nil {
		return nil, gerror.Newf("未找到驱动文件[%s]", driverPath)
	}

	driverInfo := &c_base.SDriverInfo{
		Name:      driverName,
		Available: true,
	}

	defer func() {
		if r := recover(); r != nil {
			driverInfo.Available = false
		}
	}()

	config := NewConfigInstance("", "")
	pluginSymbol, fullPath, err := config.OpenPlugin(ctx, driverPath)
	if err != nil {
		return nil, gerror.Newf("加载驱动[%s]失败: %v", driverPath, err)
	}

	// 记录路径并补全大小与哈希
	driverInfo.Path = fullPath
	if stat, statErr := os.Stat(fullPath); statErr == nil {
		driverInfo.FileSizeByte = stat.Size()
	}
	if f, openErr := os.Open(fullPath); openErr == nil {
		defer f.Close()
		h := sha256.New()
		if _, copyErr := io.Copy(h, f); copyErr == nil {
			driverInfo.HashCode = hex.EncodeToString(h.Sum(nil))
		}
	}

	if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDevice); ok {
		driver := driverNewFunction(ctx)
		if driver != nil {
			driverInfo.Type = driver.GetDriverType()
			driverInfo.Description = driver.GetDriverDescription()
		}
	} else {
		driverInfo.Available = false
		return nil, gerror.Newf("驱动[%s]不包含有效的NewPlugin函数", driverPath)
	}

	return driverInfo, nil
}

// GetDriversByType 根据设备类型获取驱动信息
func GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []c_base.SDriverInfo {
	allDrivers := GetAllDriversInfo(ctx)
	var filteredDrivers []c_base.SDriverInfo

	for _, driver := range allDrivers {
		if driver.Type == deviceType {
			filteredDrivers = append(filteredDrivers, driver)
		}
	}

	return filteredDrivers
}

// 生产环境下使用驱动加载的方式进行加载
func getDriver(ctx context.Context, driver string) c_base.IDevice {
	if driver == "" {
		panic(gerror.Newf("驱动名称为空！"))
	}

	// 获取最新的驱动文件路径
	latestDriverPath := fmt.Sprintf("%s.driver", driver)

	// 获取驱动的类型
	driverGroups := strings.Split(driver, "_")
	if driverGroups == nil || len(driverGroups) == 0 {
		panic(gerror.Newf("驱动名称错误！%s", driver))
	}

	//ctx = context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, deviceConfig.Id)

	config := NewConfigInstance("", "")
	pluginSymbol, _, err := config.OpenPlugin(ctx, latestDriverPath)
	if err != nil {
		panic(gerror.Newf("加载插件[%s]失败！原因：%v", latestDriverPath, err))
	}

	var dv c_base.IDevice
	if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDevice); ok {
		dv = driverNewFunction(ctx)

		if dv.GetDriverType() != c_base.EDeviceType(driverGroups[0]) {
			panic(gerror.Newf("%s 中驱动类型不匹配！期望类型：%s, 实际类型：%s", driver, dv.GetDriverType(), driverGroups[0]))
		}
	} else {
		panic(gerror.Newf("加载插件[%s]失败！原因：未找到函数：%s", latestDriverPath, c_base.ConstNewPluginFunctionName))
	}

	return dv
}

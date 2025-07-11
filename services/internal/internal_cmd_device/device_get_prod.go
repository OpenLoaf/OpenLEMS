//go:build !dev && !windows
// +build !dev,!windows

package internal_cmd_device

import (
	"common"
	"common/c_base"
	"context"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
)

func init() {
	//g.Log().Noticef(context.Background(), "当前环境为生产环境，从driver文件中获取驱动！")
}

// DriverInfo 驱动信息结构
type DriverInfo struct {
	Name        string               `json:"name"`        // 驱动名称
	Type        c_base.EDeviceType   `json:"type"`        // 驱动类型
	Description *c_base.SDescription `json:"description"` // 驱动描述
	Available   bool                 `json:"available"`   // 是否可用
}

// GetAllDriverNames 获取所有驱动名称
func GetAllDriverNames() []string {
	var driverNames []string

	// 扫描当前目录下的所有.driver文件
	files, err := ioutil.ReadDir(".")
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
func GetAllDriversInfo(ctx context.Context) []DriverInfo {
	var driversInfo []DriverInfo
	driverNames := GetAllDriverNames()

	for _, driverName := range driverNames {
		driverInfo := DriverInfo{
			Name:      driverName,
			Available: true,
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

			pluginSymbol, err := common.OpenPlugin(ctx, driverPath)
			if err != nil {
				driverInfo.Available = false
				return
			}

			if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDriver); ok {
				driver := driverNewFunction(ctx)
				if driver != nil {
					driverInfo.Type = driver.GetDriverType()
					driverInfo.Description = driver.GetDescription()
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
func GetDriverInfo(ctx context.Context, driverName string) (*DriverInfo, error) {
	driverPath := fmt.Sprintf("%s.driver", driverName)

	// 检查驱动文件是否存在
	if _, err := ioutil.ReadFile(driverPath); err != nil {
		return nil, gerror.Newf("未找到驱动文件[%s]", driverPath)
	}

	driverInfo := &DriverInfo{
		Name:      driverName,
		Available: true,
	}

	defer func() {
		if r := recover(); r != nil {
			driverInfo.Available = false
		}
	}()

	pluginSymbol, err := common.OpenPlugin(ctx, driverPath)
	if err != nil {
		return nil, gerror.Newf("加载驱动[%s]失败: %v", driverPath, err)
	}

	if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDriver); ok {
		driver := driverNewFunction(ctx)
		if driver != nil {
			driverInfo.Type = driver.GetDriverType()
			driverInfo.Description = driver.GetDescription()
		}
	} else {
		driverInfo.Available = false
		return nil, gerror.Newf("驱动[%s]不包含有效的NewPlugin函数", driverPath)
	}

	return driverInfo, nil
}

// GetDriversByType 根据设备类型获取驱动信息
func GetDriversByType(ctx context.Context, deviceType c_base.EDeviceType) []DriverInfo {
	allDrivers := GetAllDriversInfo(ctx)
	var filteredDrivers []DriverInfo

	for _, driver := range allDrivers {
		if driver.Type == deviceType {
			filteredDrivers = append(filteredDrivers, driver)
		}
	}

	return filteredDrivers
}

// 生产环境下使用驱动加载的方式进行加载
func (d *SDeviceCmd) getDriver(ctx context.Context, deviceConfig *c_base.SDriverConfig) c_base.IDriver {
	if deviceConfig.Driver == "" {
		panic(gerror.Newf("设备[%s]%s驱动名称为空！", deviceConfig.Id, deviceConfig.Name))
	}

	// 获取最新的驱动文件路径
	latestDriverPath := fmt.Sprintf("%s.driver", deviceConfig.Driver)

	// 获取驱动的类型
	driverGroups := strings.Split(deviceConfig.Driver, "_")
	if driverGroups == nil || len(driverGroups) == 0 {
		panic(gerror.Newf("驱动名称错误！%s", deviceConfig.Driver))
	}

	ctx = context.WithValue(ctx, c_base.ConstCtxKeyDeviceId, deviceConfig.Id)

	pluginSymbol, err := common.OpenPlugin(ctx, latestDriverPath)
	if err != nil {
		panic(gerror.Newf("加载插件[%s]失败！原因：%v", latestDriverPath, err))
	}

	var dv c_base.IDriver
	if driverNewFunction, ok := pluginSymbol.(func(ctx context.Context) c_base.IDriver); ok {
		dv = driverNewFunction(ctx)

		if dv.GetDriverType() != c_base.EDeviceType(driverGroups[0]) {
			panic(gerror.Newf("%s 中驱动类型不匹配！期望类型：%s, 实际类型：%s", deviceConfig.Driver, dv.GetDriverType(), driverGroups[0]))
		}
	} else {
		panic(gerror.Newf("加载插件[%s]失败！原因：未找到函数：%s", latestDriverPath, c_base.ConstNewPluginFunctionName))
	}

	return dv
}

package c_base

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	emsConfig    = "config.yaml"
	deviceConfig = "device.yaml"
)

var (
	once         sync.Once
	systemConfig *SSystemConfig
)

func GetSystemConfig() *SSystemConfig {
	once.Do(func() {
		cmd, err := g.Cfg(emsConfig).GetWithCmd(context.Background(), "system")
		if err != nil {
			panic(fmt.Errorf("获取system配置文件失败:%v", err))
		}
		err = gconv.Scan(cmd, &systemConfig)
		if err != nil {
			panic(fmt.Errorf("解析system配置文件失败:%v", err))
		}
	})
	return systemConfig
}

func GetConfigList[T any](ctx context.Context, key string) ([]*T, string, error) {
	var configPath = GetSystemConfig().ConfigPath

	configPath += "/"
	configPath += deviceConfig

	data := g.Cfg(configPath).MustData(ctx)[key]
	if data == nil {
		return nil, configPath, fmt.Errorf("配置文件:%s中没有devices字段", configPath)
	}
	// 解析列表

	clientConfigList := make([]*T, len(data.([]interface{})))
	err := gconv.Scan(data, &clientConfigList)
	if err != nil {
		return nil, configPath, err
	}

	return clientConfigList, configPath, err
}

// GetLatestDriver 获取最新的driver
func GetLatestDriver(ctx context.Context, name string) (string, int, error) {
	return getLatestDriver(ctx, GetSystemConfig().DriverPath, name)
}

func getLatestDriver(ctx context.Context, configPath string, prefix string) (string, int, error) {
	f, err := os.Open(configPath)
	if err != nil {
		return "", 0, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			g.Log().Error(ctx, err)
		}
	}(f)

	files, err := f.Readdir(-1)
	if err != nil {
		return "", 0, err
	}

	var matchDrivers []string
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), prefix) && filepath.Ext(file.Name()) == ".driver" {
			matchDrivers = append(matchDrivers, file.Name())
		}
	}

	// 筛选出

	return getLatestVersionFile(matchDrivers)
}

// 获取最新版本文件的函数
func getLatestVersionFile(files []string) (string, int, error) {
	if len(files) == 0 {
		return "", 0, fmt.Errorf("没有匹配的文件！")
	}

	// 正则表达式匹配版本号
	re := regexp.MustCompile(`_v(\d+)\.driver$`)

	// 版本号和文件名的映射
	versionMap := make(map[int]string)

	// 提取版本号并存入映射
	for _, file := range files {
		matches := re.FindStringSubmatch(file)
		if len(matches) > 1 {
			version, err := strconv.Atoi(matches[1])
			if err != nil {
				return "", 0, err
			}
			versionMap[version] = file
		}
	}

	if len(versionMap) == 0 {
		return "", 0, fmt.Errorf("no valid versioned files found")
	}

	// 提取所有版本号并排序
	versions := make([]int, 0, len(versionMap))
	for version := range versionMap {
		versions = append(versions, version)
	}
	sort.Ints(versions)

	// 获取最新版本号对应的文件
	latestVersion := versions[len(versions)-1]
	return versionMap[latestVersion], latestVersion, nil
}

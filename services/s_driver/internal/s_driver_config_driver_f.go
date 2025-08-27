package internal

import (
	"context"
	"regexp"
	"sort"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/pkg/errors"
)

// 获取最新版本文件的函数
func getLatestVersionFile(ctx context.Context, driverName string, files []string) string {
	if len(files) == 0 {
		panic(errors.Errorf("%s 没有匹配的文件！", driverName))
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
				panic(errors.Errorf("%s 转换版本号失败:%v", driverName, err))
			}
			versionMap[version] = file
		}
	}

	if len(versionMap) == 0 {
		panic(errors.Errorf("%s没有找到有效的版本文件", driverName))
	}

	// 提取所有版本号并排序
	versions := make([]int, 0, len(versionMap))
	for version := range versionMap {
		versions = append(versions, version)
	}
	sort.Ints(versions)

	// 获取最新版本号对应的文件
	latestVersion := versions[len(versions)-1]

	g.Log().Infof(ctx, "设备加载驱动：%v 共%d个驱动 最新版本为v%v", driverName, len(versionMap), latestVersion)
	return versionMap[latestVersion]
}

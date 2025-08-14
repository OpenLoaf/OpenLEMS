package manifest

import (
	"bytes"
	"embed"
	"fmt"
	"os"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcfg"
	"gopkg.in/yaml.v3"
)

//go:embed config/*.yaml
var configFS embed.FS

// LoadEmbeddedConfig 将嵌入的配置写入默认配置适配器，支持 profile：
// 优先读取 config-<profile>.yaml，不存在则回退 config.yaml。
func LoadEmbeddedConfig(profile string) {
	path := "config/config.yaml"
	if profile != "" && profile != "default" {
		candidate := fmt.Sprintf("config/config-%s.yaml", profile)
		if _, err := configFS.Open(candidate); err == nil {
			path = candidate
		}
	}

	baseData, err := configFS.ReadFile("config/config.yaml")
	if err != nil {
		return
	}
	// 若有 profile 覆盖，进行深度合并
	var merged = map[string]any{}
	_ = yaml.Unmarshal(baseData, &merged)

	if path != "config/config.yaml" {
		if overlayData, err := configFS.ReadFile(path); err == nil {
			var overlay = map[string]any{}
			if err := yaml.Unmarshal(overlayData, &overlay); err == nil {
				merged = deepMerge(merged, overlay)
			}
		}
	}

	var out bytes.Buffer
	_ = yaml.NewEncoder(&out).Encode(merged)

	// 使用新的 AdapterFile 承载内存配置，并替换默认配置适配器，避免磁盘路径查找
	if newCfg, err := gcfg.New(); err == nil {
		if adapter, ok := newCfg.GetAdapter().(*gcfg.AdapterFile); ok {
			adapter.SetContent(out.String())
			g.Cfg().SetAdapter(adapter)
		}
	}
}

// 提前注入一个默认的内置配置，保证最早期（例如 g.Log() 初始化）也不依赖磁盘文件。
// 默认使用 APP_PROFILE，未设置时使用 prod。
func init() {
	profile := os.Getenv("APP_PROFILE")
	if profile == "" || profile == "default" {
		profile = "prod"
	}
	LoadEmbeddedConfig(profile)
}

// deepMerge 递归将 b 合并到 a，b 的值优先生效
func deepMerge(a, b map[string]any) map[string]any {
	for k, vb := range b {
		if mb, ok := vb.(map[string]any); ok {
			if va, ok := a[k].(map[string]any); ok {
				a[k] = deepMerge(va, mb)
			} else {
				a[k] = deepCopyMap(mb)
			}
			continue
		}
		a[k] = vb
	}
	return a
}

func deepCopyMap(m map[string]any) map[string]any {
	cp := make(map[string]any, len(m))
	for k, v := range m {
		if mv, ok := v.(map[string]any); ok {
			cp[k] = deepCopyMap(mv)
		} else {
			cp[k] = v
		}
	}
	return cp
}

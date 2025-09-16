//go:build linux

package cmd

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/gogf/gf/v2/frame/g"
)

// startGui 启动GUI界面（Linux环境下的替代实现）
func startGui(ctx context.Context) {
	startGuiWithOptions(ctx, false)
}

// startGuiFullscreen 启动全屏GUI界面（Linux环境下的替代实现）
func startGuiFullscreen(ctx context.Context) {
	startGuiWithOptions(ctx, true)
}

// startGuiWithOptions 启动GUI界面，支持全屏选项（Linux环境下的替代实现）
func startGuiWithOptions(ctx context.Context, fullscreen bool) {
	if fullscreen {
		g.Log().Infof(ctx, "启动GUI界面（全屏模式）...")
	} else {
		g.Log().Infof(ctx, "启动GUI界面...")
	}

	// 获取服务器地址配置
	serverAddress := g.Config().MustGet(ctx, "server.address").String()
	if serverAddress == "" {
		serverAddress = ":80" // 默认端口
	}

	// 构建完整的服务器URL
	serverURL := "http://localhost" + serverAddress
	g.Log().Infof(ctx, "GUI将连接到服务器: %s", serverURL)

	// 尝试启动浏览器
	if err := openBrowser(serverURL, fullscreen); err != nil {
		g.Log().Errorf(ctx, "无法启动浏览器: %+v", err)
		g.Log().Infof(ctx, "请手动打开浏览器访问: %s", serverURL)
		return
	}

	if fullscreen {
		g.Log().Infof(ctx, "GUI界面已启动（全屏模式）")
	} else {
		g.Log().Infof(ctx, "GUI界面已启动")
	}
}

// GetGuiInstance 获取GUI实例（Linux环境下返回nil）
func GetGuiInstance() interface{} {
	return nil
}

// openBrowser 在Linux环境下打开浏览器
func openBrowser(url string, fullscreen bool) error {
	var cmd *exec.Cmd

	// 检测可用的浏览器
	browsers := []string{
		"google-chrome",
		"chromium-browser",
		"chromium",
		"firefox",
		"opera",
		"konqueror",
		"epiphany",
		"midori",
		"xdg-open",
	}

	var browser string
	for _, b := range browsers {
		if _, err := exec.LookPath(b); err == nil {
			browser = b
			break
		}
	}

	if browser == "" {
		return fmt.Errorf("未找到可用的浏览器")
	}

	// 构建启动命令
	switch browser {
	case "google-chrome", "chromium-browser", "chromium":
		if fullscreen {
			cmd = exec.Command(browser, "--kiosk", "--disable-web-security", "--disable-features=VizDisplayCompositor", url)
		} else {
			cmd = exec.Command(browser, "--new-window", "--disable-web-security", "--disable-features=VizDisplayCompositor", url)
		}
	case "firefox":
		if fullscreen {
			cmd = exec.Command(browser, "--kiosk", url)
		} else {
			cmd = exec.Command(browser, "--new-window", url)
		}
	case "opera":
		if fullscreen {
			cmd = exec.Command(browser, "--kiosk", url)
		} else {
			cmd = exec.Command(browser, "--new-window", url)
		}
	case "konqueror":
		cmd = exec.Command(browser, url)
	case "epiphany":
		cmd = exec.Command(browser, url)
	case "midori":
		cmd = exec.Command(browser, url)
	case "xdg-open":
		cmd = exec.Command(browser, url)
	default:
		return fmt.Errorf("不支持的浏览器: %s", browser)
	}

	// 设置环境变量
	cmd.Env = os.Environ()
	cmd.Env = append(cmd.Env, "DISPLAY=:0") // 确保有显示环境

	// 启动浏览器
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("启动浏览器失败: %w", err)
	}

	// 在后台运行，不等待退出
	go func() {
		if err := cmd.Wait(); err != nil {
			g.Log().Warningf(context.Background(), "浏览器进程退出: %+v", err)
		}
	}()

	return nil
}

// checkDisplayEnvironment 检查显示环境是否可用
func checkDisplayEnvironment() bool {
	// 检查DISPLAY环境变量
	if os.Getenv("DISPLAY") == "" {
		return false
	}

	// 检查是否在图形环境中
	if os.Getenv("XDG_SESSION_TYPE") == "tty" {
		return false
	}

	// 检查是否有可用的窗口管理器
	windowManagers := []string{
		"gnome-session",
		"kde-session",
		"xfce4-session",
		"lxde-session",
		"mate-session",
		"i3",
		"openbox",
		"fluxbox",
		"blackbox",
	}

	for _, wm := range windowManagers {
		if _, err := exec.LookPath(wm); err == nil {
			return true
		}
	}

	return false
}

// getLinuxDesktopEnvironment 获取Linux桌面环境
func getLinuxDesktopEnvironment() string {
	desktop := os.Getenv("XDG_CURRENT_DESKTOP")
	if desktop != "" {
		return desktop
	}

	desktop = os.Getenv("DESKTOP_SESSION")
	if desktop != "" {
		return desktop
	}

	desktop = os.Getenv("GNOME_DESKTOP_SESSION_ID")
	if desktop != "" {
		return "GNOME"
	}

	desktop = os.Getenv("KDE_FULL_SESSION")
	if desktop != "" {
		return "KDE"
	}

	return "Unknown"
}

// logSystemInfo 记录系统信息
func logSystemInfo(ctx context.Context) {
	g.Log().Infof(ctx, "系统信息:")
	g.Log().Infof(ctx, "  操作系统: %s", runtime.GOOS)
	g.Log().Infof(ctx, "  架构: %s", runtime.GOARCH)
	g.Log().Infof(ctx, "  桌面环境: %s", getLinuxDesktopEnvironment())
	g.Log().Infof(ctx, "  显示环境: %s", os.Getenv("DISPLAY"))
	g.Log().Infof(ctx, "  会话类型: %s", os.Getenv("XDG_SESSION_TYPE"))

	if checkDisplayEnvironment() {
		g.Log().Infof(ctx, "  图形环境: 可用")
	} else {
		g.Log().Warningf(ctx, "  图形环境: 不可用")
	}
}

//go:build !linux || (linux && amd64)

package cmd

import (
	"context"
	"runtime"
	"sync"
	"unsafe"

	"github.com/gogf/gf/v2/frame/g"
	webview "github.com/webview/webview_go"
)

/*
#cgo darwin CFLAGS: -x objective-c
#cgo darwin LDFLAGS: -framework Cocoa
#cgo windows CFLAGS: -D_WIN32_WINNT=0x0601
#cgo linux pkg-config: gtk+-3.0 webkit2gtk-4.0

#ifdef __APPLE__
#import <Cocoa/Cocoa.h>
void setFullscreen(void* window) {
    NSWindow* nsWindow = (NSWindow*)window;
    [nsWindow setCollectionBehavior:NSWindowCollectionBehaviorFullScreenPrimary];
    [nsWindow toggleFullScreen:nil];
}
#endif

#ifdef _WIN32
#include <windows.h>
void setFullscreen(void* window) {
    HWND hwnd = (HWND)window;
    SetWindowLong(hwnd, GWL_STYLE, WS_POPUP | WS_VISIBLE);
    SetWindowPos(hwnd, HWND_TOP, 0, 0, GetSystemMetrics(SM_CXSCREEN), GetSystemMetrics(SM_CYSCREEN), SWP_FRAMECHANGED);
}
#endif

#ifdef __linux__
#include <gtk/gtk.h>
void setFullscreen(void* window) {
    GtkWindow* gtkWindow = (GtkWindow*)window;
    gtk_window_fullscreen(gtkWindow);
}
#endif
*/
import "C"

var (
	guiInstance webview.WebView
	guiOnce     sync.Once
)

// startGui 启动GUI界面
func startGui(ctx context.Context) {
	startGuiWithOptions(ctx, false)
}

// startGuiFullscreen 启动全屏GUI界面
func startGuiFullscreen(ctx context.Context) {
	startGuiWithOptions(ctx, true)
}

// startGuiWithOptions 启动GUI界面，支持全屏选项
func startGuiWithOptions(ctx context.Context, fullscreen bool) {
	guiOnce.Do(func() {
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

		// 创建webview实例
		w := webview.New(false)
		defer w.Destroy()

		// 设置窗口属性
		w.SetTitle("EMS Plan - 能源管理系统")
		if fullscreen {
			// 全屏模式：设置一个较大的初始尺寸，然后立即全屏
			w.SetSize(1920, 1080, webview.HintNone)
		} else {
			// 窗口模式：设置固定尺寸
			w.SetSize(1280, 720, webview.HintNone)
		}

		// 设置默认HTML内容
		w.SetHtml(`
			<!DOCTYPE html>
			<html>
			<head>
				<title>EMS Plan</title>
				<meta charset="utf-8">
				<style>
					body {
						font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
						margin: 0;
						padding: 20px;
						background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
						color: white;
						display: flex;
						align-items: center;
						justify-content: center;
						height: 100vh;
						text-align: center;
					}
					.container {
						background: rgba(255, 255, 255, 0.1);
						padding: 40px;
						border-radius: 20px;
						backdrop-filter: blur(10px);
						box-shadow: 0 8px 32px rgba(0, 0, 0, 0.1);
					}
					h1 {
						font-size: 2.5em;
						margin-bottom: 20px;
						text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.3);
					}
					p {
						font-size: 1.2em;
						margin-bottom: 30px;
						opacity: 0.9;
					}
					.loading {
						display: inline-block;
						width: 40px;
						height: 40px;
						border: 3px solid rgba(255, 255, 255, 0.3);
						border-radius: 50%;
						border-top-color: #fff;
						animation: spin 1s ease-in-out infinite;
					}
					@keyframes spin {
						to { transform: rotate(360deg); }
					}
				</style>
			</head>
			<body>
				<div class="container">
					<h1>🚀 EMS Plan</h1>
					<p>能源管理系统正在启动...</p>
					<div class="loading"></div>
					<p style="margin-top: 20px; font-size: 0.9em; opacity: 0.7;">
						正在连接到本地服务器...
					</p>
				</div>
				<script>
					// 尝试连接到本地服务器
					setTimeout(() => {
						window.location.href = '` + serverURL + `';
					}, 2000);
				</script>
			</body>
			</html>
		`)

		// 导航到本地服务器
		w.Navigate(serverURL)

		// 如果启用全屏模式，设置全屏
		if fullscreen {
			setFullscreen(w)
			g.Log().Infof(ctx, "GUI界面已启动（全屏模式）")
		} else {
			g.Log().Infof(ctx, "GUI界面已启动，窗口大小: 1440x900")
		}

		// 保存实例引用
		guiInstance = w

		// 运行GUI
		w.Run()
	})
}

// GetGuiInstance 获取GUI实例
func GetGuiInstance() webview.WebView {
	return guiInstance
}

// setFullscreen 设置窗口全屏（平台特定实现）
func setFullscreen(w webview.WebView) {
	window := w.Window()
	if window == nil {
		g.Log().Warningf(context.Background(), "无法获取窗口句柄")
		return
	}

	switch runtime.GOOS {
	case "darwin":
		// macOS 平台：使用 Cocoa API 实现原生全屏
		g.Log().Infof(context.Background(), "使用 Cocoa API 设置全屏")
		C.setFullscreen(unsafe.Pointer(window))
	case "windows":
		// Windows 平台：使用 Win32 API
		g.Log().Infof(context.Background(), "使用 Win32 API 设置全屏")
		C.setFullscreen(unsafe.Pointer(window))
	case "linux":
		// Linux 平台：使用 GTK API
		g.Log().Infof(context.Background(), "使用 GTK API 设置全屏")
		C.setFullscreen(unsafe.Pointer(window))
	default:
		g.Log().Warningf(context.Background(), "不支持的操作系统: %s", runtime.GOOS)
		return
	}

	// 同时设置页面样式，确保内容填满整个窗口
	w.Eval(`
		// 隐藏滚动条
		document.body.style.overflow = 'hidden';
		document.documentElement.style.overflow = 'hidden';
		
		// 设置全屏样式
		document.body.style.margin = '0';
		document.body.style.padding = '0';
		document.documentElement.style.margin = '0';
		document.documentElement.style.padding = '0';
		
		// 确保内容填满整个窗口
		document.body.style.width = '100vw';
		document.body.style.height = '100vh';
		document.documentElement.style.width = '100vw';
		document.documentElement.style.height = '100vh';
	`)
}

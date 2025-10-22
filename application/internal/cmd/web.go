package cmd

import (
	"application/internal/consts"
	"application/internal/controller/alarm"
	appauth "application/internal/controller/auth"
	"application/internal/controller/automation"
	"application/internal/controller/control"
	"application/internal/controller/device"
	"application/internal/controller/driver"
	"application/internal/controller/log"
	"application/internal/controller/network"
	"application/internal/controller/policy"
	"application/internal/controller/price"
	"application/internal/controller/protocol"
	"application/internal/controller/remote"
	"application/internal/controller/setting"
	"application/internal/controller/system"
	"application/internal/utils"
	"application/manifest"
	"common/c_enum"
	"common/c_log"
	"context"
	"io"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
)

// startWeb 启动Web服务
func startWeb(ctx context.Context) *ghttp.Server {
	return startWebWithBinding(ctx, false)
}

// startWebWithBinding 启动Web服务，可选择是否只绑定到本地地址
func startWebWithBinding(ctx context.Context, localOnly bool) *ghttp.Server {
	ctx = context.WithValue(ctx, c_enum.ELogTypeEms, "Web")
	c_log.Infof(ctx, "准备启动web程序！")

	s := g.Server()
	// 中间件顺序：响应封装 -> 永不超时上下文 -> 请求ID -> 错误处理 -> 访问日志
	// 关闭框架内置访问/错误日志，避免与自定义日志重复
	s.SetAccessLogEnabled(false)
	s.SetErrorLogEnabled(false)
	s.Use(
		ghttp.MiddlewareHandlerResponse,
		ghttp.MiddlewareNeverDoneCtx,
		utils.MiddlewareRequestID,
		utils.MiddlewareErrorHandler,
		utils.MiddlewareAccessLog,
	)

	// 配置 Session（默认2小时），禁用将SessionId回写到Cookie，仅从Header注入
	s.SetSessionMaxAge(2 * time.Hour)
	s.SetSessionIdName("ems_session_id")
	s.SetSessionCookieOutput(true)

	// 设置API路由
	setupAPIRoutes(s)

	// 自定义增强API文档
	enhanceOpenAPIDoc(s)

	// 设置静态文件服务
	setupStaticFiles(s, ctx)

	// 设置本地静态文件服务
	setupLocalStaticFiles(s, ctx)

	// 启动服务器并打印地址信息
	go func() {
		// 获取服务器地址信息
		serverAddress := g.Config().MustGet(ctx, "server.address").String()
		if serverAddress == "" {
			serverAddress = ":80" // 默认端口
		}

		// 如果只允许本地访问，修改绑定地址
		if localOnly {
			if strings.HasPrefix(serverAddress, ":") {
				// 如果地址以冒号开头（如 :15880），则添加 127.0.0.1
				serverAddress = "127.0.0.1" + serverAddress
			} else if !strings.Contains(serverAddress, "127.0.0.1") && !strings.Contains(serverAddress, "localhost") {
				// 如果地址不包含本地地址，则替换为 127.0.0.1
				if strings.Contains(serverAddress, ":") {
					parts := strings.Split(serverAddress, ":")
					if len(parts) >= 2 {
						serverAddress = "127.0.0.1:" + parts[len(parts)-1]
					}
				}
			}
			c_log.Infof(ctx, "GUI模式：Web服务仅绑定到本地地址 %s", serverAddress)
		}

		utils.PrintWebServerInfo(ctx, serverAddress)
	}()

	return s
}

// setupAPIRoutes 设置API路由
func setupAPIRoutes(s *ghttp.Server) {
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(
			ghttp.MiddlewareCORS,
			func(r *ghttp.Request) {
				httpLanguage := r.GetHeader("Accept-Language")
				if httpLanguage == "" {
					httpLanguage = "zh-CN"
				}
				r.SetCtx(gi18n.WithLanguage(r.Context(), httpLanguage))
				r.Middleware.Next()
			},
			utils.MiddlewareAuth,
		)

		// 绑定所有控制器
		group.Bind(appauth.NewV1())
		group.Bind(device.NewV1())
		group.Bind(driver.NewV1())
		group.Bind(network.NewV1())
		group.Bind(system.NewV1())
		group.Bind(protocol.NewV1())
		group.Bind(control.NewV1())
		group.Bind(log.NewV1())
		group.Bind(alarm.NewV1())
		group.Bind(policy.NewV1())
		group.Bind(price.NewV1())
		group.Bind(setting.NewV1())
		group.Bind(automation.NewV1())
		group.Bind(remote.NewV1())
	})
}

// setupStaticFiles 设置静态文件服务
func setupStaticFiles(s *ghttp.Server, ctx context.Context) {
	// 静态站点：将 `application/manifest/web` 打包进可执行文件并作为根路径提供
	webfs, err := manifest.WebFS()
	if err != nil {
		c_log.Warningf(ctx, "Web 静态资源初始化失败: %+v", err)
		return
	}

	// 输出嵌入的静态资源加载情况（统计与示例）
	{
		var (
			totalFiles   int
			visibleFiles int
			exampleFiles []string
		)
		_ = fs.WalkDir(webfs, ".", func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return nil
			}
			if d.IsDir() {
				return nil
			}
			totalFiles++
			base := filepath.Base(path)
			lower := strings.ToLower(base)
			// 过滤隐藏文件与 .gitignore
			if strings.HasPrefix(base, ".") || lower == ".gitignore" || lower == "gitignore" {
				return nil
			}
			visibleFiles++
			if len(exampleFiles) < 20 { // 仅示例前20个
				exampleFiles = append(exampleFiles, path)
			}
			return nil
		})
		c_log.Infof(ctx, "Web 静态资源统计: 总文件=%d, 可见文件=%d", totalFiles, visibleFiles)
		if visibleFiles == 0 {
			c_log.Warningf(ctx, "Web 静态资源为空或仅包含隐藏文件/.gitignore，请检查 application/manifest/web 目录")
		} else if len(exampleFiles) > 0 {
			c_log.Infof(ctx, "Web 静态资源示例: %s", strings.Join(exampleFiles, ", "))
		}
	}

	fileServer := http.FileServer(http.FS(webfs))

	// 静态资源路径 - 优先处理静态资源
	s.BindHandler("GET:/assets/*", func(r *ghttp.Request) {
		fileServer.ServeHTTP(r.Response.Writer, r.Request)
	})
	s.BindHandler("GET:/images/*", func(r *ghttp.Request) {
		fileServer.ServeHTTP(r.Response.Writer, r.Request)
	})
	s.BindHandler("GET:/favicon.ico", func(r *ghttp.Request) {
		fileServer.ServeHTTP(r.Response.Writer, r.Request)
	})
	s.BindHandler("GET:/config/*", func(r *ghttp.Request) {
		fileServer.ServeHTTP(r.Response.Writer, r.Request)
	})
	s.BindHandler("GET:/demo/*", func(r *ghttp.Request) {
		fileServer.ServeHTTP(r.Response.Writer, r.Request)
	})

	// 根路径：直接输出 index.html
	s.BindHandler("GET:/", func(r *ghttp.Request) {
		f, err := webfs.Open("index.html")
		if err != nil {
			r.Response.WriteStatus(http.StatusNotFound)
			return
		}
		defer f.Close()
		// 明确 Content-Type
		r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := io.Copy(r.Response.Writer, f); err != nil {
			c_log.Warningf(ctx, "写入 index.html 失败: %+v", err)
		}
	})

	// 支持Vue history模式：所有其他GET请求都返回index.html
	s.BindHandler("GET:/*", func(r *ghttp.Request) {
		// 跳过API路由
		if strings.HasPrefix(r.URL.Path, "/api/") {
			r.Response.WriteStatus(http.StatusNotFound)
			return
		}
		// 跳过WebSocket路由
		if strings.HasPrefix(r.URL.Path, "/station") || strings.HasPrefix(r.URL.Path, "/telemetry") {
			r.Response.WriteStatus(http.StatusNotFound)
			return
		}
		// 跳过静态资源路由 - 这些应该由具体的静态资源处理器处理
		if strings.HasPrefix(r.URL.Path, "/assets/") ||
			strings.HasPrefix(r.URL.Path, "/images/") ||
			strings.HasPrefix(r.URL.Path, "/config/") ||
			strings.HasPrefix(r.URL.Path, "/demo/") ||
			r.URL.Path == "/favicon.ico" {
			r.Response.WriteStatus(http.StatusNotFound)
			return
		}

		f, err := webfs.Open("index.html")
		if err != nil {
			r.Response.WriteStatus(http.StatusNotFound)
			return
		}
		defer f.Close()
		// 明确 Content-Type
		r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
		if _, err := io.Copy(r.Response.Writer, f); err != nil {
			c_log.Warningf(ctx, "写入 index.html 失败: %+v", err)
		}
	})
}

// setupLocalStaticFiles 设置本地静态文件服务
func setupLocalStaticFiles(s *ghttp.Server, ctx context.Context) {
	// 本地静态文件服务：处理 /static 路径
	staticPath := g.Cfg().MustGet(ctx, "static.path", "out/static").String()
	if staticPath == "" {
		return
	}

	// 检查静态文件目录是否存在，如果不存在则创建
	if _, err := os.Stat(staticPath); os.IsNotExist(err) {
		c_log.Infof(ctx, "静态文件目录不存在，正在创建: %s", staticPath)
		if err := os.MkdirAll(staticPath, 0755); err != nil {
			c_log.Errorf(ctx, "创建静态文件目录失败: %s, 错误: %+v", staticPath, err)
			return
		} else {
			c_log.Infof(ctx, "静态文件目录创建成功: %s", staticPath)
		}
	}

	// 再次检查目录是否存在（创建后或原本就存在）
	if _, err := os.Stat(staticPath); err != nil {
		c_log.Errorf(ctx, "静态文件目录创建失败或无法访问: %s", staticPath)
		return
	}

	c_log.Infof(ctx, "启用本地静态文件服务，路径: %s", staticPath)

	// 静态文件路由处理
	s.BindHandler("GET:/static/*", func(r *ghttp.Request) {
		// 获取请求的文件路径
		requestPath := strings.TrimPrefix(r.URL.Path, "/static/")
		if requestPath == "" {
			requestPath = "index.html"
		}

		// 构建完整的文件路径
		filePath := filepath.Join(staticPath, requestPath)

		// 安全检查：确保文件路径在静态目录内
		absStaticPath, _ := filepath.Abs(staticPath)
		absFilePath, _ := filepath.Abs(filePath)
		if !strings.HasPrefix(absFilePath, absStaticPath) {
			r.Response.WriteStatus(http.StatusForbidden)
			return
		}

		// 检查文件是否存在
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			r.Response.WriteStatus(http.StatusNotFound)
			return
		}

		// 设置适当的Content-Type
		ext := filepath.Ext(filePath)
		switch ext {
		case ".html":
			r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
		case ".css":
			r.Response.Header().Set("Content-Type", "text/css; charset=utf-8")
		case ".js":
			r.Response.Header().Set("Content-Type", "application/javascript; charset=utf-8")
		case ".json":
			r.Response.Header().Set("Content-Type", "application/json; charset=utf-8")
		case ".png":
			r.Response.Header().Set("Content-Type", "image/png")
		case ".jpg", ".jpeg":
			r.Response.Header().Set("Content-Type", "image/jpeg")
		case ".gif":
			r.Response.Header().Set("Content-Type", "image/gif")
		case ".svg":
			r.Response.Header().Set("Content-Type", "image/svg+xml")
		case ".ico":
			r.Response.Header().Set("Content-Type", "image/x-icon")
		default:
			r.Response.Header().Set("Content-Type", "application/octet-stream")
		}

		// 提供文件
		http.ServeFile(r.Response.Writer, r.Request, filePath)
	})
}

// enhanceOpenAPIDoc 自定义增强API文档
func enhanceOpenAPIDoc(s *ghttp.Server) {
	openapi := s.GetOpenApi()
	openapi.Config.CommonResponse = ghttp.DefaultHandlerResponse{}
	openapi.Config.CommonResponseDataField = `Data`

	// API description.
	openapi.Info = goai.Info{
		Title:       consts.OpenAPITitle,
		Description: consts.OpenAPIDescription,
		Contact: &goai.Contact{
			Name: "Hex-Ems",
			URL:  "",
		},
	}
}

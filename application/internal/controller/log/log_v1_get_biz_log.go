package log

import (
	apiv1 "application/api/log/v1"
	applog "application/internal/log"
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// ControllerV1 由 log_new.go 的 NewV1 返回
// func New() *ControllerV1 { return &ControllerV1{} }

func (c *ControllerV1) GetBizLog(ctx g.Ctx, req *apiv1.GetBizLogReq) (res *apiv1.GetBizLogRes, err error) {
	// 处理默认日期
	if req.Date == "" {
		req.Date = time.Now().Format("20060102")
	}

	// 解析基础路径与文件名
	basePath, filePattern, err := resolveBase(ctx, req.Type)
	if err != nil {
		return nil, err
	}
	// 构造目录
	dir := basePath
	if req.Type != "ems" {
		if req.Id == "" {
			return nil, errors.New("id required for non-ems type")
		}
		dir = filepath.Join(basePath, req.Id)
	}
	// 文件模式：优先使用 pattern，否则默认 {Ymd}.log
	name := filePattern
	if name == "" {
		name = "{Ymd}.log"
	}
	name = replaceYmd(name, req.Date)
	fpath := filepath.Join(dir, name)

	lines, total, e := readFileTailPaged(fpath, req.Page, req.PageSize)
	if e != nil {
		return nil, e
	}

	// 解析日志行为结构化数据
	structuredLines := parseLogLines(lines)

	return &apiv1.GetBizLogRes{Total: total, Lines: structuredLines}, nil
}

// parseLogLines 解析日志行为结构化数据
func parseLogLines(rawLines []string) []apiv1.LogLine {
	var result []apiv1.LogLine

	// GoFrame日志格式正则：支持多种时间格式
	// 格式1: 2025-08-21T23:13:41.223+08:00 [INFO] 内容
	// 格式2: 2025-01-12 10:30:45.123 [INFO] 内容
	logPattern := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}[T\s]\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:[+-]\d{2}:\d{2})?)\s+\[(\w+)\]\s+(.+)$`)

	for _, line := range rawLines {
		matches := logPattern.FindStringSubmatch(line)
		if len(matches) == 4 {
			// 成功解析
			result = append(result, apiv1.LogLine{
				Timestamp: matches[1],
				Level:     matches[2],
				Content:   matches[3],
			})
		} else {
			// 无法解析的日志行，作为纯内容处理
			result = append(result, apiv1.LogLine{
				Timestamp: "",
				Level:     "UNKNOWN",
				Content:   line,
			})
		}
	}

	return result
}

// resolveBase 读取配置里 logger.<key> 的路径与文件名
func resolveBase(ctx g.Ctx, t string) (basePath string, filePattern string, err error) {
	var key string
	switch t {
	case "ems":
		key = "biz_ems"
	case "device":
		key = "biz_device"
	case "protocol":
		key = "biz_protocol"
	case "policy":
		key = "biz_policy"
	default:
		return "", "", fmt.Errorf("unknown type: %s", t)
	}
	m := g.Cfg().MustGet(ctx, fmt.Sprintf("logger.%s", key)).Map()
	if v, ok := m["path"].(string); ok {
		basePath = v
	}
	if v, ok := m["file"].(string); ok {
		filePattern = v
	}
	// 兜底路径，与 helpers 中一致
	if basePath == "" {
		basePath = filepath.Join("logs", "biz", key)
	}
	return
}

func replaceYmd(pattern, ymd string) string {
	// 简化：仅替换 {Ymd}
	return strings.ReplaceAll(pattern, "{Ymd}", ymd)
}

// 读取文件并倒序分页（逐行扫描，再倒序切片）
func readFileTailPaged(path string, page, size int) ([]string, int, error) {
	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = 100
	}
	if size > 1000 {
		size = 1000
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, 0, err
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	buf := make([]byte, 0, 1024*1024)
	s.Buffer(buf, 1024*1024)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	if err := s.Err(); err != nil {
		return nil, 0, err
	}
	total := len(lines)
	// 倒序
	sort.SliceStable(lines, func(i, j int) bool { return i > j })

	start := (page - 1) * size
	if start >= total {
		return []string{}, total, nil
	}
	end := start + size
	if end > total {
		end = total
	}
	return lines[start:end], total, nil
}

func init() {
	// 引用 applog 防止编译器移除
	_ = applog.BizEMS
}

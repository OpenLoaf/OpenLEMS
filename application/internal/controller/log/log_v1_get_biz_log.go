package log

import (
	apiv1 "application/api/log/v1"
	applog "application/internal/log"
	"bufio"
	"encoding/json"
	"os"
	"path/filepath"
	"regexp"
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

	// 统一读取 EMS 业务日志单文件
	basePath, filePattern, err := resolveBase(ctx, "ems")
	if err != nil {
		return nil, err
	}
	// 文件模式：优先使用 pattern，否则默认 {Ymd}.log
	name := filePattern
	if name == "" {
		name = "{Ymd}.log"
	}
	name = replaceYmd(name, req.Date)
	fpath := filepath.Join(basePath, name)

	// 为确保过滤后分页准确，读取全部再按需过滤+分页
	lines, _, e := readAllFileLines(fpath)
	if e != nil {
		// 文件不存在或读取失败时返回空数组
		return &apiv1.GetBizLogRes{Total: 0, Lines: []apiv1.LogLine{}}, nil
	}

	filtered := make([]apiv1.LogLine, 0)
	// 倒序遍历，保证最新在前
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if ok, jl := tryParseBizJson(line); ok {
			// JSON 行包含 type/id，按需过滤
			if matchTypeAndId(req.Type, req.Id, jl.Type, jl.Id) {
				filtered = append(filtered, apiv1.LogLine{Timestamp: jl.Time, Level: jl.Level, Content: jl.Content})
			}
			continue
		}
		// 旧格式：当请求为空/all/ems时纳入
		if req.Type == "" || strings.EqualFold(req.Type, "all") || req.Type == "ems" {
			pl := parseLogLine(line)
			filtered = append(filtered, pl)
		}
	}

	total := len(filtered)
	start := (req.Page - 1) * req.PageSize
	end := start + req.PageSize
	if start >= total {
		return &apiv1.GetBizLogRes{Total: total, Lines: []apiv1.LogLine{}}, nil
	}
	if end > total {
		end = total
	}

	return &apiv1.GetBizLogRes{Total: total, Lines: filtered[start:end]}, nil
}

// BizJsonLine 统一的业务JSON日志结构
type BizJsonLine struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
	Type    string `json:"type"`
	Id      string `json:"id"`
}

// 解析JSON业务日志
func tryParseBizJson(line string) (bool, BizJsonLine) {
	var jl BizJsonLine
	if err := json.Unmarshal([]byte(line), &jl); err == nil && (jl.Time != "" || jl.Content != "") {
		return true, jl
	}
	return false, jl
}

func matchTypeAndId(reqType, reqId, lineType, lineId string) bool {
	if reqType == "" || strings.EqualFold(reqType, "all") {
		return true
	}
	if reqType == "ems" {
		return lineType == "ems"
	}
	if reqType == "device" || reqType == "protocol" || reqType == "policy" {
		if reqId == "" {
			return false
		}
		return lineType == reqType && lineId == reqId
	}
	return false
}

// readAllFileLines 读取文件的所有行
func readAllFileLines(path string) ([]string, int, error) {
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
		line := strings.TrimSpace(s.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	if err := s.Err(); err != nil {
		return nil, 0, err
	}

	return lines, len(lines), nil
}

// resolveBase 读取配置里 logger.biz_ems 的路径与文件名（已合并单文件）
func resolveBase(ctx g.Ctx, _ string) (basePath string, filePattern string, err error) {
	m := g.Cfg().MustGet(ctx, "logger.biz_ems").Map()
	if v, ok := m["path"].(string); ok {
		basePath = v
	}
	if v, ok := m["file"].(string); ok {
		filePattern = v
	}
	// 兜底路径，与 helpers 中一致
	if basePath == "" {
		basePath = filepath.Join("logs", "biz")
	}
	return
}

func replaceYmd(pattern, ymd string) string {
	// 简化：仅替换 {Ymd}
	return strings.ReplaceAll(pattern, "{Ymd}", ymd)
}

// parseLogLine 解析旧格式单行日志（时间 等级 内容）
func parseLogLine(line string) apiv1.LogLine {
	logPattern := regexp.MustCompile(`^(\d{4}-\d{2}-\d{2}[T\s]\d{2}:\d{2}:\d{2}(?:\.\d+)?(?:[+-]\d{2}:\d{2})?)\s+\[(\w+)\]\s+(.+)$`)
	matches := logPattern.FindStringSubmatch(line)
	if len(matches) == 4 {
		return apiv1.LogLine{
			Timestamp: matches[1],
			Level:     matches[2],
			Content:   matches[3],
		}
	}
	return apiv1.LogLine{Timestamp: "", Level: "UNKNOWN", Content: line}
}

func init() {
	// 引用 applog 防止编译器移除
	_ = applog.BizEMS
}

package log

import (
	"bufio"
	"common/c_base"
	"common/c_log"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/text/gstr"
)

// BizJsonHandler 输出统一的 JSON 格式，并在每条日志中附带分类字段(type/id)
var BizJsonHandler glog.Handler = func(ctx context.Context, in *glog.HandlerInput) {
	typ := "ems"
	id := ""
	if ctx != nil {
		if v := ctx.Value(c_base.ConstCtxKeyDeviceId); v != nil {
			if s, ok := v.(string); ok && s != "" {
				typ = "device"
				id = s
			}
		}
		if typ == "ems" { // 仅当未匹配设备时再判断协议
			if v := ctx.Value(c_base.ConstCtxKeyProtocolId); v != nil {
				if s, ok := v.(string); ok && s != "" {
					typ = "protocol"
					id = s
				}
			}
		}
		if typ == "ems" { // 兜底策略ID（约定键名）
			if v := ctx.Value("PolicyId"); v != nil {
				if s, ok := v.(string); ok && s != "" {
					typ = "policy"
					id = s
				}
			}
		}
	}
	content := in.Content
	if gstr.Trim(content) == "" {
		// 回退使用 values 内容
		content = fmt.Sprint(in.Values...)
	}
	content = gstr.Trim(content)
	payload := struct {
		Time    string `json:"time"`
		Level   string `json:"level"`
		Content string `json:"content"`
		Type    string `json:"type"`
		Id      string `json:"id"`
	}{
		Time:    in.TimeFormat,
		Level:   gstr.Trim(in.LevelFormat, "[]"),
		Content: content,
		Type:    typ,
		Id:      id,
	}
	b, err := json.Marshal(payload)
	if err != nil {
		in.Buffer.WriteString(err.Error())
		in.Next(ctx)
		return
	}
	in.Buffer.Write(b)
	in.Buffer.WriteString("\n")
	in.Next(ctx)
}

// BizEMS 获取 EMS 业务日志（固定单文件，按天切分）
func BizEMS() *glog.Logger {
	l := g.Log("biz_ems")
	// 统一业务日志：JSON格式 + 异步输出
	l.SetHandlers(BizJsonHandler)
	l.SetAsync(true)
	return l
}

// BizJsonLine 统一的业务JSON日志结构
type BizJsonLine struct {
	Time    string `json:"time"`
	Level   string `json:"level"`
	Content string `json:"content"`
	Type    string `json:"type"`
	Id      string `json:"id"`
}

// QueryBizLogs 查询业务日志（从文件）
func QueryBizLogs(ctx context.Context, params c_log.LogQueryParams) (*c_log.LogQueryResult, error) {
	// 处理默认日期
	date := params.Date
	if date == "" {
		date = time.Now().Format("20060102")
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
	name = replaceYmd(name, date)
	fpath := filepath.Join(basePath, name)

	// 为确保过滤后分页准确，读取全部再按需过滤+分页
	lines, _, e := readAllFileLines(fpath)
	if e != nil {
		// 文件不存在或读取失败时返回空数组
		return &c_log.LogQueryResult{Total: 0, Lines: []c_log.LogLine{}}, nil
	}

	filtered := make([]c_log.LogLine, 0)
	// 倒序遍历，保证最新在前
	for i := len(lines) - 1; i >= 0; i-- {
		line := lines[i]
		if ok, jl := tryParseBizJson(line); ok {
			// JSON 行包含 type/id，按需过滤
			if matchTypeAndId(params.Type, params.Id, jl.Type, jl.Id) && matchLevel(params.Level, jl.Level) && isAllowedLevel(jl.Level) {
				filtered = append(filtered, c_log.LogLine{
					Timestamp: jl.Time,
					Level:     toUpperNormalized(jl.Level),
					Content:   jl.Content,
					Id:        jl.Id,
					Type:      jl.Type,
				})
			}
			continue
		}
	}

	total := len(filtered)
	start := (params.Page - 1) * params.PageSize
	end := start + params.PageSize
	if start >= total {
		return &c_log.LogQueryResult{Total: total, Lines: []c_log.LogLine{}}, nil
	}
	if end > total {
		end = total
	}

	return &c_log.LogQueryResult{Total: total, Lines: filtered[start:end]}, nil
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

func matchLevel(reqLevel, lineLevel string) bool {
	// 空或 all 不做过滤
	if reqLevel == "" || strings.EqualFold(reqLevel, "all") {
		return true
	}
	rl := normalizeLevel(reqLevel)
	ll := normalizeLevel(lineLevel)
	return rl == ll
}

func normalizeLevel(s string) string {
	v := strings.ToLower(strings.TrimSpace(s))
	switch v {
	case "warning":
		return "warn"
	}
	return v
}

func toUpperNormalized(s string) string {
	return gstr.ToUpper(normalizeLevel(s))
}

func isAllowedLevel(level string) bool {
	l := normalizeLevel(level)
	switch l {
	case "debug", "info", "warn", "error":
		return true
	default:
		return false
	}
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
func resolveBase(ctx context.Context, _ string) (basePath string, filePattern string, err error) {
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

package manifest

import (
	"embed"
	"io/fs"
)

//go:embed web/*
var web embed.FS

// WebFS 返回裁剪后的只读文件系统，根目录为 `web/`。
func WebFS() (fs.FS, error) {
	return fs.Sub(web, "web")
}

package manifest

import (
	"archive/zip"
	"bytes"
	_ "embed"
	"fmt"
	"io/fs"
	"sync"
)

//go:embed web.zip
var webZip []byte

var (
	webFSOnce sync.Once
	webFS     fs.FS
	webFSErr  error
)

// WebFS returns a read-only file system backed by the embedded web.zip bundle.
func WebFS() (fs.FS, error) {
	webFSOnce.Do(func() {
		if len(webZip) == 0 {
			webFSErr = fmt.Errorf("embedded web bundle is empty")
			return
		}

		reader, err := zip.NewReader(bytes.NewReader(webZip), int64(len(webZip)))
		if err != nil {
			webFSErr = fmt.Errorf("open embedded web bundle: %w", err)
			return
		}
		webFS = reader
	})

	return webFS, webFSErr
}

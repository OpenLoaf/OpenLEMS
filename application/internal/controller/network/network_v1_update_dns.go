package network

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"os/exec"
	"runtime"
	"strings"

	v1 "application/api/network/v1"
)

// UpdateDNS 单独更新系统 DNS
func (c *ControllerV1) UpdateDNS(ctx context.Context, req *v1.UpdateDNSReq) (res *v1.UpdateDNSRes, err error) {
	if len(req.DNS) == 0 {
		return nil, errors.Errorf("DNS不能为空")
	}
	switch runtime.GOOS {
	case "linux":
		lines := make([]string, 0, len(req.DNS))
		for _, d := range req.DNS {
			lines = append(lines, "nameserver "+d)
		}
		content := strings.Join(lines, "\n") + "\n"
		if out, err := exec.Command("sh", "-c", fmt.Sprintf("echo '%s' > /etc/resolv.conf", content)).CombinedOutput(); err != nil {
			return nil, errors.Errorf("write resolv.conf failed: %v, %s", err, string(out))
		}
	case "darwin":
		// macOS 需要针对每个服务设置；这里使用 networksetup -setdnsservers for all 'en*'. 简化方案：en0/en1/en2...
		for i := 0; i < 10; i++ {
			dev := fmt.Sprintf("en%d", i)
			args := append([]string{"-setdnsservers", dev}, req.DNS...)
			exec.Command("networksetup", args...).CombinedOutput()
		}
	default:
		return nil, errors.Errorf("unsupported OS")
	}
	return &v1.UpdateDNSRes{}, nil
}

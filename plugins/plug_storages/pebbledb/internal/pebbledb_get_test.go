package internal

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"testing"

	"github.com/cockroachdb/pebble"
)

// DeviceRecord 用于解析设备库中的 JSON 结构
type DeviceRecord struct {
	DeviceID  string                 `json:"device_id"`
	Timestamp int64                  `json:"timestamp"`
	Metrics   map[string]interface{} `json:"metrics"`
}

// TestGetPagesByTimeRange_DeviceSoc 仅测试 device 库，id=2，key=soc
func TestGetPagesByTimeRange_DeviceSoc(t *testing.T) {
	// 基础库地址（用户提供的 pebbledb 路径）
	// NOTE: 如需调整，可将 basePath 改为实际机器上的 pebbledb 基础目录
	basePath := "/Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/pebbledb"
	devicePath := filepath.Join(basePath, "device")

	db, err := pebble.Open(devicePath, &pebble.Options{})
	if err != nil {
		t.Fatalf("打开 device 库失败: %v (path=%s)", err, devicePath)
	}
	defer func() {
		_ = db.Close()
	}()

	prefix := fmt.Sprintf("%s%s%s", DevicePrefix, KeySeparator, "2")

	items, total, err := GetPagesByTimeRange(db, prefix, nil, nil, 1, 1000, SortOrderAsc, 0)
	if err != nil {
		t.Fatalf("查询失败: %v", err)
	}
	if total == 0 || len(items) == 0 {
		t.Fatalf("未查询到任何数据: total=%d len=%d (prefix=%s)", total, len(items), prefix)
	}

	// 校验时间戳升序
	for i := 1; i < len(items); i++ {
		if items[i-1].Timestamp > items[i].Timestamp {
			t.Errorf("时间戳未按升序排序: %d > %d (index %d)", items[i-1].Timestamp, items[i].Timestamp, i)
		}
	}

	// 查找包含 metrics.soc 的记录
	foundSoc := false
	for _, it := range items {
		var rec DeviceRecord
		if err := json.Unmarshal(it.Value, &rec); err != nil {
			t.Errorf("JSON 解析失败: %v, key=%s", err, it.Key)
			continue
		}
		if rec.Metrics != nil {
			if _, ok := rec.Metrics["soc"]; ok {
				foundSoc = true
				break
			}
		}
		fmt.Println(rec)
	}

	if !foundSoc {
		t.Fatalf("在 device/2 的返回结果中未找到 metrics.soc 字段")
	}
}

// TestGetFirstRecord_DeviceDB 获取 device 库中的第一条记录（按键的字典序）
func TestGetFirstRecord_DeviceDB(t *testing.T) {
	basePath := "/Users/zhao/Documents/01.Code/Hex/ems-plan/application/out/pebbledb"
	devicePath := filepath.Join(basePath, "device")

	db, err := pebble.Open(devicePath, &pebble.Options{})
	if err != nil {
		t.Fatalf("打开 device 库失败: %v (path=%s)", err, devicePath)
	}
	defer func() { _ = db.Close() }()

	iter, err := db.NewIter(nil)
	if err != nil {
		t.Fatalf("创建迭代器失败: %v", err)
	}
	defer iter.Close()

	if !iter.First() || !iter.Valid() {
		t.Fatalf("device 库为空或无有效记录")
	}

	key := string(iter.Key())
	val := make([]byte, len(iter.Value()))
	copy(val, iter.Value())

	t.Logf("first key: %s", key)

	var rec DeviceRecord
	if err := json.Unmarshal(val, &rec); err != nil {
		t.Fatalf("JSON 解析失败: %v, key=%s", err, key)
	}

	if rec.Timestamp == 0 {
		t.Errorf("记录的 timestamp 异常: %d", rec.Timestamp)
	}
	if rec.Metrics == nil {
		t.Errorf("记录的 metrics 为空")
	}

	t.Logf("first record: device_id=%s timestamp=%d", rec.DeviceID, rec.Timestamp)
}

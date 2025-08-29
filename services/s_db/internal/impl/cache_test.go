package impl

import (
	"context"
	"testing"
	"time"

	"github.com/gogf/gf/v2/test/gtest"
)

func TestAlarmServiceCache(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		ctx := context.Background()
		service := GetAlarmService()

		// 测试计数缓存
		t.Assert(service.GetAlarmHistoryCount(ctx), 0)
		t.Assert(service.GetAlarmIgnoreCount(ctx), 0)

		// 测试忽略状态缓存
		isIgnored, err := service.IsAlarmIgnored(ctx, "test_device", "test_source", "test_point")
		t.AssertNil(err)
		t.Assert(isIgnored, false) // 默认应该未被忽略

		// 再次调用以测试缓存是否生效
		t.Assert(service.GetAlarmHistoryCount(ctx), 0)
		t.Assert(service.GetAlarmIgnoreCount(ctx), 0)
	})
}

func TestCacheDuration(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		service := GetAlarmService()

		// 验证缓存过期时间设置
		impl := service.(*sAlarmServiceImpl)
		t.Assert(impl.cacheDuration, 5*time.Minute)

		// 验证缓存实例已初始化
		t.AssertNE(impl.ignoreCache, nil)
		t.AssertNE(impl.countCache, nil)
	})
}

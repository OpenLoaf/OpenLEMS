package c_util

import (
	"fmt"
	"time"
)

// FormatDuration 格式化时间间隔为可读的字符串
// 规则：
// - 小于1分钟：显示秒数
// - 小于60分钟：显示分钟数
// - 小于1天：显示小时和分钟
// - 大于等于1天：显示天、小时和分钟
func FormatDuration(d time.Duration) string {
	totalSeconds := int64(d.Seconds())

	if totalSeconds < 60 {
		// 小于1分钟，显示秒数
		return fmt.Sprintf("%d秒", totalSeconds)
	}

	totalMinutes := totalSeconds / 60
	if totalMinutes < 60 {
		// 小于60分钟，显示分钟数
		return fmt.Sprintf("%d分钟%d秒", totalMinutes, totalSeconds)
	}

	totalHours := totalMinutes / 60
	if totalHours < 24 {
		// 小于1天，显示小时和分钟
		remainingMinutes := totalMinutes % 60
		if remainingMinutes == 0 {
			return fmt.Sprintf("%d小时", totalHours)
		}
		return fmt.Sprintf("%d小时%d分钟", totalHours, remainingMinutes)
	}

	// 大于等于1天，显示天、小时和分钟
	days := totalHours / 24
	remainingHours := totalHours % 24
	remainingMinutes := totalMinutes % 60

	var result string
	if days > 0 {
		result = fmt.Sprintf("%d天", days)
	}
	if remainingHours > 0 {
		if result != "" {
			result += fmt.Sprintf("%d小时", remainingHours)
		} else {
			result = fmt.Sprintf("%d小时", remainingHours)
		}
	}
	if remainingMinutes > 0 {
		if result != "" {
			result += fmt.Sprintf("%d分钟", remainingMinutes)
		} else {
			result = fmt.Sprintf("%d分钟", remainingMinutes)
		}
	}

	return result
}

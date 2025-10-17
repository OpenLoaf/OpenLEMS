package internal

import (
	"time"

	"github.com/pkg/errors"
)

// ValidateDateRange 验证日期范围配置
func ValidateDateRange(dr *SDateRange) error {
	if dr == nil {
		return nil
	}

	if dr.StartDate == "" {
		return errors.New("开始日期不能为空")
	}

	// 验证日期格式
	if _, err := time.Parse("2006-01-02", dr.StartDate); err != nil {
		return errors.New("开始日期格式错误，应为 YYYY-MM-DD")
	}

	if !dr.IsLongTerm && dr.EndDate != "" {
		if _, err := time.Parse("2006-01-02", dr.EndDate); err != nil {
			return errors.New("结束日期格式错误，应为 YYYY-MM-DD")
		}
		if dr.EndDate < dr.StartDate {
			return errors.New("结束日期不能早于开始日期")
		}
	}

	return nil
}

// ValidateTimeRange 验证时间范围配置
func ValidateTimeRange(tr *STimeRange) error {
	if tr == nil {
		return nil
	}

	switch tr.Type {
	case "weekday":
		if tr.WeekdayType != "workday" && tr.WeekdayType != "weekend" && tr.WeekdayType != "all" {
			return errors.New("工作日类型必须是 workday、weekend 或 all")
		}
	case "custom":
		if len(tr.CustomDays) == 0 {
			return errors.New("自定义周内日不能为空")
		}
		for _, day := range tr.CustomDays {
			if day < 0 || day > 6 {
				return errors.New("自定义周内日必须在 0-6 之间")
			}
		}
	case "monthly":
		if len(tr.CustomMonths) == 0 {
			return errors.New("指定月份不能为空")
		}
		for _, month := range tr.CustomMonths {
			if month < 1 || month > 12 {
				return errors.New("指定月份必须在 1-12 之间")
			}
		}
	default:
		return errors.New("时间类型必须是 weekday、custom 或 monthly")
	}

	return nil
}

// ValidateStrategyConfig 验证策略配置
func ValidateStrategyConfig(cfg *SStrategyConfig) error {
	if cfg == nil {
		return errors.New("策略配置不能为空")
	}

	if cfg.SocMinRatio < 0 || cfg.SocMinRatio > 100 {
		return errors.New("最小SOC比例必须在 0-100 之间")
	}

	if cfg.SocMaxRatio < 0 || cfg.SocMaxRatio > 100 {
		return errors.New("最大SOC比例必须在 0-100 之间")
	}

	if cfg.SocMinRatio >= cfg.SocMaxRatio {
		return errors.New("最小SOC比例必须小于最大SOC比例")
	}

	if cfg.EnableHealthOptimization && (cfg.MonthlyChargeDay < 1 || cfg.MonthlyChargeDay > 28) {
		return errors.New("月度充电日必须在 1-28 之间")
	}

	if len(cfg.Points) > 0 {
		seen := make(map[int]bool)
		for _, point := range cfg.Points {
			if point[0] < 0 || point[0] > 24 {
				return errors.New("小时必须在 0-24 之间")
			}
			if point[1] < 0 || point[1] > 100 {
				return errors.New("百分比必须在 0-100 之间")
			}
			if seen[point[0]] {
				return errors.New("小时目标点位中小时不能重复")
			}
			seen[point[0]] = true
		}
	}

	return nil
}

// ValidateStrategy 验证完整策略（供外部调用）
func ValidateStrategy(dateRange *SDateRange, timeRange *STimeRange, config *SStrategyConfig, essDeviceIds []string) error {
	if err := ValidateDateRange(dateRange); err != nil {
		return errors.Wrap(err, "日期范围验证失败")
	}

	if err := ValidateTimeRange(timeRange); err != nil {
		return errors.Wrap(err, "时间范围验证失败")
	}

	if err := ValidateStrategyConfig(config); err != nil {
		return errors.Wrap(err, "策略配置验证失败")
	}

	if len(essDeviceIds) == 0 {
		return errors.New("储能设备ID列表不能为空")
	}

	return nil
}

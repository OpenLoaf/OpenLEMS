package alarm

import (
	v1 "application/api/alarm/v1"
	"common"
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"strings"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

// ClearAlarm 清除告警
func (c *ControllerV1) ClearAlarm(ctx context.Context, req *v1.ClearAlarmReq) (res *v1.ClearAlarmRes, err error) {
	// 参数验证
	if strings.TrimSpace(req.Level) == "" {
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "告警类型不能为空")
	}

	// 验证告警类型是否有效
	var alarmLevel c_enum.EAlarmLevel
	switch strings.ToUpper(req.Level) {
	case "NONE":
		alarmLevel = c_enum.EAlarmLevelNone
	case "WARN":
		alarmLevel = c_enum.EAlarmLevelWarn
	case "ALERT":
		alarmLevel = c_enum.EAlarmLevelAlert
	case "ERROR":
		alarmLevel = c_enum.EAlarmLevelError
	default:
		return nil, gerror.NewCode(gcode.CodeInvalidParameter, "无效的告警类型")
	}

	// 统计清除的告警数量
	clearedCount := 0

	// 遍历设备清除告警
	common.GetDeviceManager().IteratorChildDevicesById(req.DeviceId, func(config *c_base.SDeviceConfig, device c_base.IDevice) bool {
		if device == nil {
			return true
		}

		// 获取当前设备的告警列表
		alarmList := device.GetAlarmList()
		if len(alarmList) == 0 {
			return true
		}

		// 遍历告警列表，清除指定类型的告警
		for _, alarm := range alarmList {
			if alarm.GetLevel() == alarmLevel {
				// 调用设备的忽略清除告警方法
				device.IgnoreClearAlarm(alarm.GetDeviceId(), alarm.IPoint.GetKey())
				clearedCount++

				c_log.BizInfof(ctx, "清除告警 - 设备: %s, 点位: %s, 类型: %s",
					config.Name, alarm.IPoint.GetKey(), alarm.GetLevel().String())
			}
		}

		return true
	})

	// 记录清除结果日志
	if req.DeviceId != "" {
		c_log.BizInfof(ctx, "清除设备[%s]的%s类型告警完成，共清除%d条",
			req.DeviceId, req.Level, clearedCount)
	} else {
		c_log.BizInfof(ctx, "清除所有设备的%s类型告警完成，共清除%d条",
			req.Level, clearedCount)
	}

	return &v1.ClearAlarmRes{}, nil
}

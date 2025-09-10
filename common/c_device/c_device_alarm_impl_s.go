package c_device

import (
	"common"
	"common/c_alarm"
	"common/c_base"
	"common/c_enum"
	"common/c_log"
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/pkg/errors"
)

type sAlarmHandler struct {
	action  c_enum.EAlarmAction
	sort    int
	handler []func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isHandler bool)
}

type sAlarmImpl struct {
	ctx              context.Context
	rwMutex          sync.RWMutex
	deviceId         string                         // 当前设备的ID
	parentDeviceId   string                         // 父设备ID
	maxLevel         c_enum.EAlarmLevel             // 最高等级
	cache            map[string]*c_base.SPointValue // 缓存
	alarmHandlerList []*sAlarmHandler               // 告警处理器列表
}

func NewAlarmImpl(ctx context.Context, deviceId string, parentDeviceId string) c_base.IAlarm {
	return &sAlarmImpl{
		ctx:              ctx,
		rwMutex:          sync.RWMutex{},
		deviceId:         deviceId,
		parentDeviceId:   parentDeviceId,
		maxLevel:         c_enum.EAlarmLevelNone,
		cache:            make(map[string]*c_base.SPointValue),
		alarmHandlerList: make([]*sAlarmHandler, 0),
	}
}

func (s *sAlarmImpl) ResetAlarm() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// 调用注册的处理器，通知告警重置
	s.callHandlers(nil, s.maxLevel, c_enum.EAlarmActionReset)

	// 重置缓存和最高级别
	s.cache = make(map[string]*c_base.SPointValue)
	s.maxLevel = c_enum.EAlarmLevelNone
}

func (s *sAlarmImpl) RegisterAlarmHandlerFunc(alarmAction c_enum.EAlarmAction, handler func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isFirstHandler bool), sortValue ...int) {
	if handler == nil {
		c_log.Errorf(s.ctx, "RegisterAlarmHandlerFunc: handler不能为空")
		return
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// 设置默认排序值
	sortOrder := 0
	if len(sortValue) > 0 {
		sortOrder = sortValue[0]
	}

	// 创建新的处理器
	newHandler := &sAlarmHandler{
		action:  alarmAction,
		sort:    sortOrder,
		handler: []func(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, isHandler bool){handler},
	}

	// 添加到列表中
	s.alarmHandlerList = append(s.alarmHandlerList, newHandler)

	// 按sort和action进行排序
	sort.Slice(s.alarmHandlerList, func(i, j int) bool {
		// 首先按sort值排序（升序，小的在前面）
		if s.alarmHandlerList[i].sort != s.alarmHandlerList[j].sort {
			return s.alarmHandlerList[i].sort < s.alarmHandlerList[j].sort
		}

		// sort相同时，按action值排序（升序，小的在前面）
		return s.alarmHandlerList[i].action < s.alarmHandlerList[j].action
	})
}

func (s *sAlarmImpl) IgnoreClearAlarm(deviceId string, point string) {
	key := s.GetAlarmKey(deviceId, point)
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if alarm, ok := s.cache[key]; ok {
		delete(s.cache, key)
		s.updateMaxLevel()

		// 执行告警忽略
		s.callHandlers(alarm, s.maxLevel, c_enum.EAlarmActionIgnore)

		// 保存到告警历史中
		historyMessage := fmt.Sprintf("手动屏蔽告警，触发值为:%v", alarm.GetValue())
		err := c_alarm.GetAlarmManager().CreateAlarmHistory(s.ctx, s.deviceId, deviceId, alarm.IPoint, alarm.GetLevel(), historyMessage, alarm.GetHappenTime())
		if err != nil {
			c_log.Errorf(s.ctx, "保存告警记录失败！%+v", err)
		}
	}
}

func (s *sAlarmImpl) UpdateAlarm(deviceId string, point c_base.IPoint, value any) {
	if point == nil {
		c_log.Errorf(s.ctx, "告警元数据不能为空")
		return
	}

	// 提前点位是否是告警，避免不必要的处理
	if point.IsNotAlarm() {
		return
	}

	// 判断一下这个点位是否被忽略，忽略的不用触发
	if isAlarmIgnore, err := c_alarm.GetAlarmManager().IsAlarmIgnored(s.ctx, s.deviceId, deviceId, point.GetName()); err == nil && isAlarmIgnore {
		c_log.Debugf(s.ctx, "忽略设备[%s]的告警[%s]", deviceId, point.GetName())
		return
	}

	// 先获取当前Trigger的返回值，true代表触发，false代表清除
	isCurrentlyTriggered, level, err := point.AlarmTrigger(value)
	if err != nil {
		c_log.BizErrorf(s.ctx, "告警触发判断失败！%+v", err)
		return
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// 生成告警key
	alarmKey := s.GetAlarmKey(deviceId, point.GetName())

	// 获取缓存中的旧告警，判断之前是否已经触发过
	oldAlarm, wasPreviouslyTriggered := s.cache[alarmKey]

	// 记录更新前的最高告警级别
	oldMaxLevel := s.maxLevel

	// 根据当前状态和之前状态来决定动作
	var alarmAction c_enum.EAlarmAction
	var alarm *c_base.SPointValue

	deviceConfig := common.GetDeviceManager().GetDeviceConfigById(deviceId)
	alarmDeviceName := deviceConfig.Name

	if isCurrentlyTriggered {
		// 当前需要触发告警
		alarm = c_base.NewPointValue(deviceId, point, level, value)

		if !wasPreviouslyTriggered {
			// 之前没有触发过，现在是首次触发
			alarmAction = c_enum.EAlarmActionFirstTrigger
			s.cache[alarmKey] = alarm

			// 记录触发日志
			switch level {
			case c_enum.EAlarmLevelWarn:
				c_log.BizWarningf(s.ctx, fmt.Sprintf("触发%s[%s]警告！值为: %v", alarmDeviceName, alarm.IPoint.GetKey(), value))
			case c_enum.EAlarmLevelAlert:
				c_log.BizWarningf(s.ctx, fmt.Sprintf("触发%s[%s]警报！值为: %v", alarmDeviceName, alarm.IPoint.GetKey(), value))
			case c_enum.EAlarmLevelError:
				c_log.BizWarningf(s.ctx, fmt.Sprintf("触发%s[%s]故障！值为: %v", alarmDeviceName, alarm.IPoint.GetKey(), value))
			case c_enum.EAlarmLevelNone:
			}
		} else {
			// 之前已经触发过，现在继续触发（非首次）
			alarmAction = c_enum.EAlarmActionNotFirstTrigger
		}
	} else {
		// 当前需要清除告警
		if wasPreviouslyTriggered {
			// 之前触发过，现在需要清除
			alarmAction = c_enum.EAlarmActionFirstClear
			alarm = oldAlarm // 使用旧的告警信息用于日志和处理器

			var historyMessage string

			newValueExplain, _ := alarm.IPoint.ValueExplain(value)
			oldValueExplain, _ := oldAlarm.IPoint.ValueExplain(oldAlarm.GetValue())

			if newValueExplain != "" && oldValueExplain != "" {
				historyMessage = fmt.Sprintf("触发值为:%v，告警清除后值为:%v", oldValueExplain, newValueExplain)
			} else {
				historyMessage = fmt.Sprintf("触发值为:%v，告警清除后值为:%v", oldAlarm.GetValue(), value)
			}

			err := c_alarm.GetAlarmManager().CreateAlarmHistory(s.ctx, s.deviceId, deviceId, point, oldAlarm.GetLevel(), historyMessage, oldAlarm.GetHappenTime())
			if err != nil {
				c_log.Errorf(s.ctx, "保存告警记录失败！%+v", err)
			}

			// 从缓存中删除告警
			delete(s.cache, alarmKey)

			// 记录清除日志
			switch level {
			case c_enum.EAlarmLevelWarn:
				c_log.BizInfof(s.ctx, fmt.Sprintf("清除%s[%s]警告！值为: %v", alarmDeviceName, alarm.IPoint.GetKey(), value))
			case c_enum.EAlarmLevelAlert:
				c_log.BizInfof(s.ctx, fmt.Sprintf("清除%s[%s]警报！值为: %v", alarmDeviceName, alarm.IPoint.GetKey(), value))
			case c_enum.EAlarmLevelError:
				c_log.BizInfof(s.ctx, fmt.Sprintf("清除%s[%s]故障！值为: %v", alarmDeviceName, alarm.IPoint.GetKey(), value))
			case c_enum.EAlarmLevelNone:
			}
		} else {
			// 之前没有触发过，现在也不需要触发，无需处理
			return
		}
	}

	// 更新最高告警级别
	s.updateMaxLevel()

	// 调用注册的处理器
	s.callHandlers(alarm, s.maxLevel, alarmAction)

	// 基于maxLevel的变化判断告警级别上升或下降
	if s.maxLevel > oldMaxLevel {
		// 告警级别上升
		s.callHandlers(alarm, s.maxLevel, c_enum.EAlarmActionLevelUp)
	} else if s.maxLevel < oldMaxLevel {
		// 告警级别下降
		s.callHandlers(alarm, s.maxLevel, c_enum.EAlarmActionLevelDown)
	}

	// 触发父设备告警
	parentDevice := common.GetDeviceManager().GetDeviceById(s.parentDeviceId)
	if parentDevice != nil {
		// 新开一个协程去通知父节点告警, 注意这里传递的还是告警设备的deviceId
		go parentDevice.UpdateAlarm(deviceId, point, value)
	}

}

// validateAlarmInput 验证告警输入参数
func (s *sAlarmImpl) validateAlarmInput(alarm *c_base.SPointValue) error {
	if alarm == nil {
		return errors.Errorf("alarm参数不能为空")
	}

	// 注意：Level验证已在UpdateAlarm方法中提前处理，这里不再重复验证

	if alarm.GetValue() == nil {
		return errors.Errorf("告警值不能为空")
	}

	if alarm.IPoint == nil {
		return errors.Errorf("告警元数据不能为空")
	}

	if strings.TrimSpace(alarm.GetDeviceId()) == "" {
		return errors.Errorf("设备ID不能为空")
	}

	if strings.TrimSpace(alarm.IPoint.GetKey()) == "" {
		return errors.Errorf("点位名称不能为空")
	}

	// 验证设备ID长度，防止过长的ID导致内存问题
	if len(alarm.GetDeviceId()) > 256 {
		return errors.Errorf("设备ID长度不能超过256字符")
	}

	// 验证点位名称长度
	if len(alarm.IPoint.GetKey()) > 256 {
		return errors.Errorf("点位名称长度不能超过256字符")
	}

	return nil
}

// updateMaxLevel 更新最高告警级别
func (s *sAlarmImpl) updateMaxLevel() {
	s.maxLevel = c_enum.EAlarmLevelNone
	for _, alarm := range s.cache {
		if alarm.GetLevel() > s.maxLevel {
			s.maxLevel = alarm.GetLevel()
		}
	}
}

// callHandlers 调用注册的处理器
func (s *sAlarmImpl) callHandlers(alarm *c_base.SPointValue, currentMaxAlarmLevel c_enum.EAlarmLevel, action c_enum.EAlarmAction) {
	isFirstHandler := true
	for _, handler := range s.alarmHandlerList {
		if handler.action == action || handler.action == c_enum.EAlarmActionEvery { // 执行相同的action，或者是c_enum.EAlarmActionEvery
			for _, h := range handler.handler {
				// 添加panic恢复机制
				func() {
					defer func() {
						if r := recover(); r != nil {
							c_log.Errorf(s.ctx, "告警处理器panic: %v", r)
						}
					}()
					h(alarm, currentMaxAlarmLevel, isFirstHandler)
				}()
				isFirstHandler = false
			}
		}
	}

}

func (s *sAlarmImpl) GetAlarmList() []*c_base.SPointValue {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	// 从缓存中获取所有告警数据
	alarmList := make([]*c_base.SPointValue, 0, len(s.cache))
	for _, alarm := range s.cache {
		alarmList = append(alarmList, alarm)
	}

	// 根据设备ID和发生时间倒序进行排序
	sort.Slice(alarmList, func(i, j int) bool {
		// 首先按设备ID排序
		if alarmList[i].GetDeviceId() != alarmList[j].GetDeviceId() {
			return alarmList[i].GetDeviceId() < alarmList[j].GetDeviceId()
		}

		// 设备ID相同时，按发生时间倒序排序（最新的在前）
		return alarmList[i].GetHappenTime().After(alarmList[j].GetHappenTime())
	})

	return alarmList
}

func (s *sAlarmImpl) GetAlarmLevel() c_enum.EAlarmLevel {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()
	return s.maxLevel
}

// GetAlarmKey 拼接告警key
func (s *sAlarmImpl) GetAlarmKey(deviceId, pointName string) string {
	return fmt.Sprintf("%s:%s", deviceId, pointName)
}

// CleanupHandlers 清理处理器列表（可选的方法，用于防止内存泄漏）
func (s *sAlarmImpl) CleanupHandlers() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// 这里可以实现清理逻辑，比如移除长时间未使用的处理器
	// 目前保持简单实现，实际使用时可以根据需要扩展
}

package c_device

import (
	"common"
	"common/c_alarm"
	"common/c_base"
	"common/c_log"
	"context"
	"fmt"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/shockerli/cvt"
)

type sAlarmHandler struct {
	action  c_base.EAlarmAction
	sort    int
	handler []func(alarm *c_base.MetaValueWrapper, currentMaxAlarmLevel c_base.EAlarmLevel, isHandler bool)
}

type sAlarmImpl struct {
	ctx              context.Context
	rwMutex          sync.RWMutex
	deviceId         string                              // 当前设备的ID
	parentDeviceId   string                              // 父设备ID
	maxLevel         c_base.EAlarmLevel                  // 最高等级
	cache            map[string]*c_base.MetaValueWrapper // 缓存
	alarmHandlerList []*sAlarmHandler                    // 告警处理器列表
}

func NewAlarmImpl(ctx context.Context, deviceId string, parentDeviceId string) c_base.IAlarm {
	return &sAlarmImpl{
		ctx:              ctx,
		rwMutex:          sync.RWMutex{},
		deviceId:         deviceId,
		parentDeviceId:   parentDeviceId,
		maxLevel:         c_base.EAlarmLevelNone,
		cache:            make(map[string]*c_base.MetaValueWrapper),
		alarmHandlerList: make([]*sAlarmHandler, 0),
	}
}

func (s *sAlarmImpl) ResetAlarm() {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// 调用注册的处理器，通知告警重置
	s.callHandlers(nil, s.maxLevel, c_base.EAlarmActionReset)

	// 重置缓存和最高级别
	s.cache = make(map[string]*c_base.MetaValueWrapper)
	s.maxLevel = c_base.EAlarmLevelNone
}

func (s *sAlarmImpl) RegisterAlarmHandlerFunc(alarmAction c_base.EAlarmAction, handler func(alarm *c_base.MetaValueWrapper, currentMaxAlarmLevel c_base.EAlarmLevel, isFirstHandler bool), sortValue ...int) {
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
		handler: []func(alarm *c_base.MetaValueWrapper, currentMaxAlarmLevel c_base.EAlarmLevel, isHandler bool){handler},
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

func (s *sAlarmImpl) UpdateAlarm(deviceId string, deviceType c_base.EDeviceType, meta *c_base.Meta, value any) {
	if meta == nil {
		c_log.Errorf(s.ctx, "告警元数据不能为空")
		return
	}

	// 提前判断告警级别，如果是None则直接退出，避免不必要的处理
	if meta.Level == c_base.EAlarmLevelNone {
		return
	}

	// 先获取当前Trigger的返回值，true代表触发，false代表清除
	var isCurrentlyTriggered bool
	if meta.Trigger != nil {
		isCurrentlyTriggered = meta.Trigger(value)
	} else {
		// 如果没有Trigger函数，默认根据布尔值判断
		isCurrentlyTriggered = value != nil && cvt.Bool(value) != false
	}

	if isCurrentlyTriggered {
		// 如果触发了，就判断一下这个点位是否被忽略，忽略的不用触发
		isAlarmIgnore, err := c_alarm.GetAlarmManager().IsAlarmIgnored(s.ctx, deviceId, meta.Name)
		if err != nil {
			c_log.Errorf(s.ctx, "IsAlarmIgnored err:%v", err)
		}
		if isAlarmIgnore {
			return
		}
	}

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	// 生成告警key
	alarmKey := s.GetAlarmKey(deviceId, meta.Name)

	// 获取缓存中的旧告警，判断之前是否已经触发过
	oldAlarm, wasPreviouslyTriggered := s.cache[alarmKey]

	// 记录更新前的最高告警级别
	oldMaxLevel := s.maxLevel

	// 根据当前状态和之前状态来决定动作
	var alarmAction c_base.EAlarmAction
	var alarm *c_base.MetaValueWrapper

	alarmDeviceId := deviceId // 告警的设备ID，如果是父设备收到子设备的，那么alarmDeviceId应该是父设备
	alarmDeviceName := ""
	if deviceId != s.deviceId {
		// 说明是从下级设备处传递过来的
		deviceConfig := common.GetDeviceManager().GetDeviceConfigById(deviceId)
		if deviceConfig != nil {
			alarmDeviceId = s.deviceId // 使用当前设备的ID
			alarmDeviceName = deviceConfig.Name
		}
	}

	if isCurrentlyTriggered {
		// 当前需要触发告警
		alarm = &c_base.MetaValueWrapper{
			DeviceId:   alarmDeviceId,
			DeviceType: deviceType,
			Level:      meta.Level,
			Meta:       meta,
			Value:      value,
			HappenTime: nil,
		}

		if !wasPreviouslyTriggered {
			// 之前没有触发过，现在是首次触发
			alarmAction = c_base.EAlarmActionFirstTrigger
			now := time.Now()
			alarm.HappenTime = &now // 告警触发，记录触发时间
			s.cache[alarmKey] = alarm

			// 记录触发日志
			switch alarm.Level {
			case c_base.EAlarmLevelWarn:
				c_log.BizWarningf(s.ctx, fmt.Sprintf("触发%s[%s]警告！值为: %v", alarmDeviceName, alarm.Meta.Cn, value))
			case c_base.EAlarmLevelAlarm:
				c_log.BizWarningf(s.ctx, fmt.Sprintf("触发%s[%s]警报！值为: %v", alarmDeviceName, alarm.Meta.Cn, value))
			case c_base.EAlarmLevelError:
				c_log.BizWarningf(s.ctx, fmt.Sprintf("触发%s[%s]故障！值为: %v", alarmDeviceName, alarm.Meta.Cn, value))
			case c_base.EAlarmLevelNone:
			}
		} else {
			// 之前已经触发过，现在继续触发（非首次）
			alarmAction = c_base.EAlarmActionNotFirstTrigger
			// 保持原有的触发时间
			alarm.HappenTime = oldAlarm.HappenTime
			s.cache[alarmKey] = alarm
		}
	} else {
		// 当前需要清除告警
		if wasPreviouslyTriggered {
			// 之前触发过，现在需要清除
			alarmAction = c_base.EAlarmActionFirstClear
			alarm = oldAlarm // 使用旧的告警信息用于日志和处理器

			historyMessage := fmt.Sprintf("触发于:%s，触发值为:%v，告警清除后值为:%v", oldAlarm.HappenTime.Format("2006-01-02 15:04:05.000"), oldAlarm.Value, value)
			err := c_alarm.GetAlarmManager().CreateAlarmHistory(s.ctx, alarmDeviceId, meta.Name, s.maxLevel.String(), meta.Cn, historyMessage, oldAlarm.HappenTime)
			if err != nil {
				c_log.Errorf(s.ctx, "保存告警记录失败！%+v", err)
			}

			// 从缓存中删除告警
			delete(s.cache, alarmKey)

			// 记录清除日志
			switch alarm.Level {
			case c_base.EAlarmLevelWarn:
				c_log.BizInfof(s.ctx, fmt.Sprintf("清除%s[%s]警告！值为: %v", alarmDeviceName, alarm.Meta.Cn, value))
			case c_base.EAlarmLevelAlarm:
				c_log.BizInfof(s.ctx, fmt.Sprintf("清除%s[%s]警报！值为: %v", alarmDeviceName, alarm.Meta.Cn, value))
			case c_base.EAlarmLevelError:
				c_log.BizInfof(s.ctx, fmt.Sprintf("清除%s[%s]故障！值为: %v", alarmDeviceName, alarm.Meta.Cn, value))
			case c_base.EAlarmLevelNone:
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
		s.callHandlers(alarm, s.maxLevel, c_base.EAlarmActionLevelUp)
	} else if s.maxLevel < oldMaxLevel {
		// 告警级别下降
		s.callHandlers(alarm, s.maxLevel, c_base.EAlarmActionLevelDown)
	}

	// 触发父设备告警
	parentDevice := common.GetDeviceManager().GetDeviceById(s.parentDeviceId)
	if parentDevice != nil {
		// 新开一个协程去通知父节点告警
		go parentDevice.UpdateAlarm(alarmDeviceId, deviceType, meta, value)
	}

}

// validateAlarmInput 验证告警输入参数
func (s *sAlarmImpl) validateAlarmInput(alarm *c_base.MetaValueWrapper) error {
	if alarm == nil {
		return errors.Errorf("alarm参数不能为空")
	}

	// 注意：Level验证已在UpdateAlarm方法中提前处理，这里不再重复验证

	if alarm.Value == nil {
		return errors.Errorf("告警值不能为空")
	}

	if alarm.Meta == nil {
		return errors.Errorf("告警元数据不能为空")
	}

	if strings.TrimSpace(alarm.DeviceId) == "" {
		return errors.Errorf("设备ID不能为空")
	}

	if strings.TrimSpace(alarm.Meta.Name) == "" {
		return errors.Errorf("点位名称不能为空")
	}

	// 验证设备ID长度，防止过长的ID导致内存问题
	if len(alarm.DeviceId) > 256 {
		return errors.Errorf("设备ID长度不能超过256字符")
	}

	// 验证点位名称长度
	if len(alarm.Meta.Name) > 256 {
		return errors.Errorf("点位名称长度不能超过256字符")
	}

	return nil
}

// updateMaxLevel 更新最高告警级别
func (s *sAlarmImpl) updateMaxLevel() {
	s.maxLevel = c_base.EAlarmLevelNone
	for _, alarm := range s.cache {
		if alarm.Level > s.maxLevel {
			s.maxLevel = alarm.Level
		}
	}
}

// callHandlers 调用注册的处理器
func (s *sAlarmImpl) callHandlers(alarm *c_base.MetaValueWrapper, currentMaxAlarmLevel c_base.EAlarmLevel, action c_base.EAlarmAction) {
	isFirstHandler := true
	for _, handler := range s.alarmHandlerList {
		if handler.action == action || handler.action == c_base.EAlarmActionEvery { // 执行相同的action，或者是EAlarmActionEvery
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

func (s *sAlarmImpl) GetAlarmList() []*c_base.MetaValueWrapper {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	// 从缓存中获取所有告警数据
	alarmList := make([]*c_base.MetaValueWrapper, 0, len(s.cache))
	for _, alarm := range s.cache {
		alarmList = append(alarmList, alarm)
	}

	// 根据设备ID和发生时间倒序进行排序
	sort.Slice(alarmList, func(i, j int) bool {
		// 首先按设备ID排序
		if alarmList[i].DeviceId != alarmList[j].DeviceId {
			return alarmList[i].DeviceId < alarmList[j].DeviceId
		}

		// 设备ID相同时，按发生时间倒序排序（最新的在前）
		if alarmList[i].HappenTime != nil && alarmList[j].HappenTime != nil {
			return alarmList[i].HappenTime.After(*alarmList[j].HappenTime)
		}

		// 如果其中一个时间为nil，将nil时间排在后面
		if alarmList[i].HappenTime == nil {
			return false
		}
		if alarmList[j].HappenTime == nil {
			return true
		}

		return false
	})

	return alarmList
}

func (s *sAlarmImpl) GetAlarmLevel() c_base.EAlarmLevel {
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

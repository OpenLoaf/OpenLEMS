package sess_basic_v1

import (
	"context"
	_ "embed"
	common "ems-plan"
	"ems-plan/c_base"
	"ems-plan/c_device"
	"ems-plan/c_error"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtimer"
	"plug_protocol_gpio_sysfs/p_gpio_sysfs"
	"time"
)

type sStationEnergyStore struct {
	*c_base.SAlarmHandler
	*c_base.SDescription
	deviceConfig *c_base.SDriverConfig

	ctx          context.Context
	rootAmmeter  c_device.IAmmeter       // 储能总电表
	energyStores []c_device.IEnergyStore // 储能列表

	targetPower         int32   // 目标功率
	targetReactivePower int32   // 目标无功功率
	targetPowerFactor   float32 // 目标功率因数

	ledRunning  c_device.IGpio // 运行灯
	ledFault    c_device.IGpio // 故障灯
	buttonScram c_device.IGpio // 急停按钮
}

func (s *sStationEnergyStore) Destroy() {

}

func (s *sStationEnergyStore) Init(protocol c_base.IProtocol, deviceConfig *c_base.SDriverConfig) {
	s.deviceConfig = deviceConfig

	for _, deviceChild := range deviceConfig.DeviceChildren {
		// 从缓存中获取
		if deviceChild.Id == "" {
			// 如果ID都为空，就报错
			panic(gerror.Newf("设备ID为空！设备名称：%s", deviceChild.Name))
		}
	}
	for _, child := range deviceConfig.DeviceChildren {
		dv := common.GetDeviceById(child.Id)
		if dv == nil {
			panic(gerror.Newf("设备Id: %s 不存在！", child.Id))
		}
		if dv.GetDriverType() == c_base.EDeviceAmmeter {
			s.rootAmmeter = dv.(c_device.IAmmeter)
			g.Log().Infof(s.ctx, "注册总电表成功！")
		}
		if dv.GetDriverType() == c_base.EDeviceEnergyStore {
			s.energyStores = append(s.energyStores, dv.(c_device.IEnergyStore))
			g.Log().Infof(s.ctx, "注册储能柜 %s 到 % s成功！", dv.GetDeviceConfig().Name, s.deviceConfig.Name)
		}

		if dv.GetDeviceConfig().Id == p_gpio_sysfs.IdLedFault && dv.GetDriverType() == c_base.EDeviceGpio {
			s.ledFault = dv.(c_device.IGpio)

			gtimer.SetInterval(s.ctx, time.Second, func(ctx context.Context) {
				status, err := s.GetStatus()
				level := s.GetAlarmLevel()
				if err == nil && status == c_base.EPcsStatusFault {
					_ = s.ledFault.SetHigh()
				} else {
					if level != c_base.EError {
						_ = s.ledFault.SetLow()
					}
				}
			})
			g.Log().Infof(s.ctx, "注册故障灯成功！")
		}
		if dv.GetDeviceConfig().Id == p_gpio_sysfs.IdLedRunning && dv.GetDriverType() == c_base.EDeviceGpio {
			s.ledRunning = dv.(c_device.IGpio)

			gtimer.SetInterval(s.ctx, time.Second, func(ctx context.Context) {
				// 如果功率大于5，就亮灯
				power, err := s.GetPower()
				if err == nil && power > 5 {
					_ = s.ledRunning.SetHigh()
				} else {
					_ = s.ledRunning.SetLow()
				}
			})

			g.Log().Infof(s.ctx, "注册运行灯成功！")
		}

		if dv.GetDeviceConfig().Id == p_gpio_sysfs.IdButtonScram {
			gpioDv := dv.(c_device.IGpio)
			gpioDv.RegisterHandler(func(ctx context.Context, status bool, isChange bool) {
				if status {
					// 设置pcs状态为停机
					//_ = s.SetStatus(c_base.EPcsStatusOff)

					//// 紧急停机
					//if s.ledFault != nil && isChange {
					//	_ = s.ledFault.SetHigh()
					//	g.Log().Warningf(s.ctx, "储能场站触发急停！触发储能场站故障灯！")
					//} else {
					//	g.Log().Warningf(s.ctx, "储能场站触发急停！")
					//}
				}

			})
			s.buttonScram = gpioDv
			g.Log().Infof(s.ctx, "场站注册急停按钮成功！")
		}
	}
	if len(s.energyStores) == 0 {
		panic(gerror.Newf("场站储能组：%s 未注册任何储能柜！", deviceConfig.Name))
	}

	g.Log().Noticef(s.ctx, "场站储能初始化成功！")
}

func (s *sStationEnergyStore) getValue(getValueFunc func(store c_device.IEnergyStore) (float32, error), needUpdate func(value, temp float32) bool) (float32, error) {
	var minTemp float32
	for _, ess := range s.energyStores {
		value, err := getValueFunc(ess)
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 获取数据失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		if minTemp == 0 || value < minTemp {
			minTemp = value
		}
	}
	return minTemp, nil
}

func (s *sStationEnergyStore) GetCellMinTemp() (float32, error) {
	return s.getValue(func(store c_device.IEnergyStore) (float32, error) {
		return store.GetCellMinTemp()
	}, func(value, temp float32) bool {
		return value < temp
	})
}

func (s *sStationEnergyStore) GetCellMaxTemp() (float32, error) {
	return s.getValue(func(store c_device.IEnergyStore) (float32, error) {
		return store.GetCellMaxTemp()
	}, func(value, temp float32) bool {
		return value > temp
	})
}

func (s *sStationEnergyStore) GetCellAvgTemp() (float32, error) {
	// 把所有的温度加起来，然后除以个数
	var temp float32
	var count int
	for _, ess := range s.energyStores {
		value, err := ess.GetCellAvgTemp()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 平均温度失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}

		temp += value
		count++
	}
	if count == 0 {
		return 0, gerror.New("统计平均温度失败！")
	}
	return temp / float32(count), nil
}

func (s *sStationEnergyStore) GetCellMinVoltage() (float32, error) {
	return s.getValue(func(store c_device.IEnergyStore) (float32, error) {
		return store.GetCellMinVoltage()
	}, func(value, temp float32) bool {
		return value < temp
	})
}

func (s *sStationEnergyStore) GetCellMaxVoltage() (float32, error) {
	return s.getValue(func(store c_device.IEnergyStore) (float32, error) {
		return store.GetCellMinVoltage()
	}, func(value, temp float32) bool {
		return value > temp
	})
}

func (s *sStationEnergyStore) GetCellAvgVoltage() (float32, error) {
	// 把所有的电压加起来，然后除以个数
	var voltage float32
	var count int
	for _, ess := range s.energyStores {
		value, err := ess.GetCellAvgVoltage()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 平均电压失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}

		voltage += value
		count++
	}
	if count == 0 {
		return 0, gerror.New("统计平均电压失败！")
	}
	return voltage / float32(count), nil
}

func (s *sStationEnergyStore) GetDriverType() c_base.EDeviceType {
	return c_base.EStationEnergyStore
}

func (s *sStationEnergyStore) GetMetaValueList() []*c_base.MetaValueWrapper {
	var metaValueList []*c_base.MetaValueWrapper
	if s.rootAmmeter != nil {
		metaValueList = append(metaValueList, s.rootAmmeter.GetMetaValueList()...)
	}
	for _, store := range s.energyStores {
		metaValueList = append(metaValueList, store.GetMetaValueList()...)
	}
	return metaValueList
}

func (s *sStationEnergyStore) GetDeviceConfig() *c_base.SDriverConfig {
	return s.deviceConfig
}

func (s *sStationEnergyStore) GetAllowControl() bool {
	// TODO 后面改成只要有一个允许控制就返回true
	return len(s.energyStores) != 0
}

func (s *sStationEnergyStore) GetChildren() []c_base.IDriver {
	var children []c_base.IDriver
	if s.rootAmmeter != nil {
		children = append(children, s.rootAmmeter)
	}
	for _, store := range s.energyStores {
		children = append(children, store)
	}
	return children
}

func (s *sStationEnergyStore) GetLastUpdateTime() *time.Time {
	var lastUpdateTime *time.Time
	for _, ess := range s.energyStores {
		essTime := ess.GetLastUpdateTime()
		if essTime == nil {
			continue
		}
		if lastUpdateTime == nil || essTime.After(*lastUpdateTime) {
			lastUpdateTime = essTime
		}
	}
	if s.rootAmmeter != nil {
		ammeterTime := s.rootAmmeter.GetLastUpdateTime()
		if ammeterTime != nil && (lastUpdateTime == nil || ammeterTime.After(*lastUpdateTime)) {
			lastUpdateTime = ammeterTime
		}
	}
	return lastUpdateTime
}

func (s *sStationEnergyStore) GetSoc() (float32, error) {
	// 取soc平均值
	var (
		soc   float32
		count int
	)
	for _, ess := range s.energyStores {
		value, err := ess.GetSoc()
		if err != nil {
			g.Log().Debugf(s.ctx, "获取储能柜:%s SOC失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		soc += value
		count++
	}
	if count == 0 {
		return 0, gerror.New("设备离线")
	}
	return soc / float32(count), nil
}

func (s *sStationEnergyStore) GetSoh() (float32, error) {
	// 取soh平均值
	var (
		soh   float32
		count int
	)
	for _, ess := range s.energyStores {
		value, err := ess.GetSoh()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s SOH失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		soh += value
		count++
	}
	if count == 0 {
		return 0, gerror.New("设备离线")
	}
	return soh / float32(count), nil
}

func (s *sStationEnergyStore) GetCapacity() (uint32, error) {
	// 取累计容量
	var (
		capacity uint32
		count    uint
	)
	for _, ess := range s.energyStores {
		value, err := ess.GetCapacity()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 容量失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		capacity += value
		count++
	}
	if count == 0 {
		return 0, gerror.New("设备离线")
	}
	return capacity, nil
}

func (s *sStationEnergyStore) GetCycleCount() (uint, error) {
	// 取循环次数
	var cycleCount uint
	for _, ess := range s.energyStores {
		value, err := ess.GetCycleCount()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 循环次数失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		cycleCount += value
	}
	// TODO: 此次以后先缓存到本地，如果获取到的数值小于本地的值，就先返回本地的值，防止数据异常
	return cycleCount, nil
}

func (s *sStationEnergyStore) GetDcPower() (float64, error) {
	// 取直流功率和
	var dcPower float64
	for _, ess := range s.energyStores {
		value, err := ess.GetDcPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 直流功率失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		dcPower += value
	}
	return dcPower, nil
}

func (s *sStationEnergyStore) SetReset() error {
	g.Log().Debugf(s.ctx, "设置储能场站复位")
	var err error
	for _, store := range s.energyStores {
		_err := store.SetReset()
		if _err != nil {
			g.Log().Errorf(s.ctx, "储能柜:%s 复位失败！err:%v", store.GetDeviceConfig().Name, _err)
			err = _err
		}
	}
	return err
}

func (s *sStationEnergyStore) SetStatus(status c_base.EEnergyStoreStatus) error {
	g.Log().Debugf(s.ctx, "设置储能场站状态：%s", status)
	var err error
	for _, store := range s.energyStores {
		_err := store.SetStatus(status)
		if _err != nil {
			g.Log().Errorf(s.ctx, "储能柜:%s 设置状态失败！err:%v", store.GetDeviceConfig().Name, _err)
			err = _err
		}
	}
	return err
}

func (s *sStationEnergyStore) SetGridMode(mode c_base.EGridMode) error {
	var err error
	for _, store := range s.energyStores {
		_err := store.SetGridMode(mode)
		if _err != nil {
			g.Log().Errorf(s.ctx, "储能柜:%s 设置模式失败！err:%v", store.GetDeviceConfig().Name, _err)
			err = _err
		}
	}
	return err
}

func (s *sStationEnergyStore) GetStatus() (c_base.EEnergyStoreStatus, error) {
	// 取状态
	var status c_base.EEnergyStoreStatus
	for _, ess := range s.energyStores {
		singleStatus, err := ess.GetStatus()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 状态失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		if status == c_base.EPcsStatusUnknown {
			status = singleStatus
		} else if status != singleStatus {
			// 如果两个状态不想等
			return c_base.EPcsStatusSync, nil
		}

	}
	return status, nil
}

func (s *sStationEnergyStore) GetGridMode() (c_base.EGridMode, error) {
	// 取电网模式
	var mode c_base.EGridMode
	for _, ess := range s.energyStores {
		singleMode, err := ess.GetGridMode()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 电网模式失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		if mode == c_base.EGridUnknown {
			mode = singleMode
		} else if mode != singleMode {
			// 如果两个状态不想等
			return c_base.EGridSync, nil
		}
	}
	return mode, nil
}

func (s *sStationEnergyStore) SetPower(power int32) error {
	g.Log().Debugf(s.ctx, "设置储能场站功率：%d", power)
	if power == 0 {
		// 功率为0直接下发
		for _, store := range s.energyStores {
			_ = store.SetPower(0)
		}
		return nil
	}

	// 判断一下防止超限
	if power > 0 {
		maxOutputPower, err := s.GetMaxOutputPower()
		g.Log().Debugf(s.ctx, "储能场站最大输出功率：%f", maxOutputPower)
		if err != nil {
			return err
		}
		if power > int32(maxOutputPower) {
			return c_error.OverLimitError
		}
	} else {
		maxInputPower, err := s.GetMaxInputPower()
		g.Log().Debugf(s.ctx, "储能场站最大输入功率：%f", maxInputPower)
		if err != nil {
			return err
		}
		if power < int32(-maxInputPower) {
			return c_error.OverLimitError
		}
	}

	s.targetPower = power
	// TODO选择算法

	essList := make([]*sSessBasic, 0)
	for _, store := range s.energyStores {
		var err error
		soc, err := store.GetSoc()
		if err != nil {
			g.Log().Warningf(s.ctx, "储能柜:%s SOC失败！统计时跳过该柜 err:%v", store.GetDeviceConfig().Name, err)
			continue
		}
		ratedPower := store.GetRatedPower()
		if ratedPower < 0 {
			g.Log().Errorf(s.ctx, "储能柜:%s 额定功率小于0！", store.GetDeviceConfig().Name)
			continue
		}
		inputPower, err := store.GetMaxInputPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "储能柜:%s 最大输入功率失败！", store.GetDeviceConfig().Name)
			continue
		}

		outputPower, err := store.GetMaxOutputPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "储能柜:%s 最大输出功率失败！", store.GetDeviceConfig().Name)
			continue
		}

		count, err := store.GetCycleCount()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 循环次数失败！统计时跳过该柜 err:%v", store.GetDeviceConfig().Name, err)
			continue
		}

		getPower, err := store.GetPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 有功功率失败！统计时跳过该柜 err:%v", store.GetDeviceConfig().Name, err)
			continue
		}
		ess := &sSessBasic{
			Id:                store.GetDeviceConfig().Id,
			Name:              store.GetDeviceConfig().Name,
			CurrentPower:      getPower,
			Soc:               int(soc),
			RatedPower:        int(ratedPower),
			MaxDischargePower: float64(outputPower),
			MaxChargePower:    float64(inputPower),
			CycleCount:        int(count),
			EfficiencyCurve:   []float64{1},
		}
		g.Log().Noticef(s.ctx, "储能柜:%s SOC:%d 额定功率:%d 最小功率:%f 最大功率:%f 循环次数:%d", ess.Name, ess.Soc, ess.RatedPower, ess.MaxDischargePower, ess.MaxChargePower, ess.CycleCount)
		essList = append(essList, ess)
	}

	// 计算功率，先不执行
	allocatePower, err := AllocatePower(float64(power), 200, 0, true, essList)
	if err != nil {
		g.Log().Errorf(s.ctx, "储能柜功率分配失败！使用平均功率分配算法！err:%v", err)
		var singlePower = power / int32(len(s.energyStores))
		for _, store := range s.energyStores {
			_ = store.SetPower(singlePower)
		}
	} else {
		for i, f := range allocatePower {
			for _, store := range s.energyStores {
				if store.GetDeviceConfig().Id == essList[i].Id {
					value := int32(0)
					if power < 0 {
						// 功率小于0的时候，设置功率也是小于0的
						value = -int32(f)
					} else {
						value = int32(f)
					}
					err := store.SetPower(value)
					if err != nil {
						g.Log().Errorf(s.ctx, "储能柜:%s 设置功率失败！err:%v", store.GetDeviceConfig().Name, err)
					} else {
						g.Log().Noticef(s.ctx, "储能柜:%s 分配功率：%f", essList[i].Name, f)
					}
					break
				}
			}

		}
	}

	return nil
}

func (s *sStationEnergyStore) SetReactivePower(power int32) error {
	// TODO，设置有功功率， 先临时这样写，后面使用算法设置
	var singlePower = power / int32(len(s.energyStores))
	for _, store := range s.energyStores {
		_ = store.SetReactivePower(singlePower)
	}
	return nil
}

func (s *sStationEnergyStore) SetPowerFactor(factor float32) error {
	return c_error.NonSupportError
}

func (s *sStationEnergyStore) GetTargetPower() int32 {
	return s.targetPower
}

func (s *sStationEnergyStore) GetTargetReactivePower() int32 {
	return s.targetReactivePower
}

func (s *sStationEnergyStore) GetTargetPowerFactor() float32 {
	return s.targetPowerFactor
}

func (s *sStationEnergyStore) GetPower() (float64, error) {
	// 取有功功率和
	var power float64
	for _, ess := range s.energyStores {
		value, err := ess.GetPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 有功功率失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		power += value
	}
	return power, nil
}

func (s *sStationEnergyStore) GetApparentPower() (float64, error) {
	// 取视在功率和
	var apparentPower float64
	for _, ess := range s.energyStores {
		value, err := ess.GetApparentPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 视在功率失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		apparentPower += value
	}
	return apparentPower, nil
}

func (s *sStationEnergyStore) GetReactivePower() (float64, error) {
	// 取无功功率和
	var reactivePower float64
	for _, ess := range s.energyStores {
		value, err := ess.GetReactivePower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 无功功率失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		reactivePower += value
	}
	return reactivePower, nil
}

func (s *sStationEnergyStore) GetRatedPower() uint32 {
	// 额定功率累计
	var ratedPower uint32
	for _, ess := range s.energyStores {
		value := ess.GetRatedPower()
		ratedPower += value
	}
	return ratedPower
}

func (s *sStationEnergyStore) GetMaxInputPower() (float32, error) {
	// 最大输入功率和
	var maxInputPower float32
	for _, ess := range s.energyStores {
		value, err := ess.GetMaxInputPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 最大输入功率失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		maxInputPower += value
	}
	return maxInputPower, nil
}

func (s *sStationEnergyStore) GetMaxOutputPower() (float32, error) {
	// 最大输出功率和
	var maxOutputPower float32
	for _, ess := range s.energyStores {
		value, err := ess.GetMaxOutputPower()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 最大输出功率失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		maxOutputPower += value
	}
	return maxOutputPower, nil
}

func (s *sStationEnergyStore) GetTodayIncomingQuantity() (float64, error) {
	// 今日充电量和
	var todayIncomingQuantity float64
	for _, ess := range s.energyStores {
		value, err := ess.GetTodayIncomingQuantity()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 今日充电量失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		todayIncomingQuantity += value
	}
	//TODO 以后先保存在本地，先判断一下数值大小，如果就用本地缓存的，防止出现数据异常
	return todayIncomingQuantity, nil
}

func (s *sStationEnergyStore) GetHistoryIncomingQuantity() (float64, error) {
	// 历史充电量和
	var historyIncomingQuantity float64
	for _, ess := range s.energyStores {
		value, err := ess.GetHistoryIncomingQuantity()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 历史充电量失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		historyIncomingQuantity += value
	}
	//TODO 以后先保存在本地，先判断一下数值大小，如果就用本地缓存的，防止出现数据异常
	return historyIncomingQuantity, nil
}

func (s *sStationEnergyStore) GetTodayOutgoingQuantity() (float64, error) {
	// 今日放电量和
	var todayOutgoingQuantity float64
	for _, ess := range s.energyStores {
		value, err := ess.GetTodayOutgoingQuantity()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 今日放电量失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		todayOutgoingQuantity += value
	}
	//TODO 以后先保存在本地，先判断一下数值大小，如果就用本地缓存的，防止出现数据异常
	return todayOutgoingQuantity, nil
}

func (s *sStationEnergyStore) GetHistoryOutgoingQuantity() (float64, error) {
	// 历史放电量和
	var historyOutgoingQuantity float64
	for _, ess := range s.energyStores {
		value, err := ess.GetHistoryOutgoingQuantity()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 历史放电量失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		historyOutgoingQuantity += value
	}
	//TODO 以后先保存在本地，先判断一下数值大小，如果就用本地缓存的，防止出现数据异常
	return historyOutgoingQuantity, nil
}

func (s *sStationEnergyStore) GetFireEnvTemperature() (float64, error) {
	// 取温度平均值
	var (
		temperature float64
		count       int
	)
	for _, ess := range s.energyStores {
		value, err := ess.GetFireEnvTemperature()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 环境温度失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		temperature += value
		count++
	}
	return temperature / float64(count), nil
}

func (s *sStationEnergyStore) GetCarbonMonoxideConcentration() (float64, error) {
	// 取最大值
	var carbonMonoxideConcentration float64
	for _, ess := range s.energyStores {
		value, err := ess.GetCarbonMonoxideConcentration()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 一氧化碳浓度失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		if value > carbonMonoxideConcentration {
			carbonMonoxideConcentration = value
		}
	}
	return carbonMonoxideConcentration, nil
}

func (s *sStationEnergyStore) HasSmoke() (bool, error) {
	// 如果有一个有烟雾，就返回true
	for _, ess := range s.energyStores {
		value, err := ess.HasSmoke()
		if err != nil {
			g.Log().Warningf(s.ctx, "获取储能柜:%s 烟雾失败！统计时跳过该柜 err:%v", ess.GetDeviceConfig().Name, err)
			continue
		}
		if value {
			return true, nil
		}
	}
	return false, nil
}

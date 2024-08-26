package collect

import (
	"context"
	"ems-plan/c_base"
	"ems-plan/c_device"
)

type tmpStation struct {
	Ammeters map[c_base.EGroupType][]c_device.IAmmeter

	Pv   []c_device.IPv
	Load []c_device.ILoad
	Ess  []c_device.IEnergyStore

	cabinetEss map[uint8]*tmpCabinet
}

type tmpCabinet struct {
	Ammeter  c_device.IAmmeter
	Pcs      []c_device.IPcs // 一个柜子会有一个多个PCS
	Bms      c_device.IBms
	Fire     c_device.IFire
	Humiture c_device.IHumiture
	Cooling  c_device.ICoolingBasic
}

var _tempInstanceCache = &tmpStation{
	cabinetEss: make(map[uint8]*tmpCabinet),
}

func (t *tmpStation) GetCabinetEss(cabinetId uint8) *tmpCabinet {
	cabinet := t.cabinetEss[cabinetId]
	if cabinet == nil {
		cabinet = &tmpCabinet{}
		t.cabinetEss[cabinetId] = cabinet
		return cabinet
	}
	return cabinet
}

func (t *tmpStation) Init(ctx context.Context) {
	// 先封装cabinet
	//for cabinetId, value := range t.CabinetEss {
	// 先把PCS之类的变成 CabinetPcs
	//master, slaves := getMasterAndList[c_device.IPcs](value.Pcs)
	//pcs := common_cabinet.NewPcs(ctx, cabinetId, master, slaves)
	//bms := common_cabinet.NewBms(ctx, cabinetId, value.Bms)

	/*	var (
			fire     *cabinet.Fire
			cooling  *cabinet.Cooling
			humiture *cabinet.Humiture
		)
		if value.Fire != nil {
			fire = cabinet.NewFire(ctx, cabinetId, value.Fire)
		}
		if value.Cooling != nil {
			cooling = cabinet.NewCooling(ctx, cabinetId, value.Cooling)
		}
		if value.Humiture != nil {
			humiture = cabinet.NewHumiture(ctx, cabinetId, value.Humiture)
		}*/

	//_ess := pylon_checkwatt_v1.CreateEss(ctx, cabinetId, value.Ammeter, pcs, bms, fire, cooling, humiture)
	//
	//err := _ess.Init(ctx, nil, nil)
	//if err != nil {
	//	panic(err)
	//}
	//t.Ess = append(t.Ess, _ess)
	//g.Log().Noticef(ctx, "初始化柜子成功！DeviceId: %v", _ess.GetInfo().Id)
	//}

	// TODO config没初始化！！
	//
	//// 封装station
	//if len(t.Ammeters[config.PvAmmeter]) != 0 || len(t.Pv) != 0 {
	//	ammeterMaster, ammeterSlaves := getMasterAndList[c_device.IIAmmeter](t.Ammeters[config.PvAmmeter])
	//	pvMaster, pvSlaves := getMasterAndList[c_device.IIPv](t.Pv)
	//	station.NewPv(ctx, ammeterMaster, ammeterSlaves, pvMaster, pvSlaves)
	//	g.Log().Noticef(ctx, "初始化场站光伏成功！加载了%d个主设备，%d个从设备", len(t.Ammeters[config.PvAmmeter]), len(t.Pv))
	//}
	//
	//if len(t.Ammeters[config.LoadAmmeter]) != 0 || len(t.Load) != 0 {
	//	ammeterMaster, ammeterSlaves := getMasterAndList[c_device.IIAmmeter](t.Ammeters[config.LoadAmmeter])
	//	loadMaster, loadSlaves := getMasterAndList[c_device.IILoad](t.Load)
	//	station.NewLoad(ctx, ammeterMaster, ammeterSlaves, loadMaster, loadSlaves)
	//	g.Log().Noticef(ctx, "初始化场站负载成功！加载了%d个主设备，%d个从设备", len(t.Ammeters[config.LoadAmmeter]), len(t.Load))
	//}
	//
	//if len(t.Ammeters[config.EssAmmeter]) != 0 || len(t.Ess) != 0 {
	//	ammeterMaster, ammeterSlaves := getMasterAndList[c_device.IIAmmeter](t.Ammeters[config.EssAmmeter])
	//	station.NewEss(ctx, ammeterMaster, ammeterSlaves, t.Ess)
	//	g.Log().Noticef(ctx, "初始化场站储能成功！加载了%d个主设备，%d个从设备", len(t.Ammeters[config.EssAmmeter]), len(t.Ess))
	//}
	//
	//if len(t.Ammeters[config.ChargeAmmeter]) != 0 {
	//	ammeterMaster, ammeterSlaves := getMasterAndList[c_device.IIAmmeter](t.Ammeters[config.ChargeAmmeter])
	//	station.NewCharge(ctx, ammeterMaster, ammeterSlaves)
	//	g.Log().Noticef(ctx, "初始化场站充电成功！加载了%d个主设备，%d个从设备", len(t.Ammeters[config.ChargeAmmeter]), len(t.Ess))
	//}
	//
	//if len(t.Ammeters[config.GeneratorAmmeter]) != 0 {
	//	ammeterMaster, ammeterSlaves := getMasterAndList[c_device.IIAmmeter](t.Ammeters[config.GeneratorAmmeter])
	//	station.NewGenerator(ctx, ammeterMaster, ammeterSlaves)
	//	g.Log().Noticef(ctx, "初始化场站发电机成功！加载了%d个主设备，%d个从设备", len(t.Ammeters[config.GeneratorAmmeter]), len(t.Ess))
	//}

}

//
//func getMasterAndList[T device.IDriver](list []T) (master T, slaves []T) {
//	for _, v := range list {
//		if v.GetInfo().IsMaster {
//			master = v
//		} else {
//			slaves = append(slaves, v)
//		}
//	}
//	return master, slaves
//
//}

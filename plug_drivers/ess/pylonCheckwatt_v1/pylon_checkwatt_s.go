package pylonCheckwatt_v1

import (
	"context"
	"ems-plan/c_device"
)

type PylonCheckwattEss struct {
	//*c_base.SConfigImpl // 配置信息
	ctx    context.Context
	unitId uint8 // modbus转发的id

	ammeter c_device.IAmmeter

	//bms      *common_cabinet.CabinetBms      // 电池
	//pcs      *common_cabinet.CabinetPcs      // 逆变器
	//fire     *common_cabinet.CabinetFire     // 消防
	//cooling  *common_cabinet.CabinetCooling  // 制冷
	//humidity *common_cabinet.CabinetHumidity // 温湿度
}

/*
func CreateEss(ctx context.Context, cabinetId uint8, params map[string]string,

		ammeter c_base.IAmmeter,
		pcs *common_cabinet.CabinetPcs,
		bms *common_cabinet.CabinetBms,
		fire *common_cabinet.CabinetFire,
		cooling *common_cabinet.CabinetCooling,
		humidity *common_cabinet.CabinetHumidity) c_base.IEnergyStore {
		_ess := &PylonCheckwattEss{
			ctx:         ctx,
			unitId:      cabinetId,
			ammeter:     ammeter,
			bms:         bms,
			pcs:         pcs,
			fire:        fire,
			cooling:     cooling,
			humidity:    humidity,
			SConfigImpl: c_base.NewConfig(cabinetId, c_base.EDeviceEnergyStore, c_group.EGroupEnergyStore, false, params),
		}

		if _ess.bms == nil || _ess.pcs == nil {
			panic(fmt.Sprintf("一个柜子需要电池和PCS组成，但是现在有一项不存在，请检查配置！ cabinetId:%d", cabinetId))
		}
		return _ess
	}
*/

package bms_pylon_tech_us108_v1

import (
	"common/c_log"
	"fmt"
)

func (p *sBmsPylonTechUs108) CustomDcOffService() error {
	fmt.Println("CustomDcOffService")
	c_log.BizInfo(p.ctx, "触发直流下电指令")
	return nil
}

func (p *sBmsPylonTechUs108) CustomDcOnService() error {
	fmt.Println("CustomDcOnService")
	c_log.BizInfo(p.ctx, "触发直流上电指令")
	return nil
}

func (p *sBmsPylonTechUs108) CustomSyncTimeService() {
	c_log.BizInfo(p.ctx, "触发时间同步指令")

}

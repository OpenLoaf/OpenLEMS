package bms_pylon_tech_us108_v1

import "fmt"

func (p *sBmsPylonTechUs108) CustomDcOffService() error {
	fmt.Println("CustomDcOffService")

	return nil
}

func (p *sBmsPylonTechUs108) CustomStandbyService() error {
	fmt.Println("CustomStandbyService")

	return nil
}

func (p *sBmsPylonTechUs108) CustomSyncTimeService() {

}

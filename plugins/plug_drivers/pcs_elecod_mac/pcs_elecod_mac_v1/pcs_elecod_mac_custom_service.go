package pcs_elecod_mac_v1

import "fmt"

func (s *sPcsElecodMac) CustomShutdownService() {

	fmt.Printf("停机按钮！")

}

func (s *sPcsElecodMac) CustomStartupService() {

	fmt.Printf("开机按钮！")

}

func (s *sPcsElecodMac) CustomStandbyService() {

	fmt.Printf("待机按钮！")

}

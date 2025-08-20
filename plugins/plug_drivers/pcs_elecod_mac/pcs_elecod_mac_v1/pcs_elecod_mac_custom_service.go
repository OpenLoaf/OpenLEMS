package pcs_elecod_mac_v1

import "fmt"

func (s *sPcsElecodMacConfig) CustomShutdownService() {

	fmt.Printf("停机按钮！")

}

func (s *sPcsElecodMacConfig) CustomStartupService() {

	fmt.Printf("开机按钮！")

}

func (s *sPcsElecodMacConfig) CustomStandbyService() {

	fmt.Printf("待机按钮！")

}

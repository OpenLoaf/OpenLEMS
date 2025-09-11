package pcs_elecod_mac_v1

import (
	"common/c_proto"
	"fmt"
	"pcs_elecod/pcs_elecod_mac_v1/elecod_mac_defined"
)

func (s *sPcsElecodMac) CustomShutdownService() error {
	fmt.Printf("停机按钮！被触发")
	return s.ExecuteProtocolMethod(func(protocol c_proto.ICanbusProtocol) error {
		return protocol.SendMessage(elecod_mac_defined.CmdShutdown, nil)
	})
}

func (s *sPcsElecodMac) CustomStartupService() error {
	fmt.Printf("开机按钮！被触发")
	return s.ExecuteProtocolMethod(func(protocol c_proto.ICanbusProtocol) error {
		return protocol.SendMessage(elecod_mac_defined.CmdStart, nil)
	})
}

func (s *sPcsElecodMac) CustomStandbyService() error {
	fmt.Printf("待机按钮！被触发")
	return s.ExecuteProtocolMethod(func(protocol c_proto.ICanbusProtocol) error {
		return protocol.SendMessage(elecod_mac_defined.CmdStandby, nil)
	})
}

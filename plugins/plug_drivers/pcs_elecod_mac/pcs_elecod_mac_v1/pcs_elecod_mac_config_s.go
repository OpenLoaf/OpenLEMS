package pcs_elecod_mac_v1

type sPcsElecodMacConfig struct {
	SelfAddress *uint8 `json:"selfAddress" dc:"本机地址"`
	MacAddress  *uint8 `json:"macAddress" dc:"mac地址"`
}

package internal

type sInfluxdb2Config struct {
	Token string `json:"token,omitempty" dc:"密钥"`
	Org   string `json:"org,omitempty" dc:"组织"`
}

package c_base

type SMqttConfig struct {
	Enable   bool              `json:"enable,omitempty" orm:"enable"`     // 是否启用
	Url      string            `json:"url,omitempty" orm:"url"`           // 地址
	Username string            `json:"username,omitempty" orm:"username"` // 用户名
	Password string            `json:"password,omitempty" orm:"password"` // 密码
	Timeout  int               `json:"timeout,omitempty" orm:"timeout"`   // 超时时间
	Params   map[string]string `json:"params,omitempty" orm:"params"`
}

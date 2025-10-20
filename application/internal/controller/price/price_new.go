package price

import (
	"application/api/price"

	"github.com/gogf/gf/v2/net/ghttp"
)

type Controller struct {
	Request  *ghttp.Request  // 请求数据对象
	Response *ghttp.Response // 返回数据对象
	Server   *ghttp.Server   // WebServer对象
	Cookie   *ghttp.Cookie   // COOKIE操作对象
	Session  *ghttp.Session  // SESSION操作对象
	View     *View           // 视图对象
}

type View struct{}

// NewV1 创建电价控制器
func NewV1() price.IPriceV1 {
	return &Controller{}
}

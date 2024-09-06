package ws

import (
	"application/internal/service"
	"context"
	"ems-plan/c_base"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtimer"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

type StationWebsocket struct {
	upGrader *websocket.Upgrader
}

type StationRes struct {
	Code    int            `json:"code,omitempty" dc:"状态码"`
	Message string         `json:"message,omitempty" dc:"消息"`
	Time    string         `json:"time,omitempty"  dc:"时间"`
	Data    map[string]any `json:"data,omitempty" dc:"数据"`
}

func NewStationWebsocket() *StationWebsocket {
	return &StationWebsocket{
		upGrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // 默认允许所有请求升级，如果有安全需求可以进行更严格的检查
			},
		},
	}
}

func (w *StationWebsocket) StationInfoWebsocket(r *ghttp.Request) {
	ctx, cancelFunc := context.WithCancel(r.Context())

	conn, err := w.upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		cancelFunc()
		r.Response.WriteHeader(http.StatusInternalServerError)
		return
	}

	conn.SetCloseHandler(func(code int, text string) error {
		g.Log().Noticef(ctx, "Connection closed: %v, %s", code, text)
		return nil
	})

	defer func(conn *websocket.Conn) {
		g.Log().Debugf(ctx, "关闭Station Websocket连接")
		cancelFunc()
		_ = conn.Close()
	}(conn)
	// 上面都一样

	w.sendData(conn)
	// 每秒发送数据
	gtimer.SetInterval(ctx, time.Second, func(ctx context.Context) {
		w.sendData(conn)
	})

	for {
		message, _, err := conn.ReadMessage()
		if err != nil {
			return
		}
		if message == websocket.CloseMessage {
			return
		}
	}

}

func (w *StationWebsocket) sendData(conn *websocket.Conn) {
	_ = conn.WriteJSON(&RegisterTelemetryQueryRes{
		Code:    200,
		Message: "",
		Time:    time.Now().Format("2006-01-02 15:04:05"),
		Data: map[string]any{
			string(c_base.EStationEnergyStore): service.Station().GetEnergyStoreStatus(),
		},
	})
}

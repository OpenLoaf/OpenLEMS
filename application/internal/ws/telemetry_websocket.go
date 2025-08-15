package ws

import (
	"common"
	"common/c_util"
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"math"
	"math/big"

	"net/http"
	"strings"
	"time"
)

type TelemetryWebsocket struct {
	upGrader *websocket.Upgrader
}

type RegisterTelemetryQuery struct {
	Millisecond int      `json:"millisecond,omitempty" v:"required|between:1000,10000#请输入时间间隔|时间间隔范围为:min到:max" dc:"时间间隔"`
	Keys        []string `json:"keys,omitempty" v:"required|length:1,6#请输入遥测列表|遥测列表长度为:min到:max位" dc:"遥测列表（group:deviceKey:telemetryKey）"`
}

type RegisterTelemetryQueryRes struct {
	Code    int    `json:"code,omitempty" dc:"状态码"`
	Message string `json:"message,omitempty" dc:"消息"`
	Time    string `json:"time,omitempty"  dc:"时间"`
	//DeviceId  string    `json:"key,omitempty" dc:"group:deviceKey:telemetryKey"`
	Data map[string]any `json:"data,omitempty" dc:"数据"`
}

func NewTelemetryWebsocket() *TelemetryWebsocket {
	return &TelemetryWebsocket{
		upGrader: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true // 默认允许所有请求升级，如果有安全需求可以进行更严格的检查
			},
		},
	}
}

func (w *TelemetryWebsocket) TelemetryWebsocket(r *ghttp.Request) {
	ctx, cancelFunc := context.WithCancel(r.Context())

	conn, err := w.upGrader.Upgrade(r.Response.Writer, r.Request, nil)
	if err != nil {
		cancelFunc()
		r.Response.WriteHeader(http.StatusInternalServerError)
		return
	}

	var (
		cancel context.CancelFunc
	)

	conn.SetCloseHandler(func(code int, text string) error {
		cancelFunc()
		_ = conn.Close()
		return nil
	})

	g.Log().Debugf(ctx, "遥测连接成功")

	for {

		select {
		case <-ctx.Done():
			if cancel != nil {
				cancel()
			}
			g.Log().Debugf(ctx, "通道关闭！")
			return
		default:

			var query RegisterTelemetryQuery
			err := conn.ReadJSON(&query)
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					g.Log().Noticef(ctx, "Unexpected close error: %v", err)
				} else {
					_ = conn.WriteJSON(&RegisterTelemetryQueryRes{
						Code:    500,
						Message: err.Error(),
						Time:    c_util.GetNow(),
					})
					continue
				}
				if cancel != nil {
					cancel()
				}
				return
			}

			g.Log().Debugf(ctx, "-------recv: %+v", query)
			if query.Millisecond == 0 {
				if cancel != nil {
					cancel()
				}
				continue
			}

			// 验证一下数据是否正确
			if query.Millisecond < 100 || query.Millisecond > 1000*60*60*24 {
				_ = conn.WriteJSON(&RegisterTelemetryQueryRes{
					Code:    500,
					Message: "时间间隔范围为200到86400000",
					Time:    c_util.GetNow(),
				})
				continue
			}

			if query.Keys == nil || len(query.Keys) == 0 {
				_ = conn.WriteJSON(&RegisterTelemetryQueryRes{
					Code:    500,
					Message: "遥测列表不能为空",
					Time:    c_util.GetNow(),
				})
				continue
			}

			if cancel != nil {
				// 取消之前的定时器，根据新的规则重新生成
				cancel()
				cancel = nil
			}

			var newCtx context.Context
			newCtx, cancel = context.WithCancel(ctx)

			go func(ctx context.Context, conn *websocket.Conn, query RegisterTelemetryQuery) {
				// 先执行了，再等待定时器的
				_ = writeValue(ctx, conn, query)

				ticker := time.NewTicker(time.Duration(query.Millisecond) * time.Millisecond)
				defer ticker.Stop()
				for {
					select {
					case <-ctx.Done():
						g.Log().Debugf(ctx, "遥测取消发送数据的Goroutine")
						return
					case <-ticker.C:
						_ = writeValue(ctx, conn, query)
					}
				}
			}(newCtx, conn, query)
		}
	}

}

func writeValue(ctx context.Context, conn *websocket.Conn, query RegisterTelemetryQuery) error {
	var dataMap = make(map[string]any)
	var errMap = make(map[string]string)
	for _, key := range query.Keys {
		// 解析key group:deviceKey:telemetryKey
		values := strings.SplitN(key, ":", 2)
		instance := common.GetRunningDeviceById(values[0])
		if instance == nil {
			errMap[key] = fmt.Sprintf("设备 %s 不存在", values[0])
			continue
		}

		value, err := instance.GetTelemetry(values[1], instance)
		if err != nil {
			errMap[key] = err.Error()
			continue
		}

		if value != "" && value != nil {
			if ft, ok := value.(float32); ok {
				if !math.IsNaN(float64(ft)) {
					value = big.NewFloat(float64(ft)).Text('f', 2)
					dataMap[key] = value
				}
			} else if ft, ok := value.(float64); ok {
				if !math.IsNaN(ft) {
					value = big.NewFloat(ft).Text('f', 4)
					dataMap[key] = value
				} else {
					dataMap[key] = nil
				}
			} else {
				dataMap[key] = gconv.String(value)
			}

		}
	}
	// 写入错误信息
	if len(errMap) > 0 {
		var errorMessage = ""
		for _, v := range errMap {
			errorMessage += v
		}
		_ = conn.WriteJSON(&RegisterTelemetryQueryRes{
			Code:    500,
			Message: errorMessage,
			Time:    c_util.GetNow(),
		})
	}

	// 写入数据
	_ = conn.WriteJSON(&RegisterTelemetryQueryRes{
		Code: 200,
		Time: c_util.GetNow(),
		Data: dataMap,
	})

	return nil
}

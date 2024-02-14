package servers

import (
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	jwtv4 "github.com/golang-jwt/jwt/v4"
	"github.com/hertz-contrib/websocket"

	mvjwt "qnc/biz/mw/jwt"
)

const (
	// 最大的消息大小
	maxMessageSize = 8192
)

type Controller struct {
}

// type renderData struct {
// 	ClientId string `json:"clientId"`
// }

var upgrader = websocket.HertzUpgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(ctx *app.RequestContext) bool {
		return true
	},
}

func StartServeWs(ctx *app.RequestContext) {
	hlog.Debug("start connect")
	err := upgrader.Upgrade(ctx, func(conn *websocket.Conn) {
		conn.SetReadLimit(maxMessageSize)

		_, message, err := conn.ReadMessage()
		if err != nil {
			hlog.Debug("read:", err)
			return
		}

		token, err := mvjwt.JwtMiddleware.ParseTokenString(string(message))
		if err != nil {
			hlog.Debug("parse token err:", err)
			conn.WriteJSON(RetData{
				Code:       0,
				MessageId:  "100001",
				SendUserId: "101",
				MsgType:    "nologin",
				Msg:        "token invalid",
				Data:       nil,
			})
			return
		}

		clientId := ""
		for key, value := range token.Claims.(jwtv4.MapClaims) {
			if key == "user_id" {
				uid := value.(float64)
				clientId = strconv.FormatFloat(uid, 'f', 0, 64)
			}
		}
		if clientId == "" {
			hlog.Error("no login")
			conn.WriteJSON(RetData{
				Code:       0,
				MessageId:  "100001",
				SendUserId: "101",
				MsgType:    "nologin",
				Msg:        "token invalid",
				Data:       nil,
			})
			return
		}

		clientSocket := NewClient(clientId, "1001", conn)

		Manager.Connect <- clientSocket
		Manager.AddClient(clientSocket)

		go WriteMessage()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				hlog.Debug("read:", err)
				break
			}
			hlog.Debug("recv: %s, type: %s", message, mt)
		}

	})
	if err != nil {
		hlog.Errorf("upgrade error: %v", err)
		return
	}
}

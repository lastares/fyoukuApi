package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
	"strings"
	"time"
)

type WsController struct {
	CommonController
}

type OnlineList struct {
	register chan *http.Client
}

type ArticleReadData struct {
	logicType int
	Username string `json:"username"`
	ArticleId int `json:"articleId"`
	//CurrentTime *time.Time `json:"currentTime"`
}

type ReplyJson struct {
	Msg string `json:"msg"`
}

var (
	upgrader2 = websocket.Upgrader{
		ReadBufferSize:  4096,//指定读缓存区大小
		WriteBufferSize: 1024,// 指定写缓存区大小
		// 检测请求来源
		CheckOrigin: func(r *http.Request) bool {
			requestHost := beego.AppConfig.String("testWsHost")
			if r.Header.Get("Origin") != requestHost {
				return false
			}

			// ws client method check
			requestMethod := strings.ToUpper(r.Method)
			if requestMethod != "GET" {
				return false
			}
			return true
		},
	}
)

// @router /test/ws [get]
func (this *WsController) TestWs() {
	// 收到 http 请求后 升级 协议
	connection, err := upgrader2.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		goto WS_ERROR
		this.Response(Failed(-1, "WS连接失败"))
	}

	// 心跳检测
	go func() {
		for {
			if err = connection.WriteJSON(ReplyJson{"心跳检测"}); err != nil {
				fmt.Println("心跳检测error: ", err.Error())
				return
			}
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		_, data, err := connection.ReadMessage()
		if err != nil {
			goto WS_ERROR
			this.Response(Failed(-1, "WS读取客户端消息失败"))
		}
		var articleReadData ArticleReadData
		err = json.Unmarshal([]byte(data), &articleReadData)
		if err != nil {
			connection.WriteJSON(Failed(-1, err.Error()))
		}

		err = connection.WriteJSON(ReplyJson{"消息已收到，用户名：" + articleReadData.Username})
		if err != nil {
			goto WS_ERROR
			this.Response(Failed(-1, "WS读取客户端消息失败"))
		}
	}
WS_ERROR:
	connection.Close()
}


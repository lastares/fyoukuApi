package controllers

import (
	"encoding/json"
	"fyoukuApi/models"
	"github.com/gorilla/websocket"
	"net/http"
)

type BarrageController struct {
	CommonController
}

type WsData struct {
	CurrentTime int
	EpisodesId int
}

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

// @router /test/barrages [get]
func (this *BarrageController) TestBarrages()  {
	episodesId, _ := this.GetInt("episodesId", 0)
	currentTime, _ := this.GetInt("currentTime", 0)
	endTime := currentTime + 60
	barrages, _, _ := models.GetBarrageList(episodesId, currentTime, endTime)
	this.Response(ResponseNoMeta(0, barrages))
}

// 获取弹幕websocket
// @router /barrage/ws [get]
func (this *BarrageController) BarrageWs() {
	var (
		conn *websocket.Conn
		err  error
		data []byte
	)

	conn, err = upgrader.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil)
	if err != nil {
		goto ERR
	}

	for {
		_, data, err = conn.ReadMessage()
		if err != nil {
			goto ERR
		}
		var wsData WsData
		json.Unmarshal([]byte(data), &wsData)
		endTime := wsData.CurrentTime + 60
		barrages, _, err := models.GetBarrageList(wsData.EpisodesId, wsData.CurrentTime, endTime)
		if err != nil {
			goto ERR
		}
		if err = conn.WriteJSON(barrages); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()
}

// 保存弹幕
// @router /barrage/save [post]
func (this *BarrageController) BarrageSave()  {
	episodesId, _ := this.GetInt("episodesId", 0)
	videoId, _ := this.GetInt("videoId", 0)
	uid, _ := this.GetInt("uid", 0)
	content := this.GetString("content", "")
	currentTime, _ := this.GetInt("currentTime", 1)
	if uid <= 0 {
		this.Response(Failed(-1, "亲，您还未登录~"))
	}

	if episodesId <= 0 || videoId <= 0 || content == "" {
		this.Response(Failed(-1, "params error"))
	}


	errorMsg := models.CreateBarrage(content, currentTime, uid, episodesId, videoId)
	if errorMsg != nil {
		this.Response(Failed(-1, "发布弹幕失败"))
	}
	this.Response(Ok(0))
}

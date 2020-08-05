package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Barrage struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	AddTime     int64  `json:"addTime"`
	UserId      int    `json:"userId"`
	Status      int    `json:"status"`
	CurrentTime int    `json:"currentTime"`
	EpisodesId  int    `json:"episodesId"`
	VideoId     int    `json:"videoId"`
}

type BarrageData struct {
	Id          int    `json:"id"`
	Content     string `json:"content"`
	CurrentTime int    `json:"currentTime"`
}

func init() {
	orm.RegisterModel(new(Barrage))
}

func GetBarrageList(episodesId, startTime, endTime int) (barrages []Barrage, num int64, err error) {
	num, err = orm.NewOrm().QueryTable("barrage").
		Filter("episodes_id", episodesId).
		Filter("status", 1).
		Filter("current_time__gte", startTime).
		Filter("current_time__lt", endTime).
		OrderBy("current_time").
		All(&barrages, "Id", "Content", "CurrentTime")
	return
}

func CreateBarrage(
	content string, currentTime, uid, episodesId, videoId int,
) (errorMsg error) {
	barrage := Barrage{
		Content: content,
		CurrentTime: currentTime,
		UserId: uid,
		EpisodesId: episodesId,
		VideoId: videoId,
		Status: 1,
		AddTime: time.Now().Unix(),
	}
	_, errorMsg = orm.NewOrm().Insert(&barrage)
	return
}

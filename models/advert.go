package models

import "github.com/astaxie/beego/orm"

type Advert struct {
	Id        int
	Title     string
	ChannelId int
	SubTitle  string
	AddTime   int64
	Img       string
	Url       string
}

func init() {
	orm.RegisterModel(new(Advert))
}

func GetChanelAdvert(channelId int) (advert []Advert) {
	newOrm := orm.NewOrm()
	newOrm.QueryTable("advert").
		Filter("channel_id", channelId).
		OrderBy("-id").
		One(&advert)

	return
}

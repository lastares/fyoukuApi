package models

import (
	"github.com/astaxie/beego/orm"
)

type ChannelRegion struct {
	Id   int
	Name string
	ChannelId int
	AddTime int
	Status int
	Sort int
}

func init() {
	orm.RegisterModel(new(ChannelRegion))
}

func GetChannelRegion(channelId int) (regions []ChannelRegion)  {
	orm.NewOrm().QueryTable("channel_region").
		Filter("channel_id", channelId).
		Filter("status", 1).
		OrderBy("-add_time").All(&regions, "id", "name")

	return
}

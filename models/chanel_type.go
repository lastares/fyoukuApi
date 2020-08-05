package models

import (
	"github.com/astaxie/beego/orm"
)

type ChannelType struct {
	Id   int
	Name string
	ChannelId int
	AddTime int
	Status int
}

func init() {
	orm.RegisterModel(new(ChannelType))
}

func GetChannelType(channelId int)  (channelTypes []ChannelType){
	orm.NewOrm().
		QueryTable("channel_type").
		Filter("channelId", channelId).
		Filter("status", 1).
		OrderBy("-addTime").
		All(&channelTypes, "Id", "Name")
	return
}

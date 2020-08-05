package controllers

import (
	"fyoukuApi/models"
)

type TopController struct {
	CommonController
}

// 根据频道ID获取排行榜
// @router /channel/top [get]
func (this *TopController) ChannelTop() {
	channelId, _ := this.GetInt("channelId", 0)
	if channelId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	data, err := models.GetChannelTop(channelId)
	if err != nil {
		this.Response(Failed(-1, "获取排行榜失败"))
	}
	this.Response(ResponseNoMeta(0, data))
}

// 根据类型ID获取排行榜
// @router /type/top [get]
func (this *TopController) TypeTop()  {
	typeId, _ := this.GetInt("typeId", 0)
	if typeId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	data, err := models.GetTypeTop(typeId)
	if err != nil {
		this.Response(Failed(-1, "获取排行榜失败"))
	}

	this.Response(ResponseNoMeta(0, data))
}

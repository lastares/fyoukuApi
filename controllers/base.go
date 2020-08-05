package controllers

import "fyoukuApi/models"

type BaseController struct {
	CommonController
}

// 获取频道地区列表
// @router /channel/region [get]
func (this * BaseController) ChannelRegion()  {
	channelId, _ := this.GetInt("channelId", 0)

	if channelId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	data := models.GetChannelRegion(channelId)
	this.Response(ResponseNoMeta(0, data))
}

// 获取频道下的类型
// @router /channel/type [get]
func (this *BaseController)  ChannelType() {
	channelId, _ := this.GetInt("channelId", 0)
	if channelId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	data := models.GetChannelType(channelId)
	this.Response(ResponseNoMeta(0, data))
}
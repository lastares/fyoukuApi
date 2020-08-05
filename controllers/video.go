package controllers

import (
	"fyoukuApi/models"
	"github.com/astaxie/beego/orm"
)

type VideoController struct {
	CommonController
}

//获取首页公告图片
// @router /channel/advert [get]
func (this *VideoController) ChannelAdvert() {
	channelId, _ := this.GetInt("channelId", 0)
	if channelId == 0 {
		this.Response(Failed(0, "参数错误"))
	}

	video := models.GetChanelAdvert(channelId)
	this.Response(ResponseNoMeta(0, video))
}

// @router /channel/hot [get]
func (this *VideoController) ChannelHotList() {
	channelId, _ := this.GetInt("channelId", 0)
	if channelId == 0 {
		this.Response(Failed(-1, "params error"))
	}

	data, total := models.GetChannelHotList(channelId)
	this.Response(ResponseData(0, data, Meta{Total: total}))
}

// @router /channel/recommend/region [get]
func (this *VideoController) ChannelRegionCommendVideo() {
	channelId, _ := this.GetInt("channelId", 0)
	regionId, _ := this.GetInt("regionId", 0)

	if channelId == 0 || regionId == 0 {
		this.Response(Failed(0, "Params error"))
	}

	data := models.GetChanelRegionVideo(channelId, regionId)
	this.Response(ResponseNoMeta(0, data))
}

// @router /channel/recommend/type [get]
func (this *VideoController) ChannelTypeRecommendList() {
	channelId, _ := this.GetInt("channelId", 0)
	typeId, _ := this.GetInt("typeId", 0)
	if channelId == 0 || typeId == 0 {
		this.Response(Failed(0, "Params error"))
	}

	data := models.ChannelTypeRecommendList(channelId, typeId)
	this.Response(ResponseNoMeta(0, data))
}

// 获取频道下的视频
// @router /channel/video [get]
func (this *VideoController) ChannelVideo() {
	channelId, _ := this.GetInt("channelId", 0)
	regionId, _ := this.GetInt("regionId", 0)
	typeId, _ := this.GetInt("typeId", 0)
	end := this.GetString("end", "")
	sort := this.GetString("sort", "")

	limit, _ := this.GetInt("limit", 12)
	pageSize, _ := this.GetInt("offset", 0)

	if channelId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	if limit == 0 {
		limit = 12
	}

	data, num := models.GetChannelVideoList(
		channelId,
		regionId,
		typeId,
		limit,
		pageSize,
		sort,
		end,
	)

	this.Response(ResponseData(0, data, Meta{num}))
}

// 获取视频详情
// @router /video/info [get]
func (this *VideoController) VideoInfo() {
	videoId, _ := this.GetInt("videoId", 0)
	if videoId <= 0 {
		this.Response(Failed(0, "params error"))
	}

	data, err := models.RedisGetVideoInfo(videoId)
	//data, err := models.GetVideoInfo(videoId)
	if err == orm.ErrNoRows {
		this.Response(Failed(-1, "视频不存在"))
	}
	this.Response(ResponseNoMeta(0, data))
}

// 获取视频剧集列表
// @router /video/episodes/list [get]
func (this *VideoController) VideoEpisodesList() {
	videoId, _ := this.GetInt("videoId", 0)
	if videoId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	data := models.GetVideoEpisodesList(videoId)
	this.Response(ResponseNoMeta(0, data))
}

// 我的视频接口
// @router /user/video [get]
func (this *VideoController) UserVideo() {
	uid, _ := this.GetInt("uid", 0)
	if uid <= 0 {
		this.Response(Failed(-1, "必须指定用户"))
	}

	videos, err := models.GetUserVideo(uid)
	if err != nil {
		this.Response(Failed(-1, "服务器错误，请稍后重试"))
	}

	this.Response(ResponseNoMeta(0, videos))
}

// 保存视频
// @router /video/save [*]
func (this *VideoController) VideoSave() {
	playUrl := this.GetString("playUrl")
	title := this.GetString("title")
	subTitle := this.GetString("subTitle")
	channelId, _ := this.GetInt("channelId")
	typeId, _ := this.GetInt("typeId")
	regionId, _ := this.GetInt("regionId")
	uid, _ := this.GetInt("uid")

	if uid == 0 {
		this.Response(Failed(-1, "请先登陆"))
	}

	if playUrl == "" {
		this.Response(Failed(-1, "视频地址不能为空"))
	}

	err := models.VideoSave(
		title,
		subTitle,
		playUrl,
		channelId,
		regionId,
		typeId,
		uid,
	)

	if err != nil {
		this.Response(Failed(-1, err.Error()))
	}

	this.Response(Ok(0))
}

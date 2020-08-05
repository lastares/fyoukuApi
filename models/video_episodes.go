package models

import (
	"github.com/astaxie/beego/orm"
)

type VideoEpisodes struct {
	Id            int
	Title         string
	AddTime       int64
	Num           int
	VideoId       int
	PlayUrl       string
	Status        int
	Comment       int
	AliyunVideoId string
}

func init() {
	orm.RegisterModel(new(VideoEpisodes))
}

func GetVideoEpisodesList(videoId int) (videoEpisodes []VideoEpisodes) {
	orm.NewOrm().
		QueryTable("video_episodes").
		Filter("video_id", videoId).
		OrderBy("num").
		All(&videoEpisodes)

	return
}

func GetVideoEpisodes(episodesId int) (VideoEpisodes, error) {
	videoEpisodes := VideoEpisodes{Id: episodesId}
	err := orm.NewOrm().Read(&videoEpisodes)
	return videoEpisodes, err
}

// 修改视频剧集的评论数
func IncrementVideoEpisodesComment(videoEpisodes VideoEpisodes) (err error) {
	videoEpisodes.Comment = videoEpisodes.Comment + 1
	_, err = orm.NewOrm().Update(&videoEpisodes)
	return
}

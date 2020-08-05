package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Comment struct {
	Id          int
	Content     string
	AddTime     int64
	Status      int
	UserId      int
	Stamp       int
	PraiseCount int
	EpisodesId  int
	VideoId     int
}

func init() {
	orm.RegisterModel(new(Comment))
}

func GetCommentList(episodesId, limit, offset int) (comments []Comment, total int64) {
	total, _ = orm.NewOrm().
		QueryTable("comment").
		Filter("episodes_id", episodesId).
		Filter("status", 1).
		OrderBy("-add_time").
		Limit(limit, offset).
		All(&comments)
	return
}

// 保存评论
func CommentCreate(
	episodesId, userId, videoId int,
	content string,
) (errors []error) {
	newOrm := orm.NewOrm()
	newOrm.Begin()

	// 获取视频信息
	video, err := GetVideoInfo(videoId)
	if err == orm.ErrNoRows {
		errors = append(errors, err)
		return
	}

	// 获取视频剧集信息
	videoEpisodes, err := GetVideoEpisodes(episodesId)
	if err == orm.ErrNoRows {
		errors = append(errors, err)
		return
	}

	comment := Comment{
		Content:    content,
		EpisodesId: episodesId,
		UserId:     userId,
		VideoId:    videoId,
		Status:     1,
		Stamp:      0,
		AddTime:    time.Now().Unix(),
	}
	// 保存评论
	_, err = newOrm.Insert(&comment)
	if err != nil {
		errors = append(errors, err)
		newOrm.Rollback()
		return
	}

	// 修改视频的总评论数
	err = IncrementVideoComment(video)
	if err != nil {
		errors = append(errors, err)
		newOrm.Rollback()
		return
	}

	// 修改视频剧集的评论数
	err = IncrementVideoEpisodesComment(videoEpisodes)
	if err != nil {
		errors = append(errors, err)
		newOrm.Rollback()
		return
	}
	newOrm.Commit()
	return
}

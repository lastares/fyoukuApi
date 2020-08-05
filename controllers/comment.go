package controllers

import (
	"fyoukuApi/models"
	"fyoukuApi/util"
)

type CommentController struct {
	CommonController
}

type CommentInfo struct {
	Id           int             `json:"id"`
	Content      string          `json:"content"`
	AddTime      int64           `json:"addTime"`
	AddTimeTitle string          `json:"addTimeTitle"`
	UserId       int             `json:"userId"`
	Stamp        int             `json:"stamp"`
	PraiseCount  int             `json:"praiseCount"`
	UserInfo     models.UserInfo `json:"userinfo"`
}

// 视频评论列表
// @router /comment/list [get]
func (this *CommonController) CommentList() {
	episodesId, _ := this.GetInt("episodesId", 0)
	limit, _ := this.GetInt("limit", 12)

	if limit == 0 {
		limit = 12
	}
	offset, _ := this.GetInt("offset", 0)
	if episodesId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	comments, total := models.GetCommentList(episodesId, limit, offset)

	var data []CommentInfo
	for _, comment := range comments {
		user := models.GetUserInfo(comment.UserId)
		commentInfo := CommentInfo{
			Id:           comment.Id,
			Content:      comment.Content,
			AddTime:      comment.AddTime,
			AddTimeTitle: util.DateFormat(comment.AddTime),
			UserId:       comment.UserId,
			Stamp:        comment.Stamp,
			PraiseCount:  comment.PraiseCount,
			UserInfo: models.UserInfo{
				Id:      user.Id,
				Name:    user.Name,
				AddTime: user.AddTime,
				Status:  user.Status,
				Avatar:  user.Avatar,
			},
		}
		data = append(data, commentInfo)
	}
	this.Response(ResponseData(0, data, Meta{total}))
}

// 保存评论
// @router /comment/save [post]
func (this *CommonController) CommentSave() {
	//获取视频剧集ID
	episodesId, _ := this.GetInt("episodesId", 0)
	//获取视频ID
	videoId, _ := this.GetInt("videoId", 0)
	//获取用户ID
	uid, _ := this.GetInt("uid", 0)
	//获取内容
	content := this.GetString("content", "")

	if content == "" || uid <= 0 || episodesId <= 0 || videoId <= 0 {
		this.Response(Failed(-1, "params error"))
	}

	// 保存评论
	errors := models.CommentCreate(
		episodesId,
		uid,
		videoId,
		content,
	)

	if len(errors) > 0 {
		this.Response(Failed(-1, "评论失败，请稍后重试"))
	}
	this.Response(Ok(0))
}

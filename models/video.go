package models

import (
	"fmt"
	"fyoukuApi/services"
	"github.com/astaxie/beego/orm"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"time"
)

type Video struct {
	Id                 int
	Title              string
	SubTitle           string
	AddTime            int64
	Status             int
	Img                string
	Img1               string
	ChannelId          int
	TypeId             int
	IsHot              int
	EpisodesUpdateTime int64
	RegionId           int
	IsRecommend        int
	IsEnd              int
	Comment            int
	EpisodesCount      int
	UserId             int
}

func init() {
	orm.RegisterModel(new(Video))
}

func GetChannelHotList(channelId int) (videoList []Video, num int64) {
	newOrm := orm.NewOrm()
	qs := newOrm.QueryTable("video")
	num, _ = qs.
		Filter("channel_id", channelId).
		Filter("status", 1).
		Filter("is_hot", 1).
		OrderBy("-id").
		All(&videoList)
	return
}

func GetChanelRegionVideo(channelId, regionId int) (videos []Video) {
	newOrm := orm.NewOrm()
	newOrm.QueryTable("video").
		Filter("channel_id", channelId).
		Filter("region_id", regionId).
		Filter("is_recommend", 1).
		Filter("status", 1).
		OrderBy("-episodes_update_time").
		Limit(9).
		All(&videos)
	return
}

func ChannelTypeRecommendList(channelId, typeId int) (videos []Video) {
	orm.NewOrm().QueryTable("video").
		Filter("channel_id", channelId).
		Filter("type_id", typeId).
		Filter("is_recommend", 1).
		Filter("status", 1).
		OrderBy("-episodes_update_time").
		Limit(9).
		All(&videos)
	return
}

// 获取渠道下的视频
func GetChannelVideoList(
	channelId, regionId, typeId, limit, pageSize int,
	sort, end string,
) (videos []Video, num int64) {
	qs := orm.NewOrm().
		QueryTable("video").
		Filter("channel_id", channelId).
		Filter("status", 1)

	if regionId > 0 {
		qs = qs.Filter("region_id", regionId)
	}

	if typeId > 0 {
		qs = qs.Filter("type_id", typeId)
	}

	if end == "n" {
		qs = qs.Filter("is_end", 0)
	} else if end == "y" {
		qs = qs.Filter("is_end", 1)
	}

	if sort == "episodesUpdateTime" {
		qs = qs.OrderBy("-episodes_update_time")
	} else if sort == "comment" {
		qs = qs.OrderBy("-comment")
	} else if sort == "addTime" {
		qs = qs.OrderBy("-add_time")
	} else {
		qs = qs.OrderBy("-add_time")
	}

	num, _ = qs.Limit(limit, pageSize).All(&videos)

	return
}

func GetVideoInfo(videoId int) (Video, error) {
	video := Video{Id: videoId}
	err := orm.NewOrm().Read(&video)
	return video, err
}

// 增加Redis缓存 - 获取视频详情
func RedisGetVideoInfo(videoId int) (Video, error) {
	var video Video
	conn := redisClient.PoolConnect()
	defer conn.Close()
	redisKey := "video:id:" + strconv.Itoa(videoId)
	// 判断Redis key是否存在
	isExist, err := redis.Bool(conn.Do("exists", redisKey))
	if isExist {
		fmt.Println("redis中取出的数据")
		// 取出video数据，返回
		result, _ := redis.Values(conn.Do("hgetall", redisKey))
		err = redis.ScanStruct(result, &video)
	} else {
		fmt.Println("从数据库中取出再放入Redis")
		video, err = GetVideoInfo(videoId)
		if err == nil {
			// 将视频信息写入到Redis中
			_, err = conn.Do("hmset", redis.Args{redisKey}.AddFlat(video)...)
			if err == nil {
				conn.Do("expire", redisKey, 86400)
			}
		}
	}
	return video, err
}

// 修改视频总评论数
func IncrementVideoComment(video Video) (err error) {
	video.Comment = video.Comment + 1
	_, err = orm.NewOrm().Update(&video)
	return
}

// 根据频道获取排行榜
func GetChannelTop(channelId int) (videos []Video, err error) {
	_, err = orm.NewOrm().
		QueryTable("video").
		Filter("channel_id", channelId).
		Filter("status", 1).
		OrderBy("-comment").
		Limit(10).
		All(&videos)

	return
}

// 获取类型下排行榜
func GetTypeTop(typeId int) (videos []Video, err error) {
	_, err = orm.NewOrm().
		QueryTable("video").
		Filter("type_id", typeId).
		Filter("status", 1).
		OrderBy("-comment").
		Limit(10).
		All(&videos)

	return
}

// 我的视频接口
func GetUserVideo(uid int) (videos []Video, err error) {
	_, err = orm.NewOrm().QueryTable("video").Filter("user_id", uid).All(&videos)
	return
}

// 保存我的视频
func VideoSave(
	title, subTitle, playUrl string,
	channelId, regionId, typeId, userId int,
) error {
	nowTime := time.Now().Unix()
	video := Video{
		Title:              title,
		SubTitle:           subTitle,
		AddTime:            nowTime,
		Img:                "",
		Img1:               "",
		EpisodesCount:      1,
		IsEnd:              1,
		ChannelId:          channelId,
		Status:             1,
		RegionId:           regionId,
		TypeId:             typeId,
		EpisodesUpdateTime: nowTime,
		Comment:            0,
		UserId:             userId,
	}

	videoId, err := orm.NewOrm().Insert(&video)
	if err == nil {
		orm.NewOrm().Raw("insert into video_episodes (title, add_time, num, video_id, play_url, status, comment) values (?, ?, ?, ?, ?, ?, ?)", subTitle, nowTime, 1, videoId, playUrl, 1, 0).Exec()
	}

	return err
}

package userDao

import (
	"backend/dao/sql"
	"backend/models/userInfo"
)

func BeforeStar() {
	conn = sql.GetConn()
}

/*
Star
点赞
*/
func Star(videoId, username, author string) error {
	BeforeStar()
	var starModel userInfo.Star
	starModel.VideoId = videoId
	starModel.Username = username
	starModel.Author = author

	err := conn.Model(&userInfo.Star{}).Create(&starModel).Error
	return err
}

/*
WhetherStar
点赞
*/
func WhetherStar(videoId, username, author string) bool {
	BeforeStar()
	var starModel userInfo.Star
	starModel.VideoId = videoId
	starModel.Username = username
	starModel.Author = author
	var count int
	conn.Model(&userInfo.Star{}).Where(&starModel).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

/*
UnStar
删除点赞
*/
func UnStar(videoId, username string) error {
	BeforeStar()
	err := conn.Where("video_id = ? && username = ?", videoId, username).Delete(&userInfo.Star{}).Error
	return err
}

/*
UnStar
删除点赞
*/
func GetStar(videoId, username string) int {
	BeforeStar()
	var starModel userInfo.Star
	starModel.VideoId = videoId
	starModel.Username = username
	var result int
	conn.Model(&userInfo.Star{}).Where("video_id = ? && username = ?", videoId, username).Count(&result)
	return result
}

/*
GetStarCount
删除点赞
*/
func GetStarCount(videoId string) int {
	BeforeStar()
	var starModel userInfo.Star
	starModel.VideoId = videoId
	var result int
	conn.Model(&userInfo.Star{}).Where("video_id = ?", videoId).Count(&result)
	return result
}

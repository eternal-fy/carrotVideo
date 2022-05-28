package videoDao

import (
	"backend/dao/sql"
	"backend/models/videoInfo"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func Before() {
	conn = sql.GetConn()
	conn.AutoMigrate(&videoInfo.Video{})
}

/*
CreateVideo
根据用户名，更新user信息
*/
func CreateVideo(video *videoInfo.Video) error {
	Before()
	err := conn.Model(video).Create(video).Error
	return err
}

/*
SelectVideoByUsername，offsite查询
根据用户名、chang查询
*/
func SelectVideoByUsername(username, videoType string, page, size int) ([]videoInfo.Video, error) {
	Before()
	var video videoInfo.Video
	var result []videoInfo.Video
	video.Username = username
	video.Type = videoType
	typelike := "%" + videoType + "%"
	var err error
	if username != "" {
		err = conn.Model(&videoInfo.Video{}).Offset(size*(page-1)).Limit(size).Where("username = ? && type like ?", username, typelike).Order("created_at").Find(&result).Error
	} else {
		err = conn.Model(&videoInfo.Video{}).Offset(size*(page-1)).Limit(size).Where("type like ?", typelike).Order("created_at").Find(&result).Error
	}

	return result, err
}

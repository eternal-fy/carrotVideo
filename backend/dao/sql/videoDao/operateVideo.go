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

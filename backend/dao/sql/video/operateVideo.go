package video

import (
	"backend/dao/sql"
	"backend/models/videoInfo"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func Before() {
	if conn == nil {
		conn = sql.GetConn()
		conn.AutoMigrate(&videoInfo.Video{})
	}
}

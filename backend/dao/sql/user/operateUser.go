package user

import (
	"backend/dao/sql"
	. "backend/models/userInfo"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func init() {
	conn = sql.GetConn()
	conn.AutoMigrate(&User{})
}

func AddUser(user User) {
	conn.Model(&User{}).Save(&user)
}

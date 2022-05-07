package user

import (
	"backend/dao/sql"
	. "backend/models/userInfo"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func Before() {
	if conn == nil {
		conn = sql.GetConn()
		conn.AutoMigrate(&User{})
	}
}

func SaveUser(user User) {
	Before()
	conn.Model(&User{}).Save(&user)
}

/*
UserNameValidated
查看用户名是否可以被注册，如果数据库已经存在，则不可以
*/
func UserNameValidated(userName string) bool {
	Before()
	var num int
	conn.Model(&User{}).Where("userName = ?", userName).Count(&num)
	if num == 0 {
		return true
	}
	return false
}

package user

import (
	"backend/dao/bosService"
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

/*
SaveUser
根据用户名，更新user信息
*/
func SaveUser(user User, username string) {
	Before()
	conn.Model(&user).Where("username=?", username).Updates(user)
}

/*
CreateUser
根据用户名，更新user信息
*/
func CreateUser(user User) {
	Before()
	conn.Model(&user).Create(&user)
}

/*
GetUser
根据用户名，查询user信息
*/
func GetUser(username string) User {
	Before()
	var user User
	conn.Model(&User{}).Select("*").Where("username=?", username).Find(&user)
	return user
}

/*
GetUserByAppid
根据用户名，查询user信息
*/
func GetUserByAppid(appid string) User {
	Before()
	var user User
	conn.Model(&User{}).Select("*").Where("appid=?", appid).Find(&user)
	return user
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

/*
GetProfileImgUrl
获取用户头像url
*/
func GetProfileImgUrl(userName string) string {
	Before()
	user := GetUser(userName)
	bosName := user.Profileimgname
	url := bosService.BosGetUrl(bosName)
	return url
}

package userDao

import (
	"backend/dao/bosService"
	"backend/dao/sql"
	"backend/models/userInfo"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func Before() {
	conn = sql.GetConn()
}

/*
SaveUser
根据用户名，更新user信息
*/
func SaveUser(user userInfo.User, username string) {
	Before()
	conn.Model(&user).Where("username=?", username).Updates(user)
}

/*
CreateUser
根据用户名，更新user信息
*/
func CreateUser(user userInfo.User) error {
	Before()
	err := conn.Model(&user).Create(&user).Error
	return err
}

/*
GetUser
根据用户名，查询user信息
*/
func GetUser(username string) userInfo.User {
	Before()
	var user userInfo.User
	conn.Model(&userInfo.User{}).Select("*").Where("username=?", username).Find(&user)
	return user
}

/*
GetUserByAppid
根据用户名，查询user信息
*/
func GetUserByAppid(appid string) userInfo.User {
	Before()
	var user userInfo.User
	conn.Model(&userInfo.User{}).Select("*").Where("appid=?", appid).Find(&user)
	return user
}

/*
UserNameValidated
查看用户名是否可以被注册，如果数据库已经存在，则不可以
*/
func UserNameValidated(userName string) bool {
	Before()
	var num int
	conn.Model(&userInfo.User{}).Where("userName = ?", userName).Count(&num)
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
	if bosName == "" {
		bosName = "defaultIcon"
	}
	url := bosService.BosGetUrl(bosName)
	return url
}

/*
ChangeProfileImgUrl
获取用户头像url
*/
func ChangeProfileImgUrl(userName, url string) error {
	Before()
	var user userInfo.User
	user.Profileimgname = url
	err := conn.Model(&userInfo.User{}).Where("userName = ?", userName).Updates(user).Error
	return err
}

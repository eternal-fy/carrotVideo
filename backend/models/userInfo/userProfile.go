package userInfo

import "github.com/jinzhu/gorm"

//用户头像其他信息表
type UserProfile struct {
	gorm.Model
	Username   string `gorm:"unique;not null;index:addr"`
	ProfilePic string `gorm:"not null"` //用户头像图片的url
}

func (*UserProfile) TableName() string {
	return "user_profile"
}

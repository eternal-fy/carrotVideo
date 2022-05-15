package userInfo

import (
	"github.com/jinzhu/gorm"
)

//用户信息表
type User struct {
	gorm.Model
	Username       string `gorm:"unique;not null;index:addr"`
	Appid          string `gorm:"unique;"` //APPid对应的username生成为带有时间戳和appid的加密值
	Password       string `gorm:"not null"`
	Name           string
	Age            int32 `gorm:"default:120"`
	Gender         int32 `gorm:"default:0"`
	Level          int32 `gorm:"default:10"` //普通用户10、超级用户20以及管理员30
	Address        string
	Phone          string
	Profileimgname string
}

func (*User) TableName() string {
	return "user"
}

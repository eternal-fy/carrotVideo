package userInfo

import (
	"github.com/jinzhu/gorm"
)

//用户信息表
type User struct {
	gorm.Model
	Username  string `gorm:"unique;not null;index:addr"`
	Appid     string `gorm:"unique;"`
	Password  string `gorm:"not null"`
	Name      string
	Age       *int32
	Gender    int32  `gorm:"default:0"`
	Lever     *int32 //游客为空、普通用户10、超级用户20以及管理员30
	Address   string
	Phone     string
	StarCount int64 `gorm:"-"` //点赞数
}

func (*User) TableName() string {
	return "user"
}

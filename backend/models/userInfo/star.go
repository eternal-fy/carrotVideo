package userInfo

import "github.com/jinzhu/gorm"

//点赞表
type Star struct {
	gorm.Model
	VideoId  string `gorm:"not null"`
	Username string `gorm:"not null"` //点赞人的username
	Author   string `gorm:"not null"` //作者的username
}

func (*Star) TableName() string {
	return "star"
}

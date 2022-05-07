package userInfo

import "github.com/jinzhu/gorm"

// Follow 用户关注表
type Follow struct {
	gorm.Model
	Username string `gorm:"not null"` //点赞人的username
	Author   string `gorm:"not null"` //作者的username
}

package videoInfo

import "github.com/jinzhu/gorm"

// Message  留言、回复信息
type Message struct {
	gorm.Model        //其中video+ID为BOS中的对象名,image+ID为封面名称
	VideoId    string `gorm:"not null"`
	Username   string `gorm:"not null"`
	Message    string `gorm:"not null"`
	FollowId   string //如果为空，表示为1级留言
}

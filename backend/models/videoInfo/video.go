package videoInfo

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model         //其中ID为BOS中的对象名
	UserName    string `gorm:"not null"`
	Type        string //视频类型
	Title       string //视频名
	Description string //视频描述
	VideoLever  int32  `gorm:"default:0"` //看视频的权限，任何人都可以看为0、普通用户10、超级用户20
}

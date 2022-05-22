package videoInfo

import "github.com/jinzhu/gorm"

// Video 视频信息
type Video struct {
	gorm.Model
	Username    string `gorm:"not null"`
	VideoId     string `gorm:"not null"`
	ImageId     string `gorm:"not null"`
	Type        string //视频类型
	Title       string //视频名
	Description string //视频描述
	VideoLevel  int32  `gorm:"default:0"` //看视频的权限，任何人都可以看为0、普通用户10、超级用户20
}

func (*Video) TableName() string {
	return "video"
}

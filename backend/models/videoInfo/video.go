package videoInfo

import "github.com/jinzhu/gorm"

type Video struct {
	gorm.Model
	UserName    string
	Title       string //视频名
	Description string //视频描述
	BosName     string //Bos服务中存储的名称，一般是带时间的随机生成
	VideoLever  int32  `gorm:"default:0"` //看视频的权限，任何人都可以看为0、普通用户10、超级用户20
}

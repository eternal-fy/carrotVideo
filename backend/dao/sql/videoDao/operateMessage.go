package videoDao

import (
	"backend/dao/sql"
	dao "backend/dao/sql/userDao"
	"backend/models/videoInfo"
)

func BeforeMsg() {
	conn = sql.GetConn()
}

type MsgInfo struct {
	NickName string
	Desc     string
}

/*
PostMsg
添加留言
*/
func PostMsg(videoid, username, msg string) error {
	BeforeMsg()
	var msgModel videoInfo.Message
	msgModel.VideoId = videoid
	msgModel.Username = username
	msgModel.Message = msg
	err := conn.Model(&videoInfo.Message{}).Create(&msgModel).Error
	return err
}

/*
GetMsg
添加留言
*/
func GetMsg(videoid string) ([]MsgInfo, error) {
	BeforeMsg()
	var list []videoInfo.Message
	err := conn.Model(&videoInfo.Message{}).Order("created_at").Where("video_id = ?", videoid).Find(&list).Error
	result := make([]MsgInfo, len(list))
	for i, _ := range list {
		name := dao.GetUser(list[i].Username).Name
		result[i].NickName = name
		result[i].Desc = list[i].Message
	}
	return result, err
}

package test

import (
	"backend/dao/sql"
	"backend/dao/sql/userDao"
	"backend/models/userInfo"
	"testing"
)

func TestMaria(t *testing.T) {
	conn := sql.GetConn()
	println(conn)
}

func TestInsert(t *testing.T) {
	var userinf userInfo.User
	userinf.Username = "10"
	userinf.Password = "10"
	userDao.CreateUser(userinf)
}

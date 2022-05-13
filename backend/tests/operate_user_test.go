package test

import (
	dao "backend/dao/sql/user"
	"backend/models/userInfo"
	"testing"
)

func TestUpdateUser(t *testing.T) {
	var user userInfo.User
	user.ID = 6
	user.Age = 10
	dao.SaveUser(user, "ld")

}

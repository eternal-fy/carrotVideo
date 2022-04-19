package test

import (
	"backend/dao/sql"
	"testing"
)

func TestMaria(t *testing.T) {
	conn := sql.GetConn()
	println(conn)
}

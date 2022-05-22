package sql

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func GetConn() *gorm.DB {
	if Db != nil {
		return Db
	}
	host, _ := beego.AppConfig.String("maria-host")
	username, _ := beego.AppConfig.String("maria-username")
	password, _ := beego.AppConfig.String("maria-password")
	database, _ := beego.AppConfig.String("maria-database")
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, database)
	db, err := gorm.Open("mysql", source)
	if err != nil {
		panic("链接数据库失败")
	}
	return db
}

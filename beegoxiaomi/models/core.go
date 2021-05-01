package models

import (
	"github.com/astaxie/beego"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func init() {
	mysqladmin := beego.AppConfig.String("mysqladmin")
	mysqlpwd := beego.AppConfig.String("mysqlpwd")
	mysqldb := beego.AppConfig.String("mysqldb")
	dsn := mysqladmin + ":" + mysqlpwd + "@tcp(127.0.0.1:3306)/" + mysqldb + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		beego.Error()
	}
}

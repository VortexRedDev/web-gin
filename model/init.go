package model

import (
	"github/xiaoda-ye/web-gin/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB = Init()

func Init() *gorm.DB {
	conf := config.Getsetting().Database
	//url := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	url := conf.User + ":" + conf.Password + "@tcp(" + conf.Host + ")/" + conf.DBName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常")
	}
	_ = db.AutoMigrate(&User{})
	return db
}

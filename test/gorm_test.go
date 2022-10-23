package test

import (
	"fmt"
	"github/xiaoda-ye/web-gin/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	url := "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(url), &gorm.Config{})
	if err != nil {
		panic("数据库连接异常")
	}
	data := make([]*model.User, 0)
	err = db.Find(&data).Error
	//err = db.Exec("update user set create_date =now() ").Error
	if err != nil {
		panic("查询异常")
	}
	for _, v := range data {
		fmt.Printf("user==> %v \n\n", v)
	}
}

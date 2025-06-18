package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	CreateDate time.Time `gorm:"column:create_date;type:datetime;" json:"create_date"`
	Username   string `gorm:"column:username;type:varchar(255);" json:"username"`
	Email      string `gorm:"column:email;type:varchar(255);" json:"email"`
	Id 	   int `gorm:"column:id;type:int;primaryKey;autoIncrement;" json:"id"`
}

func (table *User) TableName() string {
	return "user"
}

func List(keyword string) *gorm.DB {
	return DB.Model(new(User)).
		Where(" username like ? or email like ?", "%"+keyword+"%", "%"+keyword+"%")
}
func Save(user User) *gorm.DB {
	return DB.Model(new(User)).Create(&user)
}

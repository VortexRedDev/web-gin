package model

import "gorm.io/gorm"

type User struct {
	ID         uint   `gorm:"primarykey;" json:"id"`
	Username   string `gorm:"column:username;type:varchar(255);" json:"username"`
	Email      string `gorm:"column:email;type:varchar(255);" json:"email"`
	CreateDate MyTime `gorm:"column:create_date;type:datetime;" json:"create_date"`
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

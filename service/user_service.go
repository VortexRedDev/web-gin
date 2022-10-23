package service

import (
	"github.com/gin-gonic/gin"
	model "github/xiaoda-ye/web-gin/model"
	"net/http"
	"strconv"
	"time"
)

func Ping(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(200, gin.H{"code": "500", "msg": err})
		return
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "2"))
	if err != nil {
		c.JSON(200, gin.H{"code": "500", "msg": err})
		return
	}
	keyword := c.Query("keyword")
	page = (page - 1) * size
	tx := model.List(keyword)
	var count int64
	data := make([]*model.User, 0)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		c.JSON(200, gin.H{"code": "500", "msg": err})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": data})
}
func Save(c *gin.Context) {
	c.Redirect(http.StatusFound, "")
	username := c.PostForm("username")
	email := c.PostForm("email")

	user := model.User{
		Username:   username,
		Email:      email,
		CreateDate: model.MyTime(time.Now()),
	}

	tx := model.Save(user)
	if tx.Error != nil {
		c.JSON(200, gin.H{"code": "500", "msg": tx.Error})
		return
	}
	c.JSON(200, gin.H{"code": 200, "data": tx.RowsAffected, "id": user.ID})
}

package router

import (
	"github.com/gin-gonic/gin"
	"github/xiaoda-ye/web-gin/service"
)

func Router(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(service.Auth)
	user := api.Group("/user")
	user.GET("/list", service.Ping)
	user.POST("/save", service.Save)
}

package router

import (
	"github.com/gin-gonic/gin"
	"github/xiaoda-ye/web-gin/service"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/statics", http.Dir("statics"))
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")

	api := r.Group("/api")

	user := api.Group("/user")
	user.GET("/list", service.Ping)
	user.POST("/save", service.Save)
	return r
}

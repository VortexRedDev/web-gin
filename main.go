package main

import (
	"github.com/gin-gonic/gin"
	"github/xiaoda-ye/web-gin/config"
	"github/xiaoda-ye/web-gin/log"
	"github/xiaoda-ye/web-gin/router"
	"net/http"
)

func main() {
	r := gin.New()
	r.Use(gin.LoggerWithWriter(log.LoggerWriter()), gin.Recovery())
	r.StaticFS("/statics", http.Dir("statics"))
	r.StaticFile("/favicon.ico", "./statics/favicon.ico")
	
	router.Router(r)
	r.Run(":" + config.Getsetting().Server.Port)
}
